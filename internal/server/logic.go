package server

import (
	"crypto/sha1"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-pg/pg"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {
	return
}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	fmt.Println(q.FormattedQuery())
}

type LogicManager struct {
	db *pg.DB
}

func NewLogicManager() (*LogicManager, error) {
	var db *pg.DB
	// Postgres connection parameters
	opts := pg.Options{
		Database: os.Getenv("DATABASE_NAME"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASS"),
		Addr:     os.Getenv("DATABASE_URL"),
	}

	// Give the proxy time to wake up so we don't out-race it and fail to connect
	log.Printf("waiting %d seconds to connect...", 5)
	time.Sleep(time.Duration(5) * time.Second)
	db = pg.Connect(&opts)
	db.AddQueryHook(dbLogger{})

	// ping database to check connection
	_, err := db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	if os.Getenv("DATABASE_MIGRATE") != "" {
		err := createSchema(db)
		if err != nil {
			return nil, err
		}
	}
	return &LogicManager{db}, nil
}

func (lm *LogicManager) ReadUser(login, password string) (*User, error) {
	spass := calculateHash(password)

	var user User
	err := lm.db.Model(&user).
		Where("login = ? AND password = ?", login, spass).
		Select()

	return &user, err
}

func (lm *LogicManager) CreateUser(login, password string) (*User, error) {
	user := User{
		Login:    login,
		Password: calculateHash(password),
	}

	err := lm.db.Insert(&user)

	return &user, err
}

func (lm *LogicManager) GetUser(id int) (*User, error) {
	return &User{}, nil
}

func (lm *LogicManager) GetFirstPoll() (*Poll, error) {
	return &Poll{}, nil
}

func (lm *LogicManager) GetSecondPoll() (*Poll, error) {
	return &Poll{}, nil
}

func (lm *LogicManager) ProcessPoll(poll RequestPoll) (int, error) {
	return 0, nil
}

func (lm *LogicManager) GetRoadmap(id int) (*Roadmap, error) {
	return &Roadmap{}, nil
}

func calculateHash(base string) string {
	h := sha1.New()
	h.Write([]byte(base))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x\n", bs)
}
