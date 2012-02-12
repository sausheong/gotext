package main

import (
	"fmt"
	"rand"
	"time"
	"os"
	"bufio"
	"http"
  "strings"
  "strconv"
)

type Book struct {
  Title string
  Content []string
}

var data = make(map[string]Book)
var dictionary = make([]string,1)

func ReadBook(name string) ([]string) {
  slice := []string{""}
  file, file_err := os.Open(name)
  if file_err != nil { fmt.Printf("Error in reading %s", name) }  
  file_buffer := bufio.NewReader(file)
  for {
    line, err := file_buffer.ReadString(byte('.'))
    slice = append(slice, line)
    if err != nil { break }  
  }
  return slice
}

func ReadDict(name string) ([]string) {
  slice := []string{""}
  file, file_err := os.Open(name)
  if file_err != nil { fmt.Printf("Error in reading %s", name) }  
  file_buffer := bufio.NewReader(file)
  for {
    line, err := file_buffer.ReadString(byte('\n'))
    slice = append(slice, line)
    if err != nil { break }  
  }
  return slice
}

func WordsHandler(w http.ResponseWriter, r *http.Request) {
  splits := strings.Split(r.URL.Path[1:], "/")
  num := splits[1]
  var n int = 4
  if num != "" {
    var err os.Error
    n, err = strconv.Atoi(num)
    if err != nil {}
  }
  sentence := make([]string,1)
  for i := 0; i < n; i++ {
    sentence = append(sentence, strings.TrimSpace(dictionary[randInt(len(dictionary))]))
  }
  fmt.Fprintln(w, strings.Join(sentence, " "))
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
  splits := strings.Split(r.URL.Path[1:], "/")
  bookName := splits[1]
  var n int = 1
  if len(splits) > 2 {
    num := splits[2]
    if num != "" {
      var err os.Error
      n, err = strconv.Atoi(num)
      if err != nil {}
    }  
  }
  book := data[bookName]
  if book.Title != "" {
    rnd := randInt(len(book.Content))  
    limit := rnd+n  
    slice := book.Content[rnd:limit]
    var para = make([]string, n)
    for i := 0; i < n; i++ {
      string := strings.Join(strings.Split(slice[i],"\r\n")," ")
      para[i] = string
    }
    paragraph := strings.Join(para, "")
    fmt.Fprintln(w, paragraph)
  }
}

func randInt(length int) int {
  rand.Seed(time.Nanoseconds())
  return rand.Intn(length+1)+1
}

func NoHandler(w http.ResponseWriter, r *http.Request) {}

func main() {
  data["barsoom"] = Book{Title: "Barsoom", Content: ReadBook("barsoom1.txt")}
  data["crusoe"] = Book{Title: "Robinson Crusoe", Content: ReadBook("crusoe.txt")}
  data["holmes"] = Book{Title: "Sherlock Holmes", Content: ReadBook("holmes1.txt")}
  dictionary = ReadDict("dictionary.txt")

  http.HandleFunc("/words/", WordsHandler)
	http.HandleFunc("/book/", BookHandler)	
	http.HandleFunc("/", NoHandler)	
	http.ListenAndServe(":5678", nil)
}