package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

// TODO: make this a goroutine
func SendScore(score int) {
	apiPass, exists := os.LookupEnv("API_PASSWORD")
	if !exists {
		return
	}

	jsonBody := []byte(fmt.Sprintf(`{"name": "%v", "score": %v, "version": "%v"}`, UserName, score, VERSION))
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, "https://aluraflix-chufretalas.vercel.app/api/sg", bodyReader)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("authorization", fmt.Sprintf("hackersfedem %v", apiPass))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Println(resp.Status)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

}
