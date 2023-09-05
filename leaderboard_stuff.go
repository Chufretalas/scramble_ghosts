package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func CheckLDConnection() {
	if UInfo.LD_URL == "" || UInfo.LD_Pass == "" {
		LDConnection = "no leaderboard URL"
		return
	}

	req, err := http.NewRequest("GET", UInfo.LD_URL+"/api/test", nil)
	if err != nil {
		fmt.Println(err)
		LDConnection = "could not test the connection"
		return
	}

	req.Header.Add("authorization", fmt.Sprintf("BEARER %v", UInfo.LD_Pass))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		LDConnection = "could not test the connection"
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		LDConnection = "ok"
		return
	case 401:
		LDConnection = "invalid password"
		return
	default:
		LDConnection = "not connected"
		return
	}
}

func SendScore(score int) {
	if UInfo.LD_URL == "" || UInfo.LD_Pass == "" {
		return
	}

	jsonBody := []byte(fmt.Sprintf(`{"name": %q, "score": %v, "version": %q}`, UInfo.Name, UInfo.Highscore, VERSION))
	bodyReader := bytes.NewReader(jsonBody)

	fmt.Println(string(jsonBody))
	fmt.Println(UInfo.Name)

	req, err := http.NewRequest(http.MethodPost, UInfo.LD_URL+"/api/scores", bodyReader)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("authorization", fmt.Sprintf("BEARER %v", UInfo.LD_Pass))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Println(resp.Status)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

}
