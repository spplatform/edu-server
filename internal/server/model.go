package server

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
	}

	Poll struct {
		PollID      int
		Description string
		Questions   []PollQuestion
	}

	PollQuestion struct {
		QuestionID  int
		Description string
		Answers     []PollAnswer
	}

	PollAnswer struct {
		QuestionID  int
		AnswerID    int
		Description string
	}

	Roadmap struct {
		ID                  int
		Description         string
		Status              int
		MainMilestones      []Milestone
		SecondaryMilestones []Milestone
		SortedMilestones    []Milestone
	}

	Milestone struct {
		RoadmapID   int
		ID          int
		Main        bool
		Status      int
		Order       int
		Description string
		Link        string
		Steps       []Step
	}

	Step struct {
		RoadmapID   int
		MilestoneID int
		ID          int
		Status      int
		Description string
		Link        string
	}
)
