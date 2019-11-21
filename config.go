package main

import(
  "bufio"
  "log"
  "os"
  "strings"
  "fmt"
)

var config = map[string]string {
    "app_port":"APP_PORT",
    "db_user":"DB_USER",
    "db_pass":"DB_PASS",
    "db_host":"DB_HOST",
    "db_db":"DB_DATABASE",
    "db_port":"DB_PORT",
}

func readFile(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()	
}

func configureApp() {
  lines, err := readFile("app.ini")
  if err != nil {
    log.Fatalf("readConfig: %s", err)
  }
  
  for _, line := range lines {
    for k, v := range config {
      if (strings.HasPrefix(line, v)) {
	    config[k] = line[strings.Index(line, "=")+1:]
	    fmt.Println(k, config[k])
	  }
    }
  }
}
