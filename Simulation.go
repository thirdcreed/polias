package polias

//Simulation contains all agents, all relationships, and all settings of polias model
type Simulation struct {
	Name         string
	SocialGroups []SocialGroup
	// aTable|SG|,|SG| with values in [−1,1]
	SocialGroupAttitudes [][]int
	Individuals          []Individual
}

func MakeIndividual() Individual {
	return Individual{SocialGroup: SocialGroup{Name: 2}, Beliefs: []Belief{}, Contacts: []Individual{}, Broadcast: make(chan Message)}
}
