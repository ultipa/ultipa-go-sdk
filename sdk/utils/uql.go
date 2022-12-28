package utils

import (
	"regexp"
	"strings"
)

/**
check if uql has update delete or insert operations
*/

type UqlItem struct {
	Uql []byte
}

var WriteUqlCommandKeys = []string{
	"create", "alter", "drop", "grant", "revoke",
	"LTE", "UFE", "truncate", "compact",
	"insert", "update", "delete", "upsert",
	"clear", "stop", "pause", "resume",
	"top", "kill",
	`mount\(\).graph`, `unmount\(\).graph`,
}

var GlobalUqlCommandKeys = []string{
	`show\(\).user`,
	`get\(\).user`,
	`create\(\).user`,
	`delete\(\).user`,
	`grant\(\).user`,
	`revoke\(\).user`,
	`show\(\).policy`,
	`get\(\).policy`,
	`create\(\).policy`,
	`delete\(\).policy`,
	`show\(\).privilege`,
	`stats\(\)`,
	`show\(\).graph`,
	`get\(\).graph`,
	`create\(\).graph`,
	`alter\(\).graph`,
	`drop\(\).graph`,
	`kill\(\).graph`,
	`top\(\).graph`,
}

func GetUqlRegExpMatcher(fnNames []string) *regexp.Regexp {
	return regexp.MustCompile(`(?i)(\s*|^|\n)(` + strings.Join(fnNames, "|") + `)\(`)
}

func NewUql(uql string) *UqlItem {
	return &UqlItem{
		Uql: []byte(uql),
	}
}

func (t *UqlItem) HasWith() bool {
	matcher := GetUqlRegExpMatcher([]string{"with"})
	return matcher.Match(t.Uql)
}

func (t *UqlItem) HasWrite() bool {
	matcher := GetUqlRegExpMatcher(WriteUqlCommandKeys)
	return matcher.Match(t.Uql)
}

func (t *UqlItem) HasExecTask() bool {
	matcher := GetUqlRegExpMatcher([]string{`exec task`})
	return matcher.Match(t.Uql)
}

//IsGlobal check the uql needs global graphset
func (t *UqlItem) IsGlobal() bool {
	matcher := GetUqlRegExpMatcher(GlobalUqlCommandKeys)
	return matcher.Match(t.Uql)
}
