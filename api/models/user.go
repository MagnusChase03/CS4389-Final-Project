/* =========================================================================
*  File Name: models/user.go
*  Description: Handles database interactions related to users.
*  Author: MagnusChase03
*  =======================================================================*/
package models;

import (
    "fmt"
    "github.com/MagnusChase03/CS4389-Project/db"
)

type User struct {
    UserID int
    Username string
    PasswordHash string
    PublicKey string
};

/*
*  Returns a user with given UserID
*
*  Arguments:
*      - id (int): The UserID to find.
*
*  Returns:
*      - User: The user information.
*      - error: An error if any occurred.
*/
func GetUserByID(id int) (User, error) {
    var user User;
    instance, err := db.GetMariaDB();
    if err != nil {
        return user, fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err);
    }

    query, err := instance.Connection.Prepare("SELECT * FROM Users WHERE UserID = ?")
    if err != nil {
        return user, fmt.Errorf("[ERROR] Failed to get parse SQL query. %w", err);
    }
    defer query.Close();

    err = query.QueryRow(id).Scan(
        &user.UserID, 
        &user.Username, 
        &user.PasswordHash, 
        &user.PublicKey,
    );
    if err != nil {
        return user, fmt.Errorf("[ERROR] Failed to find user with UserID %d. %w", id, err);
    }

    return user, nil;
}
