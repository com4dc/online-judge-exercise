/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    
    root := &ListNode{}
    start := root
    
    target := head
    tmp := -99
    for target != nil {
        
        if target.Val != tmp {
            start.Next = target
            start = start.Next
            tmp = target.Val
        } else {
            start.Next = (*ListNode)(nil)
        }
        target = target.Next
    }
    
    return root.Next
}

