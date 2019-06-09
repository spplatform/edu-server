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
	var user User

	err := lm.db.Model(&user).
		Where("id = ?", id).
		Select()

	return &user, err
}

func (lm *LogicManager) GetFirstPoll() (*Poll, error) {
	var (
		poll      Poll
		questions []*PollQuestion
		answers   []PollAnswer
	)

	err := lm.db.Model(&poll).
		Where("poll_id = 1").
		Select()
	if err != nil {
		return nil, err
	}

	err = lm.db.Model(&questions).
		Where("poll_id = 1").
		Select()
	if err != nil {
		return nil, err
	}

	questionIDs := make([]int, 0, len(questions))
	for _, q := range questions {
		questionIDs = append(questionIDs, q.QuestionID)
	}

	err = lm.db.Model(&answers).
		Where("poll_id = 1 AND question_id IN (?)", pg.In(questionIDs)).
		Select()

	if err != nil {
		return nil, err
	}

	// type answerKey struct {
	// 	question, answer int
	// }

	// namedAnswers := make(map[answerKey]*PollAnswer)
	// for idx, a := range answers {
	// 	namedAnswers[answerKey{a.QuestionID, a.AnswerID}] = &answers[idx]
	// }

	for _, q := range questions {
		for idx, a := range answers {
			if a.QuestionID == q.QuestionID {
				q.Answers = append(q.Answers, answers[idx])
			}
		}
	}

	poll.Questions = questions

	return &poll, err
}

func (lm *LogicManager) GetSecondPoll() (*Poll, error) {
	var (
		poll      Poll
		question  PollQuestion
		interests []Interest
	)

	err := lm.db.Model(&poll).
		Where("poll_id = 2").
		Select()
	if err != nil {
		return nil, err
	}

	err = lm.db.Model(&question).
		Where("poll_id = 2").
		Select()
	if err != nil {
		return nil, err
	}

	err = lm.db.Model(&interests).
		Select()
	if err != nil {
		return nil, err
	}

	for _, intr := range interests {
		question.Answers = append(question.Answers, PollAnswer{
			PollID:      question.PollID,
			QuestionID:  question.QuestionID,
			AnswerID:    intr.ID,
			Description: intr.Description,
		})
	}

	poll.Questions = []*PollQuestion{&question}

	return &poll, nil
}

func (lm *LogicManager) CreateRoadmap(userID int, poll RequestPoll) (int, error) {
	var interests []Interest

	rm := Roadmap{
		UserID: userID,
		Status: StatusNotStarted,
	}

	switch poll.AnswersFirst[1][0] {
	case 1:
		rm.Description = "Путь программиста"
		rm.SortedMilestones = presetRoadmaps[1]
	case 2:
		rm.Description = "Путь дизайнера"
		rm.SortedMilestones = presetRoadmaps[2]
	}

	err := lm.db.Model(&interests).
		Select()
	if err != nil {
		return 0, err
	}

	selectedInterests := make(map[int]struct{})
	for _, interestID := range poll.AnswersFirst[2] {
		selectedInterests[interestID] = struct{}{}
	}

	for _, intr := range interests {
		if _, ok := selectedInterests[intr.ID]; ok {
			if intPreset, ok := presetInterests[intr.ID]; ok {
				rm.SortedMilestones = append(rm.SortedMilestones, intPreset...)
			}
		}
	}

	err = lm.db.Insert(&rm)

	return rm.ID, err
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
