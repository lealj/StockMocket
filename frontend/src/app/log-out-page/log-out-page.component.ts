import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { LogOutPageService } from './log-out-page.service'


@Component({
  selector: 'app-log-out-page',
  templateUrl: './log-out-page.component.html',
  styleUrls: ['./log-out-page.component.scss']
})

export class LogOutPageComponent implements OnInit
{
  
  title = 'Stock Mock-et';
  usernameL = '';
  passwordL = '';
  usernameD = '';
  deleteMessageToPrint = '';
  logOutMessageToPrint = '';
  resetMessageToPrint = '';
  public response: any;


  constructor(private accountInfo: LogOutPageService, private Routing: Router, private route: ActivatedRoute) {
    this.checkAuthorization();
  }

  checkAuthorization() {
    this.accountInfo.verify().then((isLoggedIn) => {
      if (isLoggedIn) {
        this.Routing.navigate(['/logout']);
      }
      else {
        this.Routing.navigate(['/account']);
      }
    });
  }

  ngOnInit() {
    usernameL: ''
    passwordL: ''
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
          console.log("The account username and password successfully deleted")
          this.deleteMessageToPrint = "Deletion Successful!";
          //added to make the input boxes blank after a delete attempt with valid input
          this.usernameD = '';

          //now that the account is successfully deleted, it should check if the account deleted was the current one (if it is log out) 

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


  AttemptLogOut()
  {

    this.accountInfo.logout().then((response) => {
      
      //if (response.status === 200) {
        console.log("Logged out of the account successfully")
        this.logOutMessageToPrint = "Logged Out Successfully!";
        //added to make the input boxes blank after a delete attempt with valid input
        //this.usernameD = '';
         
        // reload page so that user funds disappaer since user is not logged in
          this.Routing.navigate(['.'], { relativeTo: this.route }).then(() => {
            location.reload();
          });
        this.Routing.navigate(['/account']);
      //}
      
    }).catch((error) => {
        if (error.error === null)
        {
          //console.log("Correct response, error body is just not empty")
          console.log("Log out failed")
          this.deleteMessageToPrint = "Log Out failed";
          console.log(error);
        } else {
          console.log(error);
        }
      }
    );

  }

  ResetAccount()
  {
    this.accountInfo.resetAccount().then((response) => {
      console.log("Here we are")
      if (response.status === 200) {
        console.log("Reset account successfully")
        this.resetMessageToPrint = "Reset Successfully!";
      }
      
    }).catch((error) => {
        if (error.error === null)
        {
          //console.log("Correct response, error body is just not empty")
          console.log("Reset failed")
          this.resetMessageToPrint = "Reset failed";
          console.log(error);
        } else {
          console.log("Hello, reset failed")
          this.resetMessageToPrint = "Reset failed";
          console.log(error);
        }
      }
    );
  }

}
