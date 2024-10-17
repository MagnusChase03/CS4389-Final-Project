/* =========================================================================
*  File Name: controller/userController/getUser.go
*  Description: Controller for getting user public key.
*  Author: MagnusChase03
*  =======================================================================*/
package userControllers

import (
    "fmt"

    "github.com/MagnusChase03/CS4389-Project/models"
    "github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to get the users public key.
*
*  Arguments:
*      - username (string): The username of the user.
* 
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*      - error: An error if any occurred.
*
*/
func GetUserController(username string) (utils.JSONResponse, error) { 
    user, err := models.GetUserByUsername(username);
    if err != nil {
        return utils.JSONResponse{
            StatusCode: 400,
            Data: "Failed to get user.",
        }, fmt.Errorf("[ERROR] Failed to get user. %w", err);
    }

    var responseStruct struct {
        PublicKey string
    };
    responseStruct.PublicKey = user.PublicKey;

    return utils.JSONResponse{
        StatusCode: 200,
        Data: responseStruct,
    }, nil;
}
