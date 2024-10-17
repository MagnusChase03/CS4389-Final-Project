USE CS4389;

CREATE TABLE Users(
    UserID int AUTO_INCREMENT,
    Username varchar(255) UNIQUE,
    PasswordHash varchar(255),
    PublicKey TEXT(20000) UNIQUE,
    PRIMARY KEY (UserID)
);

CREATE TABLE Groups(
    GroupID int AUTO_INCREMENT,
    CreatorID int,
    GroupName varchar(255),
    PRIMARY KEY (GroupID),
    FOREIGN KEY (CreatorID) REFERENCES Users(UserID)
        ON DELETE CASCADE
);

CREATE TABLE UserGroup(
    UserID int,
    GroupID int,
    PRIMARY KEY(UserID, GroupID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
        ON DELETE CASCADE,
    FOREIGN KEY (GroupID) REFERENCES Groups(GroupID)
        ON DELETE CASCADE
);

CREATE TABLE GroupInvites(
    UserID int,
    GroupID int,
    EncryptedKey TEXT(20000),
    PRIMARY KEY(UserID, GroupID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
        ON DELETE CASCADE,
    FOREIGN KEY (GroupID) REFERENCES Groups(GroupID)
        ON DELETE CASCADE
);

CREATE TABLE FriendInvites(
    UserID int,
    User2ID int,
    PRIMARY KEY(UserID, User2ID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
        ON DELETE CASCADE,
    FOREIGN KEY (User2ID) REFERENCES Users(UserID)
        ON DELETE CASCADE
);

CREATE TABLE Friends(
    UserID int,
    User2ID int,
    PRIMARY KEY(UserID, User2ID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
        ON DELETE CASCADE,
    FOREIGN KEY (User2ID) REFERENCES Users(UserID)
        ON DELETE CASCADE
);

CREATE TABLE Messages(
    MessageID int AUTO_INCREMENT,
    UserID int,
    GroupID int,
    Timestamp datetime,
    EncryptedMessage TEXT(20000),
    PRIMARY KEY(MessageID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
        ON DELETE CASCADE,
    FOREIGN KEY (GroupID) REFERENCES Groups(GroupID)
        ON DELETE CASCADE
);

INSERT INTO Users(
    Username,
    PasswordHash,
    PublicKey
) VALUES(
    'root',
    'a7fa53aec98816e82823c7c04e6b01cc67e585213c5d71f2d58326e7513e4688',
    'rootpublickey'
);
