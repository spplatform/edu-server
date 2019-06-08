package server

type LogicManager struct {
}

func (lm *LogicManager) ReadUser(login, password string) (*User, error) {
	return &User{}, nil
}

func (lm *LogicManager) CreateUser(login, password string) (*User, error) {
	return &User{}, nil
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
