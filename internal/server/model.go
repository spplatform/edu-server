package server

import (
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

const (
	StatusNotStarted = iota
	StatusInProgress
	StatusFinished
	StatusFailed
)

type (
	User struct {
		ID       int
		Login    string
		Password string
		Name     string
		Badges   []*Badge   `pg:"many2many:user_badges"`
		Roadmaps []*Roadmap `sql:"-"`
	}

	Poll struct {
		PollID      int `sql:",pk"`
		Description string
		Questions   []*PollQuestion `sql:"-"`
	}

	PollQuestion struct {
		PollID      int `sql:",pk"`
		QuestionID  int `sql:",pk"`
		Description string
		Answers     []PollAnswer `sql:"-"`
	}

	PollAnswer struct {
		PollID      int `sql:",pk"`
		QuestionID  int `sql:",pk"`
		AnswerID    int `sql:",pk"`
		Description string
	}

	Roadmap struct {
		ID               int `sql:",pk"`
		Description      string
		UserID           int
		Status           int
		SortedMilestones []Milestone `sql:"-"`
	}

	Milestone struct {
		RoadmapID   int `sql:",pk"`
		ID          int `sql:",pk"`
		Main        bool
		Status      int
		Order       int
		Description string
		CourseID    int
		Link        string
		Steps       []Step `sql:"-"`
	}

	Step struct {
		RoadmapID   int `sql:",pk"`
		MilestoneID int `sql:",pk"`
		ID          int `sql:",pk"`
		Status      int
		Description string
		Link        string
	}

	Course struct {
		ID     int `sql:",pk"`
		Name   string
		Link   string
		UserID string
	}

	CourseInterest struct {
		CourseID   int `sql:",pk"`
		Course     *Course
		InterestID int `sql:",pk"`
		Interest   *Interest
	}

	Interest struct {
		ID          int
		Description string
	}

	Badge struct {
		ID       int `sql:",pk"`
		Name     string
		DateTime time.Time
		UserID   int
	}

	Certificate struct {
		ID       int `sql:",pk"`
		Name     string
		DateTime time.Time
		UserID   int
	}
)

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{
		(*User)(nil),
		(*Poll)(nil),
		(*PollQuestion)(nil),
		(*PollAnswer)(nil),
		(*Roadmap)(nil),
		(*Milestone)(nil),
		(*Step)(nil),
		(*Course)(nil),
		(*CourseInterest)(nil),
		(*Interest)(nil),
		(*Badge)(nil),
		(*Certificate)(nil),
	} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
