package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func roll(count, max int) []string {
	rand.Seed(time.Now().Unix())
	var results []string
	for i := 0; i < count && i < 10; i++ {
		results = append(results, strconv.Itoa(random(1, max+1)))
	}
	return results
}

func execute(w http.ResponseWriter, r *http.Request) {
	args := strings.Split(r.URL.Query()["args"][0], "d")
	var results []string
	if len(args) == 2 {
		count, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(w, "Invalid number of rolls.")
			return
		}
		sides, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Fprintf(w, "Invalid number of dice sides.")
			return
		}
		results = roll(count, sides)
	} else {
		results = roll(1, 20)
	}

	fmt.Fprintf(w, strings.Join(results, ",")) // send data to client side
}

func main() {
	http.HandleFunc("/execute", execute) // set router
	http.ListenAndServe(":80", nil)      // set listen port
}
