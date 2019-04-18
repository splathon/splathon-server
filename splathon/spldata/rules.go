// Ref: https://splatoon2.ink/data/locale/ja.json
package spldata

var rules map[string]string = map[string]string{
	"tower_control": "ガチヤグラ",
	"splat_zones":   "ガチエリア",
	"rainmaker":     "ガチホコバトル",
	"turf_war":      "ナワバリバトル",
	"clam_blitz":    "ガチアサリ",
}

var rulesByID map[int]string = map[int]string{
	0: "turf_war",
	1: "splat_zones",
	2: "tower_control",
	3: "rainmaker",
	4: "clam_blitz",
}

func GetRuleNameByKey(key string) (string, bool) {
	n, b := rules[key]
	return n, b
}

func GetRuleByID(id int) (key string, name string, ok bool) {
	k, b := rulesByID[id]
	if !b {
		return "", "", b
	}
	n, b := GetRuleNameByKey(k)
	return k, n, b
}
