# Sprint II

Frontend Demo: https://youtu.be/5fnr_Sdbr-U

Backend Demo: https://youtu.be/PWXH4vVxg20

**[Backend Documentation](https://uflorida-my.sharepoint.com/:w:/g/personal/saikumar_malve_ufl_edu/Ec3aXKEQioVDtZ4B1u-44EEBdyqKsbeOWPcsdxbbxFbiOA?e=dSX1bv)**


# Project Accomplishments:
# Backend:
1. Wrote unit test cases for the following packages:
    - bd
    - jwt
    - models
    - routers
    - middlew
2. Built readPosts API which takes id and page as parameters to return all the posts on the user's feed. 
3. To generate a build after every merge or commit to main branch an action was created with the code in go-build.yml
4. The unit test cases written for the backend in Golang are located under Gator-News/UnitTests and to run the test cases after every commit to the main branch an action was created with the code in go-test.yml
# Documentation of API Services:
<details>
  <summary>Registration API</summary>
  
  - request :  **POST** <br />
  
  - [https://gatornews.herokuapp.com/registration/ ](https://gatornews.herokuapp.com/registration/)
  
  - **BODY** raw
  ``` json
  {
    "name": "Test",
    "lastname": "User",
    "email": "testuser@se.com",
    "password": "123456"
   }
   ```
  - **OUTPUT** 
  ``` json
    "Status : 201 Created"
  ```  

 
</details>

<details>
  <summary>Login API</summary>
  
  - request :  **POST** <br />
  
  - [https://gatornews.herokuapp.com/login/ ](https://gatornews.herokuapp.com/login/)
  
  - **BODY** raw
  ``` json
    { 
          "email": "testuser@se.com", 
          "password": "123456" 
     } 

   ```
  - **OUTPUT** 
  ``` json
    "Status : 201 Created"
  ```  
  ``` json
    {
    "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2MjIwMzM4YzlmMTg1MzQxY2U4YTdkOTMiLCJiaW9ncmFwaHkiOiIiLCJiaXJ0aGRhdGUiOiIyMDIyLTAzLTA0VDAxOjA3OjI1LjEyM1oiLCJlbWFpbCI6InRlc3R1c2VyQHNlLmNvbSIsImV4cCI6MTY0NjQ0Mjc0NiwibGFzdG5hbWUiOiJVc2VyIiwibG9jYXRpb24iOiIiLCJuYW1lIjoiVGVzdCIsIndlYnNpdGUiOiIifQ.7tD_XrxyEOLUkZXN-T4Az0yYTuzCW4bomwxouzRXsQM"
    }
   ```
 
</details>

<details>
 
  <summary>viewProfile API</summary>
  
  - request :  **GET** <br />
  
  - [https://gatornews.herokuapp.com/viewProfile?id=6220338c9f185341ce8a7d93 ](https://gatornews.herokuapp.com/viewProfile?id=6220338c9f185341ce8a7d93)
  
  - **OUTPUT** 
  ``` json
    "Status : 201 Created"
  ```  
  ``` json
  
    {"id":"6220338c9f185341ce8a7d93","name":"Test","lastname":"User","BirthDate":"2022-03-04T01:07:25.123Z","email":"testuser@se.com"}
   
   ```
 

 
</details>

<details>
  <summary>modifyProfile API</summary>
  
  - request :  **PUT** <br />
  
  - [https://gatornews.herokuapp.com/modifyProfile ](https://gatornews.herokuapp.com/modifyProfile)
  
  - **BODY** raw
  ``` json
       { 
            "name":"Test" , 
            "lastname": "User", 
            "birthDate": "1995-02-23T00:00:00Z", 
            "location": "Gainesville",
            "biography": "I am MS CS student", 
            "website": "https://github.com/SaiKumarMalve/Gator-News"
        }  

   ```
  - **OUTPUT** 
  ``` json
    "Status : 201 Created"
  ``` 
 
</details>

<details>
  <summary>Post API</summary>
  
  - request :  **POST** <br />
  
  - [https://gatornews.herokuapp.com/post ](https://gatornews.herokuapp.com/post)
  
  - **BODY** raw
  ``` json
       { 
            "message":"This is the Demo for testing purpose"
       }  

   ```
  - **OUTPUT** 
  ``` json
    "Status : 201 Created"
  ``` 
 
</details>

<details>
  <summary>readPosts API</summary>
  
  - request :  **GET** <br />
  
  - [https://gatornews.herokuapp.com/readPosts?id=6220338c9f185341ce8a7d93&page=1 ](https://gatornews.herokuapp.com/readPosts?id=6220338c9f185341ce8a7d93&page=1)
  - **OUTPUT** 
  ``` json
    "Status : 201 Created"
  ``` 
   ``` json
    [ {
        "_id": "622167e0a389bd529161b5db",
        "userID": "6220338c9f185341ce8a7d93",
        "message": "Test User Message",
        "date": "2022-03-04T01:14:08.847Z"
    },
    {
        "_id": "622167cceaa1be5c027fd250",
        "userID": "6220338c9f185341ce8a7d93",
        "message": "This is the Demo for testing purpose",
        "date": "2022-03-04T01:13:48.978Z"
    } ]
  ``` 
 
</details>



# Front-End:
- Created a Modal to display Tweets by User under their Profile Page 

- Added Pagination Mechanism under Basic Modal Component to Structure Tweets made using Buzz along with specific details (I.e username, time, tweet content) 

- Invoking the /readPosts API to fetch all the tweets sent by logged-In user 

- Validation based on API Endpoint Payload and displaying appropriate Alert text for User Logged In Session 

# Documentation for Cypress Testing:

- Incorporated Cypress framework for Automation Testing to write tests for web development in React
  <details>
	<summary>Login Form</summary>
	
	  Login_Cypress_Test.js:
	  - Contains tests to check User Login based on REST API Endpoint status codes using demo credentials
	
  </details>
  <details>
	<summary>Registration Form</summary>
	
	  Registration_Cypress_Testing.js 
	  - Contains tests with New User details which triggers the /register endpoint on submission. Checks if Registration is Successful with status code =“2XX” or Failed with status code = “4XX” || “5XX”

	
  </details>
   <details>
	<summary>User Profile</summary>
	
	  Buzz_Cypress_Testing.js 
	  - Validates if the Tweet Entered by User on Buzz Form Submission has invoked the “/post” Endpoint API successfully making the User Tweet Payload displayed as part of Profile along with previous entered ones using “/readPosts” API endpoints ok successful hits. 

	
  </details>

    

# Unit Testing using JEST for Front End:
- Implemented Unit Test Cases using Jest for testing JavaScript code written as part of React framework . 

- Test Cases are written for each of the below Components to test their functionality which is written as part of the “*.js” files as part of /sec/components/…”$ComponentName.test.js” <br />
- Jest Unit test cases:
	
	- Left Menu : Testing functions for Left Menu based on Individual links for Home, Users, Profile, Logout and Buzz <br />
	- ListTweets : Testing Functions for User Profile including Personal Details, Tweet data and Time zone to be displayed based on successful tweet   	 			submissions <br />
	- Modal : Contains two files one for BasicModal and TweetModal containing Unit Test Functions for Testing Functionalities on User Page. Testing Buzz		   	  Functionalities along with API Endpoints for sending and loading tweets as part of Profile Page <br />
	- SignInForm : Contains Unit Tests to check Successful Login based on correct credentials and vice versa <br />
	- SignUpForm : Contains Unit Tests to check Successful Registration based on proper validation cases for User name and email entered by the user
	- User : Contains two files one for BannerInfo and InfoUser. <br />
		- The Unit Test Cases part of Banner Info checks for the functionality of Individual Components such as “Edit Profile” and “Follow”. It also         	             includes  cases to check if the images for profile have been loaded successfully or not <br />
		- InfoUser Unit Test cases contains User Profile details Validation <br />

