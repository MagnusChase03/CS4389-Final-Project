/* =========================================================================
*  File Name: routes/login.go
*  Description: Handler for logging in users.
*  Author: MagnusChase03
*  =======================================================================*/
package routes

import (
    "os"
    "fmt"
    "net/http"

    "github.com/MagnusChase03/CS4389-Project/utils"
    "github.com/MagnusChase03/CS4389-Project/controllers"
    "github.com/MagnusChase03/CS4389-Project/session"
)

/*
*  Handles the control flow for the login route.
*
*  Arguments:
*      - w (http.ResponseWriter): The object that is used to write a response.
*      - r (*http.Request): The request being made from the client.
* 
*  Returns:
*      - N/A
*/
func LoginHandler(w http.ResponseWriter, r *http.Request) { 
    if r.Method != "POST" {
        utils.SendBadRequest(w);
        return;
    }

    err := r.ParseForm();
    if err != nil {
        utils.SendBadRequest(w);
        return;
    }

    username := r.FormValue("username");
    password := r.FormValue("password");
    if username == "" || password == "" {
        utils.SendBadRequest(w);
        return;
    }

    resp, user, err := controllers.LoginController(username, password);
    if err != nil {
        fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err);
    } else {
        cookie, err := session.CreateUserCookie(user);
        if err != nil {
            utils.SendInternalServerError(w, err);
            return;
        }
        http.SetCookie(w, &cookie);
    }

    if err := utils.SendResponse(w, resp); err != nil {
        utils.SendInternalServerError(w, err);
    }
}
