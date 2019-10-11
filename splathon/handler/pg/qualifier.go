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
		nextQualifierRound); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create next qualifier round: %v", err)
	}
	return nil
}

func (h *Handler) createNextQualifierRound(teams []*Team,
	ranking *models.Ranking, completedMatches []*Match,
	eventID int64, nextQualifierRound int) error {
	// Fill in Team.Points from ranking.
	teamMap := make(map[int64]*Team)
	for _, t := range teams {
		teamMap[t.Id] = t
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

	// NOTE(haya14busa): use bulk insert?
	// https://github.com/t-tiger/gorm-bulk-insert doesn't quote column name.
	for _, pair := range pairs {
		if err := h.db.Create(&Match{
			TeamId:      pair.a,
			OpponentId:  pair.b,
			QualifierId: sql.NullInt64{Int64: nextQ.GetID(), Valid: true},
			// TODO(haya14busa): assign room and order.
		}).Error; err != nil {
			return err
		}
	}
	return nil
}
