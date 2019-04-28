package pg

import (
	"database/sql"
	"fmt"
	"time"
)

type Round interface {
	GetID() int64
	GetEventID() int64
	GetRoundNumber() int32
	GetName() string
}

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

func (q *Qualifier) GetID() int64          { return q.Id }
func (q *Qualifier) GetEventID() int64     { return q.EventId }
func (q *Qualifier) GetRoundNumber() int32 { return q.Round }
func (q *Qualifier) GetName() string       { return fmt.Sprintf("予選第%dラウンド", q.Round) }

// Tournament round.
type Tournament struct {
	Id        int64 `gorm:"primary_key"`
	EventId   int64
	Round     int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Tournament) GetID() int64          { return t.Id }
func (t *Tournament) GetEventID() int64     { return t.EventId }
func (t *Tournament) GetRoundNumber() int32 { return t.Round }
func (t *Tournament) GetName() string       { return t.Name }

type Team struct {
	Id           int64 `gorm:"primary_key"`
	EventId      int64
	Name         string
	CompanyName  string
	Rank         int32
	ShortComment string
	CreatedAt    time.Time
	UpdatedAt    time.Time

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
	QualifierId    sql.NullInt64
	TournamentId   sql.NullInt64
	TeamPoints     int64
	OpponentPoints int64
	RoomId         int64
	Order          int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Battle struct {
	Id      int64 `gorm:"primary_key"`
	MatchId int64
	// RuleId/StageId are actually non-null but use NullInt64 to update zero
	// value.
	RuleId    sql.NullInt64
	StageId   sql.NullInt64
	WinnerId  sql.NullInt64
	Order     int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Participant struct {
	Id               int64 `gorm:"primary_key"`
	EventId          int64
	SlackUserName    string
	SlackUserId      string
	Nickname         string
	FullnameKana     string
	CompanyName      string
	RawPassword      string
	IconUrl          string
	IsStaff          bool
	JoinParty        bool
	HasCompanion     bool
	Fee              int32
	TeamId           sql.NullInt64
	RankSplatZones   string
	RankTowerControl string
	RankRainmaker    string
	RankClamBlitz    string
	MainWeapon       string
	ShortComment     string
	HasSwitchDock    bool
	OrderInTeam      sql.NullInt64

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Reception struct {
	Id            int64 `gorm:"primary_key"`
	ParticipantId int64
	Memo          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
