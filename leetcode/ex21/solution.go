
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 * https://leetcode.com/problems/merge-two-sorted-lists/
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    rootNode := &ListNode{}
    
    start := rootNode
    t1 := l1
    t2 := l2
    
    // どっちかが空になるまでぶん回す
    for t1 != nil && t2 != nil {
        
        if t1.Val < t2.Val {
            start.Next = t1
            t1 = t1.Next
        } else {
            start.Next = t2
            t2 = t2.Next
        }
        start = start.Next
    }
    
    // 最後残ってる方のNodeを全部紐付ける
    if t1 != nil {
        start.Next = t1
    } else {
        start.Next = t2
    }
    return rootNode.Next 
    // startはガンガン書き換わってる。rootNodeはアドレスなのでそれのNext取れば計算後の値が取れる
}


