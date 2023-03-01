describe('CEN Sprint 2 Test', () => {
  it('incorrectly inputs valid username and password for login', () => {
    cy.visit('localhost:4200') //visits the local host where the login page info is located
    cy.get('input[id="loginUsername"]').type('testFail') //this should find just the "Enter your username" box
    cy.get('input[id="loginPassword"]').type('testFail') //this should find just the "Enter your username" box
    
    //following the cypress-recommended method of finding the components
    //cy.get('[data-cy="loginUNBox').type('testFail') //this should find just the "Enter your username" box
    //cy.get('[data-cy="loginPWBox').type('testFail') //this should find just the "Enter your username" box
    cy.get('button').contains('Log In').click() //this presses the button
    
    //this prints out POST 200 /login meaning that it is incorrect!!
      //this needs to print out POST 401 to show it doesn't find the account in the backend
  })
 })
 

describe('CEN Sprint 2 Test', () => {
  it('correctly inputs valid username and password for login', () => {
    cy.visit('localhost:4200') //visits the local host where the login page info is located
    //cy.get('input[id="loginUsername"]').type('test') //this should find just the "Enter your username" box
    //cy.get('input[id="loginPassword"]').type('test') //this should find just the "Enter your username" box
    
    //following the cypress-recommended method of finding the components
    cy.get('[data-cy="loginUNBox').type('test') //this should find just the "Enter your username" box
    cy.get('[data-cy="loginPWBox').type('test') //this should find just the "Enter your username" box
    cy.get('button').contains('Log In').click() //this presses the button
    //this prints out POST 200 /login meaning that it is correct!!
  })
 })
 




describe('CEN Sprint 2 Test', () => {
  it('inputs valid username and password for sign up', () => {
    cy.visit('localhost:4200') //visits the local host where the login page info is located
    //cy.get('input[id="loginUsername"]').type('test') //this should find just the "Enter your username" box
    //cy.get('input[id="loginPassword"]').type('test') //this should find just the "Enter your username" box
    
    //following the cypress-recommended method of finding the components
    cy.get('[data-cy="loginUNBox').type('testFail') //this should find just the "Enter your username" box
    cy.get('[data-cy="loginPWBox').type('testFail') //this should find just the "Enter your username" box
    cy.get('button').contains('Log In').click() //this presses the button
    //this prints out POST 200 /login meaning that it is correct!!
  })
 })
 