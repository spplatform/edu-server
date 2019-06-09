package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// RequestHandler is a main HTTP request handler
type RequestHandler struct {
	lm *LogicManager
}

// NewRequestHandler creates a new request handler with given set of dependencies
func NewRequestHandler(lm *LogicManager) *RequestHandler {
	return &RequestHandler{lm}
}

// HandleHello handles hello requests (health check)
func (rh *RequestHandler) HandleHello(w http.ResponseWriter, r *http.Request) {
	log.Println("hello request")
	fmt.Fprint(w, "hello")
	w.WriteHeader(http.StatusOK)
}

// Login handles user login in one of several cases
// - user exists: returns a user data
// - new user: registers a user and returns a user data
func (rh *RequestHandler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		login    RequestLogin
		roadmaps []int
		newUser  bool
	)

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := rh.lm.ReadUser(login.Login, login.Password)
	if err != nil {
		user, err = rh.lm.CreateUser(login.Login, login.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		newUser = true
	} else {
		roadmaps, err = rh.lm.GetUserRoadmaps(user.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	pollFirst, err1 := rh.lm.GetFirstPoll()
	pollSecond, err2 := rh.lm.GetSecondPoll()
	if err1 != nil || err2 != nil {
		http.Error(w, "fails on poll fetch", http.StatusInternalServerError)
		return
	}

	resp := ResponceLogin{
		User: ResponseUser{
			ID:          user.ID,
			Name:        user.Name,
			RoadmapsIDs: roadmaps,
		},
		New: newUser,
		FirstPoll: &ResponsePoll{
			ID:          pollFirst.PollID,
			Description: pollFirst.Description,
			Questions:   make([]ResponseQuestion, 0, len(pollFirst.Questions)),
		},
		SecondPoll: &ResponsePoll{
			ID:          pollSecond.PollID,
			Description: pollSecond.Description,
			Questions:   make([]ResponseQuestion, 0, len(pollSecond.Questions)),
		},
	}

	for _, q := range pollFirst.Questions {
		answers := make([]ResponseAnswer, 0, len(q.Answers))
		for _, a := range q.Answers {
			answers = append(answers, ResponseAnswer{
				ID:          a.AnswerID,
				Description: a.Description,
			})
		}

		resp.FirstPoll.Questions = append(resp.FirstPoll.Questions, ResponseQuestion{
			ID:          q.QuestionID,
			Description: q.Description,
			Answers:     answers,
		})
	}

	for _, q := range pollSecond.Questions {
		answers := make([]ResponseAnswer, 0, len(q.Answers))
		for _, a := range q.Answers {
			answers = append(answers, ResponseAnswer{
				ID:          a.AnswerID,
				Description: a.Description,
			})
		}

		resp.SecondPoll.Questions = append(resp.SecondPoll.Questions, ResponseQuestion{
			ID:          q.QuestionID,
			Description: q.Description,
			Answers:     answers,
		})
	}

	payload, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (rh *RequestHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	rvars := mux.Vars(r)
	id, _ := strconv.Atoi(rvars["id"])

	user, err := rh.lm.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	roadmaps, _ := rh.lm.GetUserRoadmaps(user.ID)

	resp := ResponseUser{
		ID:          user.ID,
		Name:        user.Name,
		RoadmapsIDs: roadmaps,
	}
	payload, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (rh *RequestHandler) ProcessPolls(w http.ResponseWriter, r *http.Request) {
	rvars := mux.Vars(r)
	id, _ := strconv.Atoi(rvars["id"])

	var pollResult RequestPoll

	err := json.NewDecoder(r.Body).Decode(&pollResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	roadmapID, err := rh.lm.CreateRoadmap(id, pollResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := ResponsePollProcess{
		RoadmapID: roadmapID,
	}

	payload, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (rh *RequestHandler) GetRoadmap(w http.ResponseWriter, r *http.Request) {
	rvars := mux.Vars(r)

	id, _ := strconv.Atoi(rvars["id"])

	roadmap, err := rh.lm.GetRoadmap(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := ResponseRoadmap{
		ID:              roadmap.ID,
		Description:     roadmap.Description,
		MilestonesMain:  make([]ResponseMilestone, 0),
		MilestonesOther: make([]ResponseMilestone, 0),
	}

	for _, m := range roadmap.SortedMilestones {
		rm := ResponseMilestone{
			ID:          m.ID,
			Description: m.Description,
			CourseLink:  m.Link,
			Status:      m.Status,
			Steps:       make([]ResponseStep, 0, len(m.Steps)),
		}
		for _, s := range m.Steps {
			rm.Steps = append(rm.Steps, ResponseStep{
				ID:          s.ID,
				Description: s.Description,
				StepLink:    s.Link,
				Status:      s.Status,
			})
		}
		if m.Main {
			resp.MilestonesMain = append(resp.MilestonesMain, rm)
		} else {
			resp.MilestonesOther = append(resp.MilestonesOther, rm)
		}
	}

	payload, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

// Request structures
type (
	RequestLogin struct {
		Login    string `json:"login"`
		Password string `json:"passwors"`
	}

	RequestUserKey struct {
		UserID int `json:"user-id"`
	}

	RequestPoll struct {
		AnswersFirst  map[int][]int `json:"answers-first"`
		AnswersSecond map[int][]int `json:"answers-second"`
	}
)

// Responce structures
type (
	ResponceLogin struct {
		User       ResponseUser  `json:"user"`
		New        bool          `json:"new"`
		FirstPoll  *ResponsePoll `json:"first-poll,omitempty"`
		SecondPoll *ResponsePoll `json:"second-poll,omitempty"`
	}

	ResponseUser struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		RoadmapsIDs []int  `json:"roadmap-ids,omitempty"`
	}

	ResponseRoadmap struct {
		ID              int                 `json:"id"`
		Description     string              `json:"description"`
		Status          int                 `json:"status"`
		MilestonesMain  []ResponseMilestone `json:"milestones-main"`
		MilestonesOther []ResponseMilestone `json:"milestones-other"`
	}

	ResponseMilestone struct {
		ID          int            `json:"id"`
		Description string         `json:"description"`
		CourseLink  string         `json:"course-link"`
		Status      int            `json:"status"`
		Steps       []ResponseStep `json:"steps"`
	}

	ResponseStep struct {
		ID          int    `json:"id"`
		Description string `json:"description"`
		StepLink    string `json:"step-link,omitempty"`
		Status      int    `json:"status"`
	}

	ResponseBadge struct {
		ID          int    `json:"id"`
		Description string `json:"description"`
	}

	ResponceCertificate struct {
		ID            int       `json:"id"`
		Description   string    `json:"description"`
		IssueDateTime time.Time `json:"issue-date-time"`
	}

	ResponseAnswer struct {
		ID          int    `json:"id"`
		Description string `json:"description"`
	}

	ResponseQuestion struct {
		ID          int              `json:"id"`
		Description string           `json:"description"`
		Answers     []ResponseAnswer `json:"answers"`
	}

	ResponsePoll struct {
		ID          int                `json:"id"`
		Description string             `json:"description"`
		Questions   []ResponseQuestion `json:"questions"`
	}

	ResponsePollProcess struct {
		RoadmapID int `json:"roadmap-id"`
	}
)
