package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Log struct {
	Timestamp  time.Time `json:"timestamp"`
	Duration   float64   `json:"duration"`
	HTTPMethod string    `json:"httpMethod"`
	RequestURI string    `json:"requestURI"`
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		l := Log{
			Timestamp:  start.UTC(),
			Duration:   (time.Since(start)).Seconds(),
			HTTPMethod: r.Method,
			RequestURI: r.RequestURI,
		}

		j, err := json.Marshal(&l)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(j))
	})
}

func (l *Log) MarshalJSON() ([]byte, error) {
	t := l.Timestamp.UTC().Format("2006-02-01T15:04:05.999Z")
	d := fmt.Sprintf("%.3fs", l.Duration)

	type Alias Log
	return json.Marshal(&struct {
		Timestamp string `json:"timestamp"`
		Duration  string `json:"duration"`
		*Alias
	}{
		Timestamp: t,
		Duration:  d,
		Alias:     (*Alias)(l),
	})
}
