package db

import (
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

func CreateNote(userID int, title, content string) (*Note, error) {
	// Insertion d'une note dans la table 'notes'
	note := Note{
		UserID: userID,
		Title: title,
		Content: content,
	}
	result := DbInstance.Create(&note)
	return &note, result.Error
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
