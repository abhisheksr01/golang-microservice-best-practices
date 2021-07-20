package main

import "github.com/cucumber/godog"

func afternoonTimeIsBetweenAnd(arg1, arg2 string) error {
	return godog.ErrPending
}

func eveningTimeIsBetweenAnd(arg1, arg2 string) error {
	return godog.ErrPending
}

func morningTimeIsBetweenAnd(arg1, arg2 string) error {
	return godog.ErrPending
}

func nightTimeIsBetweenAnd(arg1, arg2 string) error {
	return godog.ErrPending
}

func theUserInvokesTheAPI() error {
	return godog.ErrPending
}

func theUserShouldBeGreetedWithEitherOrOrOr(arg1, arg2, arg3, arg4 string) error {
	return godog.ErrPending
}

func theUserWantsToBeGreeted(arg1 string) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^Afternoon time is between "([^"]*)" and "([^"]*)"$`, afternoonTimeIsBetweenAnd)
	ctx.Step(`^Evening time is between "([^"]*)" and "([^"]*)"$`, eveningTimeIsBetweenAnd)
	ctx.Step(`^Morning time is between "([^"]*)" and "([^"]*)"$`, morningTimeIsBetweenAnd)
	ctx.Step(`^Night time is between "([^"]*)" and "([^"]*)"$`, nightTimeIsBetweenAnd)
	ctx.Step(`^the user invokes the API$`, theUserInvokesTheAPI)
	ctx.Step(`^the user should be greeted with either "([^"]*)" or "([^"]*)" or "([^"]*)" or "([^"]*)"$`, theUserShouldBeGreetedWithEitherOrOrOr)
	ctx.Step(`^the user "([^"]*)" wants to be greeted$`, theUserWantsToBeGreeted)
}
