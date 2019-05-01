package pg

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/haya14busa/secretbox"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/splathon/splathon-server/swagger/models"
)

// Handler is splathon API handler backed by PostgreSQL.
type Handler struct {
	db *gorm.DB

	eventCacheMu sync.Mutex
	eventCache   map[int64]int64

	tm      *TokenManager
	adminID string
	adminPW string

	rankingCacheMu sync.Mutex
	rankingCache   map[int64]*rankingCache // key: event ID

	teamCacheMu sync.Mutex
	teamCache   map[int64]*teamCache // key: event ID
}

type rankingCache struct {
	ranking   *models.Ranking
	timestamp time.Time
}

type teamCache struct {
	teams     *models.Teams
	timestamp time.Time
}

type Option struct {
	Host            string
	Port            string
	User            string
	DBName          string
	Password        string
	SSLMode         string
	ApplicationName string
}

func (opt *Option) DBArg() string {
	appname := "splathon-api-server"
	if opt.ApplicationName != "" {
		appname = opt.ApplicationName
	}
	arg := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s fallback_application_name=%s",
		opt.Host, opt.Port, opt.User, opt.DBName, opt.Password, appname)
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
	opt.ApplicationName = os.Getenv("DB_APPLICATION_NAME")
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
	handler := &Handler{
		eventCache:   make(map[int64]int64),
		teamCache:    make(map[int64]*teamCache),
		rankingCache: make(map[int64]*rankingCache),
	}

	// Setup DB.
	db, err := gorm.Open("postgres", opt.DBArg())
	if err != nil {
		return nil, err
	}
	if os.Getenv("DB_DEBUGMODE") == "1" {
		db.LogMode(true)
	}
	if n, ok := getIntEnv("DB_MAX_IDLE_CONNS"); ok {
		log.Printf("SET DB_MAX_IDLE_CONNS=%d", n)
		db.DB().SetMaxIdleConns(n)
	}
	if n, ok := getIntEnv("DB_MAX_OPEN_CONNS"); ok {
		log.Printf("SET DB_MAX_OPEN_CONNS=%d", n)
		db.DB().SetMaxOpenConns(n)
	}
	if n, ok := getIntEnv("DB_CONN_MAX_MAX_LIFE_TIME_SEC"); ok {
		log.Printf("SET DB_CONN_MAX_MAX_LIFE_TIME_SEC=%d", n)
		db.DB().SetConnMaxLifetime(time.Duration(n) * time.Second)
	}

	handler.db = db

	// Setup admin ID/PASS.
	if handler.adminID, err = nonEmptyEnv("SPLATHON_ADMIN_ID"); err != nil {
		return nil, err
	}
	if handler.adminPW, err = nonEmptyEnv("SPLATHON_ADMIN_PASSWORD"); err != nil {
		return nil, err
	}

	// Setup API token manager.
	sec, err := nonEmptyEnv("SPLATHON_SECRETBOX_SECRET")
	if err != nil {
		return nil, err
	}
	cipher, err := secretbox.NewFromHexKey(sec)
	if err != nil {
		return nil, fmt.Errorf("failed to create secretbox: %v", err)
	}
	var env Env
	switch os.Getenv("SPLATHON_ENV") {
	case "PROD":
		env = ENV_PROD
	case "DEV":
		env = ENV_DEV
	}
	handler.tm = NewTokenManager(cipher, env)

	return handler, nil
}

// Close closes the DB connections.
func (h *Handler) Close() error {
	return h.db.Close()
}

func (h *Handler) queryInternalEventID(eventIDInPath int64) (int64, error) {
	h.eventCacheMu.Lock()
	defer h.eventCacheMu.Unlock()
	if eid, ok := h.eventCache[eventIDInPath]; ok {
		return eid, nil
	}

	var event Event
	q := fmt.Sprintf("Splathon#%d", int32(eventIDInPath)) + "%"
	if err := h.db.Where("name LIKE ?", q).Find(&event).Error; err != nil {
		return 0, fmt.Errorf("event not found (event_id=%d): %v", int32(eventIDInPath), err)
	}
	h.eventCache[eventIDInPath] = event.Id
	return event.Id, nil
}

func getIntEnv(env string) (int, bool) {
	s := os.Getenv(env)
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}
	return n, true
}
