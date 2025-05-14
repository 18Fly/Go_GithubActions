package test

import (
	"testing"

	"github.com/18Fly/mydemo/say"
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
