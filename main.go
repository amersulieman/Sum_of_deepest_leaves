package main

//binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func deepestLeavesSum(root *TreeNode) int {

	//in case only root is given :)
	currentLevelSum := root.Val

	var currentLevel = []*TreeNode{root}
	var nextLevel []*TreeNode = nil

	reachedLastLevel := false
	for !reachedLastLevel {
		nextLevel = getNextLevel(currentLevel)
		if len(nextLevel) == 0 {
			reachedLastLevel = true
		} else {
			currentLevel = make([]*TreeNode, len(nextLevel))
			copy(currentLevel, nextLevel)
			currentLevelSum = sum(currentLevel)
		}
	}
	return currentLevelSum
}

func getNextLevel(currentLevel []*TreeNode) []*TreeNode {
	var nextLevel []*TreeNode = nil
	for _, node := range currentLevel {
		if node.Left != nil {
			nextLevel = append(nextLevel, node.Left)
		}
		if node.Right != nil {
			nextLevel = append(nextLevel, node.Right)
		}
	}
	return nextLevel
}

func sum(slice []*TreeNode) int {

	total := 0
	for _, item := range slice {
		total += item.Val
	}
	return total
}
