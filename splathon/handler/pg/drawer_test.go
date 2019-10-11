package pg

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestDrawer(t *testing.T) {
	tests := []struct {
		teams   []*Team
		matches []*Match
		want    []teamPair
	}{
		{
			teams: []*Team{
				{Id: 1, Rank: 3},
				{Id: 2, Rank: 3},
				{Id: 3, Rank: 3},
				{Id: 4, Rank: 2},
				{Id: 5, Rank: 2},
				{Id: 6, Rank: 2},
				{Id: 7, Rank: 1},
				{Id: 8, Rank: 1},
				{Id: 9, Rank: 1},
				{Id: 10, Rank: 1},
			},
			matches: nil,
			want: []teamPair{
				{a: 2, b: 1},
				{a: 3, b: 4},
				{a: 6, b: 5},
				{a: 8, b: 9},
				{a: 10, b: 7},
			},
		},
		{
			teams: []*Team{
				{Id: 1, Points: 3},
				{Id: 2, Points: 3},
				{Id: 3, Points: 3},
				{Id: 4, Points: 1},
				{Id: 5, Points: 1},
				{Id: 6, Points: 1},
				{Id: 7, Points: 1},
				{Id: 8, Points: 0},
				{Id: 9, Points: 0},
				{Id: 10, Points: 0},
			},
			matches: []*Match{
				{TeamId: 1, OpponentId: 10},
				{TeamId: 2, OpponentId: 9},
				{TeamId: 3, OpponentId: 8},
				{TeamId: 4, OpponentId: 7},
				{TeamId: 5, OpponentId: 6},
			},
			want: []teamPair{
				{a: 2, b: 1},
				{a: 3, b: 4},
				{a: 7, b: 6},
				{a: 5, b: 8},
				{a: 9, b: 10},
			},
		},
		{
			teams: []*Team{
				{Id: 1, Points: 6},
				{Id: 2, Points: 3},
				{Id: 3, Points: 4},
				{Id: 4, Points: 2},
				{Id: 5, Points: 4},
				{Id: 6, Points: 1},
				{Id: 7, Points: 4},
				{Id: 8, Points: 0},
				{Id: 9, Points: 0},
				{Id: 10, Points: 3},
			},
			matches: []*Match{
				{TeamId: 1, OpponentId: 10},
				{TeamId: 2, OpponentId: 9},
				{TeamId: 3, OpponentId: 8},
				{TeamId: 4, OpponentId: 7},
				{TeamId: 5, OpponentId: 6},
				{TeamId: 1, OpponentId: 2},
				{TeamId: 4, OpponentId: 3},
				{TeamId: 6, OpponentId: 7},
				{TeamId: 8, OpponentId: 5},
				{TeamId: 10, OpponentId: 9},
			},
			want: []teamPair{
				{a: 1, b: 7},
				{a: 3, b: 5},
				{a: 10, b: 2},
				{a: 4, b: 6},
				{a: 8, b: 9},
			},
		},
		{
			teams: []*Team{
				{Id: 1, Points: 7},
				{Id: 2, Points: 6},
				{Id: 3, Points: 5},
				{Id: 4, Points: 5},
				{Id: 5, Points: 5},
				{Id: 6, Points: 1},
				{Id: 7, Points: 5},
				{Id: 8, Points: 1},
				{Id: 9, Points: 1},
				{Id: 10, Points: 3},
			},
			matches: []*Match{
				{TeamId: 1, OpponentId: 10},
				{TeamId: 2, OpponentId: 9},
				{TeamId: 3, OpponentId: 8},
				{TeamId: 4, OpponentId: 7},
				{TeamId: 5, OpponentId: 6},
				{TeamId: 1, OpponentId: 2},
				{TeamId: 4, OpponentId: 3},
				{TeamId: 6, OpponentId: 7},
				{TeamId: 8, OpponentId: 5},
				{TeamId: 10, OpponentId: 9},
				{TeamId: 1, OpponentId: 7},
				{TeamId: 3, OpponentId: 5},
				{TeamId: 10, OpponentId: 2},
				{TeamId: 4, OpponentId: 6},
				{TeamId: 8, OpponentId: 9},
			},
			want: []teamPair{
				{a: 1, b: 4},
				{a: 2, b: 7},
				{a: 3, b: 10},
				{a: 5, b: 9},
				{a: 8, b: 6},
			},
		},
		{
			teams: []*Team{
				{Id: 1, Points: 8},
				{Id: 2, Points: 6},
				{Id: 3, Points: 6},
				{Id: 4, Points: 6},
				{Id: 5, Points: 6},
				{Id: 6, Points: 1},
				{Id: 7, Points: 8},
				{Id: 8, Points: 4},
				{Id: 9, Points: 2},
				{Id: 10, Points: 4},
			},
			matches: []*Match{
				{TeamId: 1, OpponentId: 10},
				{TeamId: 2, OpponentId: 9},
				{TeamId: 3, OpponentId: 8},
				{TeamId: 4, OpponentId: 7},
				{TeamId: 5, OpponentId: 6},
				{TeamId: 1, OpponentId: 2},
				{TeamId: 4, OpponentId: 3},
				{TeamId: 6, OpponentId: 7},
				{TeamId: 8, OpponentId: 5},
				{TeamId: 10, OpponentId: 9},
				{TeamId: 1, OpponentId: 7},
				{TeamId: 3, OpponentId: 5},
				{TeamId: 10, OpponentId: 2},
				{TeamId: 4, OpponentId: 6},
				{TeamId: 8, OpponentId: 9},
				{TeamId: 1, OpponentId: 4},
				{TeamId: 2, OpponentId: 7},
				{TeamId: 3, OpponentId: 10},
				{TeamId: 5, OpponentId: 9},
				{TeamId: 8, OpponentId: 6},
			},
			want: []teamPair{
				{a: 7, b: 3},
				{a: 2, b: 4},
				{a: 1, b: 5},
				{a: 8, b: 10},
				{a: 9, b: 6},
			},
		},
	}
	for i, tt := range tests {
		r := rand.New(rand.NewSource(14))
		got, err := NewDrawer(tt.teams, tt.matches, r).NewMatches()
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(got, tt.want, cmp.AllowUnexported(teamPair{})); diff != "" {
			t.Errorf("round %d got diff:\n%s\ngot:%v", i+1, diff, got)
		}
	}
}

