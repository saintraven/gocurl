package main

import (
  "fmt"
  "os"
  "net/http"
  "io"
  "strings"
  "errors"
)

type Config struct {
  URL string
  Method string
  ShowHeaders bool
  ShowBody bool
}

func parseArgs(args []string) (Config, error) {
  config := Config{
    Method:   "GET",
    ShowHeaders: false,
    ShowBody: true, 
  }
  
  if len(args) < 1 {
    return config, errors.New("URL was not given") 
  }
  
  config.URL= args[0]
   
  if args[0] == "-i" {
    config.ShowHeaders = true 
  
    if len(args) < 2 {
      return config, errors.New("URL was not given")
    }

    if args[1] == "" {
      return config, errors.New("URL was not given")
    } 
    
    config.URL = args[1]
      
  } else if args[0] == "-I" {
  
    config.ShowHeaders = true 
    config.ShowBody = false 
    config.Method = "HEAD"
    
    if len(args) < 2 {
      return config, errors.New("URL was not given")
    }
    config.URL = args[1] 

  } else {
    config.URL = args[0]
  } 

  return config, nil 
} 


func main () {
  
  config, err := parseArgs(os.Args[1:])
  if err != nil {
    fmt.Println(err)
    return
  } 

  if !strings.HasPrefix(config.URL, "http://") && !strings.HasPrefix(config.URL, "https://") {
    config.URL = "https://" + config.URL
  } 

  req, err := http.NewRequest(config.Method, config.URL, nil)
 
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

  if config.ShowHeaders {
    for key, values := range resp.Header {
      fmt.Printf("%s: %s\n", key, strings.Join(values, ", ")) 
    } 
  }

  if config.ShowBody {  io.Copy(os.Stdout, resp.Body) } 
}
