package printers

import (
	"github.com/fatih/color"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func PrintUqlErr(errmsg string) {

	strs := strings.Split(errmsg, "\n")

	if len(strs) < 2 {
		return
	}

	s := strs[0]
	msg := strs[1]

	r := regexp.MustCompile(`\[(\d+)-(\d+)\](.*)`)
	matches := r.FindAllStringSubmatch(s, -1)

	if len(matches) < 1 {
		return
	}

	m := matches[0]

	//log.Println(m)
	if len(m) < 4 {
		return
	}

	uql := m[3]
	startIndex, _ := strconv.ParseInt(m[1], 10, 0)
	startIndex--
	start := int(startIndex)
	endIndex, _ := strconv.ParseInt(m[2], 10, 0)
	endIndex--
	end := int(endIndex)

	log.Println(uql, startIndex, endIndex)

	color.NoColor = false
	style := color.New(color.FgHiRed).Add(color.Underline).Add(color.Bold).SprintFunc()


	log.Printf("UQL Syntax Error: %s%s%s\n", color.YellowString(uql[0:start]), style(uql[start:end]), color.YellowString(uql[end:]))
	log.Println(color.YellowString(msg))

}
