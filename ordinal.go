package main

import (
	"fmt"
	"net/http"
	"strconv"

	//"net/url"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/sbani/go-humanizer/numbers"
)

// Map represents a map with a number, an author and a finished status
type Map struct {
	Number   int    `json:"number"`
	Author   string `json:"author"`
	Finished bool   `json:"finished"`
}

// MapInfo represents a map with a number, an author, a time left, a time until, a finished status, a server and a difficulty
type MapInfo struct {
	Number     int
	Author     string
	TimeLeft   int
	TimeUntil  int
	Finished   bool
	Server     int
	Difficulty string
}

// Server represents a server with a number, a difficulty, a slice of maps, a time limit and a time left
type Server struct {
	ServerNumber     int    `json:"serverNumber"`
	ServerDifficulty string `json:"serverDifficulty"`
	Maps             []Map  `json:"maps"`
	TimeLimit        int    `json:"timeLimit"`
	TimeLeft         int    `json:"timeLeft"`
}

// Data represents the whole data with a slice of servers and a competition time left
type Data struct {
	Servers      []Server `json:"servers"`
	ComptimeLeft int      `json:"comptimeLeft"`
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		argString := r.URL.Query().Get("num")
		if argString == "" {
			fmt.Fprintf(w, "Error: No number provided")
			return
		}
		humanizeString := r.URL.Query().Get("hum")
		arg, err := strconv.Atoi(argString)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
			return
		}
		if strings.ToLower(humanizeString) == "false" {
			ordinal := getOrdinal(arg)
			fmt.Fprintf(w, ordinal)
			return
		}
		ordinal := getOrdinalHumanized(arg)
		fmt.Fprintf(w, ordinal)
	}
	fmt.Println("Listening on http://localhost:8026")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8026", nil)
}

func getOrdinalHumanized(num int) string {
	ord := numbers.Ordinal(num)
	numFormatted := humanize.Comma(int64(num))
	fullNum := fmt.Sprintf("%s%s", numFormatted, ord)
	return fullNum
}

func getOrdinal(num int) string {
	ord := numbers.Ordinal(num)
	fullNum := fmt.Sprintf("%d%s", num, ord)
	return fullNum
}
