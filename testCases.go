package main

func setupTest1() (root TreeNode) {
	num8 := TreeNode{
		Val:   8,
		Left:  nil,
		Right: nil,
	}

	num7 := TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	num6 := TreeNode{
		Val:   6,
		Left:  nil,
		Right: &num8,
	}

	num5 := TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	num4 := TreeNode{
		Val:   4,
		Left:  &num7,
		Right: nil,
	}
	num3 := TreeNode{
		Val:   3,
		Left:  nil,
		Right: &num6,
	}
	num2 := TreeNode{
		Val:   2,
		Left:  &num4,
		Right: &num5,
	}
	root = TreeNode{
		Val:   1,
		Left:  &num2,
		Right: &num3,
	}
	return root
}
