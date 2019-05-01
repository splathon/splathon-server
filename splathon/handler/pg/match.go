package pg

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"
	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations/admin"
	"github.com/splathon/splathon-server/swagger/restapi/operations/match"
	"golang.org/x/sync/errgroup"
)

const (
	qualifierMaxBattleNum  = 2
	tournamentMaxBattleNum = 3
)

func getMaxBattleNum(m Match) int {
	// TODO(haya14busa): register and get theses magic numbers from database.
	n := qualifierMaxBattleNum
	if !m.QualifierId.Valid {
		n = tournamentMaxBattleNum
	}
	return n
}

func (h *Handler) GetMatch(ctx context.Context, params match.GetMatchParams) (*models.Match, error) {
	var eg errgroup.Group

	var (
		match        Match
		teams        []*Team
		participants []*Participant
		battles      []*Battle
		roundName    string
	)

	eg.Go(func() error {
		if err := h.db.Where("id = ?", params.MatchID).Find(&match).Error; err != nil {
			return err
		}

		// Fetch team.
		eg.Go(func() error {
			return h.db.Where("id = ? OR id = ?", match.TeamId, match.OpponentId).Find(&teams).Error
		})

		// Fetch participants.
		eg.Go(func() error {
			return h.db.Where("team_id = ? OR team_id = ?", match.TeamId, match.OpponentId).Order("order_in_team asc").Find(&participants).Error
		})

		// Fetch round name.
		eg.Go(func() error {
			var err error
			roundName, err = fetchRoundName(h.db, &match)
			return err
		})

		return nil
	})

	eg.Go(func() error {
		return h.db.Where("match_id = ?", params.MatchID).Order(`"order" asc`).Find(&battles).Error
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	// team_id => Team
	teamMap := make(map[int64]*Team)
	for _, t := range teams {
		teamMap[t.Id] = t
	}
	m := convertMatch(&match, teamMap)

	// Fill in members.
	m.TeamAlpha.Members = make([]*models.Member, 0)
	m.TeamBravo.Members = make([]*models.Member, 0)
	for _, p := range participants {
		var t *models.Team
		switch p.TeamId.Int64 {
		case match.TeamId:
			t = m.TeamAlpha
		case match.OpponentId:
			t = m.TeamBravo
		}
		t.Members = append(t.Members, convertParticipant2TeamMember(p))
	}

	// TODO(haya14busa): Remove later when all participants data are in database.
	if len(m.TeamAlpha.Members) == 0 {
		fillInDummyMembers(false, m.TeamAlpha)
	}
	if len(m.TeamBravo.Members) == 0 {
		fillInDummyMembers(false, m.TeamBravo)
	}

	m.RoundName = roundName

	seenBattleOrders := make(map[int]bool)
	for _, b := range battles {
		m.Battles = append(m.Battles, convertBattle(b, m))
		seenBattleOrders[int(b.Order)] = true
	}

	// Fill in not-finished battles.
	maxBattleNum := getMaxBattleNum(match)
	for order := 1; order <= maxBattleNum; order++ {
		if seenBattleOrders[order] {
			continue
		}
		m.Battles = append(m.Battles, &models.Battle{Order: swag.Int32(int32(order))})
	}
	sort.Slice(m.Battles, func(i, j int) bool {
		return *m.Battles[i].Order < *m.Battles[j].Order
	})
	return m, nil
}

func (h *Handler) GetNextMatch(ctx context.Context, params match.GetNextMatchParams) (*models.GetNextMatchResponse, error) {
	token, err := h.getTokenSession(params.XSPLATHONAPITOKEN)
	if err != nil {
		return nil, err
	}
	teamID := token.TeamID
	if params.TeamID != nil {
		teamID = *params.TeamID
	}
	if teamID == 0 {
		return nil, errors.New("team_id is not specified or you are not a member of any teams")
	}

	var eg errgroup.Group
	var (
		match        Match
		matchFound   bool
		ownTeam      Team
		opponentTeam Team
		room         Room
		roundName    string
	)

	eg.Go(func() error {
		if err := h.db.Where(`
(team_id = ? OR opponent_id = ?) AND team_points IS NULL AND opponent_points IS NULL`,
			teamID, teamID).Order("created_at asc").Find(&match).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return nil
			}
			return err
		}
		matchFound = true

		eg.Go(func() error {
			opponentID := match.OpponentId
			if opponentID == teamID {
				opponentID = match.TeamId
			}
			return h.db.Where("id = ?", opponentID).Find(&opponentTeam).Error
		})

		eg.Go(func() error {
			return h.db.Select("id, name").Where("id = ?", match.RoomId).Find(&room).Error
		})

		// Fetch round name.
		eg.Go(func() error {
			var err error
			roundName, err = fetchRoundName(h.db, &match)
			return err
		})

		return nil
	})

	eg.Go(func() error {
		return h.db.Where("id = ?", teamID).Find(&ownTeam).Error
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}
	if !matchFound {
		// Return non-error if the next match is not created yet.
		return &models.GetNextMatchResponse{}, nil
	}

	// team_id => Team
	teamMap := map[int64]*Team{
		ownTeam.Id:      &ownTeam,
		opponentTeam.Id: &opponentTeam,
	}
	resp := &models.GetNextMatchResponse{
		NextMatch: &models.NextMatch{
			MatchDetail:      convertMatch(&match, teamMap),
			OwnTeam:          convertTeam(&ownTeam),
			OpponentTeam:     convertTeam(&opponentTeam),
			MatchOrderInRoom: int32(match.Order),
			Room: &models.NextMatchRoom{
				ID:   int32(room.Id),
				Name: swag.String(room.Name),
			},
			RoundName: roundName,
		},
	}
	return resp, nil
}

func (h *Handler) UpdateMatch(ctx context.Context, params admin.UpdateMatchParams) error {
	if err := h.checkAdminAuth(params.XSPLATHONAPITOKEN); err != nil {
		return err
	}
	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return err
	}
	fmt.Printf("UpdateMatch: remove result cache: teamID=%d", *params.Match.AlphaTeamID)
	fmt.Printf("UpdateMatch: remove result cache: teamID=%d", *params.Match.BravoTeamID)
	delete(h.resultCache, resultCacheKey{eventID: eventID, teamID: *params.Match.AlphaTeamID})
	delete(h.resultCache, resultCacheKey{eventID: eventID, teamID: *params.Match.BravoTeamID})
	delete(h.resultCache, resultCacheKey{eventID: eventID, teamID: 0})
	return h.db.Model(&Match{Id: params.MatchID}).Updates(map[string]interface{}{
		"team_id":     *params.Match.AlphaTeamID,
		"opponent_id": *params.Match.BravoTeamID,
		"room_id":     *params.Match.RoomID,
		"order":       *params.Match.OrderInRoom,
	}).Error
}

func fetchRoundName(db *gorm.DB, match *Match) (string, error) {
	if match.QualifierId.Valid {
		var q Qualifier
		if err := db.Select("round").Where("id = ?", match.QualifierId).Find(&q).Error; err != nil {
			return "", err
		}
		return q.GetName(), nil
	}
	if match.TournamentId.Valid {
		var t Tournament
		if err := db.Select("name").Where("id = ?", match.TournamentId).Find(&t).Error; err != nil {
			return "", err
		}
		return t.GetName(), nil
	}
	return "", fmt.Errorf("room not found: match_id=%d", match.Id)
}
