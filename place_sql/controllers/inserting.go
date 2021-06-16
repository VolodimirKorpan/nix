package controllers

import (
	"github/VolodimirKorpan/place_sql/database"
	"github/VolodimirKorpan/place_sql/models"
	"log"
	"math/rand"
	"time"
)

// InsertPosts ...
func InsertPosts(post *models.Post) {
	var stmt, err = database.DB.Prepare("insert into posts(id, user_id, title, body) values(?, ?, ?, ?)")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec(post.ID, post.UserID, post.Title, post.Body)

	if err != nil {
		log.Fatalln(err)
	}
	amt := time.Duration(rand.Intn(250))
	time.Sleep(time.Millisecond * amt)
}

// InsertComments ...
func InsertComments(comments *models.Comments) {
	stmt, err := database.DB.Prepare("insert into comments(id, post_id, name, email, body) values(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec(comments.ID, comments.PostID, comments.Name, comments.Email, comments.Body)

	if err != nil {
		log.Fatalln(err)
	}
	amt := time.Duration(rand.Intn(250))
	time.Sleep(time.Millisecond * amt)
}
