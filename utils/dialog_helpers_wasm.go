package utils

import (
	"fmt"
	"log"
)

func ErrorAndDie(message string) {
	log.Fatalln(message)
}

func WarnAndNotDie(message string) {
	fmt.Println(message)
}
