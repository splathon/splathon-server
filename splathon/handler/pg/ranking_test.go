package pg

import (
	"testing"

	"github.com/go-openapi/swag"
	"github.com/google/go-cmp/cmp"
	"github.com/splathon/splathon-server/swagger/models"
)

func TestBuildRanking(t *testing.T) {
	teams := []*Team{
		{Id: 1, Name: "team 1", Points: 3},
		{Id: 2, Name: "team 2", Points: 3},
		{Id: 3, Name: "team 3", Points: 1},
		{Id: 4, Name: "team 4", Points: 1},
		{Id: 5, Name: "team 5", Points: 3},
		{Id: 6, Name: "team 6", Points: 0},
	}
	matches := []*Match{
		{
			TeamId:         1,
			OpponentId:     3,
			TeamPoints:     3,
			OpponentPoints: 0,
		},
		{
			TeamId:         3,
			OpponentId:     2,
			TeamPoints:     0,
			OpponentPoints: 3,
		},
		{
			TeamId:         3,
			OpponentId:     4,
			TeamPoints:     1,
			OpponentPoints: 1,
		},
		{
			TeamId:         5,
			OpponentId:     6,
			TeamPoints:     3,
			OpponentPoints: 0,
		},
		{
			// Not done yet.
			TeamId:     1,
			OpponentId: 2,
		},
	}
	got := buildRanking(teams, matches)

	want := &models.Ranking{
		Rankings: []*models.Rank{
			{
				Rank:  swag.Int32(1),
				Point: swag.Int32(3),
				Omwp:  0.1111111111111111,
				Team: &models.Team{
					ID:   swag.Int32(1),
					Name: swag.String("team 1"),
				},
			},
			{
				Rank:  swag.Int32(1),
				Point: swag.Int32(3),
				Omwp:  0.1111111111111111,
				Team: &models.Team{
					ID:   swag.Int32(2),
					Name: swag.String("team 2"),
				},
			},
			{
				Rank:  swag.Int32(3),
				Point: swag.Int32(3),
				Omwp:  0.0,
				Team: &models.Team{
					ID:   swag.Int32(5),
					Name: swag.String("team 5"),
				},
			},
			{
				Rank:  swag.Int32(4),
				Point: swag.Int32(1),
				Omwp:  0.7777777777777778,
				Team: &models.Team{
					ID:   swag.Int32(3),
					Name: swag.String("team 3"),
				},
			},
			{
				Rank:  swag.Int32(5),
				Point: swag.Int32(1),
				Omwp:  0.1111111111111111,
				Team: &models.Team{
					ID:   swag.Int32(4),
					Name: swag.String("team 4"),
				},
			},
			{
				Rank:  swag.Int32(6),
				Point: swag.Int32(0),
				Omwp:  1,
				Team: &models.Team{
					ID:   swag.Int32(6),
					Name: swag.String("team 6"),
				},
			},
		},
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("result has diff:\n%s", diff)
	}
}
