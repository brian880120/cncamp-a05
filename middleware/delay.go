package middleware

import (
	"net/http"
	"cncamp_a04/metrics"
	"math/rand"
	"time"
	"log"
)

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max - min)
}

func Delay(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		log.Println("entering request handler")
		timer := metrics.NewTimer()
		defer timer.ObserveTotal()
	
		delay := randInt(10, 2000)
		time.Sleep(time.Millisecond * time.Duration(delay))

		next.ServeHTTP(w, r)

		log.Printf("Respond in %d ms", delay)
	})
}
