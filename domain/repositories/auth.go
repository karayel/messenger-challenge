package repositories

import (
	"errors"

	"github.com/bahattincinic/messenger-challenge/domain/models"
)

// CreateAccessToken repository returns auth access Token
func CreateAccessToken(accessToken string, user models.User) models.Accesstoken {
	InsertRow(
		"INSERT INTO AccessTokens(access_token, user_id) VALUES (?,?)",
		accessToken, user.ID,
	)

	return models.Accesstoken{
		Token:  accessToken,
		UserID: user.ID,
	}
}

// GetUser returns specific user
func GetUser(username string, password string) (user models.User, err error) {
	row := fetchRows(
		"SELECT id, username, fullname from Users WHERE username = ? AND password = ?",
		username, password,
	)

	if row.Next() {
		errRow := row.Scan(&user.ID, &user.Username, &user.FullName)
		CheckErr(errRow)
	} else {
		err = errors.New("Invalid credentials")
	}
	return
}

// CreateUser repository creates user
func CreateUser(username string, password string, fullname string) int64 {
	objID := InsertRow(
		"INSERT INTO Users(username, password, fullname) values(?,?,?)",
		username, password, fullname,
	)
	return objID
}

// CheckAccessToken checks Access Token
func CheckAccessToken(token string) (accessToken models.Accesstoken, err error) {
	row := fetchRows(
		"SELECT access_token, user_id FROM AccessTokens WHERE access_token = ? LIMIT 1",
		token,
	)

	if row.Next() {
		err := row.Scan(&accessToken.Token, &accessToken.UserID)
		CheckErr(err)
	} else {
		err = errors.New("Invalid credentials")
	}

	return
}