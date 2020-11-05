package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Result struct
type Result struct {
	Status string
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	result := Result{Status: couponLength(r.PostFormValue("coupon"))}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, string(jsonResult))
}

func couponLength(coupon string) string {
	if len(coupon) == 3 {
		return "coupon length valid"
	}
	return "coupon length invalid"
}
