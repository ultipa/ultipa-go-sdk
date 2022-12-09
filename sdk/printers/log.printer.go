package printers

import (
	"github.com/fatih/color"
	"log"
	"os"
)

func SprintError(str string) string {
	return color.RedString("[ERROR] " + str)
}

func SprintWarn(str string) string {
	return color.YellowString("[WARN] " + str)
}

func SprintInfo(str string) string {
	return color.GreenString("[INFO] ") + str
}

func PrintError(str string) {
	log.Println(SprintError(str))
}

func PrintWarn(str string) {
	log.Println(SprintWarn(str))
}

func PrintInfo(str string) {
	log.Println(SprintInfo(str))
}

func PrintErrAndExist(str string) {
	PrintError(str)
	os.Exit(1)
}
