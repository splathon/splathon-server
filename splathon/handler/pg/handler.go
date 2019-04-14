package pg

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/splathon/splathon-server/gormctx"
)

// Handler is splathon API handler backed by PostgreSQL.
type Handler struct {
	db *gormctx.DB

	eventCacheMu sync.Mutex
	eventCache   map[int64]int64
}

type Option struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
	SSLMode  string
}

func (opt *Option) DBArg() string {
	arg := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		opt.Host, opt.Port, opt.User, opt.DBName, opt.Password)
	if opt.SSLMode != "" {
		arg = fmt.Sprintf("%s sslmode=%s", arg, opt.SSLMode)
	}
	return arg
}

func NewOptionFromEnv() (*Option, error) {
	opt := &Option{}
	var err error
	if opt.Host, err = nonEmptyEnv("DB_HOST"); err != nil {
		return nil, err
	}
	if opt.Port, err = nonEmptyEnv("DB_PORT"); err != nil {
		return nil, err
	}
	if opt.User, err = nonEmptyEnv("DB_USER"); err != nil {
		return nil, err
	}
	if opt.DBName, err = nonEmptyEnv("DB_DBNAME"); err != nil {
		return nil, err
	}
	if opt.Password, err = nonEmptyEnv("DB_PASSWORD"); err != nil {
		return nil, err
	}
	opt.SSLMode = os.Getenv("DB_SSLMODE")
	return opt, nil
}

func nonEmptyEnv(envname string) (string, error) {
	v := os.Getenv(envname)
	if v == "" {
		return "", fmt.Errorf("environment variable %q is empty", envname)
	}
	return v, nil
}

func NewHandler(opt *Option) (*Handler, error) {
	sqldb, err := sql.Open("postgres", opt.DBArg())
	if err != nil {
		return nil, err
	}
	db, err := gormctx.FromDB("postgres", sqldb)
	if err != nil {
		return nil, err
	}
	if os.Getenv("DB_DEBUGMODE") == "1" {
		db.LogMode(true)
	}
	return &Handler{db: db, eventCache: make(map[int64]int64)}, nil
}

func (h *Handler) queryInternalEventID(ctx context.Context, eventIDInPath int64) (int64, error) {
	h.eventCacheMu.Lock()
	defer h.eventCacheMu.Unlock()
	if eid, ok := h.eventCache[eventIDInPath]; ok {
		return eid, nil
	}

	var event Event
	q := fmt.Sprintf("Splathon#%d", int32(eventIDInPath)) + "%"
	if err := h.db.WithContext(ctx).Where("name LIKE ?", q).Find(&event).Error; err != nil {
		return 0, fmt.Errorf("event not found (event_id=%d): %v", int32(eventIDInPath), err)
	}
	h.eventCache[eventIDInPath] = event.Id
	return event.Id, nil
}
