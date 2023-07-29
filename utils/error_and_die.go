package utils

import (
	"os"

	"github.com/sqweek/dialog"
)

func ErrorAndDie(message string) {
	dialog.Message(message).Title("Something went wrong 👻").Error()
	os.Exit(1)
}
