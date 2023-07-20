package main

import (
	"log"
	"os"
	"strings"
)

// TODO: remove the pass stuff from here?
func LoadUserInfo() {

	var contents []byte
	var err error

	_, err = os.Stat("./user_info.txt")

	if err != nil {
		err = nil
		file, createError := os.Create("./user_info.txt")
		if createError != nil {
			file.Close()
			log.Fatal("error trying to create a user_info.txt, file try creating it manually in the same folder as the game")
		}
		file.WriteString("open_user_info.txt_and_write_a_name")
		file.Close()
	}

	err = nil

	contents, err = os.ReadFile("./user_info.txt")

	if err != nil {
		log.Fatal("something went worng when reading the user_info.txt file")
	}

	lines := strings.Split(string(contents), "\n")

	split := strings.Split(lines[0], " ")

	for _, e := range split {
		e, _ = strings.CutSuffix(strings.Trim(e, " "), "\r")
		if e != "" {
			UserName = e
			break
		}
	}

	if len(lines) > 1 {
		ApiPass = lines[1]
		// fmt.Println(ApiPass)
	} else {
		ApiPass = ""
	}
}
