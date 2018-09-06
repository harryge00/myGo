package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"log"
	"math"
	"strconv"
	"time"
)

const TIME_LAYOUT = "2006-01-02T15:04:05Z"

type PayLoad struct {
	LoanAmount     string `json:"loanAmount"`
	NominalRate string `json:"nominalRate"`
	Duration int `json:"duration"`
	StartDate string `json:"startDate"`
}

type PaymentItem struct {
	BorrowerPaymentAmount     string `json:"borrowerPaymentAmount"`
	Date string `json:"date"`
	InitialOutstandingPrincipal string `json:"initialOutstandingPrincipal"`
	Interest string `json:"interest"`
	Principal string `json:"Principal"`
	RemainingOutstandingPrincipal string `json:"remainingOutstandingPrincipal"`
}


func calculateAnnuity(loanAmount, monthRate, duration float64) float64 {
	return monthRate * loanAmount /(1 - math.Pow(1 + monthRate, -duration))
}

func generatePlan(loanAmount, nominalRate float64, duration int, t time.Time) []PaymentItem {
	monthRate := nominalRate/1200.0
	var res []PaymentItem
	annuity := calculateAnnuity(loanAmount, monthRate, float64(duration))
	var interest, principal float64
	for i := 0; i < duration; i++ {
		interest = loanAmount * monthRate
		principal = annuity - interest
		payment := PaymentItem{
			BorrowerPaymentAmount: strconv.FormatFloat(annuity, 'f', 2, 64),
			InitialOutstandingPrincipal: strconv.FormatFloat(loanAmount, 'f', 2, 64),
			Interest: strconv.FormatFloat(interest, 'f', 2, 64),
			Principal: strconv.FormatFloat(principal, 'f', 2, 64),
			Date: t.String(),
		}
		loanAmount -= principal
		payment.RemainingOutstandingPrincipal = strconv.FormatFloat(loanAmount, 'f', 2, 64)
		t = t.AddDate(0, 1, 0)
		res = append(res, payment)
	}
	return res
}

// PostHandler converts post request body to string
func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
		return
	}
	var payLoad PayLoad
	err = json.Unmarshal(body, &payLoad)
	if err != nil {
		http.Error(w, "Error unmarshaling json",
			http.StatusInternalServerError)
		return
	}
	fmt.Println(payLoad)

	// Handle illegal input
	loanAmount, err := strconv.ParseFloat(payLoad.LoanAmount, 64)
	if err != nil {
		fmt.Errorf("Failed to parseFloat: %v", err)
		http.Error(w, "Illegal LoadAmount", http.StatusBadRequest)
		return
	}
	nominalRate, err := strconv.ParseFloat(payLoad.NominalRate, 64)
	if err != nil {
		fmt.Errorf("Failed to parseFloat: %v", err)
		http.Error(w, "Illegal NominalRate", http.StatusBadRequest)
		return
	}
	if payLoad.Duration <= 0 {
		http.Error(w, "Illegal Duration", http.StatusBadRequest)
		return
	}
	startTime, err := time.Parse(TIME_LAYOUT, payLoad.StartDate)
	if err != nil {
		http.Error(w, "Illegal startTime", http.StatusBadRequest)
		return
	}

	paymentItems := generatePlan(loanAmount, nominalRate, payLoad.Duration, startTime)
	jsonData, err := json.Marshal(paymentItems)
	if err != nil {
		http.Error(w, "Error marshaling json",
			http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {
	h := http.NewServeMux()

	h.HandleFunc("/generate-plan", PostHandler)

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "Only support POST /generate-plan")
	})

	err := http.ListenAndServe(":8080", h)
	log.Fatal(err)
}