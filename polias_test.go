package polias

import (
	"bytes"
	"testing"
)

func TestRun(*testing.T) {

	var model = []byte(`
SG:
  - Name: "Democrats"
    Attitudes:
      - 0
      - 0.5
      - 0.3
  - Name: "Republicans"
    Attitudes:
      - 0
      - 0.5
      - 0.3
  - Name: "Independents"
    Attitudes:
      - 0
      - 0.5
      - 0.3
alpha: .5
epsilon: .3`)

	LoadModel(bytes.NewBuffer(model))

	//Run(Simulation{"test", s, aTable, individuals})
}
