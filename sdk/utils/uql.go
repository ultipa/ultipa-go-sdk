package utils

import (
	"regexp"
	"strings"
)

/**
	check if uql has update delete or insert operations
 */

type UqlItem struct {
	Name string
	Params string
}

func (t *UqlItem) GetListParams() []string {
	var ps []string
	str := strings.TrimSpace(t.Params)
	if  !strings.HasPrefix(str,"{") {
		for _, p := range strings.Split(str, ",") {
			ps = append(ps, strings.Trim(p, ` '"`))
		}
	} else {
		if str != "" {
			ps = append(ps, str)
		}
	}
	return ps
}

var WriteUqlCommandKeys = []string{
	"alert", "create","drop", "grant", "revoke",
	"LTE","UFE","truncate","compact",
	"insert", "update", "delete",
	"clear","stop", "pause", "resume",
	"top", "kill",
}
type EasyUqlParse struct {
	Uql string
	Commands []*UqlItem
}

func (t *EasyUqlParse)Parse(uql string)  {
	r, _ := regexp.Compile("([a-z_A-Z]*)\\(([^\\(|^\\)]*)\\)")
	findAll := r.FindAllStringSubmatch(uql, -1)
	t.Uql = uql
	t.Commands = []*UqlItem{}
	for _, find := range findAll {
		name := find[1]
		value := ""
		if len(find) >= 2 {
			value = find[2]
		}
		t.Commands = append(t.Commands, &UqlItem{
			Name:   name,
			Params: value,
		})
	}
}

func (t *EasyUqlParse) GetCommand(index int) *UqlItem {
	if len(t.Commands) > index {
		return t.Commands[index]
	}
	return nil
}
func (t *EasyUqlParse) FirstCommandName() string {
	item := t.GetCommand(0)
	if item != nil {
		return item.Name
	}
	return ""
}
func (t *EasyUqlParse) SecondCommandName() string {
	item := t.GetCommand(1)
	if item != nil {
		return item.Name
	}
	return ""
}
func (t *EasyUqlParse) HasWith() bool {
	return len(strings.Split(t.Uql, "with")) > 1
}
func (t *EasyUqlParse) HasWrite() bool {
	for _, item := range t.Commands {
		if Contains(WriteUqlCommandKeys, item.Name) {
			return true
		}
	}
	return false
}
func (t *EasyUqlParse) HasAlgo() bool {
	return t.FirstCommandName() == "algo"
}
