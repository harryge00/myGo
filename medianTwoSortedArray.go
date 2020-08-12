func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1) < len(nums2) {
        tmp := nums1
        nums1 = nums2
        nums2 = tmp
    }
    m := len(nums1)
    n := len(nums2)
    halfLen := (m + n) / 2
    iMin, iMax := 0, m - 1
    var i, j int
    for iMin <= iMax {
    	i = (iMin + iMax) / 2
    	j = halfLen - i
    	if j == 0 || nums1[i] >= nums2[j - 1] {
    		break
    	}
    	if j > 0 && nums1[j - 1] > nums2[i] {
    		iMin = i + 1
    	}
    	if j < n && nums1[i - 1] > nums2[j] {
    		iMax = i - 1
    	}
    }
    if (m + n) % 2 == 0 {
    	return (math.Min(float64(nums1[i]), float64(nums2[j])) + math.Max(float64(nums1[i - 1]), float64(nums2[j - 1]))) / 2.0
    }
    return math.Min(float64(nums1[i]), float64(nums2[j]))

    return 0;
}