// Fail seeds:
// seed=1570392539561091045 round=4 team=36
// seed=1570392614919889485 round=4 team=36
// seed=1570393054555678204 round=4 team=32
// seed=1570393090001559073 round=4 team=32
func TestDrawer_random(t *testing.T) {
	for i := 0; i < 5; i++ {
		testRandomDrawer(t, 32, 4)
	}
}

func testRandomDrawer(t *testing.T, teamNum, roundNum int) {
	var teams []*Team
	teamMap := make(map[int64]*Team)
	var matches []*Match
	for i := 0; i < 36; i++ {
		t := &Team{Id: int64(i + 1)}
		teams = append(teams, t)
		teamMap[t.Id] = t
	}
	seed := time.Now().UTC().UnixNano()
	t.Logf("seed: %v", seed)
	r := rand.New(rand.NewSource(seed))
	for round := 1; round <= roundNum; round++ {
		t.Logf("Round %d", round)
		newMatches, err := NewDrawer(teams, matches, r).NewMatches()
		if err != nil {
			t.Fatalf("failed %v: seed=%v", err, seed)
		}
		for _, pair := range newMatches {
			t.Logf("new match: %d v.s. %d", pair.a, pair.b)
			switch r.Int31n(3) {
			case 0:
				teamMap[pair.a].Points += 3
			case 1:
				teamMap[pair.b].Points += 3
			case 2:
				teamMap[pair.a].Points += 1
				teamMap[pair.b].Points += 1
			}
			matches = append(matches, &Match{TeamId: pair.a, OpponentId: pair.b})
		}
	}
}
