package pg

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations/match"
)

// Handler is splathon API handler backed by PostgreSQL.
type Handler struct {
	db *gorm.DB
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
	opt.Host, err = nonEmptyEnv("DB_HOST")
	opt.Port, err = nonEmptyEnv("DB_PORT")
	opt.User, err = nonEmptyEnv("DB_USER")
	opt.DBName, err = nonEmptyEnv("DB_DBNAME")
	opt.Password, err = nonEmptyEnv("DB_PASSWORD")
	if err != nil {
		return nil, err
	}
	opt.SSLMode = os.Getenv("DB_SSLMODE")
	return opt, nil
}

func nonEmptyEnv(envname string) (string, error) {
	v := os.Getenv(envname)
	if v == "" {
		return "", fmt.Errorf("Environment variable %q is empty", envname)
	}
	return v, nil
}

func NewHandler(opt *Option) (*Handler, error) {
	db, err := gorm.Open("postgres", opt.DBArg())
	if err != nil {
		return nil, err
	}
	return &Handler{db: db}, nil
}

func (h *Handler) queryInternalEventID(eventIDInPath float64) (int64, error) {
	var event Event
	q := fmt.Sprintf("Splathon#%d", int32(eventIDInPath)) + "%"
	if err := h.db.Where("name LIKE ?", q).Find(&event).Error; err != nil {
		return 0, fmt.Errorf("event not found (event_id=%d): %v", int32(eventIDInPath), err)
	}
	return event.Id, nil
}

func (h *Handler) GetMatch(ctx context.Context, params match.GetMatchParams) (*models.Match, error) {
	return nil, errors.New("not implemented")
}
