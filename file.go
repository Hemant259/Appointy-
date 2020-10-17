package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Meeting struct {
	IDEN           string       `json:"iden"`
	Starting_time  string       `json:"Starting_time"`
	Ending_time    string       `json:"Ending_time"`
	Title          string       `json:"title"`
	Participant    *Participant `json:"participant"`
}

type Participant struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	RSVP      string `json:"rsvp"`
}

var meetings []Meeting

func getMeetings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(meetings)
}

func createMeet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var meeting Meeting
	_ = json.NewDecoder(r.Body).Decode(&meeting)
	meeting.IDEN = strconv.Itoa(rand.Intn(100000000))
	meetings = append(meetings, meeting)
	json.NewEncoder(w).Encode(meeting)
}

func updateMeet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	para := mux.Vars(r)
	for index, item := range meetings {
		if item.IDEN == para["iden"] {
			meetings = append(meetings[:index], meetings[index+1:]...)
			var meeting Meeting
			_ = json.NewDecoder(r.Body).Decode(&meeting)
			meeting.IDEN = para["iden"]
			meetings = append(meetings, meeting)
			json.NewEncoder(w).Encode(meeting)
			return
		}
	}
}

func main() {

	r := mux.NewRouter()

	meetings = append(meetings, Meeting{IDEN: "1", Starting_time: "16:00", Ending_time: "18:00", Title: "Meeting 1", Participant: &Participant{Firstname: "Harry", Lastname: "Potter"}})
	meetings = append(meetings, Meeting{IDEN: "2", Starting_time: "12:00", Ending_time: "14:00", Title: "Meeting 2", Participant: &Participant{Firstname: "Lily", Lastname: "Potter"}})
	meetings = append(meetings, Meeting{IDEN: "3", Starting_time: "18:00", Ending_time: "20:00", Title: "Meeting 3", Participant: &Participant{Firstname: "James", Lastname: "Potter"}})
	meetings = append(meetings, Meeting{IDEN: "4", Starting_time: "14:00", Ending_time: "16:00", Title: "Meeting 4", Participant: &Participant{Firstname: "Tom", Lastname: "Riddle"}})
	meetings = append(meetings, Meeting{IDEN: "5", Starting_time: "10:00", Ending_time: "11:30", Title: "Meeting 5", Participant: &Participant{Firstname: "Albus", Lastname: "Dumbledore"}})

	r.HandleFunc("/meetings", getMeetings).Methods("GET")
	r.HandleFunc("/meetings", createMeet).Methods("POST")
	r.HandleFunc("/meetings/{iden}", updateMeet).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}
