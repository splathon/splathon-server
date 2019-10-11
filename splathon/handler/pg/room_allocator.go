package pg

import (
	"errors"
	"math/rand"
	"sort"
)

// Copy of https://github.com/kawakubox/splathon/blob/649d64920a9f4f4f2212066a6e1071d8ee47c3fa/app/services/room_allocator.rb
func allocateRooms(matches []*Match, rooms []*Room, team2roomScore map[int64]int, r *rand.Rand) error {
	// Create roomQueue ordered by room priority score (descending).
	sort.SliceStable(rooms, func(i, j int) bool {
		return rooms[i].Priority > rooms[j].Priority
	})
	roomQueue := make([]Room, len(matches))
	for i := range matches {
		roomQueue[i] = *rooms[i%len(rooms)]
	}
	sort.SliceStable(roomQueue, func(i, j int) bool {
		return roomQueue[i].Priority > roomQueue[j].Priority
	})

	// Shuffle and order matches ordered by room score (ascending).
	r.Shuffle(len(matches), func(i, j int) { matches[i], matches[j] = matches[j], matches[i] })
	sort.SliceStable(matches, func(i, j int) bool {
		return matchRoomScore(matches[i], team2roomScore) < matchRoomScore(matches[j], team2roomScore)
	})

	room2order := make(map[int64]int64)
	for _, m := range matches {
		if len(roomQueue) == 0 {
			return errors.New("room queue is empty")
		}
		roomID := roomQueue[0].Id
		roomQueue = roomQueue[1:]
		m.RoomId = roomID

		room2order[roomID] += 1
		m.Order = room2order[roomID]
	}
	return nil
}

func matchRoomScore(m *Match, team2roomScore map[int64]int) int {
	return team2roomScore[m.TeamId] + team2roomScore[m.OpponentId]
}
