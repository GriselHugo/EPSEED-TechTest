package db

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func AddUser(dbConnection *sql.DB, username, password, email string) error {
	// Insertion d'un utilisateur dans la table 'users'
	query := "INSERT INTO users (username, password, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	createdAt := time.Now()
	updatedAt := createdAt
	_, err := dbConnection.Exec(query, username, password, email, createdAt, updatedAt)
	return err
}

func DeleteUser(dbConnection *sql.DB, userID int) error {
	// Suppression d'un utilisateur de la table 'users'
	query := "DELETE FROM users WHERE id = ?"
	_, err := dbConnection.Exec(query, userID)
	return err
}

func GetUser(dbConnection *sql.DB, userID int) (User, error) {
	// Récupération d'un utilisateur de la table 'users'
	query := "SELECT * FROM users WHERE id = ?"
	var user User
	err := dbConnection.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

func GetAllUsers(dbConnection *sql.DB) ([]User, error) {
	// Récupération de tous les utilisateurs de la table 'users'
	query := "SELECT * FROM users"
	rows, err := dbConnection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
