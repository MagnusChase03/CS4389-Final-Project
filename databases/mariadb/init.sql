USE CS4389;

CREATE TABLE Users(
    UserID int AUTO_INCREMENT,
    Username varchar(255) UNIQUE,
    PasswordHash varchar(255),
    PublicKey varchar(5000) UNIQUE,
    PRIMARY KEY (UserID)
);

INSERT INTO Users(
    Username,
    PasswordHash,
    PublicKey
) VALUES(
    'root',
    'supersecretpasswordhash',
    'rootpublickey'
);
