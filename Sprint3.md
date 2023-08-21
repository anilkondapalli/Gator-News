# Sprint III

Frontend Demo: https://www.youtube.com/watch?v=4HhCG85Rz2M

Backend Demo: https://youtu.be/WQyZxikPB5g

# Project Accomplishments:
# Backend:
1. Built removePosts API which takes id as parameter to delete all the user's posts.
2. highRelationship API was created to register a relationship between the users. In order to achieve this, We have a created "Relationship.go" model which records the relationship of one user with another.
3. In order to delete the relationship between the users, we have created lowRelationship API. Using this, we delete the relationship in the database.
4. userList API is used for retrieving the list of users in the database.This basically depends on 3 parameters - (type,page,search). If type is follow, it generates the list of users that the logged in user follows. If the type is new, it generates a list of all the new users. We can also search for a particular user by specifying search parameter.
5. readFollowersPosts API returns the posts of all the followers of the current user. It takes page as a parameter.
6. consultRelation API takes id(User ID) as a parameter. The API checks if the current user and the user whose id is provided are related (following) or not.
7. uploadAvatar and uploadBanner APIs take an image as input in the HTTP request body. It sets the current user's profile picture and display picture.
8. getAvatar and getBanner APIs take id as a parameter and returns the Avatar or Banner of the user whose id is provided
 
# Frontend:
1. Created Config Modal and Structure for Edit Profile under User Modal for Profile User to Update their Personal Information Content on Website. 

2. Invoked REST API calls for /getAvatar which sends the avatar as a Http Response and /getBanner  which sends the banner as part of the Http Response

3. Updating User Information in Backend using above REST Endpoints to fetch latest details on successful form submission. 

4. Created a Follow Modal for managing follow and unfollow requests made by website users to interact with each other.
 
5. Introduced 4 major Endpoints 
   <details>
        <summary>/consultRelation</summary>
     
          consultRelation API
          - To check the relationship between two users which sets a base relation for the follow/unfollow Modal
 
   </details>
   <details>
        <summary>/highRelationship</summary>
 
          highRelationship API
          - For Establishing a relationship between two users. This is a follow request API for users to be connected with each other and to be part of network
 
   </details>
   <details>
        <summary>/lowRelationship</summary>
 
          lowRelationship API
          - To disconnect the relationship between users. To Unfollow each other
 
   </details>
   <details>
        <summary>/userList</summary>
 
          userList API
          - To read list of users in the database. 
 
   </details>

6. Created a Routing System for Structuring the Users Page
7. Rendering All the followed Users which are part of the User Connection. 
8. Introduced Pagination Mechanism where each User Profile is displayed under Users Tab in a more visual way. 
9. Incorporated Error404 Page for users to get redirected when they navigate to any URL which is outside the website.
# Steps for running the project:
1. Install the node modules using "npm install" command.
2. To start the project , type "yarn start".
3. For testing unit test cases, "npm test" command is required.
4. For cypress testing, change test field in package.json to cypress open.
