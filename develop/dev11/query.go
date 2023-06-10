package main

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Query struct {
	UserID      int       `json:"user_id"`
	QueryID     int       `json:"query_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func (q *Query) Decode(r io.Reader) error {
	err := json.NewDecoder(r).Decode(&q)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (q *Query) Validate() error {
	if q.UserID <= 0 {
		return fmt.Errorf("Неверный user_id: %v;", q.UserID)
	}

	if q.QueryID <= 0 {
		return fmt.Errorf("iНеверный query_id: %v;", q.QueryID)
	}

	if q.Title == "" {
		return fmt.Errorf("title cannot be empty;")
	}

	return nil
}
