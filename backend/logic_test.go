package main

import "testing"

func TestIfItPassed(t *testing.T) {
	got := 10

	if got != 10 {
		t.Error("Not what we wanted")
	}
}

func TestPubSub(t *testing.T) {
	msgSent := publishMsg("api-project-70002766628", "test-topic", "Hello Whuniohorld:3453fd")
	print("____________")
	msgReceived := pullMsg("api-project-70002766628", "test-subscription")

	if msgSent != msgReceived {
		t.Error("Not what we wanted")
	}
}
