import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { LoginSignUpService } from "./loginsignuppage.service"
import { Router } from '@angular/router';
import { CookieServices } from "../cookie.service";

@Component({
  selector: 'app-loginsignuppage',
  templateUrl: './loginsignuppage.component.html',
  styleUrls: ['./loginsignuppage.component.scss']
})



export class LoginsignuppageComponent implements OnInit {
  title = 'Stock Mock-et';
  usernameL = '';
  passwordL = '';
  usernameSU = '';
  passwordSU = '';
  public response: any;

  //loginForm: FormGroup;
  

  constructor(private accountInfo: LoginSignUpService, private Routing: Router) {
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

          //here is where we will route to the next page because we determine that the login attempt was successful (status of POST 200)!! 
          this.Routing.navigate(['/charts']); //send it to the /charts that Freddy set up (which will be the main page)

        } 
        else if (response.status === 401) 
        {
          console.log("Wrong Credentials")
        }
      }).catch((error) => {
          if (error.error === null) {
            console.log("Correct response, error body is just not empty")
          } else {
            console.log(error);
          }
        }
      );
      //use this code to reset the values in the box back to blank and makes it empty
      //(if we get the above then catch part working, ideally I would only want to reset the values on an incorrect input and keep the values as we move onto the next page)
      this.usernameL = '';
      this.passwordL = '';

    } else if (this.usernameL.length == 0) //enter if the username field is blank
    {
      console.log("Username field is blank. Inputs will not be sent to backend until fixed");
    } else if (this.passwordL.length == 0) //enter if the username field is blank
    {
      console.log("Password field is blank. Inputs will not be sent to backend until fixed");
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
        } else if (response.status === 401) {
          console.log("Username is taken!")
        }
      }).catch((error) => {
          if (error.error === null) {
            console.log("Correct response, error body is just not empty")
            console.log(error);
          } else {
            console.log(error);
          }
        }
      );
    }
  }
}

