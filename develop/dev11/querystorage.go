package main

import (
	"fmt"
	"sync"
	"time"
)

type Store struct {
	mu     *sync.Mutex
	events map[int][]Query
}

func (s *Store) Create(q *Query) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if events, ok := s.events[q.UserID]; ok {
		for _, event := range events {
			if event.QueryID == q.QueryID {
				return fmt.Errorf("event with such id (%v) already present for this user (%v);", q.QueryID, q.UserID)
			}
		}
	}

	s.events[q.UserID] = append(s.events[q.UserID], *q)

	return nil
}

func (s *Store) Update(q *Query) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	index := -1

	events := make([]Query, 0)
	ok := false

	if events, ok = s.events[q.UserID]; !ok {
		return fmt.Errorf("Пользователь с таким id (%v) не существует", q.UserID)
	}

	for idx, event := range events {
		if event.QueryID == q.QueryID {
			index = idx
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("У пользователя с id = (%v) нет события с id = (%v)", q.UserID, q.QueryID)
	}

	s.events[q.UserID][index] = *q

	return nil
}

func (s *Store) Delete(q *Query) (*Query, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	index := -1

	events := make([]Query, 0)
	ok := false

	if events, ok = s.events[q.UserID]; !ok {
		return nil, fmt.Errorf("Пользователь с таким id (%v) не существует", q.UserID)
	}

	for idx, event := range events {
		if event.QueryID == q.QueryID {
			index = idx
			break
		}
	}

	if index == -1 {
		return nil, fmt.Errorf("У пользователя с id = (%v) нет события с id = (%v)", q.UserID, q.QueryID)
	}

	eventsLength := len(s.events[q.UserID])
	deletedEvent := s.events[q.UserID][index]
	s.events[q.UserID][index] = s.events[q.UserID][eventsLength-1]
	s.events[q.UserID] = s.events[q.UserID][:eventsLength-1]

	return &deletedEvent, nil
}

func (s *Store) GetEventsForDay(userID int, date time.Time) ([]Query, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []Query

	events := make([]Query, 0)
	ok := false

	if events, ok = s.events[userID]; !ok {
		return nil, fmt.Errorf("Пользователь с таким id (%v) не существует", userID)
	}

	for _, event := range events {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() && event.Date.Day() == date.Day() {
			result = append(result, event)
		}
	}

	return result, nil
}

func (s *Store) GetEventsForWeek(userID int, date time.Time) ([]Query, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []Query

	events := make([]Query, 0)
	ok := false

	if events, ok = s.events[userID]; !ok {
		return nil, fmt.Errorf("Пользователь с таким id (%v) не существует", userID)
	}

	for _, event := range events {
		y1, w1 := event.Date.ISOWeek()
		y2, w2 := date.ISOWeek()
		if y1 == y2 && w1 == w2 {
			result = append(result, event)
		}
	}

	return result, nil
}

func (s *Store) GetEventsForMonth(userID int, date time.Time) ([]Query, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []Query

	events := make([]Query, 0)
	ok := false

	if events, ok = s.events[userID]; !ok {
		return nil, fmt.Errorf("Пользователь с таким id (%v) не существует", userID)
	}

	for _, event := range events {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() {
			result = append(result, event)
		}
	}

	return result, nil
}
