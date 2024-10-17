/* =========================================================================
*  File Name: controller/userController/updateUser.go
*  Description: Controller for updating a user.
*  Author: MagnusChase03
*  =======================================================================*/
package userControllers

import (
    "fmt"

    "github.com/MagnusChase03/CS4389-Project/models"
    "github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Attempts to update a user with given attributes.
*
*  Arguments:
*      - username (string): The username of the user.
*      - passwordHash (string): The password hash to use for the user.
*      - publicKey (string): The public key of the user.
* 
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*      - error: An error if any occurred.
*
*/
func UpdateUserController(userID int, passwordHash string, publicKey string) (utils.JSONResponse, error) { 
    err := models.UpdateUser(userID, passwordHash, publicKey);
    if err != nil {
        return utils.JSONResponse{
            StatusCode: 400,
            Data: "Failed to update user.",
        }, fmt.Errorf("[ERROR] Failed to update user. %w", err);
    }

    return utils.JSONResponse{
        StatusCode: 200,
        Data: "Ok",
    }, nil;
}
