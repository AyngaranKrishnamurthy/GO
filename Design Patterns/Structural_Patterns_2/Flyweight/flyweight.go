package flyweight

import "time"

type Team struct { //The team struct has other structs inside, a total of 4 structs are created.
	ID             uint64           //ID of the team
	Name           string           //Name of the Team
	Shield         []byte           //Some image in a slice of bytes representing the team's shield
	Players        []Player         //A slice of Players
	HistoricalData []HistoricalData //A slice of Historical Data
}

const ( //ID of two  teams
	TEAM_A = iota //iota is an untyped integer that automatically increments its value for each new value between the parentheses
	TEAM_B
)

type Player struct { //Player struct having player data
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

type HistoricalData struct { //HistoricalData struct
	Year          uint8
	LeagueResults []Match
}

type Match struct { //Represents historical results of a match
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

type teamFlyweightFactory struct { //Factory to  create and store teams
	createdTeams map[int]*Team //Map of years including pointers to Team as values
}

func (t *teamFlyweightFactory) GetTeam(teamID int) *Team {
	if t.createdTeams[teamID] != nil {
		return t.createdTeams[teamID]
	}
	team := getTeamFactory(teamID)
	t.createdTeams[teamID] = &team

	return t.createdTeams[teamID]
}

func getTeamFactory(team int) Team {
	switch team {
	case TEAM_B:
		return Team{
			ID:   2,
			Name: "TEAM_B",
		}
	default:
		return Team{
			ID:   1,
			Name: "TEAM_A",
		}
	}
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int { //Return number of objects created
	return len(t.createdTeams)
}

func NewTeamFactory() teamFlyweightFactory { //Comment the snippet and we can face an error this is because Line49 has not found the team in the map it has created
	return teamFlyweightFactory{ //and is tring to assign it to the map, but the map is nil as we did not initialize it while creating the factory
		createdTeams: make(map[int]*Team),
	}
}
