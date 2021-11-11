package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Omikuji struct {
	Date    string `json:"date"`
	Fortune string `json:"fortune"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	o := new(Omikuji)
	o.Date = date()

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(10)

	res, err := omikuji(i, o.Date)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "もう一度アクセスしてね！\n")
	}
	o.Fortune = res
	if err := json.NewEncoder(w).Encode(o); err != nil {
		log.Println("Error:", err)
	}
}

func omikuji(i int, date string) (string, error) {
	if date == "01/01" || date == "01/02" || date == "01/03" {
		return "大吉", nil
	} else {
		switch i {
		case 0:
			return "凶", nil
		case 1, 2:
			return "小吉", nil
		case 3, 4, 5:
			return "吉", nil
		case 6, 7, 8:
			return "中吉", nil
		case 9:
			return "大吉", nil
		default:
			return "", fmt.Errorf("error omikuji")
		}
	}
}

func date() string {
	day := time.Now()
	const layout = "01/02"
	return day.Format(layout)
}
