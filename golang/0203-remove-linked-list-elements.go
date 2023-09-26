/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeElements(head *ListNode, val int) *ListNode {
    result := new(ListNode)
    ptr := result
    ptr.Next = head

    for ptr != nil && ptr.Next != nil {
        if ptr.Next.Val == val {
            ptr.Next = ptr.Next.Next
        } else {
            ptr = ptr.Next
        }
    }

    return result.Next
}