//package main
//
//import (
//	"fmt"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"strconv"
//)
//
//func f()  {
//	for i := 1; i<=100; i++ {
//		url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s", strconv.Itoa(i))
//		resp, err := http.Get(url)
//		if err != nil {
//			log.Fatalln(err)
//		}
//		body, err := ioutil.ReadAll(resp.Body)
//		if err != nil {
//			log.Fatalln(err)
//		}
//
//		fmt.Println(string(body))
//	}
//}
//
//func main() {
//	go f()
//	var input string
//	fmt.Scanln(&input)
//}
text := strconv.Itoa(post.UserId) + " " + strconv.Itoa(post.Id) + " " + post.Title + " " + post.Body