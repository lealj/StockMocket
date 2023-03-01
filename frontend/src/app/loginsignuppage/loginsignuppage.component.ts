import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { lastValueFrom } from 'rxjs'
import { last } from 'cypress/types/lodash';


//create an interface for account info
interface interfaceAccountInfo
{
  usernameForLoginAttempt: string
  passwordForLoginAttempt: string
}


@Component({
  selector: 'app-loginsignuppage',
  templateUrl: './loginsignuppage.component.html',
  styleUrls: ['./loginsignuppage.component.scss']
})


export class LoginsignuppageComponent implements OnInit {


  title = 'Stock Mock-et';
  //had to change the variables to this style because it wasn't in the proper format to be sent to the backend!
  //declaring it as "public" is not necessary bc that is implicit
  usernameForLoginAttempt = '';
  passwordForLoginAttempt = '';
  usernameForSignUpAttempt = '';
  passwordForSignUpAttempt = '';


  //keep track of account info (username, password)
  //if this doesn't work add "public" in from of "accountInfo"
  accountInfo: interfaceAccountInfo[] = []




  //create a client for http! (used to pass front end variables to backend)
  constructor(private client: HttpClient) {}


  //make a method for ngOnInit
  //may be unnecessary but keep in case I want to implement this later
  async ngOnInit() {
    //don't need the following line bc on initialization it shouldn't post anything (only post on button press)
    //this.LoadAccountInfo()
  }

//TA said the GET functionality wouldn't be necessary bc as a frontend, I am just POSTing the info to the backend
/*
  //this is for getting the info
  async LoadAccountInfo()
  {
    //had to do it this method because "toPromise()" is becoming deprecated
    //these lines cause an issue "cannot GET"
    const valInfo = this.client.get<interfaceAccountInfo[]>('/login');
    this.accountInfo = await lastValueFrom(valInfo);
  }
*/

  //this function is called on the log-in button press
  //this is for posting the info
  async AttemptLogin()
  {
    console.log("testprint1"); //used to debug and see if the error happens before this log

    if(this.usernameForLoginAttempt.length > 0 && this.passwordForLoginAttempt.length > 0) //check if the username or password field are left blank
    {
      //Josue said he is implementing this part and to just send it to '/login' (this is the POST call that sends the username and password values to backend)
      const loginInfoToSend = this.client.post('/login', {
      //push the username and password that is inputted
      usernameForLoginAttempt: this.usernameForLoginAttempt,
      passwordForLoginAttempt: this.passwordForLoginAttempt
      })
      console.log("testprint2"); //used to debug and see if the error happens before this log

      //await lastValueFrom(infoToSend)

      console.log("testprint3"); //used to debug and see if the error happens before this log

      // the below code shows that the username and password are properly accepted by the input and stored as a variable for the backend to use in their database management
      //use this until we can view the values being placed in the server!!
      console.log("Log In Account Info:");
      console.log("Username: ", this.usernameForLoginAttempt);
      console.log("Password: ", this.passwordForLoginAttempt);
      //resets the username and password back to '' to clear the input boxes and allow brand new inputs to be entered
      this.usernameForLoginAttempt = ''
      this.passwordForLoginAttempt = ''
    }
    else if (this.usernameForLoginAttempt.length == 0)
    {
      console.log("The username field was left blank...username and password NOT sent to backend");
    }
    else if (this.passwordForLoginAttempt.length == 0)
    {
      console.log("The password field was left blank...username and password NOT sent to backend");
    }

    //use this to get the info we posted to the server (use await so it doesn't erase the inputted values until after this is called)
    //await this.LoadAccountInfo()    
  }




































  //this function is called on the sign-up button press
  AttemptSignUp()
  {
    
    console.log("testprint1"); //used to debug and see if the error happens before this log

    if(this.usernameForSignUpAttempt.length > 0 && this.passwordForSignUpAttempt.length > 0) //check if the username or password field are left blank
    {
      //Josue said he is implementing this part and to just send it to '/login' (this is the POST call that sends the username and password values to backend)
      const signUpInfoToSend = this.client.post('/signup', {
      //push the username and password that is inputted
      usernameForSignUpAttempt: this.usernameForSignUpAttempt,
      passwordForSignUpAttempt: this.passwordForSignUpAttempt
      })
      console.log("testprint2"); //used to debug and see if the error happens before this log

      //await lastValueFrom(infoToSend)

      console.log("testprint3"); //used to debug and see if the error happens before this log

      // the below code shows that the username and password are properly accepted by the input and stored as a variable for the backend to use in their database management
      //use this until we can view the values being placed in the server!!
      console.log("Sign Up Account Info:");
      console.log("Username: ", this.usernameForSignUpAttempt);
      console.log("Password: ", this.passwordForSignUpAttempt);
      //resets the username and password back to '' to clear the input boxes and allow brand new inputs to be entered
      this.usernameForSignUpAttempt = ''
      this.passwordForSignUpAttempt = ''
    }
    else if (this.usernameForSignUpAttempt.length == 0)
    {
      console.log("The username field was left blank...username and password NOT sent to backend");
    }
    else if (this.passwordForSignUpAttempt.length == 0)
    {
      console.log("The password field was left blank...username and password NOT sent to backend");
    }

    //use this to get the info we posted to the server (use await so it doesn't erase the inputted values until after this is called)
    //await this.LoadAccountInfo()    
  }



}


