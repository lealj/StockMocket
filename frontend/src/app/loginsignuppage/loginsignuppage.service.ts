
//import { LoginsignuppageComponent } from "./loginsignuppage.component"
import { AccountInfo } from "./loginsignuppage"
import { HttpClient } from "@angular/common/http"

class LoginSignUpService {

    constructor(private client: HttpClient) {}

    //this one is for POSTing users on signup button press
    AddUserOnSignUp(usernameForSignUpAttempt: string, passwordForSignUpAttempt: string): Promise<interfaceAccountInfo[]>
    {
        return this.client.post<interfaceAccountInfo[]>("/signup",
        {
            usernameForSignUpAttempt, passwordForSignUpAttempt
        }).toPromise()
    }

    //this one is for POSTing users on login button press
    AddUserOnLogin(usernameForLoginAttempt: string, passwordForLoginAttempt: string): Promise<interfaceAccountInfo[]>
    {
        return this.client.post<interfaceAccountInfo[]>("/login",
        {
            usernameForLoginAttempt, passwordForLoginAttempt
        }).toPromise()
    }
}
