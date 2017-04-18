package polias

import "testing"

func TestRun(*testing.T) {
	i := MakeIndividual()
	i2 := MakeIndividual()
	s := []SocialGroup{SocialGroup{Name: 2}}
	aTable := [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
	i.Contacts = []Individual{i2}
	i2.Contacts = []Individual{i}
	go i2.listen(i.Broadcast)
	individuals := []Individual{i, i2}
	Run(Simulation{"test", s, aTable, individuals})
}
