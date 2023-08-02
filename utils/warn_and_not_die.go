package utils

import (
	"github.com/sqweek/dialog"
)

func WarnAndNotDie(message string) {
	dialog.Message(message).Title("Something went wrong 👻").Info()
}
