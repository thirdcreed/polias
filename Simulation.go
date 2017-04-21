package polias

import (
	"fmt"
	"io"

	"github.com/spf13/viper"
)

//Simulation contains all agents, all relationships, and all settings of polias model
type Simulation struct {
	Name         string
	SocialGroups []SocialGroup
	// aTable|SG|,|SG| with values in [−1,1]
	SocialGroupAttitudes [][]float64
	Individuals          []Individual
}

type socialGroupConfig struct {
	Name      string
	Attitudes []float64
}

type simulationConfig struct {
	SG      []socialGroupConfig
	Alpha   float64
	Epsilon float64
}

//MakeIndividual is a function that is temporary and stupid
func MakeIndividual() Individual {
	return Individual{SocialGroup: SocialGroup{Name: 2}, Beliefs: []Belief{}, Contacts: []Individual{}, Broadcast: make(chan Message)}
}

//LoadModel loads a Model
func LoadModel(config io.Reader) Simulation {
	viper.SetConfigType("yaml")

	var cfg simulationConfig

	if err := viper.ReadConfig(config); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file %s ", err))
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("fatal error marshaling %s ", err))
	}

	fmt.Println(cfg)

	return Simulation{}
}
