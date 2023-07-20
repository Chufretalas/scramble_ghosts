package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

// TODO: make this a goroutine
func SendScore(score int) {
	if ApiPass == "" {
		return
	}

	jsonBody := []byte(fmt.Sprintf(`{"name": %q, "score": %v, "version": %q}`, UserName, score, VERSION))
	bodyReader := bytes.NewReader(jsonBody)

	fmt.Println(string(jsonBody))
	fmt.Println(UserName)

	req, err := http.NewRequest(http.MethodPost, "https://aluraflix-chufretalas.vercel.app/api/sg", bodyReader)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("authorization", fmt.Sprintf("hackersfedem %v", ApiPass))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Println(resp.Status)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

}
