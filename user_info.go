package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	u "github.com/Chufretalas/scramble_ghosts/utils"
)

type UserInfo struct {
	LD_URL    string `json:"ld_url"`
	LD_Pass   string `json:"ld_pass"`
	Name      string `json:"name"`
	Highscore int    `json:"highscore"`
}

func checkForUIFile() {
	_, err := os.Stat("./user_info.json")

	if err != nil {
		file, createError := os.Create("./user_info.json")
		if createError != nil {
			file.Close()
			u.WarnAndNotDie("error trying to create a user_info.json, file try creating it manually in the same folder as the game: " + createError.Error())
		}
		file.WriteString("{\n\t\"ld_url\": \"\",\n\t\"ld_pass\": \"\",\n\t\"name\": \"check out the user_info.json file and write a name\",\n\t\"highscore\": 0\n}")
		file.Close()
	}
}

func LoadUserInfo() {

	checkForUIFile()

	contents, err := os.ReadFile("./user_info.json")

	if err != nil {
		u.WarnAndNotDie("something went wrong when reading the user_info.json file: " + err.Error())
	}

	json.Unmarshal([]byte(contents), &UInfo)
	UInfo.LD_URL, _ = strings.CutSuffix(UInfo.LD_URL, "/")
}

func SaveHighscore() {
	dataToWrite, marshalErr := json.MarshalIndent(UInfo, "", "\t")
	if marshalErr != nil {
		fmt.Println("could not save the new Highscore: " + marshalErr.Error())
		return
	}
	err := os.WriteFile("./user_info.json", dataToWrite, 0644)
	if err != nil {
		fmt.Println("could not save the new Highscore: " + err.Error())
	}

}
