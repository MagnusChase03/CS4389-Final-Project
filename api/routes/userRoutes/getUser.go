/* =========================================================================
*  File Name: routes/userRoutes/getUser.go
*  Description: Handler for getting public key of user.
*  Author: MagnusChase03
*  =======================================================================*/
package userRoutes

import (
    "os"
    "fmt"
    "net/http"

    "github.com/MagnusChase03/CS4389-Project/utils"
    "github.com/MagnusChase03/CS4389-Project/controllers/userControllers"
)

/*
*  Handles the control flow for the get user route.
*
*  Arguments:
*      - w (http.ResponseWriter): The object that is used to write a response.
*      - r (*http.Request): The request being made from the client.
* 
*  Returns:
*      - N/A
*/
func GetUserHandler(w http.ResponseWriter, r *http.Request) { 
    if r.Method != "POST" {
        utils.SendBadRequest(w);
        return;
    }

    err := r.ParseForm();
    if err != nil {
        fmt.Printf("[ERROR] Failed to parse form.\n");
        utils.SendBadRequest(w);
        return;
    }

    username := r.FormValue("username");
    if username == "" {
        fmt.Printf("[ERROR] username is empty.\n");
        utils.SendBadRequest(w);
        return;
    }

    resp, err := userControllers.GetUserController(username);
    if err != nil {
        fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err);
    }

    if err := utils.SendResponse(w, resp); err != nil {
        utils.SendInternalServerError(w, err);
    }
}
