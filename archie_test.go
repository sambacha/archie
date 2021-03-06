package archie

import (
	"testing"

	"github.com/briggysmalls/archie/writers"
	"gotest.tools/assert"
)

var yaml = `
elements:
  - name: user
    kind: actor
  - name: sound system
    children:
      - name: speaker
        children:
          - name: enclosure
            technology: physical
          - name: driver
            technology: electro-mechanical
          - connector
          - cable
      - name: amplifier
        children:
          - audio in connector
          - audio out connector
          - ac-dc converter
          - mixer
          - amplifier circuit
          - name: power button
            technology: electro-mechanical
          - name: input select
            technology: electro-mechanical
associations:
  - source: user
    destination: sound system/amplifier/input select
  - source: sound system/amplifier/input select
    destination: sound system/amplifier/mixer
  - source: sound system/amplifier/mixer
    destination: sound system/amplifier/audio in connector
  - source: sound system/amplifier/ac-dc converter
    destination: sound system/amplifier/mixer
  - source: sound system/amplifier/ac-dc converter
    destination: sound system/amplifier/amplifier circuit
  - source: sound system/amplifier/amplifier circuit
    destination: sound system/amplifier/audio out connector
  - source: sound system/amplifier/audio out connector
    destination: sound system/speaker/cable
  - source: sound system/speaker/cable
    destination: sound system/speaker/connector
  - source: sound system/speaker/connector
    destination: sound system/speaker/driver
  - source: sound system/speaker/driver
    destination: sound system/speaker/enclosure
  - source: sound system/amplifier/power button
    destination: sound system/amplifier/ac-dc converter
`

func TestContext(t *testing.T) {
	// Create an archie
	a, err := New(writers.MermaidStrategy{}, yaml)
	assert.NilError(t, err)

	// Create a landscape view
	_, err = a.ContextView("sound system")
	assert.NilError(t, err)

	// Create diagram but the scope does not exist
	_, err = a.ContextView("I don't exist")
	assert.ErrorContains(t, err, "I don't exist")
}

func TestExternalStrategy(t *testing.T) {
	// Create a custom strategy
	_, err := New(strategy{}, yaml)
	assert.NilError(t, err)
}

type strategy struct {
}

func (s strategy) Header(scribe writers.Scribe) {
	scribe.WriteLine("graph TD")
	scribe.UpdateIndent(1)
}

func (s strategy) Footer(scribe writers.Scribe) {
	// Do nothing
}

func (s strategy) Element(scribe writers.Scribe, element writers.Element) {
	scribe.WriteLine("%s(%s)", element.ID(), element.Name())
	// Also add a hyperlink
	scribe.WriteLine("click %s %s", element.ID(), "mermaidCallback")
}

func (s strategy) StartParentElement(scribe writers.Scribe, element writers.Element) {
	scribe.WriteLine("subgraph %s", element.Name())
}

func (s strategy) EndParentElement(scribe writers.Scribe, element writers.Element) {
	scribe.WriteLine("end")
}

func (s strategy) Association(scribe writers.Scribe, association writers.Association) {
	scribe.WriteLine("%s-->%s", association.Source().ID(), association.Destination().ID())
}
