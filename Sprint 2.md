# Stock-Mocket

Sprint 2
# Work Completed
## Front-end
In this sprint, Brian Hersh edited the code that makes up the login and sign up page. Originally, the code was not well designed to work alongside a backend component, so this was fixed. While the HTML aspect did not change much, the component.ts file was heavily edited. Firstly, I decided to change the variable names to be more intuitive for other members of the group to recognize and utilize in their own implementations. Then, I created two separate functions (similar concept to how it was previously): one for an attempt at logging into an account and one for attempting to sign up an account. Inside the AttemptLogin function, I ensured that the current usernameL and passwordL were valid inputs (meaning they were not left blank); if they were blank it would print a console log message that said if the username was blank or if the password was blank. If one of these input boxes were blank, it would simply print the previously mentioned messages and not change the variables back to empty (essentially leaving the input in the box until 2 valid inputs were detected). In the case that they were both valid inputs, it would print a console log message stating what the values were and then pass these variables to “AddOnLogin.” This is the biggest change between the first sprint and this sprint: I added a service.ts file because it is proper procedure when dealing with backend html POSTs, GETS, etc. In this file, it creates an HttpClient with a constructor that is used to send the variables to the backend using POST and promise services. When it reaches the backend, they determine if the passed in variables match an existing username-password combination. If it does, it prints a “POST 200” terminal message suggesting that it succeeded. If it did not find an instance of this combo, it would return “POST 401” suggesting that it failed. The sign up aspect of my code follows a near identical process where it utilizes html, component.ts, and service.ts to POST the input variables to the backend. Although for this portion, the backend team will store the inputted username-password combinations into the database for the future. This returns the same 200 and 401 messages on success and failure respectively.

For testing, Brian Hersh used Cypress to create end-2-end testing by opening E2E testing in chrome and creating a new spec (I called it LoginPageSpec). Then, I opened it in my editor of choice (VS Code) and began writing my tests. To test the functionality of my code, I decided on 3 complete tests to do in order to show that my code is working correctly. The first test puts in a valid username and password, but the inputs do not exist in the database already. As such, it returns POST 401 since it does not find a match. The second test also plugs in a valid username and password, but this time the inputs exist in the database already; as a result, the second test returns POST 200 because it successfully finds a match. The third test works with the sign up section to ensure that the username and password (valid inputs) are able to be added to the database correctly; since they should be added properly, it returns POST 200.

For this sprint Freddy installed cypress for the project and created some unit test using native angular tools. Spent most of the time setting up component testing for cypress but to no success. Freddy setup e2e testing successfully and aided Brian in writing some cypress test to understand and get used to cypress. Freddy decided to use the spec files in the angular side of the project to run unit tests on the login component. I created one test to make sure the login textboxes were empty when the function is initially called. And another unit test to make sure the boxes emptied the input after pressing the login button. The tests are found in loginsignuppage.component.spec.ts and the tests are in the “it should” procedures.

## Back-end
The http handler function was changed to return CORS handlers to allow communication between the frontend(FE) and backend(BE). A proxy configuration was created to send http calls from FE to the BE.  A credentials file was made, that receives a username and password from the FE. If the login button was used, it will return a status code indicating whether the login attempt was successful based on if the username and password already exist in the database. If the sign up button was used, the username and password are stored in the database. Unit test files were created for functions regarding users, stocks, credentials, for which establishing a mock database was necessary. The api currently allows for servicing requests: GET, POST, PUT, DELETE, for tables “stocks”, “users”, and “credentials”, through which we use CORS to avoid errors due to security. In the api we also establish our database using gorm to set our database variable to use the dns for an AWS hosted MySQL server. 
 
## Front-end cypress/unit tests
LoginPageSpec.cy.js
- Test login with invalid username and password (verifies entry is not in database)
- Test login with valid username and password (verifies entry is in database)
- Test sign-up (adds entry to database)
- Test login textboxes empty (when initially called)
- Test login textboxes empty after pressing login button

## Back-end unit tests
user_test.go 
- TestGetUsers - returns users in mock database
- TestGetUser - returns user information based on user id
- TestCreateUser - creates user based on information provided
- TestUpdateUser - updates user information
- TestDeleteUser - deletes user (does not currently work, despite the function working in actual application). 

stock_test.go
- TestGetStocks - returns all stocks in mock database
- TestGetStock - returns stock based on id

credentials_test.go
- TestLogin - passes if the login was successful
- TestLoginFalse - passes if the login was not successful 
- TestSignup - gets username and password from frontend and creates a new credentials object in the database
- TestSignupFail - should fail if the body of the router request is invalid. During testing, if a value was missing, it would just add values present and leave missing values blanks. Test needs an update

## Video Links
https://youtu.be/M5LFYKBEVR8



