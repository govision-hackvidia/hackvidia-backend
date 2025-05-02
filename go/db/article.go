package db

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/jackc/pgx/v5"
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

	// select random from available article
	// check the quantity of available articles
	var nCount int64
	err := d.Pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM article").Scan(&nCount)
	if err != nil {
		fmt.Println("Query row failed!")
	}

	if nCount > 0 {
		// there is article
		// randomize value
		randomSelectId := rand.Intn(int(nCount)-int(0)+1) + int(nCount)

		// temporary variable for function return
		var title, date, source, read_duration, photo, highlighted_text, description string

		// select an article with randomSelectId result
		err = d.Pool.QueryRow(context.Background(), "SELECT title, date, source, read_duration, photo, highlighted_text, description WHERE id = $1 LIMIT 1", randomSelectId).Scan(&title, &date, &source, &read_duration, &photo, &highlighted_text, &description)
		if err != pgx.ErrNoRows {
			fmt.Println("There is no row from select")
		}
	} else {
		fmt.Println("There is no article available")
	}

}

func (d *DB) SelectArticleByID(article_id int) (string, string, string, string, string, string, string) {
	// TODO: Select article by ID

	// temporary variable for function return
	var title, date, source, read_duration, photo, highlighted_text, description string

	err := d.Pool.QueryRow(context.Background(), "SELECT title, date, source, read_duration, photo, highlighted_text, description WHERE id = $1 LIMIT 1", article_id).Scan(&title, &date, &source, &read_duration, &photo, &highlighted_text, &description)
	if err != nil {
		if err == pgx.ErrNoRows {
			// There is not such article on that id
			fmt.Printf("There is no article on id: %d", article_id)
			return "", "", "", "", "", "", ""
		} else {
			// Failed to run query function
			fmt.Println("Failed to execute :(")
			return "", "", "", "", "", "", ""
		}
	} else {
		fmt.Println("FOUND!! Will sent the data")
		return title, date, source, read_duration, photo, highlighted_text, description
	}

}

func (d *DB) DeleteArticle() {
	// TODO: Delete Article
}
