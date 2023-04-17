describe('template spec', () => {
  it('passes', () => {

      //create the account and sign into it
      cy.visit('localhost:4200/account') //visits the local host where the login page info is located
      //following the cypress-recommended method of finding the components
      cy.get('[data-cy="signUpUNBox"]').type('testSignUpUN') //this should find just the "Enter a username" box
      cy.get('[data-cy="signUpPWBox"]').type('testSignUpPW') //this should find just the "Enter a password" box
      cy.get('button').contains('Sign Up').click() //this presses the button

      //enters a username and password to login that exists in the DB so should succeed
      cy.visit('localhost:4200/account') //visits the local host where the login page info is located
      //following the cypress-recommended method of finding the components
      cy.get('[data-cy="loginUNBox"]').type('testSignUpUN') //this should find just the "Enter your username" box
      cy.get('[data-cy="loginPWBox"]').type('testSignUpPW') //this should find just the "Enter your password" box
      cy.get('button').contains('Log In').click() //this presses the button
      //this prints out POST 200 /login meaning that it is correct!!



      //now click the buy button on the charts page...
      cy.visit('localhost:4200/charts')
      cy.get('button').contains('Buy').click() //this presses the button
      cy.get('[data-cy="buyButton"]').clear().type('2') //type in 2 into the quantity tab to buy 2 of the stock
      cy.get('button').contains('Confirm Purchase').click() //this presses the button


      cy.get('button').contains('Sell').click() //this presses the button
      cy.get('[data-cy="sellButton"]').clear().type('1') //type in 1 into the quantity tab to sell 1 of the stock
      cy.get('button').contains('Confirm Sale').click() //this presses the button


      //use this to bring it back to a net 0 stocks
      cy.get('button').contains('Sell').click() //this presses the button
      cy.get('[data-cy="sellButton"]').clear().type('1') //type in 1 into the quantity tab to buy 1 of the stock
      cy.get('button').contains('Confirm Sale').click() //this presses the button
      //the above cases should pass





      //the following test cases should all fail



      

      //test in the case of a value greater than the allowed 50 for each purchase
      cy.get('button').contains('Buy').click() //this presses the button
      cy.get('[data-cy="buyButton"]').clear().type('100') //type in 100 into the quantity tab to buy 100 of the stock
      cy.get('button').contains('Confirm Purchase').click() //this presses the button
      //should fail here and say insufficient funds

      //test in the case of 0 being inputted
      cy.get('button').contains('Buy').click() //this presses the button
      cy.get('[data-cy="buyButton"]').clear().type('0') //type in 0 into the quantity tab to buy 0 of the stock
      cy.get('button').contains('Confirm Purchase').click() //this presses the button
      //should fail here and say "Invalid quantity input"

      //test in the case of a negative number being inputted
      cy.get('button').contains('Buy').click() //this presses the button
      cy.get('[data-cy="buyButton"]').clear().type('-20') //type in -20 into the quantity tab to buy -20 of the stock
      cy.get('button').contains('Confirm Purchase').click() //this presses the button
        //should fail here and say "Invalid quantity input"


      //testing the sales portion:


      //test in the case of a value greater than the allowed 50 for each sale
      cy.get('button').contains('Sell').click() //this presses the button
      cy.get('[data-cy="sellButton"]').clear().type('100') //type in 100 into the quantity tab to buy 100 of the stock
      cy.get('button').contains('Confirm Sale').click() //this presses the button
      //should fail here and say "Invalid quantity input"

      //test in the case of 0 being inputted
      cy.get('button').contains('Sell').click() //this presses the button
      cy.get('[data-cy="sellButton"]').clear().type('0') //type in 0 into the quantity tab to buy 0 of the stock
      cy.get('button').contains('Confirm Sale').click() //this presses the button
      //should fail here and say "Invalid quantity input"

      //test in the case of a negative number being inputted
      cy.get('button').contains('Sell').click() //this presses the button
      cy.get('[data-cy="sellButton"]').clear().type('-20') //type in -20 into the quantity tab to buy -20 of the stock
      cy.get('button').contains('Confirm Sale').click() //this presses the button
      //should fail here and say "Invalid quantity input"
      



      //test in the case of a valid number that is still more than the current money owned
      cy.get('button').contains('Buy').click() //this presses the button
      cy.get('[data-cy="buyButton"]').clear().type('50') //type in 50 into the quantity tab to buy 50 of the stock
      cy.get('button').contains('Confirm Purchase').click() //this presses the button
      //should fail here and say "Insufficient funds"


      //test in the case of a valid number that is still more than the current stocks owned
      cy.get('button').contains('Sell').click() //this presses the button
      cy.get('[data-cy="sellButton"]').clear().type('50') //type in 50 into the quantity tab to buy 50 of the stock
      cy.get('button').contains('Confirm Sale').click() //this presses the button
      //should fail here and say "You are attempting to sell more shares than you own"

  })
})