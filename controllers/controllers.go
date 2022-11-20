package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/SilAronica/labora-proyectoFinal.git/services"
)

func CreateEventSearch(w http.ResponseWriter, r *http.Request) {
	url := "http://api.checks.truora.com/v1/checks"

	dniFromBody := "2838428283737"

	stringUrl := fmt.Sprintf("national_id=%s&country=CO&type=person&user_authorized=true&force_creation=true", dniFromBody)

	method := "POST"

	payload := strings.NewReader(stringUrl)

	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Truora-API-Key", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiIiwiYWRkaXRpb25hbF9kYXRhIjoie30iLCJjbGllbnRfaWQiOiJUQ0llZWU3NWUzMjBiODhhYTdiYTVkYTkyYjFkYWU5Y2IxMCIsImV4cCI6MzI0NTI4MDMyNiwiZ3JhbnQiOiIiLCJpYXQiOjE2Njg0ODAzMjYsImlzcyI6Imh0dHBzOi8vY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb20vdXMtZWFzdC0xX2lYcDhPTVBpRyIsImp0aSI6ImZiYjEyNWQwLWM4NWItNDc3OS05ZmQ4LTE3NjRlMTZmODUyYiIsImtleV9uYW1lIjoiYXBpa2V5Iiwia2V5X3R5cGUiOiJiYWNrZW5kIiwidXNlcm5hbWUiOiJnbWFpbHN2Z29uemFsZXpyYW1vcy1hcGlrZXkifQ.199M11RxVNkeCNRtejcifZF5_u23ERDkudm0bwSPFe4")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return

	}
	services.InsertLogInDb(dniFromBody)
	fmt.Println(string(body))
}

func GetEvent(w http.ResponseWriter, r *http.Request) {

	url := "https://api.checks.truora.com/v1/checks/CHK00734f349df8c14682adc321428a3075"

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Truora-API-Key", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiIiwiYWRkaXRpb25hbF9kYXRhIjoie30iLCJjbGllbnRfaWQiOiJUQ0llZWU3NWUzMjBiODhhYTdiYTVkYTkyYjFkYWU5Y2IxMCIsImV4cCI6MzI0NTI4MDMyNiwiZ3JhbnQiOiIiLCJpYXQiOjE2Njg0ODAzMjYsImlzcyI6Imh0dHBzOi8vY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb20vdXMtZWFzdC0xX2lYcDhPTVBpRyIsImp0aSI6ImZiYjEyNWQwLWM4NWItNDc3OS05ZmQ4LTE3NjRlMTZmODUyYiIsImtleV9uYW1lIjoiYXBpa2V5Iiwia2V5X3R5cGUiOiJiYWNrZW5kIiwidXNlcm5hbWUiOiJnbWFpbHN2Z29uemFsZXpyYW1vcy1hcGlrZXkifQ.199M11RxVNkeCNRtejcifZF5_u23ERDkudm0bwSPFe4")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
