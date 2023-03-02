import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { LoginSignUpService } from "./loginsignuppage.service"

@Component({
  selector: 'app-loginsignuppage',
  templateUrl: './loginsignuppage.component.html',
  styleUrls: ['./loginsignuppage.component.scss']
})

export class LoginsignuppageComponent implements OnInit 
{
  
  
  title = 'Stock Mock-et';
  usernameL = '';
  passwordL = '';
  usernameSU = '';
  passwordSU = '';

  //loginForm: FormGroup;

  constructor(private accountInfo: LoginSignUpService) { }

  ngOnInit() 
  {
      usernameL: ''
      passwordL: ''
  }

  AttemptLogin() {
   
    //console.log("testprint1"); //this is a test print to make sure it is entering the function correctly

    //initialize the variables with "this." to store their values for this function
    usernameL: this.usernameL; 
    passwordL: this.passwordL;

    if(this.usernameL.length > 0 && this.passwordL.length > 0) //ensure they are a valid input (this.usernameL.length > 0 && this.passwordL.length > 0)
    {
      //print the info for the console log to see if the elements are being inputted correctly
      console.log("Login Information: ");
      console.log("Username: ", this.usernameL);
      console.log("Password: ", this.passwordL);

      //this sends the username and password that is passed in to the service
      this.accountInfo.AddOnLogin(this.usernameL, this.passwordL) 
        .then(response => {
          // Handle successful login
          //console.log("successfully passed in username and password");

          //this is currently not working with the backend since they always return an error and not a response value even when it passes
          //the code actually works but its returning an error from the backend so it never enters this function (fix this next sprint if we need it)
        })
        .catch(error => {
          // Handle login error
          //console.log("username and password NOT passed to backend");
          //this is currently not working with the backend since they always return an error and not a response value even when it passes
          //the code actually works but its returning an error from the backend so it never enters this function (fix this next sprint if we need it)
        });

      //use this code to reset the values in the box back to blank and makes it empty 
      //(if we get the above then catch part working, ideally I would only want to reset the values on an incorrect input and keep the values as we move onto the next page)
      this.usernameL = '';
      this.passwordL = '';

    }
    else if(this.usernameL.length == 0) //enter if the username field is blank
    {
      console.log("Username field is blank. Inputs will not be sent to backend until fixed");
    }
    else if(this.passwordL.length == 0) //enter if the username field is blank
    {
      console.log("Password field is blank. Inputs will not be sent to backend until fixed");
    }

  }



  AttemptSignUp() {
   
    //console.log("testprint1"); //this is a test print to make sure it is entering the function correctly

    //initialize the variables with "this." to store their values for this function
    usernameSU: this.usernameSU;
    passwordSU: this.passwordSU;

    if(this.usernameSU.length > 0 && this.passwordSU.length > 0) //ensure they are a valid input (this.usernameL.length > 0 && this.passwordL.length > 0)
    {
      //print the info for the console log to see if the elements are being inputted correctly
      console.log("Sign Up Information: ");
      console.log("Username: ", this.usernameSU);
      console.log("Password: ", this.passwordSU);

      this.accountInfo.AddOnSignUp(this.usernameSU, this.passwordSU)
        .then(response => {
          // Handle successful login
          //console.log("successfully passed in username and password");
        })
        .catch(error => {
          // Handle login error
          //console.log("username and password NOT passed to backend");
        });
        
        //use this code to reset the values in the box back to blank and makes it empty 
        this.usernameSU = '';
        this.passwordSU = '';
    }
    else if(this.usernameSU.length == 0) //enter if the username field is blank
    {
      console.log("Username field is blank. Inputs will not be sent to backend until fixed");
    }
    else if(this.passwordSU.length == 0) //enter if the username field is blank
    {
      console.log("Password field is blank. Inputs will not be sent to backend until fixed");
    }

  }


}

