package services

import (
	"github.com/jedzeins/jlpt_api/userService/src/database"
	"github.com/jedzeins/jlpt_api/userService/src/models"
)

func RegisterSession(ownerId string) (*models.Session, *models.ApiError) {

	s := models.Session{OwnerId: ownerId}

	// should this be with a repo? probably
	session, err := s.CreateSession(database.DB)

	if err != nil {
		return nil, err
	}

	return session, nil

}
