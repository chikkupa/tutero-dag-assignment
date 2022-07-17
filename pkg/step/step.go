package step

import (
	"errors"
	"tutero_assignment/pkg/src/graph"
)

type Stepper interface {
	// Step returns a prediction for the correct node; or an error if a prediction cannot be made.
	Step(graph graph.Graph) (graph.Node, error)
}

func New() *stepper {
	newStepper := new(stepper)
	//* You may mutate this instantiation if necessary; but the function signature should not change.
	return newStepper
}

type stepper struct {
	//* You may add fields to this struct.
}

func (s *stepper) Step(graph graph.Graph) (graph.Node, error) {
	//* Implement the Step function.

	allNodes := graph.Nodes()

	if len(allNodes) == 0 {
		return "", errors.New("No nodes")
	}

	nodeLeastRelation := graph.Nodes()[len(allNodes)-1]
	noRelations := len(graph.Children(nodeLeastRelation)) + len(graph.Parents(nodeLeastRelation))

	for _, node := range allNodes {
		noNodeRelations := len(graph.Children(node)) + len(graph.Parents(node))
		if noNodeRelations < noRelations {
			nodeLeastRelation = node
			noRelations = noNodeRelations
		}
	}

	return nodeLeastRelation, nil //graph.Nodes()[len(allNodes)-1], nil
}
