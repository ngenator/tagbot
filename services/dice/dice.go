package main

import (
    "net/http"
    "fmt"
    "math/rand"
    "time"
    "strconv"
    "strings"
)

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func execute(w http.ResponseWriter, r *http.Request) {
    args := strings.Split(r.URL.Query()["args"][0], "d")

    num_rolls, err := strconv.Atoi(args[0])
    sides, err := strconv.Atoi(args[1])
    if err != nil {
      fmt.Fprintf(w, "Invalid dice sides.")
      return
    }

    var results []string
    for i := 0; i < num_rolls; i++ {
      results = append(results, strconv.Itoa(random(1, sides)))
    }
    fmt.Fprintf(w, strings.Join(results,",")) // send data to client side
}

func main() {
    http.HandleFunc("/execute", execute) // set router
    http.ListenAndServe(":80", nil) // set listen port
}
