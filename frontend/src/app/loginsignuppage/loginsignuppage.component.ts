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
   
    console.log("testprint1");

    usernameL: this.usernameL;
    passwordL: this.passwordL;

    if(this.usernameL.length > 0 && this.passwordL.length > 0) //ensure they are a valid input (this.usernameL.length > 0 && this.passwordL.length > 0)
    {
      console.log("Login Information: ");
      console.log("Username: ", this.usernameL);
      console.log("Password: ", this.passwordL);

      this.accountInfo.AddOnLogin(this.usernameL, this.passwordL)
        .then(response => {
          // Handle successful login
          //console.log("successfully passed in username and password");
        })
        .catch(error => {
          // Handle login error
          //console.log("username and password NOT passed to backend");
        });
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
   
    console.log("testprint1");

    usernameSU: this.usernameSU;
    passwordSU: this.passwordSU;

    if(this.usernameSU.length > 0 && this.passwordSU.length > 0) //ensure they are a valid input (this.usernameL.length > 0 && this.passwordL.length > 0)
    {
      console.log("Login Information: ");
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

