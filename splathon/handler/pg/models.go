package pg

import (
	"time"
)

type Qualifier struct {
	Id        int64 `gorm:"primary_key"`
	EventId   int64
	Round     int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Team struct {
	Id          int64 `gorm:"primary_key"`
	EventId     int64
	Name        string
	CompanyName string
	Points      int32
	Rank        int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Room struct {
	Id        int64 `gorm:"primary_key"`
	EventId   int64
	Name      string
	Priority  int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Match struct {
	Id             int64 `gorm:"primary_key"`
	TeamId         int64
	OpponentId     int64
	QualifierId    int64
	TeamPoints     int64
	OpponentPoints int64
	RoomId         int64
	Order          int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
