package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Post struct {
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

func f(i string)  {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s", i)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var post Post
	err = json.Unmarshal([]byte(body), &post)
	if err != nil {
		log.Fatalln(err)
	}
	filename := "store/post" + i + ".txt"
	text := strconv.Itoa(post.UserId) + " " + strconv.Itoa(post.Id) + " " + post.Title + " " + post.Body
	err = ioutil.WriteFile(filename, []byte(text), 0777)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("save on - %s \n", filename)
	amt := time.Duration(rand.Intn(250))
	time.Sleep(time.Millisecond * amt)
}

func main() {
	for i := 1; i<=5; i++ {
		go f(strconv.Itoa(i))
	}
	var input string
	fmt.Scanln(&input)
}
