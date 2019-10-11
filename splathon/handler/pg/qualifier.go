package pg

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations/admin"
	"golang.org/x/sync/errgroup"
)

var random = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

func (h *Handler) CreateNewQualifier(ctx context.Context, params admin.CreateNewQualifierParams) error {
	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return err
	}

	var (
		teams              []*Team
		matches            []*Match
		nextQualifierRound int
		rooms              []*Room
	)

	var eg errgroup.Group

	eg.Go(func() error {
		qids := h.db.Table("qualifiers").Select("id").Where("event_id = ?", eventID).QueryExpr()
		query := "qualifier_id in (?) AND team_points IS NOT NULL AND opponent_points IS NOT NULL"
		return h.db.Where(query, qids).Find(&matches).Error
	})

	eg.Go(func() error {
		return h.db.Where("event_id = ?", eventID).Find(&teams).Error
	})

	eg.Go(func() error {
		if err := h.db.Table("qualifiers").Select("MAX(round)").Where("event_id = ?", eventID).Row().Scan(&nextQualifierRound); err != nil {
			return err
		}
		nextQualifierRound++
		return nil
	})

	eg.Go(func() error {
		return h.db.Where("event_id = ?", eventID).Order("id asc").Find(&rooms).Error
	})

	if err := eg.Wait(); err != nil {
		return err
	}
	rankResp := buildRanking(teams, matches, make(map[int64][]*models.Member))
	tx := h.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := h.createNextQualifierRound(teams, rankResp, matches, eventID,
		nextQualifierRound, rooms); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create next qualifier round: %v", err)
	}
	return nil
}

func (h *Handler) createNextQualifierRound(teams []*Team,
	ranking *models.Ranking, completedMatches []*Match,
	eventID int64, nextQualifierRound int, rooms []*Room) error {
	// Fill in Team.Points from ranking.
	teamMap := make(map[int64]*Team)
	for _, t := range teams {
		teamMap[t.Id] = t
	}
	roomMap := make(map[int64]*Room)
	for _, r := range rooms {
		roomMap[r.Id] = r
	}
	for _, r := range ranking.Rankings {
		t := r.Team
		if team, ok := teamMap[int64(*t.ID)]; ok {
			team.Points = *r.Point
		} else {
			return fmt.Errorf("swiss draw: team id=%d not found.", t.ID)
		}
	}

	nextQ := Qualifier{EventId: eventID, Round: int32(nextQualifierRound)}
	if err := h.db.Where(nextQ).FirstOrCreate(&nextQ).Error; err != nil {
		return err
	}

	pairs, err := NewDrawer(teams, completedMatches, random).NewMatches()
	if err != nil {
		return err
	}

	// https://github.com/t-tiger/gorm-bulk-insert doesn't quote column name.
	newMatches := make([]*Match, len(pairs))
	for i, pair := range pairs {
		newMatches[i] = &Match{
			TeamId:      pair.a,
			OpponentId:  pair.b,
			QualifierId: sql.NullInt64{Int64: nextQ.GetID(), Valid: true},
		}
	}
	team2roomScore := make(map[int64]int)
	for _, m := range completedMatches {
		score := int(roomMap[m.RoomId].Priority)
		team2roomScore[m.TeamId] += score
		team2roomScore[m.OpponentId] += score
	}
	if err := allocateRooms(newMatches, rooms, team2roomScore, random); err != nil {
		return err
	}
	// NOTE(haya14busa): use bulk insert?
	for _, m := range newMatches {
		if err := h.db.Create(m).Error; err != nil {
			return err
		}
	}
	return nil
}
