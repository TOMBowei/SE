This is SE Projects By Center of the Road Group

**Group members:**
- Frontend: 
	- Yuwei Xia(XIA1900): Designed and built all webpages, implemented all required functions, integrated with backend.
	- Kaiyue Wang(kywithsweet): Cypress test.
- Backend: Bowei Wu(TOMBowei), Yingjie Chen(fongziyjun16)

**Technological stack:**
- Frontend: React etc.
- Backend: Golang, GIN,  GORM.
- Database: Mysql, Redis, ElasticSearch..


## API Documentation

- http://167.71.166.120:10010/swagger/index.html

## Wiki Documentation
- https://github.com/fongziyjun16/SE/wiki

## Project Board Link
- https://github.com/fongziyjun16/SE/projects

## Sprint 4 Deliverables
- https://github.com/fongziyjun16/SE/tree/main/Sprint4

## How to run
- https://github.com/fongziyjun16/SE/wiki#how-to-run

## Demo video functionality

https://user-images.githubusercontent.com/90939944/164087559-24c308af-6fe5-425a-a714-28f19fa8d725.mp4

## Cypress test video

## Backend test video
https://user-images.githubusercontent.com/89665680/164123715-44a495b6-2f54-42f2-8b7c-42459bd919f9.mp4



# Gator Forum

## Description
Gator-forum is similar to any other forums, like Stack Overflow, Douban, Quora, Reddit… But we hope it is more student-life-related instead of society-related or limited to a specific field. 

In order to gather more students who have interests in common, students could set up interest groups and manage the group. There would be a specific description of the group in order to make people fully clear before they join the group. They can post articles in the group and look up all other articles from different groups which may help them to find more interests.

All articles could be liked, commented and favorited by anyone. It will be a good place to make new friends, whenever someone is interested in the others, they can just follow them to get more information about others. And because of the favorited function, good articles could be saved whenever students want to review them. Besides the search function is also complete, students could search articles or groups just by a few words they want, and the forum will respond to them with the possible answers it owned back to students.

Furthermore, the file system is also complete, each student has its own file system and limited file space for them to upload their avatars as they like. And so do the group model, the group manager could also do that to update the avatar for the group, which will keep the group active. 

A few gator forums have been discovered:
-   [https://insidethegators.com/forums/1](https://insidethegators.com/forums/1)
-   [https://247sports.com/college/florida/Board/Florida-Gators-Message-Board-Forum-14/](https://247sports.com/college/florida/Board/Florida-Gators-Message-Board-Forum-14/)
-   [https://florida.forums.rivals.com/forums/the-locker-room.14/](https://florida.forums.rivals.com/forums/the-locker-room.14/)
-   [https://gatorchatter.com/forums/main-sports-forum.20/](https://gatorchatter.com/forums/main-sports-forum.20/)

But they seem to be all sports-related. Although lots of students have a huge interest in sports, we do not want to only focus on sports. We want to create a place where people can not only discuss about sports, but also lots of things including their courses, their professors, their daily life, their hobbies and so on. They can share whatever they like and legal on Gator-forum.

## Users:

- regular users (UF students/employees)
- admin (Gator Forum Administrator)


## Components and functions:

-  If **not** signed in, **visitors** can only view posts and cannot interact with other users.
- If signed in as **regular users**, he/she can do following: 
	1. reply to others' posts
	2. reply to private messages
	3. create new post in group or in his/her own page
	4. subscribe/ unsubscribe a user
	5. join groups/communities
	6. anonymous post
	7. block/report/like/dislike/follow others users/posts
	8. setting personal information
	9. receive notifications, such as new followers, new posts, new replies etc
-   If signed as **group admin**:
	1. can close/delete a post
	2. can mute a user if he/she violates certain rules
	3. post rules
	4. deal with reports from users
	5. receive notifications including reports or requests from users

## Other Functions
- Can have different groups in which people of similar interests can discuss with each other. Groups can be "sports", "music", "rate-my-professor", "daily life", "courses" ... and so on. Users can join multiple groups and discuss with others in the group. Each group may have a few managers.
- Can earn points/badge based on their activation. To encourage them to post.
- Can switch among best/most-recent/hot posts.
- Can search for discussions related to key words.
- Can search for users.
