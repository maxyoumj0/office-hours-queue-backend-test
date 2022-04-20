package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/go-resty/resty/v2"
)

type EecsohCookie struct {
	Value string `json:"value"`
}

// Run Server to handle OAuth
func run_server() {
	session_cookie := ""

	// Fetch `session` cookie from Chrome Extension
	http.HandleFunc("/send_session_eecsoh/", func(w http.ResponseWriter, r *http.Request) {
		session_cookie = r.PostFormValue("session")
	})

	// Send `session` cookie to `login` function
	http.HandleFunc("/get_session_eecsoh/", func(w http.ResponseWriter, r *http.Request) {
		data := EecsohCookie{session_cookie}
		jData, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jData)
	})

	fmt.Printf("Starting server for login...\n")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}

func login(client *resty.Client) {
	fmt.Println("Fetching login... Use the chrome extension to login")
	session := ""
	// Wait for user to login through Chrome Extension
	for session == "" {
		resp, err := client.R().
			SetHeaders(map[string]string{
				"accept": "*/*",
			}).
			Get("http://localhost:8082/get_session_eecsoh/")
		if err != nil {
			fmt.Println(err)
			continue
		}
		jsonParsed, err := gabs.ParseJSON(resp.Body())
		if err != nil {
			fmt.Println(err)
			continue
		}
		session = jsonParsed.Path("value").Data().(string)
		time.Sleep(500 * time.Millisecond)
	}
	client.SetCookie(&http.Cookie{
		Name:     "session",
		Value:    session,
		Path:     "/",
		Domain:   "lvh.me",
		MaxAge:   0,
		HttpOnly: true,
		Secure:   true,
	})
	fmt.Println("Session Fetched!")
	return
}

/*
	Run API tests for 100% line and branch coverage on code we wrote on:
	- server/api/routes.go
	- server/api/queue.go
	- server/api/types.go
	- server/db/queue.go
*/
func api_test(client *resty.Client, course_id string) {
	// test case #1
}

func main() {
	client := resty.New()
	client.SetTimeout(4 * time.Second)

	go run_server()
	login(client)

	// Queue #1 `rfd`
	queue_1 := "https://lvh.me:8080/queues/27xCqMHnGre0qrglCpa3pL1ag5Y"

	api_test(client, queue_1)
}
