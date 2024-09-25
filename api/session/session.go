/* =========================================================================
*  File Name: session/session.go
*  Description: A file containing common cookie functions for all modules.
*  Author: MagnusChase03
*  =======================================================================*/
package session

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "io"
    "net/http"
    "strconv"
    "strings"
    "time"

    "github.com/MagnusChase03/CS4389-Project/models"
    "github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Returns an encrypted user auth cookie.
*
*  Arguments:
*      - user (models.User): The user information to create the cookie.
*
*  Returns:
*      - http.Cookie: The encrypted cookie.
*      - error: An error if any occurred.
 */
func CreateUserCookie(user models.User) (http.Cookie, error) {
    env := utils.GetEnvironment(); 
    block, err := aes.NewCipher([]byte(env["COOKIE_KEY"]));
    if err != nil {
        return http.Cookie{}, fmt.Errorf("[ERROR] Failed to create AES block. %w\n", err);
    }

    gcm, err := cipher.NewGCM(block);
    if err != nil {
        return http.Cookie{}, fmt.Errorf("[ERROR] Failed to create GCM cipher. %w\n", err);
    }

    nonce := make([]byte, gcm.NonceSize());
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return http.Cookie{}, fmt.Errorf("[ERROR] Failed to create GCM nonce. %w\n", err);
    }

    plaintext := fmt.Sprintf("%d-%s", user.UserID, user.Username);
    ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil);

    return http.Cookie{
        Name: "authCookie",
        Value: base64.StdEncoding.EncodeToString(ciphertext),
        Expires: time.Now().Add(time.Hour),
    }, nil;
}

/*
*  Deleted the user auth cookie, used for logout.
*
*  Arguments:
*      - w (http.ResponseWriter): The response writer.
*
*  Returns:
*      - N/A
*/
func DeleteUserCookie(w http.ResponseWriter) {
    http.SetCookie(w, &http.Cookie{
        Name: "authCookie",
        Value: "",
        Expires: time.Now(),
    });
}

/*
*  Returns the user information from an encrypted cookie value.
*
*  Arguments:
*      - value (string): The encrypted cookie value.
*
*  Returns:
*      - int: The UserID.
*      - string: The Username.
*      - error: An error if any occurred.
*/
func ParseUserCookie(value string) (int, string, error) {
    ciphertext, err := base64.StdEncoding.DecodeString(value);
    if err != nil {
        return 0, "", fmt.Errorf("[ERROR] Failed to parse base64 value. %w\n", err);
    }

    env := utils.GetEnvironment(); 
    block, err := aes.NewCipher([]byte(env["COOKIE_KEY"]));
    if err != nil {
        return 0, "", fmt.Errorf("[ERROR] Failed to create AES block. %w\n", err);
    }

    gcm, err := cipher.NewGCM(block);
    if err != nil {
        return 0, "", fmt.Errorf("[ERROR] Failed to create GCM cipher. %w\n", err);
    }

    nonceSize := gcm.NonceSize();
    if len(ciphertext) < nonceSize {
        return 0, "", fmt.Errorf("[ERROR] Failed to decrypt ciphertext.\n");
    }

    nonce, data := ciphertext[:nonceSize], ciphertext[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, data, nil);
    if err != nil {
        return 0, "", fmt.Errorf("[ERROR] Failed to decrypt ciphertext. %w\n", err);
    }

    userFields := strings.Split(string(plaintext), "-");
    userID, err := strconv.ParseInt(userFields[0], 10, 64);
    if err != nil {
        return 0, "", fmt.Errorf("[ERROR] Failed to parse UserID from plaintext. %w\n", err);
    }

    return int(userID), userFields[1], nil;
}
