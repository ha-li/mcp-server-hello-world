package mcp

import "testing"

// mocktail:Pineapple

func TestMock(t *testing.T) {
	var s Pineapple = newPineappleMock(t).
		OnJuice("foo", Water{}).TypedReturns(Water{}).
		Once().
		Parent

	s.Juice("foo", Water{})
}
