package db

import (
	"context"
	"fmt"

	"github.com/dhiiyaur/go-mangamee/internal/models"
)

func GetUsername(username string) (models.User, error) {
	var user models.User

	if username == "" {
		return user, nil
	}

	dbUsersRef := ConDB().FbUsersRef
	dbUserRef := dbUsersRef.Child(username)

	err := dbUserRef.Get(context.Background(), &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(user models.User) (models.User, error) {

	if user.Username == "" {
		return user, nil
	}

	dbUsersRef := ConDB().FbUsersRef
	dbUserRef := dbUsersRef.Child(user.Username)

	dbUserRef.Set(context.Background(), user)

	return user, nil
}

// MangaHistory

func GetUserHistory(username string) ([]models.UserHistory, error) {

	// fmt.Println(username)
	userHistory := []models.UserHistory{}
	var tempUserHistorys map[string]models.UserHistory

	dbUsersHistoryRef := ConDB().FbUsersHistoryRef
	dbUserHistoryRef := dbUsersHistoryRef.Child(username)
	dbUserHistoryRef.Get(context.Background(), &tempUserHistorys)

	// fmt.Println(tempUserHistorys)
	for _, tempUserHistory := range tempUserHistorys {

		// fmt.Println(tempUserHistory)
		userHistory = append(userHistory, tempUserHistory)
	}

	return userHistory, nil

}

// down

func CreateUserHistory(username string, ReqHistory models.UserHistory) error {

	dbUsersHistoryRef := ConDB().FbUsersHistoryRef
	dbUserHistoryRef := dbUsersHistoryRef.Child(username).Child(ReqHistory.ID)
	dbUserHistoryRef.Set(context.Background(), ReqHistory)

	return nil
}

func DeleteUserHistory(username string, mangaID string) error {

	dbUsersHistoryRef := ConDB().FbUsersHistoryRef
	err := dbUsersHistoryRef.Child(username).Child(mangaID).Delete(context.Background())
	fmt.Println(err)

	return nil
}

func UpdateUserHistory(username string) (interface{}, error) {

	history := []string{"heelo", "world", "nihao", "gutentag"}

	dbUsersHistoryRef := ConDB().FbUsersHistoryRef
	dbUserHistoryRef := dbUsersHistoryRef.Child(username)

	dbUserHistoryRef.Set(context.Background(), history)

	return history, nil

}
