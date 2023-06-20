package datastructures

type BinaryTreeNode struct {
	Value interface{}
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type TraversalOrder uint

const (
	InOrder TraversalOrder = iota
	PreOrder
	PostOrder
)

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (t *BinaryTree) Insert(v any) {
	if t.Root == nil {
		t.Root = &BinaryTreeNode{Value: v}

		return
	}

	current := t.Root

	for current != nil {
		if v.(int) < current.Value.(int) {
			if current.Left == nil {
				current.Left = &BinaryTreeNode{Value: v}

				return
			}

			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = &BinaryTreeNode{Value: v}

				return
			}

			current = current.Right
		}
	}
}

func (t *BinaryTree) Search(v any) *BinaryTreeNode {
	current := t.Root

	for current != nil {
		if current.Value == v {
			return current
		}

		if v.(int) < current.Value.(int) {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return nil
}

func (t *BinaryTree) Delete(v any) {
	if t.Root == nil {
		return
	}

	if t.Root.Left == nil && t.Root.Right == nil {
		if t.Root.Value == v {
			t.Root = nil
		}

		return
	}

	var targetNode, temp *BinaryTreeNode

	q := NewQueue()
	q.Enqueue(t.Root)

	for q.Len() > 0 {
		temp, _ = q.Dequeue().(*BinaryTreeNode)

		if temp.Value == v {
			targetNode = temp
		}

		if temp.Left != nil {
			q.Enqueue(temp.Left)
		}

		if temp.Right != nil {
			q.Enqueue(temp.Right)
		}
	}

	if targetNode != nil {
		targetNode.Value = temp.Value
		t.deleteDeepest(temp)
	}
}

func (t *BinaryTree) deleteDeepest(node *BinaryTreeNode) {
	q := NewQueue()
	q.Enqueue(t.Root)

	for q.Len() > 0 {
		if q.Peek().(*BinaryTreeNode) == node {
			q.Dequeue()

			return
		}

		temp, _ := q.Dequeue().(*BinaryTreeNode)

		if temp.Right != nil {
			if temp.Right == node {
				temp.Right = nil

				return
			}

			q.Enqueue(temp.Right)
		}

		if temp.Left != nil {
			if temp.Left == node {
				temp.Left = nil

				return
			}

			q.Enqueue(temp.Left)
		}
	}
}

func (t *BinaryTreeNode) IsLeaf() bool {
	return t.Left == nil && t.Right == nil
}

func (t *BinaryTreeNode) IsFull() bool {
	return t.Left != nil && t.Right != nil
}

func (t *BinaryTreeNode) IsHalf() bool {
	return (t.Left != nil && t.Right == nil) || (t.Left == nil && t.Right != nil)
}

func (t *BinaryTree) DepthFirstTraversal(order ...TraversalOrder) []any {
	var result []any

	if len(order) == 0 {
		order = append(order, InOrder)
	}

	switch order[0] {
	case InOrder:
		t.inOrderTraversal(t.Root, &result)
	case PreOrder:
		t.preOrderTraversal(t.Root, &result)
	case PostOrder:
		t.postOrderTraversal(t.Root, &result)
	default:
		t.inOrderTraversal(t.Root, &result)
	}

	return result
}

func (t *BinaryTree) inOrderTraversal(node *BinaryTreeNode, result *[]any) {
	if node == nil {
		return
	}

	t.inOrderTraversal(node.Left, result)
	*result = append(*result, node.Value)
	t.inOrderTraversal(node.Right, result)
}

func (t *BinaryTree) preOrderTraversal(node *BinaryTreeNode, result *[]any) {
	if node == nil {
		return
	}

	*result = append(*result, node.Value)
	t.preOrderTraversal(node.Left, result)
	t.preOrderTraversal(node.Right, result)
}

func (t *BinaryTree) postOrderTraversal(node *BinaryTreeNode, result *[]any) {
	if node == nil {
		return
	}

	t.postOrderTraversal(node.Left, result)
	t.postOrderTraversal(node.Right, result)
	*result = append(*result, node.Value)
}

func (t *BinaryTree) LevelOrderTraversal() []any {
	if t.Root == nil {
		return nil
	}

	var result []any

	queue := NewQueue()
	queue.Enqueue(t.Root)

	for !queue.IsEmpty() {
		node, ok := queue.Dequeue().(*BinaryTreeNode)
		if !ok {
			return nil
		}

		result = append(result, node.Value)

		if node.Left != nil {
			queue.Enqueue(node.Left)
		}

		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
	}

	return result
}

func (t *BinaryTree) BreadthFirstTraversal() []any {
	return t.LevelOrderTraversal()
}
