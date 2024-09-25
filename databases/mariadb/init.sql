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
    'a7fa53aec98816e82823c7c04e6b01cc67e585213c5d71f2d58326e7513e4688',
    'rootpublickey'
);
