package service

import (
	"GFBackend/cache"
	"GFBackend/logger"
	"GFBackend/middleware/auth"
	"GFBackend/model"
	"GFBackend/model/dao"
	"GFBackend/utils"
	"errors"
	"fmt"
	"github.com/google/wire"
	"gorm.io/gorm"
	"strings"
	"sync"
)

var userManageServiceLock sync.Mutex
var userManageService *UserManageService

type IUserManageService interface {
	Register(username, password string, forAdmin bool) error
	Login(username, password string) (string, error)
	Logout(username string) error
	UpdatePassword(username, password, newPassword string) error
	Delete(username string) error
	Update(userInfo model.User) error
	Follow(followee, follower string) error
	Unfollow(followee, follower string) error
}

type UserManageService struct {
	userDAO   dao.IUserDAO
	followDAO dao.IFollowDAO
	spaceDAO  dao.ISpaceDAO
}

func NewUserManageService(userDAO dao.IUserDAO, followDAO dao.IFollowDAO, spaceDAO dao.ISpaceDAO) *UserManageService {
	if userManageService == nil {
		userManageServiceLock.Lock()
		if userManageService == nil {
			userManageService = &UserManageService{
				userDAO:   userDAO,
				followDAO: followDAO,
				spaceDAO:  spaceDAO,
			}
		}
		userManageServiceLock.Unlock()
	}
	return userManageService
}

var UserManageServiceSet = wire.NewSet(
	dao.NewUserDAO,
	wire.Bind(new(dao.IUserDAO), new(*dao.UserDAO)),
	dao.NewFollowDAO,
	wire.Bind(new(dao.IFollowDAO), new(*dao.FollowDAO)),
	dao.NewSpaceDAO,
	wire.Bind(new(dao.ISpaceDAO), new(*dao.SpaceDAO)),
	NewUserManageService,
)

