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


  //keep track of account info (username, password)
  //if this doesn't work add "public" in from of "accountInfo"
  accountInfo: interfaceAccountInfo[] = []




  //create a client for http! (used to pass front end variables to backend)
  constructor(private client: HttpClient) {}


  //make a method for ngOnInit
  async ngOnInit() {
    this.LoadAccountInfo()
  }


  //this is for getting the info
  async LoadAccountInfo()
  {
    //had to do it this method because "toPromise()" is becoming deprecated
    //these lines cause an issue "cannot GET"
    const valInfo = this.client.get<interfaceAccountInfo[]>('/login');
    this.accountInfo = await lastValueFrom(valInfo);
  }
 


  //this function is called on the log-in button press
  //this is for posting the info
  async AttemptLogin()
  {
    //(if it doesn't work maybe chnage it to '/rout/login')
    //Josue said he is implementing this part and to just send it to '/login'

    const infoToSend = this.client.post('/login', {
      //push the username and password that is inputted
      usernameForLoginAttempt: this.usernameForLoginAttempt,
      passwordForLoginAttempt: this.passwordForLoginAttempt
    })


    await lastValueFrom(infoToSend)


    // the below code shows that the username and password are properly accepted by the input and stored as a variable for the backend to use in their database management
    //use this until we can view the values being placed in the server!!
    console.log("Log In Account Info:");
    console.log("Username: ", this.usernameForLoginAttempt);
    console.log("Password: ", this.passwordForLoginAttempt);


    //use this to get the info we posted to the server (use await so it doesn't erase the inputted values until after this is called)
    await this.LoadAccountInfo()


    //resets the username and password back to '' to clear the input boxes and allow brand new inputs to be entered
    this.usernameForLoginAttempt = ''
    this.passwordForLoginAttempt = ''
  }




































  //this function is called on the sign-up button press
  AttemptSignUp()
  {
    //store the text from the input boxes in the following variables
    var usernameForSignUpAttempt = ((document.getElementById("signUpUsername") as HTMLInputElement).value);
    var passwordForSignUpAttempt = ((document.getElementById("signUpPassword") as HTMLInputElement).value);
   
    if(usernameForSignUpAttempt.length > 0 && passwordForSignUpAttempt.length > 0) //prevents a blank input
    {
      //add to the database in this if statement


      //make the input boxes return back to blank
      (document.getElementById("signUpUsername") as HTMLInputElement).value = "";
      (document.getElementById("signUpPassword") as HTMLInputElement).value = "";




      // the below code shows that the username and password are properly accepted by the input and stored as a variable for the backend to use in their database management
      console.log("Sign Up Account Info:");
      console.log("Username: ", usernameForSignUpAttempt);
      console.log("Password: ", passwordForSignUpAttempt);
    }
    else
    {
      //Don't add the username or password to the database here
      //print out "Please enter a valid username or password"
      //Work on this section when connected to backend to check for repeats
    }
  }




}


