package main

import (
	"bytes"
	"co.edu.meli/luisillera/prueba-tecnica/domain/dto"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a Application

func TestMain(m *testing.M) {
	a.initializeSatelliteProvider()
	a.initializeMessageProvider()
	a.initializeRoutes()
	a.initializeUseCases()
	code := m.Run()
	os.Exit(code)
}

func TestTopSecretRoute_OK(t *testing.T) {
	var jsonStr = []byte(`{"satellites":[{"name":"kenobi","distance":485.7,"message":["este","","","mensaje",""]},{"name":"skywalker","distance":266.1,"message":["","es","","","secreto"]},{"name":"sato","distance":600.5,"message":["este","","un","",""]}]}`)
	req, _ := http.NewRequest("POST", "/topsecret/", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	var body dto.Response
	decoder := json.NewDecoder(response.Body)

	if err := decoder.Decode(&body); err != nil {
		t.Errorf("Expected body not compliant: [%s]\n", err.Error())
	}

	if body.Position.X != -100 {
		t.Errorf("Expected Spaceship position.X [%.1f] Got [%.1f]\n", float64(-100), body.Position.X)
	}

	if body.Position.Y != 75.5 {
		t.Errorf("Expected Spaceship position.Y [%.1f] Got [%.1f]\n", 75.5, body.Position.X)
	}

	if body.Message != "este es un mensaje secreto" {
		t.Errorf("Expected Spaceship message [%s] Got [%s]\n", "este es un mensaje secreto", body.Message)
	}
}

func TestTopSecretRoute_BadRequest(t *testing.T) {
	var jsonStr = []byte(`{"satellites":[{"","mensaje",""]},{"name":"skywalker","distance":266.1,"message":["","es","","","secreto"]},{"name":"sato","distance":600.5,"message":["este","","un","",""]}]}`)
	req, _ := http.NewRequest("POST", "/topsecret/", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestTopSecretRoute_PositionCantBeFound(t *testing.T) {
	var jsonStr = []byte(`{"satellites":[{"name":"kenobi","distance":15400.7,"message":["este","","","mensaje",""]},{"name":"skywalker","distance":266.1,"message":["","es","","","secreto"]},{"name":"sato","distance":600.5,"message":["este","","un","",""]}]}`)
	req, _ := http.NewRequest("POST", "/topsecret/", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestTopSecretSplitRoute_OK(t *testing.T) {
	//kenobi
	firstCallJsonStr := []byte(`{"distance":485.7,"message":["este","","","mensaje",""]}`)
	req, _ := http.NewRequest("POST", "/topsecret_split/kenobi", bytes.NewBuffer(firstCallJsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	response := executeRequest(req)
	checkResponseIncompleteResponse(t, response, http.StatusOK, `{"error":"No hay suficiente informacion."}`)
	//skywalker
	secondCallJsonStr := []byte(`{"distance":266.1,"message":["","es","","","secreto"]}`)
	req, _ = http.NewRequest("POST", "/topsecret/skywalker", bytes.NewBuffer(secondCallJsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	checkResponseIncompleteResponse(t, response, http.StatusOK, `{"error":"No hay suficiente informacion."}`)
	//sato
	thirdCallJsonStr := []byte(`{"name":"sato","distance":600.5,"message":["este","","un","",""]}`)
	req, _ = http.NewRequest("POST", "/topsecret/sato", bytes.NewBuffer(thirdCallJsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")
	checkResponseIncompleteResponse(t, response, http.StatusOK, `{"error":"No hay suficiente informacion."}`)


}

func checkResponseIncompleteResponse(t *testing.T, response *httptest.ResponseRecorder, statusCode int, message string) {
	checkResponseCode(t, http.StatusOK, response.Code)
	body := response.Body.String()
	if body != message {
		t.Errorf("Expected Error message [%s] Got [%s]\n", message, body)
	}
}

func checkResponseCode(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	a.router.ServeHTTP(recorder, req)
	return recorder
}
