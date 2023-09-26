/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return make([][]int, 0)
    }

    q := &Queue{}

    q.PushBack(root)
    var bigList [][]int

    for q.Length() > 0 {
        var smallList []int

        size := q.Length()

        for i := 0; i < size; i++ {
            node := q.PopLeft()
            smallList = append(smallList, node.Val)

            if node.Left != nil {
                q.PushBack(node.Left)
            }

            if node.Right != nil {
                q.PushBack(node.Right)
            }
        }

        bigList = append(bigList, smallList)
    }

    return bigList
}

type Queue struct {
    Nodes []*TreeNode
}

func (q *Queue) PushBack(node *TreeNode) {
    q.Nodes = append(q.Nodes, node)
}

func (q *Queue) PopLeft() *TreeNode {
    if q.Length() == 0 {
        return nil
    }

    node := q.Nodes[0]
    q.Nodes = q.Nodes[1:]

    return node
}

func (q *Queue) Length() int {
    return len(q.Nodes)
}
