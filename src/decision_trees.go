package src

import (
	"math"
)

// data: numpy.ndarray, default=None
// children: dict(feat_value: Node), default=None
// split_on: int, default=None
// pred_class : str, default=None
// is_leaf: bool, default=False

type Node struct {
	data      string
	children  map[int]Node
	splitOn   int
	predicted int
	isLeaf    bool
}

func CreateDecisionTree() Node {
	var root Node
	return root
}

func FitDecisionTree() {

}

func PredictDecisionTree(currentNode Node) int {
	for {
		if currentNode.isLeaf {
			return currentNode.predicted
		} else {

		}
	}
}

func calculateEntropy(Y []int) float32 {
	labelCounts := make(map[int]int)
	for _, label := range Y {
		labelCounts[label]++
	}

	totalInstances := len(Y)
	entropy := 0.0

	for _, count := range labelCounts {
		probability := float64(count) / float64(totalInstances)
		entropy += probability * math.Log2(1/probability)
	}

	return float32(entropy)
}

func splitOnFunction(data string, featureIndex int) {

}
