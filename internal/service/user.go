package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/dhiiyaur/go-mangamee/internal/db"
	"github.com/dhiiyaur/go-mangamee/internal/models"
	"github.com/dhiiyaur/go-mangamee/internal/utils"
	"google.golang.org/appengine/log"
)

func Register(ReqUser models.User) (models.User, error) {

	user, err := db.GetUsername(ReqUser.Username)

	if err != nil {
		log.Errorf(context.Background(), "Error GetUserByUsername %v", err)
		return user, err
	}

	if user.Username != "" {
		return user, errors.New("username already taken")
	}

	userTokenClaim := jwt.StandardClaims{
		Issuer: ReqUser.Username,
	}

	ReqUser.AuthToken, err = utils.EncodeToken(userTokenClaim)

	if err != nil {
		log.Errorf(context.Background(), "Error EncodeToken %v", err)
		return user, err
	}

	pw, err := utils.EncodePassword(ReqUser.Password)

	if err != nil {
		log.Errorf(context.Background(), "Error EncodeToken %v", err)
		return user, err
	}

	newUser := models.User{

		Username:     ReqUser.Username,
		HashPassword: pw,
		Email:        ReqUser.Email,
		AuthToken:    ReqUser.AuthToken,
	}

	_, err = db.CreateUser(newUser)

	if err != nil {
		log.Errorf(context.Background(), "Error CreateUser %v", err)
		return user, err
	}

	return newUser, nil
}

func Login(ReqUser models.User) (models.User, error) {

	user, err := db.GetUsername(ReqUser.Username)

	if err != nil {
		log.Errorf(context.Background(), "Error GetUserByUsername %v", err)
		return user, err
	}

	if user.Username != ReqUser.Username {
		return user, errors.New("username is wrong")
	}

	isPassword := utils.DecodePassword(ReqUser.Password, user.HashPassword)

	if !isPassword {

		fmt.Println("err", isPassword)
		return user, errors.New("Invalid Password")

	} else {

		userData := models.User{

			Username:  user.Username,
			AuthToken: user.AuthToken,
		}
		return userData, nil
	}

}

func GetHistory(username string) ([]models.UserHistory, error) {

	result, err := db.GetUserHistory(username)

	if err != nil {
		log.Errorf(context.Background(), "Empty %v", err)
		return result, err
	}

	return result, nil
}

// dbawah belum

func DeleteHistory(username string, ReqUser models.UserHistory) error {

	err := db.DeleteUserHistory(username, ReqUser.ID)

	if err != nil {
		log.Errorf(context.Background(), "Empty %v", err)
		return err
	}

	return nil
}

func CreateHistory(username string, ReqHistory models.UserHistory) error {

	db.CreateUserHistory(username, ReqHistory)

	// if err != nil {

	// 	log.Errorf(context.Background(), "Empty %v", err)
	// 	return err
	// }

	return nil

}
