//Package polias - The goal of the Polias model is to model the impacts on attitudes, within a
//population, of a series of actions (e.g. commercial operation,
//change in public policy or, in our application context: medical support,
//military intervention, bombing..) by some external actors called forces
//(e.g. company, NGO or in our military application context:
//United Nation force, terrorist or others) over time. These impacts are
//(subjectively) evaluated by the individuals through an evolving attitude
//value toward each force at stake.
package polias

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

//Message - During a simulation, actors communicate on their actions to the
//population in which the information is propagated. These communications
//are emitted through messages defined by: m=⟨emitter,date,a,Adr⟩
type Message struct {
	Emitter Individual
	Date    time.Time
	//MessengersBelief Belief
	Adresses []Individual
}

//Belief is short for action belief.  It encapsulates an agents belief about
// a particular action that was taken.
type Belief struct {
	Name           string
	Actor          SocialObject
	Coresponsibles []SocialObject
	Date           time.Time
	// Benefactor describes the agent who principally benefits from the action
	Benefactor SocialObject
	//Payoff is a float value representing, the benefit of the action to the
	//benefactor. This value can be negative.
	Payoff float64
}

// Beliefs are simply a collection of Action Beliefs
type Beliefs []Belief

// Individual is defined by a tuple i=⟨sg,Blf,Cnt⟩ with sg∈SG the social group of the individual,
// Blf the set of all the beliefs on actions present in the individual's memory and Cnt ⊂ Ind−{i}
// the set of all the contacts of the individual in a social network.
type Individual struct {
	SocialGroup SocialGroup
	Beliefs     Beliefs
	Contacts    []Individual
	Broadcast   chan Message
}

//SocialGroup is a set of individuals sharing similar characteristics or goals
type SocialGroup struct {
	Name int
}

// A SocialObject is an element capable of emitting an action message (Social Group, Individual)
type SocialObject interface {
	listen(chan Message)
}

func (listener Individual) listen(msg chan Message) {
	for {
		message := <-msg
		fmt.Println(message)
		listener.Broadcast <- message
	}
}

// Run starts the Polias simulation clock atickin'
func Run(simulation Simulation) {
	fmt.Println("got this far")
	for _, i := range simulation.Individuals {
		fmt.Println("broadcasting to", i.Contacts)
		viper.AddConfigPath("test")
		i.Broadcast <- Message{Emitter: i, Date: time.Now(), Adresses: i.Contacts}
	}
}
