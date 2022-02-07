package flyweight

import "testing"

func TestTeamFlyWeightFactory_GetTeam(t *testing.T) {
	factory := teamFlyweightFactory{
		createdTeams: make(map[int]*Team, 0),
	}

	teamA1 := factory.GetTeam(TEAM_A)
	if teamA1 == nil {
		t.Errorf("Pointer to TEAM_A returned nil")
	}

	teamA2 := factory.GetTeam(TEAM_A)
	if teamA2 == nil {
		t.Errorf("Pointer to TEAM_A returned nil")
	}

	if teamA1 != teamA2 {
		t.Errorf("The pointers to TEAM_A weren't the same")
	}

	if factory.GetNumberOfObjects() != 1 {
		t.Errorf("The number of objects was not same as expected number of objects")
	}
}
