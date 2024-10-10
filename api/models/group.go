/* =========================================================================
*  File Name: models/group.go
*  Description: Handles database interactions related to groups.
*  Author: Matthew-Basinger
*  =======================================================================*/
package models;

import (
    "fmt"
    "github.com/MagnusChase03/CS4389-Project/db"
)

type Group struct {
    GroupID int
    CreatorID int
    GroupName string
};

/*
*  Returns a group with given GroupID
*
*  Arguments
*      - id (int): The GroupID to find
*
*  Returns: 
*      - Group: the group information
*      - error: An error if any occured
*/
func GetGroupByID(id int) (Group, error) {
    var group Group;
    instance, err := db.GetMariaDB();
    if err != nil {
        return group, fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err);
    }

    query, err := instance.Connection.Prepare("SELECT * FROM Groups WHERE GroupID = ?")
    if err != nil {
        return group, fmt.Errorf("[ERROR] Failed to get parse SQL query. %w", err);
    }
    defer query.Close();

    err = query.QueryRow(id).Scan(
        &group.GroupID, 
        &group.CreatorID, 
        &group.GroupName, 
    );
    if err != nil {
        return group, fmt.Errorf("[ERROR] Failed to find group with GroupID %d. %w", id, err);
    }

    return group, nil;
}

/*
*  Create a group in the database with given attributes.
*
*  Arguments:
*      - groupName (string): The name of the group.
*      - creatorID (int): The ID of the creator.
*  
*  Returns:
*      - error: An error if any occurred.
*
*/
func CreateGroup(creatorID int, groupName string) error {
    instance, err := db.GetMariaDB();
    if err != nil {
        return fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err);
    }

    insertStatement, err := instance.Connection.Prepare(
        "INSERT INTO Groups(CreatorID, GroupName) VALUES (?, ?)",
    );
    if err != nil {
        return fmt.Errorf("[ERROR] Failed to parse SQL query. %w", err);
    }
    defer insertStatement.Close();

    _, err = insertStatement.Exec(creatorID, groupName);
    if err != nil {
        return fmt.Errorf("[ERROR] Failed to create group. %w", err);
    }

    return nil;
}

/*
*  Delete a group in the database with given ID.
*
*  Arguments:
*      - groupID (int): The groupID.
*  
*  Returns:
*      - error: An error if any occurred.
*
*/
func DeleteGroup(groupname string) error {
    instance, err := db.GetMariaDB();
    if err != nil {
        return fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err);
    }

    deleteStatement, err := instance.Connection.Prepare(
        "DELETE FROM Groups WHERE GroupName = ?",
    );
    if err != nil {
        return fmt.Errorf("[ERROR] Failed to parse SQL query. %w", err);
    }
    defer deleteStatement.Close();

    _, err = deleteStatement.Exec(groupname);
    if err != nil {
        return fmt.Errorf("[ERROR] Failed to delete group. %w", err);
    }

    return nil;
}

/*
*  Locates a group in the database with given groupname.
*
*  Arguments:
*      - groupname (string): The groupname.
*  
*  Returns:
*      - error: An error if any occurred.
*
*/
func GetCreatorIDByGroupName(groupname string) (int, error) {
    var group Group;
    instance, err := db.GetMariaDB();
    if err != nil {
        return 0, fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err);
    }

    query, err := instance.Connection.Prepare("SELECT * FROM Groups WHERE GroupName = ?")
    if err != nil {
        return 0, fmt.Errorf("[ERROR] Failed to get parse SQL query. %w", err);
    }
    defer query.Close();

    err = query.QueryRow(groupname).Scan(
        &group.GroupID, 
        &group.CreatorID, 
        &group.GroupName, 
    );
    if err != nil {
        return 0, fmt.Errorf("[ERROR] Failed to find group with GroupID %d. %w", groupname, err);
    }

    return group.CreatorID, nil;
}