package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/cucumber/godog"
)

type TestState struct {
	userName           string
	morningStartTime   int
	morningEndTime     int
	afternoonStartTime int
	afternoonEndTime   int
	eveningStartTime   int
	eveningEndTime     int
	nightStartTime     int
	nightEndTime       int
	apiResponse        *http.Response
}

func (t *TestState) theUserWantsToBeGreeted(arg1 string) error {
	t.userName = arg1
	return nil
}

func (t *TestState) theUserInvokesTheAPI() (err error) {
	req, err := http.NewRequest("GET", "http://localhost:5000/greetings/greet/"+t.userName, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	t.apiResponse = resp

	return nil
}

func (t *TestState) morningTimeIsBetweenAnd(arg1, arg2 int) error {
	t.morningStartTime = arg1
	t.morningEndTime = arg2
	return nil
}

func (t *TestState) afternoonTimeIsBetweenAnd(arg1, arg2 int) error {
	t.afternoonStartTime = arg1
	t.afternoonEndTime = arg2
	return nil
}

func (t *TestState) eveningTimeIsBetweenAnd(arg1, arg2 int) error {
	t.eveningStartTime = arg1
	t.eveningEndTime = arg2
	return nil
}

func (t *TestState) nightTimeIsBetweenAnd(arg1, arg2 int) error {
	t.nightStartTime = arg1
	t.nightEndTime = arg2
	return nil
}

func (t *TestState) theUserShouldBeGreetedWithEitherOrOrOr(morningMsg, afternoonMsg, eveningMsg, nightMsg string) error {
	b, err := io.ReadAll(t.apiResponse.Body)
	if err != nil {
		log.Fatalln(err)
	}

	expectedGreeting := ""
	if current_hour := time.Now().Hour(); current_hour >= t.morningStartTime && current_hour < t.morningEndTime {
		expectedGreeting = morningMsg
	} else if current_hour >= t.afternoonStartTime && current_hour < t.afternoonEndTime {
		expectedGreeting = afternoonMsg
	} else if current_hour >= t.eveningStartTime && current_hour < t.eveningEndTime {
		expectedGreeting = eveningMsg
	} else {
		expectedGreeting = nightMsg
	}
	if expectedGreeting != string(b) {
		log.Fatalln("Response doesn't match")
	}

	return fmt.Errorf("")
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	t := &TestState{}

	ctx.Step(`^Afternoon time is between "([^"]*)" and "([^"]*)"$`, t.afternoonTimeIsBetweenAnd)
	ctx.Step(`^Evening time is between "([^"]*)" and "([^"]*)"$`, t.eveningTimeIsBetweenAnd)
	ctx.Step(`^Morning time is between "([^"]*)" and "([^"]*)"$`, t.morningTimeIsBetweenAnd)
	ctx.Step(`^Night time is between "([^"]*)" and "([^"]*)"$`, t.nightTimeIsBetweenAnd)
	ctx.Step(`^the user invokes the API$`, t.theUserInvokesTheAPI)
	ctx.Step(`^the user should be greeted with either "([^"]*)" or "([^"]*)" or "([^"]*)" or "([^"]*)"$`, t.theUserShouldBeGreetedWithEitherOrOrOr)
	ctx.Step(`^the user "([^"]*)" wants to be greeted$`, t.theUserWantsToBeGreeted)
}
