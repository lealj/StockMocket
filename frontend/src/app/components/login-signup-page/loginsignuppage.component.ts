import { Component, OnInit } from '@angular/core';
import { LoginSignUpService } from "../../services/page-services/loginsignuppage.service"
import { Router, ActivatedRoute } from '@angular/router';

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
  public response: any;

  //loginForm: FormGroup;


  constructor(private accountInfo: LoginSignUpService, private Routing: Router, private route: ActivatedRoute) {
    this.checkAuthorization();
  }
  checkAuthorization() {
    this.accountInfo.verify().then((isLoggedIn) => {
      if (isLoggedIn) {
        this.Routing.navigate(['/logout']);
      } else {
        this.Routing.navigate(['/account']);
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

          // reload page so that user funds appear since user is logged in
        this.Routing.navigate(['.'], { relativeTo: this.route }).then(() => {
          location.reload();
        });
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

  async getData() {
    const response = await this.accountInfo.claimData();
    console.log('Username:', response.username);
    console.log('Claims:', response);

    // This is an example of the data that can be obtained from HTTPOnly Cookies
    /* {username: 'test', role: 'admin', aud: 'test', exp: 1681760571, jti: '33', …}
     *  aud: "test"
     *  exp: 1681760571
     *  iat: 1681674171
     *  iss: "StockMocket"
     *  jti: "33"
     *  role: "admin"
     *  sub: "TheCODE"
     *  username: "test"
     */
  }
}
