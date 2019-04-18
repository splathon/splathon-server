package pg

import (
	"database/sql"
	"time"
)

// TODO(haya14busa): Add numbering.
type Event struct {
	Id        int64 `gorm:"primary_key"`
	Name      string
	HeldOn    time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

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
	Rank        int32
	CreatedAt   time.Time
	UpdatedAt   time.Time

	// Deprecated: Calculate by matches instead.
	Points int32
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

type Battle struct {
	Id        int64 `gorm:"primary_key"`
	MatchId   int64
	RuleId    sql.NullInt64
	StageId   sql.NullInt64
	WinnerId  int64
	Order     int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
