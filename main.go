package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chayaninpia/backend-challenge/src"
)

func main() {
	if len(os.Args) > 1 {
		input := os.Args[1]
		switch input {
		case "maxpath":
			total, err := src.MaxPath("./files/hard.json")
			if err != nil {
				panic(err)
			}
			log.Printf("Max path value : %v ", total)
		default:
			res := src.LeftRightEqual(input)
			log.Printf("LeftRightEqual : %v ", res)
		}
	} else {

		http.HandleFunc("/beef/summary", src.PieFireDire)
		log.Println("watch /beef/summary")
		log.Println("listen and serve on :9000")
		if err := http.ListenAndServe(":9000", nil); err != nil {
			panic(err)
		}

	}

}
