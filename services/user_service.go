package services

import (
	"errors"
	"go-cassandra/config"
	"go-cassandra/models"

	"github.com/gocql/gocql"
)

func CreateUser(user *models.User) error {
	user.ID = gocql.TimeUUID()
	return config.Session.Query(
		"INSERT INTO users (id, name, email) VALUES (?, ?, ?)",
		user.ID, user.Name, user.Email,
	).Exec()
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	var u models.User

	iter := config.Session.Query("SELECT id, name, email FROM users").Iter()
	for iter.Scan(&u.ID, &u.Name, &u.Email) {
		users = append(users, u)
	}
	err := iter.Close()
	return users, err
}

func GetUserByID(id gocql.UUID) (*models.User, error) {
	var u models.User
	err := config.Session.Query(
		"SELECT id, name, email FROM users WHERE id = ?",
		id,
	).Consistency(gocql.One).Scan(&u.ID, &u.Name, &u.Email)

	if err != nil {
		return nil, errors.New("user not found")
	}
	return &u, nil
}

func UpdateUser(id gocql.UUID, updated *models.User) error {
	return config.Session.Query(
		"UPDATE users SET name = ?, email = ? WHERE id = ?",
		updated.Name, updated.Email, id,
	).Exec()
}

func DeleteUser(id gocql.UUID) error {
	return config.Session.Query("DELETE FROM users WHERE id = ?", id).Exec()
}
