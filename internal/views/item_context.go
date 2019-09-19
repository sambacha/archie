package views

import (
	"github.com/briggysmalls/archie/internal/types"
)

// Create a system landscape view
func NewItemContextView(model *types.Model, scope *types.Element) types.Model {
	// Find relevant elements
	var elements []*types.Element

	// The main elements of interest are the children of the scope
	elements = append(elements, scope.Children...)

	// We also want to add elements:
	// - related to these children
	// - same depth as scope
	scopeDepth, err := model.Depth(scope)
	if err != nil {
		panic(err)
	}
	for _, rel := range model.ImplicitRelationships() {
		if linked := getLinked(rel, scope); linked != nil {
			linkedDepth, err := model.Depth(linked)
			if err != nil {
				panic(err)
			}
			if linkedDepth == scopeDepth {
				elements = append(elements, linked)
			}
		}
	}

	// Create a model from the model's root elements
	return CreateSubmodel(model, elements)
}

// Get the linked element, if the specified element is in the relationship
func getLinked(relationship types.Relationship, element *types.Element) *types.Element {
	if relationship.Source == element {
		return relationship.Destination
	}
	if relationship.Destination == element {
		return relationship.Source
	}
	return nil
}