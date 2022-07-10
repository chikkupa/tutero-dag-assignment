package step

import (
	"errors"
	"testing"
	"tutero_assignment/pkg/src/graph"
)

type TestCase struct {
	graph    graph.Graph
	expected graph.Node
	errors   error
}

func getTestCases() []TestCase {
	var testgraph1 graph.Graph
	var testgraph2 graph.Graph

	var nodes = []graph.Node{
		"A",
		"B",
		"C",
		"D",
		"E",
	}

	for _, node := range nodes {
		testgraph1.AddNode(node)
	}

	testgraph1.AddEdge(nodes[0], nodes[1])
	testgraph1.AddEdge(nodes[0], nodes[2])
	testgraph1.AddEdge(nodes[0], nodes[3])

	testcases := []TestCase{
		{testgraph1, nodes[0], nil},
		{testgraph2, nodes[0], errors.New("Empty graph")},
	}

	return testcases
}

func TestStep(t *testing.T) {
	testCases := getTestCases()

	for index, test := range testCases {
		stepper := New()
		node, err := stepper.Step(test.graph)
		if err != nil && node != test.expected {
			t.Fatalf("Test case %d: Test failed", index+1)
		} else if err != nil && test.errors == nil {
			t.Fatalf("Test case %d: Error expected!", index+1)
		}
	}
}
