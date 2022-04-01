// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package reqs

import (
	"GFBackend/controller"
	"GFBackend/model/dao"
	"GFBackend/service"
)

// Injectors from wire.go:

func InitializeUserManageController() (*controller.UserManageController, error) {
	userDAO := dao.NewUserDAO()
	followDAO := dao.NewFollowDAO()
	spaceDAO := dao.NewSpaceDAO()
	communityMemberDAO := dao.NewCommunityMemberDAO()
	userManageService := service.NewUserManageService(userDAO, followDAO, spaceDAO, communityMemberDAO)
	userManageController := controller.NewUserManageController(userManageService)
	return userManageController, nil
}

func InitializeCommunityManageController() (*controller.CommunityManageController, error) {
	communityDAO := dao.NewCommunityDAO()
	communityMemberDAO := dao.NewCommunityMemberDAO()
	communityManageService := service.NewCommunityManageService(communityDAO, communityMemberDAO)
	communityManageController := controller.NewCommunityManageController(communityManageService)
	return communityManageController, nil
}

func InitializeFileManageController() (*controller.FileManageController, error) {
	spaceDAO := dao.NewSpaceDAO()
	fileManageService := service.NewFileManageService(spaceDAO)
	fileManageController := controller.NewFileManageController(fileManageService)
	return fileManageController, nil
}

func InitializeArticleTypeManageController() (*controller.ArticleTypeManageController, error) {
	articleTypeDAO := dao.NewArticleTypeDAO()
	articleTypeManageService := service.NewArticleTypeManageService(articleTypeDAO)
	articleTypeManageController := controller.NewArticleTypeManageController(articleTypeManageService)
	return articleTypeManageController, nil
}

func InitializeArticleManageController() (*controller.ArticleManageController, error) {
	articleDAO := dao.NewArticleDAO()
	articleTypeDAO := dao.NewArticleTypeDAO()
	communityDAO := dao.NewCommunityDAO()
	articleCommentDAO := dao.NewArticleCommentDAO()
	articleLikeDAO := dao.NewArticleLikeDAO()
	articleFavoriteDAO := dao.NewArticleFavoriteDAO()
	articleManageService := service.NewArticleManageService(articleDAO, articleTypeDAO, communityDAO, articleCommentDAO, articleLikeDAO, articleFavoriteDAO)
	articleManageController := controller.NewArticleManageController(articleManageService)
	return articleManageController, nil
}

func InitializeArticleLikeController() (*controller.ArticleLikeController, error) {
	articleDAO := dao.NewArticleDAO()
	articleLikeDAO := dao.NewArticleLikeDAO()
	articleLikeService := service.NewArticleLikeService(articleDAO, articleLikeDAO)
	articleLikeController := controller.NewArticleLikeController(articleLikeService)
	return articleLikeController, nil
}

func InitializeArticleFavoriteController() (*controller.ArticleFavoriteController, error) {
	articleFavoriteDAO := dao.NewArticleFavoriteDAO()
	articleDAO := dao.NewArticleDAO()
	articleFavoriteService := service.NewArticleFavoriteService(articleFavoriteDAO, articleDAO)
	articleFavoriteController := controller.NewArticleFavoriteController(articleFavoriteService)
	return articleFavoriteController, nil
}

func InitializeArticleCommentController() (*controller.ArticleCommentController, error) {
	articleCommentDAO := dao.NewArticleCommentDAO()
	articleDAO := dao.NewArticleDAO()
	articleCommentService := service.NewArticleCommentService(articleCommentDAO, articleDAO)
	articleCommentController := controller.NewArticleCommentController(articleCommentService)
	return articleCommentController, nil
}
