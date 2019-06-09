package server

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/go-pg/pg"
)

type dbLogger struct{}

// ErrWrongPassword wong user credentials
var ErrWrongPassword = errors.New("Wrong user credentials")

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {
	return
}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	// debug SQL output
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
	// log.Printf("waiting %d seconds to connect...", 5)
	// time.Sleep(time.Duration(5) * time.Second)
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

// ReadUser reads user information by user login and password
func (lm *LogicManager) ReadUser(login, password string) (*User, error) {
	var user User
	err := lm.db.Model(&user).
		Where("login = ?", login).
		Select()

	if err != nil {
		return nil, err
	}
	if user.Password != calculateHash(password) {
		return nil, ErrWrongPassword
	}

	return &user, err
}

// CreateUser creates a new user
func (lm *LogicManager) CreateUser(login, password string) (*User, error) {
	user := User{
		Login:    login,
		Password: calculateHash(password),
	}

	err := lm.db.Insert(&user)

	return &user, err
}

// GetUser returns a user by id
func (lm *LogicManager) GetUser(id int) (*User, error) {
	var user User

	err := lm.db.Model(&user).
		Where("id = ?", id).
		Select()

	return &user, err
}

// GetFirstPoll returns a first poll
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

// GetSecondPoll returns a second poll
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

// GetUserRoadmaps returns a list of roadmap IDs of user
func (lm *LogicManager) GetUserRoadmaps(userID int) ([]int, error) {
	var roadmaps []Roadmap

	err := lm.db.Model(&roadmaps).
		Column("id").
		Where("user_id = ?", userID).
		Select()

	if err != nil {
		return nil, err
	}

	res := make([]int, 0, len(roadmaps))
	for _, r := range roadmaps {
		res = append(res, r.ID)
	}

	return res, nil
}

// CreateRoadmap generates and saves a new roadmap
func (lm *LogicManager) CreateRoadmap(userID int, poll RequestPoll) (int, error) {
	var interests []Interest

	rm := Roadmap{
		UserID: userID,
		Status: StatusNotStarted,
	}

	switch skillID := poll.AnswersFirst[1][0]; skillID {
	case 1, 2:
		rm.Description = presetTitles[skillID]
		rm.SortedMilestones = presetRoadmaps[skillID]
	default:
		return 0, fmt.Errorf("Specialization %d doesn't supported yet", skillID)
	}

	interestKeys := poll.AnswersSecond[1]
	err := lm.db.Model(&interests).
		Where("id in(?)", pg.In(interestKeys)).
		Select()
	if err != nil {
		return 0, err
	}

	for _, intr := range interests {
		if intPreset, ok := presetInterests[intr.ID]; ok {
			rm.SortedMilestones = append(rm.SortedMilestones, intPreset...)
		}
	}

	err = lm.db.Insert(&rm)
	if err != nil {
		return 0, err
	}

	milestones := make([]*Milestone, 0)
	for msIdx := range rm.SortedMilestones {
		rm.SortedMilestones[msIdx].RoadmapID = rm.ID
		milestones = append(milestones, &rm.SortedMilestones[msIdx])
	}

	err = lm.db.Insert(&milestones)
	if err != nil {
		return 0, err
	}

	steps := make([]*Step, 0)
	for _, ms := range milestones {
		for stIdx := range ms.Steps {
			ms.Steps[stIdx].RoadmapID = rm.ID
			ms.Steps[stIdx].MilestoneID = ms.ID
			steps = append(steps, &ms.Steps[stIdx])
		}
	}

	err = lm.db.Insert(&steps)
	if err != nil {
		return 0, err
	}

	return rm.ID, nil
}

// GetRoadmap returns a roadmap by id
func (lm *LogicManager) GetRoadmap(id int) (*Roadmap, error) {
	var (
		roadmap    Roadmap
		milestones []Milestone
		steps      []Step
	)

	err := lm.db.Model(&roadmap).
		Where("id = ?", id).
		Select()
	if err != nil {
		return nil, err
	}

	lm.db.Model(&milestones).
		Where("roadmap_id = ?", id).
		Select()

	lm.db.Model(&steps).
		Where("roadmap_id = ?", id).
		Select()

	for msIdx, ms := range milestones {
		for _, st := range steps {
			if st.MilestoneID == ms.ID {
				milestones[msIdx].Steps = append(milestones[msIdx].Steps, st)
			}
		}
	}

	sort.SliceStable(milestones, func(i, j int) bool {
		return milestones[i].Order < milestones[j].Order
	})

	roadmap.SortedMilestones = milestones

	return &roadmap, nil
}

// CreateBadge creates new badge based on milestone data
func (lm *LogicManager) CreateBadge(userID, roadmapID, milestoneID int) (*Badge, error) {
	m := Milestone{
		RoadmapID: roadmapID,
		ID:        milestoneID,
	}

	err := lm.db.Model(&m).
		WherePK().
		Select()

	if err != nil {
		return nil, err
	}

	badge := Badge{
		UserID:   userID,
		Name:     fmt.Sprintf(`Курс '%s' пройден`, m.Description),
		DateTime: time.Now(),
	}

	err = lm.db.Insert(&badge)

	return &badge, err
}

// GetBadge returns badge by id
func (lm *LogicManager) GetBadge(id int) (*Badge, error) {
	badge := Badge{
		ID: id,
	}

	err := lm.db.Model(&badge).
		WherePK().
		Select()

	return &badge, err
}

// CreateCertificate creates new certificate based on roadmap data
func (lm *LogicManager) CreateCertificate(userID, roadmapID int) (*Certificate, error) {
	r := Roadmap{
		ID: roadmapID,
	}

	err := lm.db.Model(&r).
		WherePK().
		Select()

	if err != nil {
		return nil, err
	}

	cert := Certificate{
		UserID:   userID,
		Name:     fmt.Sprintf(`Специализация '%s' получена`, r.Description),
		DateTime: time.Now(),
	}

	err = lm.db.Insert(&cert)

	return &cert, err
}

// GetCertificate returns certificate by di
func (lm *LogicManager) GetCertificate(id int) (*Certificate, error) {
	cert := Certificate{
		ID: id,
	}

	err := lm.db.Model(&cert).
		WherePK().
		Select()

	return &cert, err
}

// GetUserAwards returns a lists of badges and certificates of user
func (lm *LogicManager) GetUserAwards(id int) ([]Badge, []Certificate) {
	var (
		badges []Badge
		certs  []Certificate
	)

	lm.db.Model(&badges).
		Where("user_id = ?", id).
		Select()

	lm.db.Model(&certs).
		Where("user_id = ?", id).
		Select()

	return badges, certs
}

// calculateHash generates a SHA1 hash for a string
func calculateHash(base string) string {
	h := sha1.New()
	h.Write([]byte(base))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
