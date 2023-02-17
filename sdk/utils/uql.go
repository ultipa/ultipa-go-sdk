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

var ParseGraphCommandKeys = `(mount|unmount|truncate)\(\s*\)\.graph\(\s*["'](?P<graph>\w+)["']\s*\)`

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
	`drop\(\).user`,
	`grant\(\).user`,
	`revoke\(\).user`,
	`alter\(\).user`,
	`show\(\).policy`,
	`get\(\).policy`,
	`create\(\).policy`,
	`delete\(\).policy`,
	`drop\(\).policy`,
	`alter\(\).policy`,
	`show\(\).privilege`,
	`stats\(\)`,
	`show\(\).graph`,
	`get\(\).graph`,
	`create\(\).graph`,
	`alter\(\).graph`,
	`drop\(\).graph`,
	`kill\(\)`,
	`top\(\)`,
}

var ExtraUqlCommandKeys = map[string]struct{}{
	`top()`:       {},
	`kill`:        {},
	`show().task`: {},
	//`stop().task`:      {},
	//`clear().task`:     {},
	`stats()`:      {},
	`show().graph`: {},
	`show().algo`:  {},
	//`create().policy`:  {},
	//`drop().policy`:    {},
	`show().policy`: {},
	//`grant().user`:     {},
	//`revoke().user`:    {},
	`show().privilege`: {},
	`show().user`:      {},
	`show().self`:      {},
	//`create().user`:    {},
	//`alter().user`:     {},
	//`drop().user`:      {},
	`show().index`: {},
}

func GetUqlRegExpMatcher(fnNames []string) *regexp.Regexp {
	return regexp.MustCompile(`(?i)(` + strings.Join(fnNames, "|") + `)(\()?`)
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

//IsExtra check whether the uql is extra, if yes, then it should be sent to uqlEx via ControlClient
func (t *UqlItem) IsExtra() bool {
	matcher := regexp.MustCompile(`(?P<first>[a-z_A-Z]*)(?:\((?P<param>[^\(|^\)]*)\))?(?:[.]*(?P<second>[a-z_A-Z]*))*`)
	result := matcher.FindStringSubmatch(strings.TrimSpace(string(t.Uql)))
	if len(result) == 0 {
		return false
	}
	if _, ok := ExtraUqlCommandKeys[result[0]]; ok {
		return true
	}

	firstCmd := ""
	param := ""
	secondCmd := ""
	firstIdx := matcher.SubexpIndex("first")
	if firstIdx > -1 {
		firstCmd = result[firstIdx]
	}
	paramIdx := matcher.SubexpIndex("param")
	if paramIdx > -1 {
		param = result[paramIdx]
	}
	secondIdx := matcher.SubexpIndex("second")
	if secondIdx > -1 {
		secondCmd = result[secondIdx]
	}

	if firstCmd != "" && secondCmd != "" {
		if _, ok := ExtraUqlCommandKeys[firstCmd+"()."+secondCmd]; ok {
			return true
		}
	}

	//主要处理类似kill(1)命令中有参数的语句
	_, firstCmdIsExtra := ExtraUqlCommandKeys[firstCmd]
	if len(result) > 1 && result[0] != "" && firstCmdIsExtra && param != "" {
		return true
	}
	return false
}

//ParseGraph check whether fetch graph name from uql or not
func (t *UqlItem) ParseGraph() (bool, string) {
	matcher := regexp.MustCompile(ParseGraphCommandKeys)
	result := matcher.FindSubmatch(t.Uql)
	if result != nil {
		idx := matcher.SubexpIndex("graph")
		if idx > -1 {
			return true, string(result[idx])
		}
		return false, ""
	}
	return false, ""
}
