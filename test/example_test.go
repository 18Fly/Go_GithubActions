package test

import (
	"studyProject/say"
	"testing"
)

func TestHello(t *testing.T) {
	say.Hello()
}

func TestGoodBye(t *testing.T) {
	say.GoodBye()
}

func TestSay(t *testing.T) {
	say.Say()
}
