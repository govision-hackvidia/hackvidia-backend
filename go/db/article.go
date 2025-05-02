package db

import (
	"context"
	"fmt"
)

// table_name is article

/*
vvv Column for postgres vvv
id		title		date		source		read_duration		photo		highlighted_text		description
001
*/

func (d *DB) InsertArticle(id int, title string, date string, source string, read_duration string, photo string, highlighted_text string, description string) {
	// TODO: Insert New Article

	// Query function to insert the article
	_, err := d.Pool.Exec(context.Background(), "INSERT INTO article (id, title, date, source, read_duration, photo, highlighted_text, description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", id, title, date, source, read_duration, photo, highlighted_text, description)

	// Check whether the query execution is success or failed
	if err != nil {
		// If the query is failed
		fmt.Printf("Failed to insert this article with id of %d", id)
	} else {
		// If the query is succeed
		fmt.Println("Article added!!")
	}
}

func (d *DB) SelectManyArticle() {
	// TODO: Select many Article
}

func (d *DB) SelectArticleByID() {
	// TODO: Select article by ID
}

func (d *DB) DeleteArticle() {
	// TODO: Delete Article
}
