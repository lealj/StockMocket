 describe('CEN Sprint 2 Test', () => {
  it('incorrectly inputs valid username and password for login', () => {
    
    cy.visit('localhost:4200') //visits the local host where the login page info is located
    cy.get('[data-cy="loginUNBox"]').type('asdfghkj') //this should find just the "Enter your username" box
    cy.get('[data-cy="loginPWBox"]').type('tesasdflkasdhgtFail') //this should find just the "Enter your username" box
    cy.get('button').contains('Log In').click() //this presses the button
    
    //this prints out POST 200 /login meaning that it is incorrect!!
    //this needs to print out POST 401 to show it doesn't find the account in the backend
    //mention this error in the report and video (will fix it in the next sprint)
    
    //following the cypress-recommended method of finding the components
    cy.get('[data-cy="loginUNBox"]').type('test') //this should find just the "Enter your username" box
    cy.get('[data-cy="loginPWBox"]').type('test') //this should find just the "Enter your username" box
    cy.get('button').contains('Log In').click() //this presses the button

    //this prints out POST 200 /login meaning that it is correct!!
 
    
    //following the cypress-recommended method of finding the components
    cy.get('[data-cy="signUpUNBox"]').type('testSignUpUN') //this should find just the "Enter your username" box
    cy.get('[data-cy="signUpPWBox"]').type('testSignUpPW') //this should find just the "Enter your username" box
    cy.get('button').contains('Sign Up').click() //this presses the button

    //this prints out POST 200 /login meaning that it is correct!!
  })
 })
 