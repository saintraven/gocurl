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

func normalizeURL (url string) string {

  if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
    url = "https://" + url 
  }

  return url 
}

func doRequest (method string, url string) (*http.Response, error) {
  
  request, err := http.NewRequest(method, url, nil)

  if err != nil {
    return nil, err  
  }

  client := &http.Client{}
  response, err := client.Do(request)

  if err != nil {
    return nil, err 
  }

  return response, err 
}

func printResponse (resp *http.Response, config Config) {

  defer resp.Body.Close()
  fmt.Println(resp.Status)

  if config.ShowHeaders {
    for key, values := range resp.Header {
      fmt.Printf("%s: %s\n", key, strings.Join(values, ", ")) 
    } 
  }

  if config.ShowBody {  io.Copy(os.Stdout, resp.Body) } 
}


func main () {
  
  config, err := parseArgs(os.Args[1:])
  if err != nil {
    fmt.Println(err)
    return
  } 

  config.URL =  normalizeURL(config.URL)
  resp, err := doRequest(config.Method, config.URL)

  if err != nil {
    fmt.Println(err)
    return 
  }

  printResponse(resp, config) 
}
