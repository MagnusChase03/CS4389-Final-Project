/* =========================================================================
*  File Name: routes/logout.go
*  Description: Handler for logging out users.
*  Author: MagnusChase03
*  =======================================================================*/
package routes

import (
    "net/http"

    "github.com/MagnusChase03/CS4389-Project/utils"
    "github.com/MagnusChase03/CS4389-Project/controllers"
    "github.com/MagnusChase03/CS4389-Project/session"
)

/*
*  Handles the control flow for the logout route.
*
*  Arguments:
*      - w (http.ResponseWriter): The object that is used to write a response.
*      - r (*http.Request): The request being made from the client.
* 
*  Returns:
*      - N/A
*/
func LogoutHandler(w http.ResponseWriter, r *http.Request) { 
    if r.Method != "POST" {
        utils.SendBadRequest(w);
        return;
    }

    session.DeleteUserCookie(w);
    resp := controllers.LogoutController();
    if err := utils.SendResponse(w, resp); err != nil {
        utils.SendInternalServerError(w, err);
    }
}
