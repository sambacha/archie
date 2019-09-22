package views

import (
	"github.com/briggysmalls/archie/internal/types"
)

// Create a system landscape view
func NewLandscapeView(model *types.Model) types.Model {
	// Create a model from the model's root elements
	view, err := CreateSubmodel(model, model.RootElements())
	// We shouldn't error (we've pulled elements out sensibly)
	if err != nil {
		panic(err)
	}
	return view
}
