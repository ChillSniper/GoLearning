package Selectlearning

import "testing"

func TestLearnNoDefaultAndCaseError(t *testing.T) {
	LearnNoDefaultAndCaseError()
}

func TestSeveralCaseAndDefault(t *testing.T) {
	SeveralCaseAndDefault()
}

func TestRandomChoice(t *testing.T) {
	for i := 0; i < 1000; i++ {
		RandomChoice()
	}
}
