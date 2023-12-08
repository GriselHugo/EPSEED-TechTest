package db

import (
	// "database/sql"
	"time"
)

type Note struct {
	ID        int		`gorm:"primary_key"`
	UserID    int		`gorm:"not null"`
	Title     string	`gorm:"not null"`
	Content   string	`gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func CreateNote(userID int, title, content string) error {
	// Insertion d'une note dans la table 'notes'
	note := Note{
		UserID: userID,
		Title: title,
		Content: content,
	}
	result := DbInstance.Create(&note)
	return result.Error
}

func GetNoteByID(noteID int) (*Note, error) {
	// Récupération d'une note de la table 'notes'
	var note Note
	result := DbInstance.First(&note, noteID)
	return &note, result.Error
}

func GetNotesByUserID(userID int) ([]Note, error) {
	// Récupération de toutes les notes d'un utilisateur de la table 'notes'
	var notes []Note
	result := DbInstance.Where("user_id = ?", userID).Find(&notes)
	return notes, result.Error
}

func GetAllNotes() ([]Note, error) {
	// Récupération de toutes les notes de la table 'notes'
	var notes []Note
	result := DbInstance.Find(&notes)
	return notes, result.Error
}

func UpdateNote(noteID int, newTitle, newContent string) error {
	// Mise à jour d'une note dans la table 'notes'
	note, err := GetNoteByID(noteID)
	if err != nil {
		return err
	}
	note.Title = newTitle
	note.Content = newContent
	result := DbInstance.Save(&note)
	return result.Error
}

func DeleteNoteByID(noteID int) error {
	// Suppression d'une note de la table 'notes'
	result := DbInstance.Delete(&Note{}, noteID)
	return result.Error
}


// func AddNote(dbConnection *sql.DB, userID int, title, content string) error {
// 	// Insertion d'une note dans la table 'notes'
// 	query := "INSERT INTO notes (user_id, title, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
// 	createdAt := time.Now()
// 	updatedAt := createdAt
// 	_, err := dbConnection.Exec(query, userID, title, content, createdAt, updatedAt)
// 	return err
// }

// func DeleteNote(dbConnection *sql.DB, noteID int) error {
// 	// Suppression d'une note de la table 'notes'
// 	query := "DELETE FROM notes WHERE id = ?"
// 	_, err := dbConnection.Exec(query, noteID)
// 	return err
// }

// func UpdateNote(dbConnection *sql.DB, noteID int, title, content string) error {
// 	// Mise à jour d'une note dans la table 'notes'
// 	query := "UPDATE notes SET title = ?, content = ?, updated_at = ? WHERE id = ?"
// 	updatedAt := time.Now()
// 	_, err := dbConnection.Exec(query, title, content, updatedAt, noteID)
// 	return err
// }

// func GetNote(dbConnection *sql.DB, noteID int) (Note, error) {
// 	// Récupération d'une note de la table 'notes'
// 	query := "SELECT * FROM notes WHERE id = ?"
// 	var note Note
// 	err := dbConnection.QueryRow(query, noteID).Scan(&note.ID, &note.UserID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)
// 	return note, err
// }

// func GetAllNotes(dbConnection *sql.DB) ([]Note, error) {
// 	// Récupération de toutes les notes de la table 'notes'
// 	query := "SELECT * FROM notes"
// 	rows, err := dbConnection.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var notes []Note
// 	for rows.Next() {
// 		var note Note
// 		err := rows.Scan(&note.ID, &note.UserID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)
// 		if err != nil {
// 			return nil, err
// 		}
// 		notes = append(notes, note)
// 	}

// 	return notes, nil
// }

// func GetNotesByUser(dbConnection *sql.DB, userID int) ([]Note, error) {
// 	// Récupération de toutes les notes d'un utilisateur de la table 'notes'
// 	query := "SELECT * FROM notes WHERE user_id = ?"
// 	rows, err := dbConnection.Query(query, userID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var notes []Note
// 	for rows.Next() {
// 		var note Note
// 		err := rows.Scan(&note.ID, &note.UserID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)
// 		if err != nil {
// 			return nil, err
// 		}
// 		notes = append(notes, note)
// 	}

// 	return notes, nil
// }
