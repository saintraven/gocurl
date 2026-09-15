package main

import (
  "fmt"
  "os"
  "net/http"
  "io"
)

func main () {
  
  if len(os.Args) < 2 {
    fmt.Println("URL was not given")
    return 
  } 
 
  url := os.Args[1]
  resp, err := http.Get(url)
  if err != nil {
    fmt.Println(err)
    return
  } 

  defer resp.Body.Close()
  io.Copy(os.Stdout, resp.Body)
}
