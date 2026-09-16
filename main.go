package main

import (
  "fmt"
  "os"
  "net/http"
  "io"
  "strings"
)

func main () {
  
  if len(os.Args) < 2 {
    fmt.Println("URL was not given")
    return 
  } 

  request := os.Args[1]
 
  if !strings.HasPrefix(request, "http://") && !strings.HasPrefix(request, "https://") {
    request = "https://" + request
  } 

  resp, err := http.Get(request)
 
  if err != nil {
    fmt.Println(err)
    return
  } 

  defer resp.Body.Close()
  io.Copy(os.Stdout, resp.Body)
}
