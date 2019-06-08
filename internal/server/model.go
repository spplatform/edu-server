package server

import "time"

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
		Badges   []*Badge `pg:"many2many:user_badges"`
		Roadmaps []*Roadmap
	}

	Poll struct {
		PollID      int `sql:",pk"`
		Description string
		Questions   []*PollQuestion
	}

	PollQuestion struct {
		PollID      int `sql:",pk"`
		QuestionID  int `sql:",pk"`
		Description string
		Answers     []*PollAnswer
	}

	PollAnswer struct {
		QuestionID  int `sql:",pk"`
		AnswerID    int `sql:",pk"`
		Description string
	}

	Roadmap struct {
		ID               int `sql:",pk"`
		Description      string
		Status           int
		SortedMilestones []*Milestone
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
		Steps       []*Step
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
		ID   int `sql:",pk"`
		Name string
	}

	UserBadge struct {
		UserID  int `sql:",pk"`
		User    *User
		BadgeID int `sql:",pk"`
		Badge   *Badge
	}

	Certificate struct {
		ID       int `sql:",pk"`
		Name     string
		DateTime time.Time
		UserID   int
	}
)
