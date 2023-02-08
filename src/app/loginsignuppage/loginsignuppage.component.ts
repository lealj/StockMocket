import { Component } from '@angular/core';

@Component({
  selector: 'app-loginsignuppage',
  templateUrl: './loginsignuppage.component.html',
  styleUrls: ['./loginsignuppage.component.scss']
})
export class LoginsignuppageComponent {

  title = 'CENLogInPage';

  //this function is called on the log-in button press 
  AttemptLogin()
  {
    //store the text from the input boxes in the following variables
    var usernameForLoginAttempt = ((document.getElementById("loginUsername") as HTMLInputElement).value);
    var passwordForLoginAttempt = ((document.getElementById("loginPassword") as HTMLInputElement).value);

    if(usernameForLoginAttempt.length > 0 && passwordForLoginAttempt.length > 0)
    {
      //make the input boxes return back to blank
      (document.getElementById("loginUsername") as HTMLInputElement).value = "";
      (document.getElementById("loginPassword") as HTMLInputElement).value = "";

      // the below code shows that the username and password are properly accepted by the input and stored as a variable for the backend to use in their database management
      console.log("Log In Account Info:");
      console.log("Username: ", usernameForLoginAttempt);
      console.log("Password: ", passwordForLoginAttempt);
    }
    else
    {

    }
  }

  //this function is called on the sign-up button press 
  AttemptSignUp()
  {
    //store the text from the input boxes in the following variables
    var usernameForSignUpAttempt = ((document.getElementById("signUpUsername") as HTMLInputElement).value);
    var passwordForSignUpAttempt = ((document.getElementById("signUpPassword") as HTMLInputElement).value);
    
    if(usernameForSignUpAttempt.length > 0 && passwordForSignUpAttempt.length > 0)
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
