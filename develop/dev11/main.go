package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const dateLayout = "2023-02-22"

var storage Store = Store{events: make(map[int][]Query), mu: &sync.Mutex{}}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", CreateEventHandler)
	mux.HandleFunc("/update_event", UpdateEventHandler)
	mux.HandleFunc("/delete_event", DeleteEventHandler)

	mux.HandleFunc("/events_for_day", EventsForDayHandler)
	mux.HandleFunc("/events_for_week", EventsForWeekHandler)
	mux.HandleFunc("/events_for_month", EventsForMonthHandler)

	port := "8080"

	wrappedMux := NewLogger(mux)

	log.Fatalln(http.ListenAndServe(":"+port, wrappedMux))

}

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var q Query

	if err := q.Decode(r.Body); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := q.Validate(); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := storage.Create(&q); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Событие успешно создано", []Query{q}, http.StatusCreated)

	fmt.Println(storage.events)
}

func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	var q Query

	if err := q.Decode(r.Body); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := q.Validate(); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := storage.Update(&q); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Событие было успешно обновлено", []Query{q}, http.StatusOK)

	fmt.Println(storage.events)
}

func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	var q Query

	if err := q.Decode(r.Body); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	var deletedEvent *Query
	var err error
	if deletedEvent, err = storage.Delete(&q); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Событие успешно удалено", []Query{*deletedEvent}, http.StatusOK)

	fmt.Println(storage.events)
}

func EventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := time.Parse(dateLayout, r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	var events []Query
	if events, err = storage.GetEventsForDay(userID, date); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Запрос успешно выполнен!", events, http.StatusOK)
}

func EventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := time.Parse(dateLayout, r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	var events []Query
	if events, err = storage.GetEventsForWeek(userID, date); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Request has been executed successfully!", events, http.StatusOK)
}

func EventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := time.Parse(dateLayout, r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	var events []Query
	if events, err = storage.GetEventsForMonth(userID, date); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Request has been executed successfully!", events, http.StatusOK)
}

func errorResponse(w http.ResponseWriter, e string, status int) {
	errorResponse := struct {
		Error string `json:"error"`
	}{Error: e}

	js, err := json.Marshal(errorResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func resultResponse(w http.ResponseWriter, r string, e []Query, status int) {
	resultResponse := struct {
		Result string  `json:"result"`
		Events []Query `json:"events"`
	}{Result: r, Events: e}

	js, err := json.Marshal(resultResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
