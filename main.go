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
  showHeaders := false
  showBody := true 
  method := "GET"

  if os.Args[1] == "-i" {
    showHeaders = true 
  
    if len(os.Args) < 3 {
      fmt.Println("URL was not given")
      return 
    }

    if os.Args[2] == "" {
      fmt.Println("URL was not given")
      return  
    } 
    request = os.Args[2]
      
  } else if os.Args[1] == "-I" {
  
    showHeaders = true 
    showBody = false 
    method = "HEAD"
    
    if len(os.Args) < 3 {
      fmt.Println("URL was not given")
      return 
    }

    if os.Args[2] == "" {
      fmt.Println("URL was not given")
      return  
    } 
    request = os.Args[2]
  
  } else {
    request = os.Args[1]
  } 

  if !strings.HasPrefix(request, "http://") && !strings.HasPrefix(request, "https://") {
    request = "https://" + request
  } 

  req, err := http.NewRequest(method, request, nil)
 
  if err != nil {
    fmt.Println(err)
    return
  } 

  client := &http.Client{}
  resp, err := client.Do(req)

  if err != nil {
    fmt.Println(err)
    return
  }

  defer resp.Body.Close()
  
  fmt.Println(resp.Status)

  if showHeaders {
    for key, values := range resp.Header {
      fmt.Printf("%s: %s\n", key, strings.Join(values, ", ")) 
    } 
  }

  if showBody {  io.Copy(os.Stdout, resp.Body) } 
}
