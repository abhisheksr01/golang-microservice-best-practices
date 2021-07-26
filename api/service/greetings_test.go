package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var username string = "Abhishek"

func TestGreetings_shouldReturnMorningGreeting_whenTimeIsGreaterThanEqualTo4AndLessThan12(t *testing.T) {
	
	expectedGreetingMsg := "Hi Abhishek, Good Morning."

	actualGreetingMsg := Greet(time.Date(2021, 1, 1, 5, 0, 0, 0, time.UTC), username)

	assert.Equal(t, expectedGreetingMsg, actualGreetingMsg)
}

func TestGreetings_shouldReturnAfternoonGreeting_whenTimeIsGreaterThanEqualTo12AndLessThan17(t *testing.T) {
	username := "Abhishek"
	expectedGreetingMsg := "Hi Abhishek, Good Afternoon."
	
	actualGreetingMsg := Greet(time.Date(2021, 1, 1, 14, 0, 0, 0, time.UTC), username)
	
	assert.Equal(t, expectedGreetingMsg, actualGreetingMsg)
}

func TestGreetings_shouldReturnEveningGreeting_whenTimeIsGreaterThanEqualTo17AndLessThan20(t *testing.T) {
	username := "Abhishek"
	expectedGreetingMsg := "Hi Abhishek, Good Evening."
	
	actualGreetingMsg := Greet(time.Date(2021, 1, 1, 19, 0, 0, 0, time.UTC), username)
	
	assert.Equal(t, expectedGreetingMsg, actualGreetingMsg)
}

func TestGreetings_shouldReturnNightGreeting_whenTimeIsGreaterThanEqualTo20AndLessThan4(t *testing.T) {
	username := "Abhishek"
	expectedGreetingMsg := "Hi Abhishek, Good Night."
	
	actualGreetingMsg := Greet(time.Date(2021, 1, 1, 21, 0, 0, 0, time.UTC), username)
	
	assert.Equal(t, expectedGreetingMsg, actualGreetingMsg)
}