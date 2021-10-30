package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var baseDir = flag.String("database_dir", "files/database/", "SQL files directory")
var host = flag.String("h", "127.0.0.1", "Database host")
var user = flag.String("u", "root", "Database user")
var pass = flag.String("p", "", "Database password")
var dbName = flag.String("db", "toggl", "Database name")

func main() {
	os.Exit(Main())
}
func Main() int {
	flag.Parse()

	files, err := ioutil.ReadDir(*baseDir)
	if err != nil {
		log.Fatal(err)
	}

	sqlPrefix := fmt.Sprintf("mysql -h %s -u %s", *host, *user)
	if *pass != "" {
		sqlPrefix += (" -p" + *pass)
	}
	_, err = exec.Command("sh", "-c", sqlPrefix+"< "+*baseDir+"init.sql").Output()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := file.Name()
		if fileName == "init.sql" || !strings.Contains(fileName, "sql") {
			continue
		}
		filePath := *baseDir + fileName
		log.Println(" > Executing file :", fileName, " - ", sqlPrefix+" "+*dbName+" < "+filePath)
		_, err := exec.Command("sh", "-c", sqlPrefix+" "+*dbName+" < "+filePath).Output()
		if err != nil {
			log.Println(fileName, " : ", err.Error())
		}
	}
	return 0
}
