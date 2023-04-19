# Stock-Mocket

Sprint 3
### video link: x
# Work Completed
## Front-end

Populated the portfolio page to show stocks owned by the user and relevent info regarding the stock. Displayed portfolio-value and change-in-portfolio-value and a history of the user's purchases and sales on the page as well. 
For Sprint 4, the front end team worked on the issues outlined on GitHub. One such issue was creating an “About us” page where there is information about the members of both teams. Another issue we worked on was creating a homepage that provided information about how to get started with the website, such as how to login, where to buy stocks, and how to view the aforementioned about us page.
Another issue the front end team worked on was creating a logout page (/logout) that would allow the user to log out of their account. We made the “Account” button in the top right corner route the user to the logout page when it was pressed and the user was already logged into their account. If the user is not currently logged into an account, the button routes you to the login page (/account) that was set up in the previous sprints. When routed to this page, the user has the option to logout by pressing the button on the webpage (which turns red when the mouse hovers over as a warning). Another option on this page is to reset your account by also selecting the button listed under the “Reset Account” section. Additionally, the “Delete User” option was moved to this page because the team felt it would be more appropriate for it to be in a section where a logged in user had access to the option. The format was updated where the text and button are red to deter the user from utilizing this functionality unless they truly wanted to. 
Additionally, the front end team utilized cypress end 2 end tests to test the functionality of the issues implemented. Initially, the tests created in the previous sprints were updated to match the new layout of the page, such as routing to the logout page to delete an account. One such test implemented was logging out of the account where it visits the logout page (/logout) after signing in on the login page (/account), and pressing the button, effectively logging the user out of the account.
In addition to the login tests, chart tests were also created to purchase and sell stocks. The first test simply created and logged into an account before visiting the chart page (/charts) where it pressed the buy button and purchased 2 stocks. This is successful since the user starts off with 1000 dollars and the current stock costs less than that amount. Then, it proceeds to test the sell feature by selling 1 owned stock. It succeeds because we owned 2 (now 1), then it sells another stock which succeeds for the same reason; however, when we try to sell 1 more, it fails because we do not have enough stocks to sell. Then, the following tests simply test the edge cases for the buy and sell features (more than allowed or greater than 50 and an invalid number that is either 0 or negative). These fail and send out an error message that the user needs to input a valid number of stocks. The final tests check the max amount the user is technically allowed to sell, but can’t simply because they don’t have enough money to purchase that many.

## Back-end
Implemented a reset-account feature for the user, removing their ownership of any stocks and resetting their funds amount to $1000. Stocks' prices are now updated when the program is run to ensure users view accurate prices. Created a portfolio-history file to keep records of users' purchases and sales, so a user can view this information through the website. The file also contains functions that calculate a user's portfolio value, portfolio-value change, and their gain/loss percent on each of their owned stocks, allowing the user to analyze their portfolio status. 

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
- ### TestUpdateStocks()
    - sets the prices of stocks to zero in the database then updates the prices. tests if the prices are still zero after calling the updatestock function. 
- ### TestQueryStocks()
    - tests the retrieval of real-time stock data pertaining to a ticker. 

### userstocks_test.go
- ### TestPurchaseStock()
    - tests the functionality of a user purchasing stock. 
- ### TestPurchaseStock_NoFunds()
    - tests the case where a user attempts to purchase a stock without having the funds, by checking the status returned (should be 400), and that no entries were added to database. 
- ### TestSellStock()
    - tests the functionality of a user selling stock. 
- ### TestSellStoc_NotOwned()
    - tests the case where a user attempts to sell a stock he doesn't own by checking the status code returned (should be 406). 
- ### TestSellStoc_NotOwned()
    - tests the case where a user attempts to more shares of a stock than he owns by checking the status returned (should be 408). 
- ### TestGetStocksOwned()
    - tests the retrieval of stocks owned by a particular user. 
- ### TestResetAccount()
    - tests that the users owned stocks were removed from the database and that their funds were reset to $1000. 

### credentials_test.go
- ### TestLogin()
    - tests if login attempt was successful.
- ### TestLoginFalse()
    - tests if the login attempt was rejected. 
- ### TestSignup()
    - tests the success of a user attempting to sign up. 
- ### TestSignupFail()
    - tests if the signup failed. 
- ### TestDeleteCredentials()
    - tests the deleteCredentials by requesting to delete a newly created account within delete. If signup passes, then there should be no issue.
