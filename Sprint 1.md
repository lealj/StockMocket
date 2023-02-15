# Welcome to Stock-Mocket!

Sprint 1
# Front-end user stories

 - A site visitor can view a company’s stocks on the line graph
    
-   A user can log in to save their account progress
    
-   A user can view how much money they currently have
    
-   A site visitor can search for stocks with a search bar
    
-   A visitor can click sign up to create an account on the website
    
-   A visitor can view a help page to learn about each feature of the website
    
-   A visitor can view the line graph by month, day, and year
    
-   A user can press a button that resets their money (has a confirmation box that pops up to ensure you want to actually reset it)
-   A visitor can see green text if you are earning money and red if you are losing money
- User can view stocks they own and stocks they want to keep watch on

## Back-end user stories
-   As a user, create a login to identify money and stocks associated with it
    
-   User can log in and know that their investments/other info is correct.
    
-   User knows stock information is reliable and current, as it is updated daily.
    
-   User is provided with accurate company data when user searches for it.
-  User can view accurate representation of stock market, as information is updated daily.
- As a site developer, there needs to be proper structure of source code to ensure efficient progress.
- As a user, data needs to be transferred to the system to be used for processing and system response.

## What issues your team planned to address

Backend team planned to address a login system, finding a populated stocks dataset, and implementing the necessary functions to supply the front end with stock information. To begin implementing these it was necessary to first set up a database and process json data.

The front end team planned to address setting up the login and sign up page as well as setting up the basic structure of the line graph we intend to use to display the stock data. We wanted the inputs to only be accepted if the text boxes (username and password) were not blank and were valid. For instance, we did not want the user to be able to create a new account that has the same name as one already created. To begin these sections, we first had to properly install node.js and angular cli and understand the basic style of writing in typescript and html to display the info on the screen and have buttons and textboxes interact with one another.
  

## Which ones were successfully completed

The back-end fundamentals for a login system were created, as we can create and store user information such as name and email, and provide user information to the front end upon request. Stock information can also be sent upon request. Of course, it was necessary to set up a local database and create an api to communicate with front-end requests for this.

The basic front-end structure of the login and sign up screen were implemented. The text boxes accept a user input and allow the user to type in the box. In addition, they can interact with the button that takes the input from the username and password boxes and stores them as variables that we can use to add to a database or check if the variables already exist in the database in coordination with the backend team (when we manage to connect the 2 team’s code). We successfully implemented a solution to the issue of the text boxes accepting blank inputs by adding a conditional statement that ensures the string has a length greater than 0. A simple line graph was implemented. In this graph, data points were displayed as dots at specific points on the lines (in line with the months put on the x-axis). The graphs successfully display the statistics of the stocks for the first few months.
## Which ones didn't and why?

back-end attempted to set up a mysql database hosted on AWS, and while the backend is able to connect to it successfully, there was trouble creating schema, tables etc. on it. This is why a large stock dataset was not implemented. The login system was not fully completed as getting used to new environments and structure was confusing.

Back-end still has not implemented a proper account info page with user money and current stocks, this will be achieved when a database is established to hold that information. Currently, the login and sign up attempts don't check for existing usernames in the database because the back-end and front-end code have not been linked. We have not implemented a search bar yet, it will be implemented when we have access to the stock market database. Currently the chart lacks intractability, an intricate chart system is in plan where you can see different time frames, but adding more functionality will be done when we learn more about using chart.js.

## Video Links
Front-end:
https://youtu.be/Dl2voAwoQLM

Back-end:
https://www.youtube.com/watch?v=x0YpvVrJ82Y



