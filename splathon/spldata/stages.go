// Ref: https://splatoon2.ink/data/locale/ja.json
package spldata

import "sort"

var stagesByID = map[int]string{
	0:  "バッテラストリート",
	1:  "フジツボスポーツクラブ",
	2:  "ガンガゼ野外音楽堂",
	3:  "チョウザメ造船",
	4:  "海女美術大学",
	5:  "コンブトラック",
	6:  "マンタマリア号",
	7:  "ホッケふ頭",
	8:  "タチウオパーキング",
	9:  "エンガワ河川敷",
	10: "モズク農園",
	11: "Ｂバスパーク",
	12: "デボン海洋博物館",
	13: "ザトウマーケット",
	14: "ハコフグ倉庫",
	15: "アロワナモール",
	16: "モンガラキャンプ場",
	17: "ショッツル鉱山",
	18: "アジフライスタジアム",
	19: "ホテルニューオートロ",
	20: "スメーシーワールド",
	21: "アンチョビットゲームズ",
	22: "ムツゴ楼",
	// 9999: "ミステリーゾーン",
}

var stages = make([]Stage, 0)

func init() {
	for id, name := range stagesByID {
		stages = append(stages, Stage{ID: id, Name: name})
	}
	sort.Slice(stages, func(i, j int) bool {
		return stages[i].ID < stages[j].ID
	})
}

func GetStageByID(stageID int) (string, bool) {
	n, b := stagesByID[stageID]
	return n, b
}

type Stage struct {
	ID   int
	Name string
}

func ListStages() []Stage {
	return stages
}
