package pg

import (
	"errors"
	"math/rand"
	"sort"
)

// Drawer is basially copy of drawer in original splathon app.
// https://github.com/kawakubox/splathon/blob/649d64920a9f4f4f2212066a6e1071d8ee47c3fa/app/services/drawer.rb
type Drawer struct {
	teamQueue       []*Team
	temp            []*Team
	existingMatches map[teamPair]bool
	newMatches      []teamPair
	r               *rand.Rand
}

func NewDrawer(teams []*Team, completedMatches []*Match, r *rand.Rand) *Drawer {
	d := &Drawer{
		teamQueue:       teams,
		existingMatches: make(map[teamPair]bool),
		r:               r,
	}
	// Create existingMatches table.
	for _, m := range completedMatches {
		d.existingMatches[makeTeamPair(m.TeamId, m.OpponentId)] = true
	}
	d.shuffle()
	return d
}

func (d *Drawer) shuffle() {
	// Sort team queue in ascending order.
	d.shuffleTeam(d.teamQueue)
	sort.SliceStable(d.teamQueue, func(i, j int) bool {
		if len(d.existingMatches) == 0 {
			// Use pre-defiend team "rank" for first round.
			return d.teamQueue[i].Rank < d.teamQueue[j].Rank
		}
		return d.teamQueue[i].Points < d.teamQueue[j].Points
	})
}

func (d *Drawer) NewMatches() ([]teamPair, error) {
	var ms []teamPair
	for len(d.teamQueue) > 0 {
		m, err := d.arrangeMatch()
		if err != nil {
			return nil, err
		}
		ms = append(ms, *m)
	}
	return ms, nil
}

func (d *Drawer) arrangeMatch() (*teamPair, error) {
	if len(d.teamQueue) < 2 {
		return nil, errors.New("failed to arrange new match. retry or cannot create new matches.")
	}
	alpha, bravo := d.teamQueue[len(d.teamQueue)-1], d.teamQueue[len(d.teamQueue)-2]
	d.teamQueue = d.teamQueue[:len(d.teamQueue)-2]
	if d.matched(alpha, bravo) {
		d.teamQueue = append(d.teamQueue, alpha)
		d.temp = append(d.temp, bravo)
		return d.arrangeMatch()
	}
	if len(d.temp) > 0 {
		d.teamQueue = append(d.teamQueue, d.temp...)
		d.temp = nil
	}
	return &teamPair{a: alpha.Id, b: bravo.Id}, nil
}

func (d *Drawer) matched(a, b *Team) bool {
	return d.existingMatches[makeTeamPair(a.Id, b.Id)]
}

type teamPair struct {
	a int64
	b int64
}

func makeTeamPair(a, b int64) teamPair {
	if a < b {
		return teamPair{a: a, b: b}
	}
	return teamPair{a: b, b: a}
}

func (d *Drawer) shuffleTeam(ts []*Team) {
	d.r.Shuffle(len(ts), func(i, j int) {
		ts[i], ts[j] = ts[j], ts[i]
	})
}
