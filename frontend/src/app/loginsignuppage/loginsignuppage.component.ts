import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { LoginSignUpService } from "./loginsignuppage.service"
import { Router } from '@angular/router';
import { CookieServices } from "../cookie.service";
import {catchError} from "rxjs";

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
  usernameD = '';
  loginMessageToPrint = '';
  signUpMessageToPrint = '';
  deleteMessageToPrint = '';
  public response: any;

  //loginForm: FormGroup;


  constructor(private accountInfo: LoginSignUpService, private Routing: Router) {
    this.checkAuthorization();
  }

  checkAuthorization() {
    this.accountInfo.checkAuth().then((response) => {
      if (response.status === 200) {
        this.Routing.navigate(['/charts']);
      } else {

      }
    }).catch((error) => {
      if (error.error === null) {
        //console.log("Correct response, error body is just not empty")
        //moving the above 401 checker to this section since it was entering here on a 401 status return since it is considered an error instead!!
        console.log("You have beend logged out")

        //use this code to reset the values in the box back to blank and makes it empty on a failed attempt
        this.usernameL = '';
        this.passwordL = '';

      } else {
        console.log(error);
      }
    });
  }

  ngOnInit() {
    usernameL: ''
    passwordL: ''
  }

  AttemptLogin() {
    //initialize the variables with "this." to store their values for this function
    usernameL: this.usernameL;
    passwordL: this.passwordL;

    if (this.usernameL.length > 0 && this.passwordL.length > 0) //ensure they are a valid input (this.usernameL.length > 0 && this.passwordL.length > 0)
    {
      //print the info for the console log to see if the elements are being inputted correctly
      console.log("Login Information: ");
      console.log("Username: ", this.usernameL);
      console.log("Password: ", this.passwordL);

      // this sends the username and password that is passed in to the service which returns codes
      // 200 - okay, 401 - unauthorized, 400 - bad request, 502 - bad gateway
      this.accountInfo.AddOnLogin(this.usernameL, this.passwordL).then((response) => {
        if (response.status === 200)
        {
          console.log("Correct Credentials")
          this.loginMessageToPrint = '';
          //here is where we will route to the next page because we determine that the login attempt was successful (status of POST 200)!!
          this.Routing.navigate(['/charts']); //send it to the /charts that Freddy set up (which will be the main page)

        }
        /*
        else if (response.status === 401)
        {
          console.log("Wrong Credentials")
        }
        */
      }).catch((error) => {
          if (error.error === null) {
            //console.log("Correct response, error body is just not empty")
            //moving the above 401 checker to this section since it was entering here on a 401 status return since it is considered an error instead!!
            console.log("Wrong Credentials")
            this.loginMessageToPrint = 'Username or password was incorrect';

            //use this code to reset the values in the box back to blank and makes it empty on a failed attempt
            this.usernameL = '';
            this.passwordL = '';

          } else {
            console.log(error);
          }
        }
      );
    }
    else if(this.usernameL.length == 0 && this.passwordL.length == 0) //enter if BOTH username and password fields are blank
    {
      console.log("Username and password fields are blank. Inputs will not be sent to backend until fixed");
      this.loginMessageToPrint = 'Enter a username and password!';
    }
    else if (this.usernameL.length == 0) //enter if the username field is blank
    {
      console.log("Username field is blank. Inputs will not be sent to backend until fixed");
      this.loginMessageToPrint = 'Enter a username!';
    }
    else if (this.passwordL.length == 0) //enter if the username field is blank
    {
      console.log("Password field is blank. Inputs will not be sent to backend until fixed");
      this.loginMessageToPrint = 'Enter a password!';
    }

  }


  AttemptSignUp() {

    //console.log("testprint1"); //this is a test print to make sure it is entering the function correctly

    //initialize the variables with "this." to store their values for this function
    usernameSU: this.usernameSU;
    passwordSU: this.passwordSU;

    if (this.usernameSU.length > 0 && this.passwordSU.length > 0) //ensure they are a valid input (this.usernameL.length > 0 && this.passwordL.length > 0)
    {
      //print the info for the console log to see if the elements are being inputted correctly
      console.log("Sign Up Information: ");
      console.log("Username: ", this.usernameSU);
      console.log("Password: ", this.passwordSU);

      this.accountInfo.AddOnSignUp(this.usernameSU, this.passwordSU).then((response) => {
        if (response.status === 200) {
          console.log("The account username and password successfully created")
          this.signUpMessageToPrint = 'Account successfully created!';

          //added to make the input boxes blank after a sign up attempt with valid inputs
          this.usernameSU = '';
          this.passwordSU = '';
        }
        /*
        else if (response.status === 401) {
          console.log("Username is taken!")
        }
        */
      }).catch((error) => {
          if (error.error === null) //moving the above 401 checker to this section since it was entering here on a 401 status return since it is considered an error instead!!
          {
            //console.log("Correct response, error body is just not empty")
            console.log("Username is taken!")
            this.signUpMessageToPrint = 'Username is taken. Please try again!';
            console.log(error);
          } else {
            console.log(error);
          }
        }
      );
    }
    else if(this.usernameSU.length == 0 && this.passwordSU.length == 0) //enter if BOTH username and password fields are blank
    {
      console.log("Username and password fields are blank. Inputs will not be sent to backend until fixed");
      this.signUpMessageToPrint = 'Enter a username and password!';
    }
    else if (this.usernameSU.length == 0) //enter if the username field is blank
    {
      console.log("Username field is blank. Inputs will not be sent to backend until fixed");
      this.signUpMessageToPrint = 'Enter a username!';
    }
    else if (this.passwordSU.length == 0) //enter if the username field is blank
    {
      console.log("Password field is blank. Inputs will not be sent to backend until fixed");
      this.signUpMessageToPrint = 'Enter a password!';
    }
  }


  DeleteUser() //added to delete username-password combo from database
  {
    usernameD: this.usernameD;

    if (this.usernameD.length > 0)
    {
      console.log("Deletion Information: ");
      console.log(this.usernameD);

      this.accountInfo.DeleteUNandPW(this.usernameD).then((response) => {
        if (response.status === 200) {
          console.log("The account username and password successfully created")
          this.deleteMessageToPrint = "Deletion Successful!";
          //added to make the input boxes blank after a delete attempt with valid input
          this.usernameD = '';
        }
      }).catch((error) => {
          if (error.error === null)
          {
            //console.log("Correct response, error body is just not empty")
            console.log("No username found...Deletion failed")
            this.deleteMessageToPrint = "Deletion failed: Username was not found in the database!";
            console.log(error);
          } else {
            console.log(error);
          }
        }
      );
    }
    else if (this.usernameD.length == 0) //enter if the username field is blank
    {
      console.log("Username field is blank. Input will not be sent to backend until fixed");
      this.deleteMessageToPrint = "Please enter a username!";
    }
  }

}
