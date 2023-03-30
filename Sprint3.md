# Stock-Mocket

Sprint 3
### video link: https://youtu.be/2xmIv4r1p8o
# Work Completed
## Front-end
For Sprint 3, the front end team worked on completing the issues outlined on GitHub. One such issue was the routing from the login and sign up page to the main page “/charts” which had the stock data graph in it. We set it up so that if you input a valid username and password in the input boxes and pressed the login button, it would send those values to the backend where they would determine if the username-password combination existed in the database; if it existed, it would return a status code 200 (401 if unsuccessful) which the front end team would use to determine if the page should route. On a POST 200, the page would change to /charts because that was a successful login attempt. 

Another issue the front end team worked on was allowing the users to have the option to delete their account by entering their username in an input field and selecting the “delete” button. On a successful deletion, the back end team removed the username-password combination from the database and returned a POST 200 code, which the front end team used to print a console message confirming that the deletion was successful. If it was not successful (i.e. there was no username found in the database) it returned a POST 409 code that suggests that it was not found/the deletion was unable to be done.
The front end team also did some minor fixing on the console messages to ensure they were printing in the correct situations. One of the test functions that were not working has been fixed (an attempt at logging in with incorrect username-password returns POST 401 as intended now). The team also created more Cypress e2e tests to show the functionality of the code works for the new issues that were resolved.

The front end team also worked on creating the buy and sell buttons on the chart page to work with the functionality that has already been implemented in the backend. The frontend functionality was not successful, but with the new backend API file posted it would be done soon. Apart from that the front end implemented a header for the website which currently contains links to the chart page and to the login page, it will be expanded to include the profile page as well.

Josue worked on the functionality between the frontend and backend. This includes adding more proxy connections that would allow the frontend to access the stocks and userstocks http API. Josue also worked on implementing HTTP request code checks for login(), signup(), and delete() in the frontend. This allows proper routing for correct login combinations. All return 200 if action is correct.

## Back-end
Completed the back-end implementation that allows users to purchase or sell shares of stocks, given they have the funds to purchase or the shares to sell. Fixed the issue of duplicate user accounts being created upon signup. Added the ability for front-end (FE) to query prices and dates of any stock on the market, which will be useful for creating graphs in the future, that users will be able to analyze. Also added the ability for the prices of the stocks in the database to be updated to real-time market prices, ensuring that in the future, the user will be able to view accurate stock-price data. Credentials was updated to have a new delete function that returns a status code 200 if correct or 509 if the account does not exist. For testing, Josue added tests: TestdeleteCredentials() to make sure status code 200 is received when properly deleting. TestDeleteCredentialsFail() ensures a 409 is received if the requested account is not in the database. In addition to the new tests, Sprint 2 tests were updated to work. Some tests like TestSignUpFail() were always failing. All of the tests passed. Other tests were added to tests the functions, which are listed below. 

# Unit Tests
## Front-end
- Testing delete button on a username that does NOT exist in the database
- Attempts login with a username and password that does NOT exist in the database
- Attempts login with a username that exists in the database but a password that does not match the password for the username in the database
- Attempts to sign up with a username-password combination that already exists in the database
- Attempts login with a username and password that does exist in the database
- Attempts to sign up with a username-password combination that does not exist in the database yet. Then, it tries to sign up with that same combo again, so the username-password values are now in the database on this attempt.
- Testing delete button on a username that does exist in the database
Component test for buy and sell button: shows input box when buy button is pressed
## Back-end 
### stock_test.go
- ### TestGetStocks()
    - tests the retrieval of stock data stored in the database.
- ### TestGetStock()
    - tests the retrieval of a specific stocks data in the database, using ticker as input.
- ### TestUpdateStock()
    - sets the prices of stocks to zero in the database then updates the prices. tests if the prices are still zero after calling the updatestock function. 
- ### TestQueryStocks()
    - tests the retrieval of real-time stock data pertaining to a ticker. 
### userstocks_test.go
- ### TestPurchaseStock()
    - tests the funcionality of a user purchasing stock. 
- ### TestSellStock()
    - tests the functionality of a user selling stock. 
- ### TestGetStocksOwned()
    - tests the retrieval of stocks owned by a particular user. 
### credentials_test.go
- ### TestLogin()
    - tests if login attempt was successful.
- ### TestLoginFalse()
    - tests if the login attempt was rejected. 
- ### TestSignup()
    - tests the success of a user attempting to sign up. 
- ### TestSignupFail()
    - tests if the signup failed. 

# Back-end API Documentation
## main.go
### Functions
#### httpHandler()
- Defines the API requests to be used in the project, setting the url expected, the function associated with the url, and the http request type. Returns handlers using CORS, allowing the transfer of data between two domains and circumventing security features implemented in modern browsers.
#### main()
- Calls InitialMigration() and listens for requests from frontend.  

## dns.go
### Functions
#### InitialMigration()
- Create the database variable using gorm. Automatically creates tables for the Stock, UserStocks, and Credentials structs. Connection to the database is established elsewhere in the file. 

## angular_live.go
### Functions
#### getOrigin()
- Creates an url struct using localhost:4200 and returns a pointer to this url struct. 

## credentials.go
### Types
#### type Credentials
- Defines a username, password, and funds value, and their json representations. 
### Functions
#### login()
- Receives a username and password as json input. Checks if the combination already exists in the database, and returns a "status okay" http code if this is true. 
#### signup()
- Receives a username and password as json input. Checks if the combination already exists in the database. Returns a "status unauthorized" if username already exists. Otherwise we create the credentials entry in the database, and assign the default value for funds which is $1000. We return http "status okay" in the latter case. 
#### deleteCredentials()
- Receives a username as json input. If the username exists, we delete the credentials entry in the database, returning an http status okay code. 

## stock.go
### Types
#### type Stock
- Uses gorm.Model. Contains a ticker and a price. 
#### type Query
- Contains a ticker and a start & end day, month, and year, to be used in the QueryStocks() function. 
### Functions
#### GetStocks()
- Returns all data pertaining to all stocks in json format.
#### GetStock()
- Returns all stock data pertaining to the ticker used as input in json format.
#### UpdateStocks()
- Updates stock-price data in the database for each stock stored in the database. Does not require json input. Relies on the github.com/piquette/finance-go package, which uses the yahoo finance api to gather accurate market data. 
#### QueryStocks()
- Returns daily stock price(s) and date(s) for a specific stock, over a start and end period. Requires json input data in the form of the Query struct. Returns this information in json format. Relies on the github.com/piquette/finance-go package, which uses the yahoo finance api to gather accurate market data. 

## userstocks.go
### Types
#### type UserStocks
- Uses gorm.Model. Defines a username, ticker, and shares variables, and the json representations of each. In this context, it is meant to represent a purchase order.  
### Functions
#### PurchaseStock()
- Receives input according to the UserStocks struct. Uses the username to access the funds of the user in the credentials table. Calculates the cost of the purchase order, considers if the user has enough funds. If so, the purchase order is completed, updating or creating an entry in the database containing the username, how many shares the user owns, and ticker of the owned stock. The users funds are updated as well. 
#### SellStock()
- Receives input according to the UserStocks struct. Follows the same format as the previous function, however, we check if the user owns shares of the stock they are attempting to sell. Once this is done, the sale is complete and the users funds are updated. The owned shares of the user are updated or deleted in the database. 
#### GetStocksOwned()
- Returns the stock information regarding the shares the user owns. Uses the username as an input. 