func (userManageService *UserManageService) Register(username, password string, forAdmin bool) error {
	salt := utils.GetRandomString(6)
	newUser := model.User{
		Username: username,
		Password: utils.EncodeInMD5(password + salt),
		Salt:     salt,
	}

	err := model.DB.Transaction(func(tx *gorm.DB) error {
		createUserError := userManageService.userDAO.CreateUser(newUser, tx)
		if createUserError != nil {
			logger.AppLogger.Error(fmt.Sprintf("Create User Error: %s", createUserError.Error()))
			return createUserError
		}

		registrySpaceError := userManageService.spaceDAO.CreateSpaceInfo(username, tx)
		if registrySpaceError != nil {
			logger.AppLogger.Error(fmt.Sprintf("Create Space Info Error: %s", createUserError.Error()))
			return registrySpaceError
		}

		createDirFlag := utils.CreateDir(username)
		if !createDirFlag {
			logger.AppLogger.Error("Create Dir Error for user: " + username)
			return errors.New("500")
		}

		role := "regular"
		if forAdmin {
			role = "admin"
		}
		_, CasbinAddPolicyError := auth.CasbinEnforcer.AddGroupingPolicy(username, role)
		if CasbinAddPolicyError != nil {
			logger.AppLogger.Error(fmt.Sprintf("Add New User Policy Error: %s", CasbinAddPolicyError.Error()))
			return CasbinAddPolicyError
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (userManageService *UserManageService) Login(username, password string) (string, error) {
	dbUser := userManageService.userDAO.GetUserByUsername(username)
	if dbUser.Username == "" {
		return "", errors.New("400")
	}

	inputPassword := utils.EncodeInMD5(password + dbUser.Salt)
	if inputPassword != dbUser.Password {
		return "", errors.New("400")
	}

	token, err := auth.TokenGenerate(username)
	if err != nil {
		logger.AppLogger.Error(err.Error())
		return "", errors.New("500")
	}

	sign, _ := auth.GetTokenSign(token.Token)
	err = cache.AddLoginUserWithSign(username, sign)
	if err != nil {
		logger.AppLogger.Error(err.Error())
		return "", errors.New("500")
	}

	return token.Token, nil
}

func (userManageService *UserManageService) Logout(username string) error {
	err := cache.DelLoginUserSign(username)
	if err != nil {
		if err != nil {
			logger.AppLogger.Error(err.Error())
		}
		return errors.New("")
	}
	return nil
}

func (userManageService *UserManageService) UpdatePassword(username, password, newPassword string) error {
	user := userManageService.userDAO.GetUserByUsername(username)

	if utils.EncodeInMD5(password+user.Salt) != user.Password {
		return errors.New("400")
	}

	err := userManageService.userDAO.UpdateUserPassword(username, utils.EncodeInMD5(newPassword+user.Salt))
	if err != nil {
		logger.AppLogger.Error(err.Error())
		return errors.New("500")
	}

	return nil
}

func (userManageService *UserManageService) Delete(username string) error {
	user := userManageService.userDAO.GetUserByUsername(username)
	if user.Username == "" {
		return errors.New("user does not exist")
	}

	err := model.DB.Transaction(func(tx *gorm.DB) error {
		deleteUserError := userManageService.userDAO.DeleteUserByUsername(username, tx)
		if deleteUserError != nil {
			logger.AppLogger.Error(fmt.Sprintf("Delete User Error: %s", deleteUserError.Error()))
			return deleteUserError
		}

		if !utils.DeleteDir(username) {
			logger.AppLogger.Error("Delete User Error for user: " + username)
			return errors.New("500")
		}

		DeleteSpaceError := userManageService.spaceDAO.DeleteSpaceInfo(username, tx)
		if DeleteSpaceError != nil {
			logger.AppLogger.Error(DeleteSpaceError.Error())
			return DeleteSpaceError
		}

		deleteUserFollowError := userManageService.followDAO.DeleteFollow(username, tx)
		if deleteUserFollowError != nil {
			logger.AppLogger.Error(fmt.Sprintf("Delete User Error: %s", deleteUserError.Error()))
			return deleteUserFollowError
		}

		_, CasbinAddPolicyError := auth.CasbinEnforcer.DeleteUser(username)
		if CasbinAddPolicyError != nil {
			logger.AppLogger.Error(fmt.Sprintf("Delete User Policy Error: %s", CasbinAddPolicyError.Error()))
			return CasbinAddPolicyError
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil

}

func (userManageService *UserManageService) Update(userInfo model.User) error {
	err := userManageService.userDAO.UpdateUserByUsername(userInfo)
	if err != nil {
		logger.AppLogger.Error(err.Error())
		return errors.New("500")
	}
	return nil
}

func (userManageService UserManageService) Follow(followee, follower string) error {
	follow, err1 := userManageService.followDAO.GetOneFollow(followee, follower)
	if err1 != nil {
		if !strings.Contains(err1.Error(), "record not found") {
			logger.AppLogger.Error(err1.Error())
			return errors.New("500")
		}
	}

	if follow.Followee == followee {
		return errors.New("400")
	}

	followeeUserInfo := userManageService.userDAO.GetUserByUsername(followee)
	if followeeUserInfo.Username == "" {
		return errors.New("400")
	}

	err2 := userManageService.followDAO.UserFollow(followee, follower)
	if err2 != nil {
		logger.AppLogger.Error(err2.Error())
		return errors.New("500")
	}

	return nil
}

func (userManageService UserManageService) Unfollow(followee, follower string) error {
	followeeUserInfo := userManageService.userDAO.GetUserByUsername(followee)
	if followeeUserInfo.Username == "" {
		return errors.New("400")
	}

	err1 := userManageService.followDAO.UserUnfollow(followee, follower)
	if err1 != nil {
		logger.AppLogger.Error(err1.Error())
		return errors.New("500")
	}
	return nil
}
