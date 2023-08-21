# Sprint I

Frontend Demo: https://youtu.be/Gzo_MphEFjE

Backend Demo: https://youtu.be/aLl1goWGH6w


# Project Accomplishments:
# Backend:
1. Created a database on MongoDB and established the connection in the backend

2. Hosted the backend APIs on https://gatornews.herokuapp.com/

3. Created user model that contains the user information

4. Built Registration API which takes email and password to get the user started to use the application. The data is then stored in users collection in gatorNews DB

5. Login API was created which would take email, and password. The credentials are validated with the login credentials stored in Database. A JWT (JSON Web Token) is generated which is valid through the session

6. User profile can be retreived by calling View Profile API which takes JWT and User ID. User profile is output as a JSON object with all the non-empty fields populated

7. A user's profile can be modified with a call to Modify Profile API with JWT and user updated profile as a JSON object in the HTTP body

8. The logged in user can post a message (news or gossip). Post API can be used to post a message with a Authorized JWT and HTTP body with "message" property.

# Frontend:

1.	Created User Login and Registration System

2.	Added Validation For Login and Registration using API Endpoints Created from Backend to check User successful registration and Login. 

3.	Added Logout Functionality using Token System against the Server

4.	Added Routing Mechanism for responsive UI and to shift between frames efficiently 

5.	Designed and Created Pages and Routes for User Profile 

6.	Developed Hooks to retrieve User Information from Auth and User IDs

7.	Incorporated Buzz Functionality as part of the application for Registered Gators to Post information on the Feed. 

8.	Designed a Modal for obtaining the Tweets of the User and displaying them as part of their Profile Feed.
