package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func SendScore() {
	if UInfo.LD_URL == "" || UInfo.LD_Pass == "" {
		return
	}

	jsonBody := []byte(fmt.Sprintf(`{"name": %q, "score": %v, "version": %q}`, UInfo.Name, UInfo.Highscore, VERSION))
	bodyReader := bytes.NewReader(jsonBody)

	fmt.Println(string(jsonBody))
	fmt.Println(UInfo.Name)

	req, err := http.NewRequest(http.MethodPost, UInfo.LD_URL, bodyReader)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("authorization", fmt.Sprintf("hackersfedem %v", UInfo.LD_Pass))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Println(resp.Status)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

}
