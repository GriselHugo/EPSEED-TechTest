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

func UpdateNoteForUser(UserId int, NoteId int, Title string, Content string) (*Note, error) {
	var note Note
	result := DbInstance.Where("user_id = ? AND id = ?", UserId, NoteId).First(&note)
	if result.Error != nil {
		return nil, result.Error
	}
	note.Title = Title
	note.Content = Content
	result = DbInstance.Save(&note)
	return &note, result.Error
}

func DeleteNoteForUser(UserId int, NoteId int) error {
	var note Note
	result := DbInstance.Where("user_id = ? AND id = ?", UserId, NoteId).First(&note)
	if result.Error != nil {
		return result.Error
	}
	result = DbInstance.Delete(&note)
	return result.Error
}
