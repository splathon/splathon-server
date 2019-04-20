// Ref: https://splatoon2.ink/data/locale/ja.json
package spldata

var (
	turfWar      = Rule{ID: 0, Key: "turf_war", Name: "ナワバリバトル"}
	splatZones   = Rule{ID: 1, Key: "splat_zones", Name: "ガチエリア"}
	towerControl = Rule{ID: 2, Key: "tower_control", Name: "ガチヤグラ"}
	rainmaker    = Rule{ID: 3, Key: "rainmaker", Name: "ガチホコバトル"}
	clamBlitz    = Rule{ID: 4, Key: "clam_blitz", Name: "ガチアサリ"}
)

var (
	rulesByKey = make(map[string]*Rule)
	rulesByID  = make(map[int]*Rule)
)

func init() {
	for _, r := range ListRules() {
		rulesByKey[r.Key] = &r
		rulesByID[r.ID] = &r
	}
}

type Rule struct {
	ID   int
	Key  string
	Name string
}

func ListRules() []Rule {
	return []Rule{turfWar, splatZones, towerControl, rainmaker, clamBlitz}
}

func GetRuleByKey(key string) (*Rule, bool) {
	r, ok := rulesByKey[key]
	return r, ok
}

func GetRuleByID(id int) (*Rule, bool) {
	r, ok := rulesByID[id]
	return r, ok
}
