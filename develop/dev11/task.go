package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Event struct {
	Name        string
	User_id     string
	Description string
	Date        string
}

type DeleteRequest struct {
	Name    string
	User_id string
}

type UpdateEvent struct {
	Name    string
	User_id string

	New_name    string
	Description string
	Date        string
}

func (e *Event) toString() string {
	return "Name : " + e.Name + "\tUser id : " + e.User_id + "\tDescription  : " +
		e.Description + "\tDate : " + e.Date
}

var Events []Event

func createEventHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var e Event
	if err := json.Unmarshal(content, &e); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	reg, err := regexp.MatchString(`^\d\d.\d\d.\d\d\d\d$`, e.Date)
	if !reg {
		fmt.Fprint(w, "{\"result\": \"Wrong date: event creation cancelled\"}")
		fmt.Printf("Request processed for %s %s\n", r.Method, r.URL.Path)
		return
	}
	Events = append(Events, e)

	fmt.Fprint(w, "{\"result\": \"Event succesfully created!\"}")
}

func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var ue UpdateEvent
	if err := json.Unmarshal(content, &ue); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if ue.Date != "" {
		reg, _ := regexp.MatchString(`^\d\d.\d\d.\d\d\d\d$`, ue.Date)
		if !reg {
			fmt.Fprint(w, "{\"result\": \"Wrong date: event creation cancelled\"}")
			fmt.Printf("Request processed for %s %s\n", r.Method, r.URL.Path)
			return
		}
	}

	for i, event := range Events {
		if event.Name == ue.Name && event.User_id == ue.User_id {
			var newEvent Event
			if ue.New_name != "" {
				newEvent.Name = ue.New_name
			} else {
				newEvent.Name = event.Name
			}
			if ue.Description != "" {
				newEvent.Description = ue.Description
			} else {
				newEvent.Description = event.Description
			}
			if ue.Date != "" {
				newEvent.Date = ue.Date
			} else {
				newEvent.Date = event.Date
			}
			newEvent.User_id = event.User_id
			Events[i] = newEvent
			fmt.Fprint(w, "{\"result\": \"Event succesfully updated!\"}")
			return
		}
	}

	fmt.Fprint(w, "{\"result\": \"Event not found!\"}")
}

func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	var dr DeleteRequest
	if err := json.Unmarshal(content, &dr); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	for i, event := range Events {
		if event.Name == dr.Name && event.User_id == dr.User_id {
			removeEvent(i)
			fmt.Fprint(w, "{\"result\": \"Event succesfully deleted!\"}")
			return
		}
	}
	fmt.Fprint(w, "{\"result\": \"Event not found\"}")
}

func removeEvent(i int) {
	Events = append(Events[:i], Events[i+1:]...)
}

func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	idStr := strings.TrimPrefix(r.URL.Path, "/events_for_day/")
	result := ""

	dateNow := time.Now()
	dateYear := strconv.Itoa(dateNow.Year())
	dateMonth := strconv.Itoa(int(dateNow.Month()))
	if len(dateMonth) == 1 {
		dateMonth = "0" + dateMonth
	}
	dateDay := strconv.Itoa(dateNow.Day())
	if len(dateDay) == 1 {
		dateDay = "0" + dateDay
	}
	dateNowStr := dateDay + "." + dateMonth + "." + dateYear

	for i := 0; i < len(Events); i++ {
		if Events[i].User_id == idStr && Events[i].Date == dateNowStr {
			result += Events[i].toString()
		}
	}

	w.WriteHeader((http.StatusOK))
	if len(result) != 0 {
		result = "{\"result\": \"" + result + "\"}"
		w.Write([]byte(result))
		return
	}
	w.Write([]byte("{\"result\": \"Event not found\"}"))
}

func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	idStr := strings.TrimPrefix(r.URL.Path, "/events_for_month/")
	result := ""

	dateNow := time.Now()
	dateYear := strconv.Itoa(dateNow.Year())
	dateMonth := strconv.Itoa(int(dateNow.Month()))
	if len(dateMonth) == 1 {
		dateMonth = "0" + dateMonth
	}
	dateNowStr := dateMonth + "." + dateYear

	for i := 0; i < len(Events); i++ {
		if Events[i].User_id == idStr && Events[i].Date[3:10] == dateNowStr {
			result += Events[i].toString()
		}
	}

	w.WriteHeader((http.StatusOK))
	if len(result) != 0 {
		result = "{\"result\": \"" + result + "\"}"
		w.Write([]byte(result))
		return
	}
	w.Write([]byte("{\"result\": \"Event not found\"}"))
}

func eventsForYearHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	idStr := strings.TrimPrefix(r.URL.Path, "/events_for_year/")
	result := ""

	dateNow := time.Now()
	dateYear := strconv.Itoa(dateNow.Year())

	for i := 0; i < len(Events); i++ {
		x := Events[i]
		if x.User_id == idStr && x.Date[6:10] == dateYear {
			result += x.toString()
		}
	}

	w.WriteHeader((http.StatusOK))
	if len(result) != 0 {
		result = "{\"result\": \"" + result + "\"}"
		w.Write([]byte(result))
		return
	}
	w.Write([]byte("{\"result\": \"Event not found\"}"))
}

func main() {

	loadConfig()

	port := viper.GetString("port")
	if port == "" {
		log.Fatal("Port not specified in the configuration")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", createEventHandler)
	mux.HandleFunc("/update_event", updateEventHandler)
	mux.HandleFunc("/delete_event", deleteEventHandler)
	mux.HandleFunc("/events_for_day/", eventsForDayHandler)
	mux.HandleFunc("/events_for_month/", eventsForMonthHandler)
	mux.HandleFunc("/events_for_year/", eventsForYearHandler)
	mux.HandleFunc("/hello", helloHandler)

	serverAddr := ":" + port
	fmt.Printf("Server is running on http://localhost%s\n", serverAddr)

	http.ListenAndServe(serverAddr, mux)

}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
	fmt.Printf("Request processed for %s %s\n", r.Method, r.URL.Path)

	w.Write([]byte("events"))
	w.WriteHeader((http.StatusOK))
}

/*
Requests examples:
curl -X POST -d " { \"user_id\" : \"e2\" , \"name\" : \"Whop\" ,\"description\" : \"zis is discr\", \"date\" : \"04.12.2023\"} " http://localhost:8081/create_event
curl -X GET  http://localhost:8081/events_for_day/e2
curl -X POST -d " { \"user_id\" : \"e2\" , \"name\" : \"Whop\" } " http://localhost:8081/delete_event
curl -X POST -d " { \"user_id\" : \"e2\" , \"name\" : \"Whop\", \"new_name\" : \"A0Rop\",\"description\" : \"nwe is discr\", \"date\" : \"09.12.2023\"} " http://localhost:8081/update_event
*/

func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config not found, using defaults")
	}

	viper.SetDefault("port", "8080")
}
