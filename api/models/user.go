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

/*
*  Returns a user with given Username.
*
*  Arguments:
*      - username (string): The Username to find.
*
*  Returns:
*      - User: The user information.
*      - error: An error if any occurred.
*/
func GetUserByUsername(username string) (User, error) {
    var user User;
    instance, err := db.GetMariaDB();
    if err != nil {
        return user, fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err);
    }

    query, err := instance.Connection.Prepare("SELECT * FROM Users WHERE Username = ?")
    if err != nil {
        return user, fmt.Errorf("[ERROR] Failed to get parse SQL query. %w", err);
    }
    defer query.Close();

    err = query.QueryRow(username).Scan(
        &user.UserID,
        &user.Username,
        &user.PasswordHash,
        &user.PublicKey,
    );
    if err != nil {
        return user, fmt.Errorf("[ERROR] Failed to find user with Username %d. %w", username, err);
    }

    return user, nil;
}

/*
*  Returns a user with given username and password.
*
*  Arguments:
*      - username (string): The username to match against.
*      - passwordHash (string): The password for the given user.
*
*  Returns:
*      - User: The user information.
*      - error: An error if any occurred.
*/
func GetUserByCreds(username string, passwordHash string) (User, error) {
    var user User;
    instance, err := db.GetMariaDB();
    if err != nil {
        return user, fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err);
    }

    query, err := instance.Connection.Prepare(
        "SELECT * FROM Users WHERE Username = ? AND PasswordHash = ?",
    )
    if err != nil {
        return user, fmt.Errorf("[ERROR] Failed to get parse SQL query. %w", err);
    }
    defer query.Close();

    err = query.QueryRow(username, passwordHash).Scan(
        &user.UserID,
        &user.Username,
        &user.PasswordHash,
        &user.PublicKey,
    );
    if err != nil {
        return user, fmt.Errorf(
            "[ERROR] Failed to find user %s with matching credentials. %w",
            username,
            err,
        );
    }

    return user, nil;
}

/*
*  Create a user in the database with given attributes.
*
*  Arguments:
*      - username (string): The username of the user.
*      - passwordHash (string): The hashed password for the user.
*      - publicKey (string): The public key of the user.
*  
*  Returns:
*      - error: An error if any occurred.
*
*/
func CreateUser(username string, passwordHash string, publicKey string) error {
    instance, err := db.GetMariaDB();
    if err != nil {
        return fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err);
    }

    insertStatement, err := instance.Connection.Prepare(
        "INSERT INTO Users(Username, PasswordHash, PublicKey) VALUES (?, ?, ?)",
    );
    if err != nil {
        return fmt.Errorf("[ERROR] Failed to parse SQL query. %w", err);
    }
    defer insertStatement.Close();

    _, err = insertStatement.Exec(username, passwordHash, publicKey);
    if err != nil {
        return fmt.Errorf("[ERROR] Failed to create user. %w", err);
    }

    return nil;
}

/*
*  Delete a user in the database with given ID.
*
*  Arguments:
*      - userID (int): The userID.
*  
*  Returns:
*      - error: An error if any occurred.
*
*/
func DeleteUser(userID int) error {
    instance, err := db.GetMariaDB();
    if err != nil {
        return fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err);
    }

    deleteStatement, err := instance.Connection.Prepare(
        "DELETE FROM Users WHERE UserID = ?",
    );
    if err != nil {
        return fmt.Errorf("[ERROR] Failed to parse SQL query. %w", err);
    }
    defer deleteStatement.Close();

    _, err = deleteStatement.Exec(userID);
    if err != nil {
        return fmt.Errorf("[ERROR] Failed to delete user. %w", err);
    }

    return nil;
}
