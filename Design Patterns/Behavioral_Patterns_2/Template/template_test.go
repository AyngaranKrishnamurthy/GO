package template

import (
	"strings"
	"testing"
)

type TestStruct struct {
	Template //It embeds the Template to call the ExecuteAlgorithm method
}

func (m *TestStruct) Message() string { //TestStruct type implementing(implicit) MessageRetriever interface returning the string "world"
	return "world"
}

func TestTemplate_ExecuteAlgorithm(t *testing.T) {
	t.Run("Using Interfaces", func(t *testing.T) {
		s := &TestStruct{}
		res := s.ExecuteAlgorithm(s)
		expected := "world"
		if !strings.Contains(res, expected) {
			t.Errorf("Expected string was not found on returned string")
		}
	})
}
