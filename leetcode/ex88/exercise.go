// https://leetcode.com/problems/merge-sorted-array/
func merge(nums1 []int, m int, nums2 []int, n int)  {
    t1 := []int{}
    t2 := []int{}
    
    t1 = append(t1, nums1[:m]...)
    t2 = append(t2, nums2[:n]...)
    
    i := 0
    for i < m + n {
        
        // t1 がなくなったら探索終了
        if len(t1) == 0 {
            break
        }
        
        // t2が空だったら問答無用で t1 を優先
        if len(t2) == 0 || t1[0] < t2[0] {
            nums1[i] = t1[0]
            t1 = t1[1:]
        } else {
            nums1[i] = t2[0]
            t2 = t2[1:]
        }
        
        i++
    }
    
    // t2 が残ってたら追加
    for len(t2) > 0 {
        nums1[i] = t2[0]
        t2 = t2[1:]
        i++
    }
}
