package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func heure(w http.ResponseWriter, req *http.Request) {
	rand.Seed(time.Now().UnixNano())
	fmt.Fprintf(w, "il est %d h %d\n", time.Now().Hour(), time.Now().Minute())
}

func random(w http.ResponseWriter, req *http.Request) {
	c := rand.Intn(1000) + 1
	fmt.Fprintf(w, "le résultat d’un dé à 1000 faces est= %04d", c)
}

func randoms(w http.ResponseWriter, req *http.Request) {
	s := make([]int, 8)
	s[0] = rand.Intn(2) + 1
	s[1] = rand.Intn(4) + 1
	s[2] = rand.Intn(6) + 1
	s[3] = rand.Intn(8) + 1
	s[4] = rand.Intn(10) + 1
	s[5] = rand.Intn(12) + 1
	s[6] = rand.Intn(20) + 1
	s[7] = rand.Intn(100) + 1

	for i := 0; i < 15; i++ {
		if i <= 7 {
			if i == 7 {
				fmt.Fprintf(w, " %03d", s[i])
			}
			if i >= 4 {
				fmt.Fprintf(w, " %02d", s[i])
			} else {
				fmt.Fprintf(w, " %d", s[i])
			}
		} else {
			c := rand.Intn(7)
			if c == 7 {
				fmt.Fprintf(w, " %03d", s[c])
			}
			if c >= 4 {
				fmt.Fprintf(w, " %02d", s[c])
			} else {
				fmt.Fprintf(w, " %d", s[c])
			}

		}
	}

}

func main() {

	http.HandleFunc("/", heure)
	http.HandleFunc("/dice", random)
	http.HandleFunc("/dices", randoms)

	http.ListenAndServe(":4567", nil)
}
