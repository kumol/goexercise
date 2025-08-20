package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) Note {
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
}

func (n *Note) Display() {
	fmt.Print("Title: ", n.Title, "\nContent: ", n.Content, "\nCreated At: ", n.CreatedAt.Format(time.RFC3339), "\n")
}

func (n *Note) Save() error {
	title := strings.ReplaceAll(n.Title, " ", "_")
	title = strings.ToLower(title)
	title += ".json"
	data, err := json.Marshal(n)
	if err != nil {
		return err
	}
	return os.WriteFile(title, data, 0644)
}

// func (n *Note) GetNote() *
