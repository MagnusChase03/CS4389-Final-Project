/* =========================================================================
*  File Name: models/friend.go
*  Description: Handles database interactions related to friends.
*  Author: MagnusChase03
*  =======================================================================*/
package models;

import (
    "fmt"
    "github.com/MagnusChase03/CS4389-Project/db"
)

type FriendInvite struct {
    UserID int
    User2ID int
};

/*
*  Send an invite request to a specified user.
*
*  Arguments:
*      - userID (int): The UserID of the sender.
*      - id (int): The Username of the receiver.
*
*  Returns:
*      - error: An error if any occurred.
*/
func FriendRequestUser(userID int, username string) error {
    instance, err := db.GetMariaDB();
    if err != nil {
        return fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err);
    }

    insertStatement, err := instance.Connection.Prepare(
        `INSERT INTO FriendInvites(UserID, User2ID) Values(
            ?,
            (SELECT UserID FROM Users WHERE Username = ?) 
        )`,
    );
    if err != nil {
        return fmt.Errorf("[ERROR] Failed to parse SQL query. %w", err);
    }
    defer insertStatement.Close();

    _, err = insertStatement.Exec(userID, username);
    if err != nil {
        return fmt.Errorf("[ERROR] Failed to send friend request. %w", err);
    }

    return nil;
}
