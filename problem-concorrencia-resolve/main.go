package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

// Modo 2 de resolver usando soma atomic
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//number++
		atomic.AddUint64(&number, 1)
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Visitas: %d", number)))
	})
	http.ListenAndServe(":8080", nil)
}

/*
// Modo 1 de resolver utilizando mutex
func main() {
	m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		number++
		m.Unlock()
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Visitas: %d", number)))
	})
	http.ListenAndServe(":8080", nil)
}
*/
