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

    console.log("testprint2: ", this.usernameL, " ", this.passwordL);


    this.accountInfo.AddOnLogin(this.usernameL, this.passwordL)
      .then(response => {
        // Handle successful login
        console.log("successfully passed in username and password");
      })
      .catch(error => {
        // Handle login error
        console.log("username and password NOT passed to backend");
      });
  }
}

