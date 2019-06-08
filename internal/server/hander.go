package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// RequestHandler is a main HTTP request handler
type RequestHandler struct {
	lm *LogicManager
}

// NewRequestHandler creates a new request handler with given set of dependencies
func NewRequestHandler() *RequestHandler {
	return &RequestHandler{}
}

// HandleHello handles hello requests
func (rh *RequestHandler) HandleHello(w http.ResponseWriter, r *http.Request) {
	log.Println("hello request")
	fmt.Fprint(w, "Hello")
	w.WriteHeader(http.StatusOK)
}

func (rh *RequestHandler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		login   RequestLogin
		newUser bool
	)

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	user, err := rh.lm.ReadUser(login.Login, login.Password)
	if err != nil {
		user, err = rh.lm.CreateUser(login.Login, login.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
			return
		}
		newUser = true
	}
	pollFirst, err1 := rh.lm.GetFirstPoll()
	pollSecond, err2 := rh.lm.GetSecondPoll()
	if err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "failes on poll fetch")
		return
	}

	resp := ResponceLogin{
		User: ResponseUser{
			ID:   user.ID,
			Name: user.Name,
		},
		New: newUser,
		FirstPoll: &ResponsePoll{
			ID:          pollFirst.PollID,
			Description: pollFirst.Description,
			Questions:   []ResponseQuestion{},
		},
		SecondPoll: &ResponsePoll{
			ID:          pollSecond.PollID,
			Description: pollSecond.Description,
			Questions:   []ResponseQuestion{},
		},
	}

	payload, err := json.Marshal(&resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (rh *RequestHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	var userID RequestUserKey

	err := json.NewDecoder(r.Body).Decode(&userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	user, err := rh.lm.GetUser(userID.UserID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err)
		return
	}

	payload, err := json.Marshal(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (rh *RequestHandler) ProcessPolls(w http.ResponseWriter, r *http.Request) {
	var pollResult RequestPoll

	err := json.NewDecoder(r.Body).Decode(&pollResult)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	// user, err := rh.lm.ProcessPoll(pollRequest)
	// if err != nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	fmt.Fprint(w, err)
	// 	return
	// }

	// payload, err := json.Marshal(&user)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	fmt.Fprint(w, err)
	// 	return
	// }

	// w.WriteHeader(http.StatusOK)
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(payload)
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
		AnswersFirst  map[int][]int `json:"answers"`
		AnswersSecond map[int][]int `json:"answers"`
	}

	// RequestAnswer struct {
	// 	QuestionID int   `json:"question-id"`
	// 	Answers    []int `json:"answer-ids"`
	// }
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
		ID       int               `json:"id"`
		Name     string            `json:"name"`
		Roadmaps []ResponseRoadmap `json:"roadmaps,omitempty"`
	}

	ResponseRoadmap struct {
		ID              int                 `json:"id"`
		Description     string              `json:"description"`
		MilestonesMain  []ResponseMilestone `json:"milestones"`
		MilestonesOther []ResponseMilestone `json:"milestones"`
	}

	ResponseMilestone struct {
		ID          int            `json:"id"`
		Description string         `json:"description"`
		CourseLink  string         `json:"course-link"`
		Steps       []ResponseStep `json:"steps"`
	}

	ResponseStep struct {
		ID          int    `json:"id"`
		Description string `json:"description"`
		StepLink    string `json:"step-link,omitempty"`
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
)
