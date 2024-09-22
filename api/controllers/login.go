/* =========================================================================
*  File Name: controller/login.go
*  Description: Controller for the login route.
*  Author: MagnusChase03
*  =======================================================================*/
package controllers

import (
    "fmt"

    "github.com/MagnusChase03/CS4389-Project/models"
    "github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Determines if a user with given credentials is correct.
*
*  Arguments:
*      - N/A
* 
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*/
func LoginController(username string, passwordHash string) (utils.JSONResponse, models.User, error) { 
    user, err := models.GetUserByCreds(username, passwordHash); 
    if err != nil {
        return utils.JSONResponse{
            StatusCode: 401,
            Data: "Invalid username or password.",
        }, user, fmt.Errorf("[ERROR] Invalid username or password. %w", err);
    }

    return utils.JSONResponse{
        StatusCode: 200,
        Data: "ok",
    }, user, nil;
}
