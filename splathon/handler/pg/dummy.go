package pg

import (
	"fmt"
	"os"

	"github.com/go-openapi/swag"
	"github.com/splathon/splathon-server/swagger/models"
)

func fillInDummyMembers(detail bool, team *models.Team) {
	if os.Getenv("SPLATHON_USE_DUMMY_MEMBER") != "1" {
		return
	}
	if len(team.Members) > 0 {
		// If members are already filled, do nothing and return.
		return
	}
	team.Members = make([]*models.Member, 4)
	for i := 0; i <= 3; i++ {
		member := &models.Member{
			Name: swag.String(fmt.Sprintf("Dummy member %d", i+1)),
			// https://www.irasutoya.com/2017/08/blog-post_313.html
			Icon: "https://3.bp.blogspot.com/-cyF7po_IL_4/WYAx9EHo9DI/AAAAAAABFxA/zuPU_uVv-EIQm5uPgwgl4nGyDO1ZOysRwCLcBGAs/s400/character_fish_ika.png",
		}
		if detail {
			member.Detail = &models.MemberDetail{
				ShortComment:     "くまさんぶきは最強だくまぁ!!!",
				MainWeapon:       "くまさんぶき",
				RankSplatZones:   "X (9999)",
				RankTowerControl: "X (9999)",
				RankClamBlitz:    "-",
			}
		}
		team.Members[i] = member
	}
}