- ### TestDeleteCredentialsFail()
    - tests deleteCredentials and expects a failed status since the test tries to delete an account with random typed letters, it should not be in the database.

### portfoliohistory_test.go
- ### TestGetLogs()
    - tests the retrieval of logs or records of a user's purchases and sales.
- ### TestGetUserPortfolioInfo()
    - tests the retrieval of a user's portfolio value and change in portfolio value. 
# Back-end API Documentation

## main.go
### Functions
#### httpHandler()
- Defines the API requests to be used in the project, setting the url expected, the function associated with the url, and the http request type. Returns handlers using CORS, allowing the transfer of data between two domains and circumventing security features implemented in modern browsers.
#### main()
- Calls InitialMigration() and UpdateStocks() and listens for requests from frontend.  


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
- Like the rest of our databases, credentials uses gorm.Model and contains just username, password and funds.
### Functions
#### login()
- When login() is called, it verifies that the username and password combination exits in the database. If it does, then 
it returns and http status code 200. If it does not, it will return code 401 for unauthorized
#### signup()
- This functions will create a new user account with a username and password. If the username is not take, it will return a 200 code. 
If the username is taken, then the code will return a 401 for unauthorized.
#### deleteCredentials()
- This will check if the username give is in the database, if it is, it will delete and return a 200 code. If the username
is not in the database, then the code returned is 409.


## stock.go
### Types
#### type Stock
- Uses gorm.Model. Contains a ticker and a price. 
#### type Query
- Contains a ticker and a start & end day, month, and year, to be used in the QueryStocks() function. 
### Functions
#### GetStockPrice()
- Returns price associated with stock ticker.
#### GetStocks()
- Returns all data pertaining to all stocks in json format.
#### GetStock()
- Returns all stock data pertaining to the ticker used as input in json format.
#### UpdateStocks()
- Updates stock-price data in the database for each stock stored in the database. Does not require json input.
#### QueryStocks()
- Returns daily stock price(s) and date(s) for a specific stock, over a start and end period. Requires json input data in the form of the Query struct. Returns this information in json format. 


## userstocks.go
### Types
#### type UserStocks
- Uses gorm.Model. Defines a username, ticker, and shares variables, and the json representations of each. In this context, it is meant to represent a purchase order. 
#### type StockTickerShares
- Used to store ticker, shares, price, and change variables pertaining to user's owned stock.  
### Functions
#### PurchaseStock()
- Receives input according to the UserStocks struct. Uses the username to access the funds of the user in the credentials table. Calculates the cost of the purchase order, considers if the user has enough funds. If so, the purchase order is completed, updating or creating an entry in the database containing the username, how many shares the user owns, and ticker of the owned stock. The users funds are updated as well. 
#### SellStock()
- Receives input according to the UserStocks struct. Follows the same format as the previous function, however, we check if the user owns shares of the stock they are attempting to sell. Once this is done, the sale is complete and the users funds are updated. The owned shares of the user are updated or deleted in the database. 
#### GetStocksOwned()
- Returns stock information regarding the stock the user owns. Uses the username as an input. Uses the StockTickerShares struct.
#### GetUserStocksArray()
- returns an array of stocks owned by the user. Uses the UserStocks struct. 
#### ResetAccount()
- Takes username as input. Removes records of ownership of any stocks, and from the database. Also removes logs associated with the account. User funds is reset to default value. 


## portfoliohistory.go
### Types
#### type PortfolioHistory
- Struct for storing username, ticker, shares, ordertype, and ordervalue
#### type PortfolioInfo
- Struct for storing portfolio value and change in portfolio value. 
### Functions
#### CreateLog()
- Inputs are ordertype string, UserStocks object, and ordervalue. Creates a log in the database recording a purchase/sell. 
#### DeleteAllLogs()
- Deletes all records in the database pertaining to the user. 
#### GetLogs()
- Returns all logs pertaining to the user in JSON format. 
#### GetUserPortfolioInfo()
- Encodes information into the PorfolioInfo struct and returns it in JSON format. Uses functions GetUserPortfolioValue().
#### GetUserPortfolioValue()
- Takes username as input. Sums the value of each stock per share that the user owns. Adds the user's funds to this amount and returns the total value. 
#### GetIndividualStockChange()
- Takes username, ticker, and shares as input. Gets the average purchase price or cost for each stock the user owns. Compares these values to current value of stock. Returns the change as a percent. 
