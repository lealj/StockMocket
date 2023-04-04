 describe('CEN Sprint 2 Test', () => {
  it('incorrectly inputs valid username and password for login', () => {
     
    
    
    //test the delete button on a username that does NOT exist in the DB already
     cy.visit('localhost:4200/account');
     cy.get('[data-cy="loginToDelete"]').type('bananas') //this should find just the "Enter a username to delete" box
     cy.get('button').contains('Delete').click() //this presses the button
     //should return 409 bc the username is NOT found in the DB so it should fail


    
    //enters a username and password to login that does NOT exists in the DB so should fail
    cy.visit('localhost:4200/account') //visits the local host where the login page info is located
    cy.get('[data-cy="loginUNBox"]').type('asdfghkj') //this should find just the "Enter your username" box
    cy.get('[data-cy="loginPWBox"]').type('tesasdflkasdhgtFail') //this should find just the "Enter your password" box
    cy.get('button').contains('Log In').click() //this presses the button
    //this prints out POST 401 /login meaning that it is correct!!
    //this needs to print out POST 401 to show it doesn't find the account in the backend
    //mention this error in the report and video (will fix it in the next sprint) <--error resolved !!



    //enters a username that exists and password to login that does NOT exists in the DB so should fail
    cy.visit('localhost:4200/account') //visits the local host where the login page info is located
    cy.get('[data-cy="loginUNBox"]').type('test') //this should find just the "Enter your username" box
    cy.get('[data-cy="loginPWBox"]').type('tesasdflkasdhgtFail') //this should find just the "Enter your password" box
    cy.get('button').contains('Log In').click() //this presses the button
    //this prints out POST 401 /login meaning that it is correct!!



    //enters a username and password to signup that already exists in the DB so should fail
    cy.visit('localhost:4200/account') //visits the local host where the login page info is located
    //following the cypress-recommended method of finding the components
    cy.get('[data-cy="signUpUNBox"]').type('test') //this should find just the "Enter a username" box
    cy.get('[data-cy="signUpPWBox"]').type('test') //this should find just the "Enter a password" box
    cy.get('button').contains('Sign Up').click() //this presses the button
    //this prints out POST 401 /login meaning that it is correct!! <-- returns 401 bc test-test is already in DB





    //enters a username and password to login that exists in the DB so should succeed
    cy.visit('localhost:4200/account') //visits the local host where the login page info is located
    //following the cypress-recommended method of finding the components
    cy.get('[data-cy="loginUNBox"]').type('test') //this should find just the "Enter your username" box
    cy.get('[data-cy="loginPWBox"]').type('test') //this should find just the "Enter your password" box
    cy.get('button').contains('Log In').click() //this presses the button
    //this prints out POST 200 /login meaning that it is correct!!

    


    /* ===== the following 2 e2e tests are used together to test one specific function of the code!! ===== */

    //enter a username and password to sign up that does NOT already exist in the DB so it should add it
    cy.visit('localhost:4200/account') //visits the local host where the login page info is located
    //following the cypress-recommended method of finding the components
    cy.get('[data-cy="signUpUNBox"]').type('testSignUpUN') //this should find just the "Enter a username" box
    cy.get('[data-cy="signUpPWBox"]').type('testSignUpPW') //this should find just the "Enter a password" box
    cy.get('button').contains('Sign Up').click() //this presses the button
    //this prints out POST 200 /login meaning that it is correct!!

    
    //enter a username and password to sign up that does already exist in the DB (already added from above test)
    cy.visit('localhost:4200/account') //visits the local host where the login page info is located
    //following the cypress-recommended method of finding the components
    cy.get('[data-cy="signUpUNBox"]').type('testSignUpUN') //this should find just the "Enter a username" box
    cy.get('[data-cy="signUpPWBox"]').type('testSignUpPW') //this should find just the "Enter a password" box
    cy.get('button').contains('Sign Up').click() //this presses the button
    //this prints out POST 401 /login meaning that it is correct!! 
    //the reason it prints 401 versus the above one is bc it is trying to add the un-pw to the database when it was just stored



    //test the delete button
    cy.visit('localhost:4200/account');
    cy.get('[data-cy="loginToDelete"]').type('testSignUpUN') //this should find just the "Enter a username to delete" box
    cy.get('button').contains('Delete').click() //this presses the button
    //should return 200 since the username is found in DB and successfully removed!


  })
 })
 