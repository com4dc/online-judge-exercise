func merge(nums1 []int, m int, nums2 []int, n int)  {
    t1 := nums1[0:m]
    t2 := nums2[0:n]
      
    for m > 0 && n > 0{

        if t1[m-1] >= t2[n-1] {
            nums1[m+n-1] = t1[m-1]
            m--
        } else {
            nums1[m+n-1] = t2[n-1]
            n--
        }
    }
    
    if n > 0 {
        nums1[n-1] = t2[n-1]
        n--
    }
 }

