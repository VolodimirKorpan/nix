package main

import (
	"encoding/json"
	"fmt"
	"github/VolodimirKorpan/place_sql/controllers"
	"github/VolodimirKorpan/place_sql/database"
	"github/VolodimirKorpan/place_sql/models"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// GetComments ...
func GetComments(id string) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId=%s", id)
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var comments []models.Comments
	err = json.Unmarshal([]byte(body), &comments)
	if err != nil {
		log.Fatalln(err)
	}
	for _, i := range comments {
		go controllers.InsertComments(&i)
	}
	fmt.Printf("%+v", comments)
	amt := time.Duration(rand.Intn(250))
	time.Sleep(time.Millisecond * amt)

}

func main() {
	database.Connect()

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts?userId=2")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var user models.User
	err = json.Unmarshal([]byte(body), &user.Post)
	if err != nil {
		log.Fatalln(err)
	}
	for _, i := range user.Post {
		go controllers.InsertPosts(&i)
		go GetComments(strconv.Itoa(i.ID))
	}
	
	var input string
	fmt.Scanln(&input)
}
