// Solution in GO
// Define a custom struct to hold a value and its index
type data struct {
	val int
	idx int
}

// Function to find the maximum distance between two elements in different arrays
func maxDistance(arrs [][]int) int {
	// Get the number of arrays
	m := len(arrs)
	
	// Initialize the first two arrays
	arr0 := arrs[0]
	arr1 := arrs[1]
	
	// Initialize the minimum and maximum values with their corresponding indices
	var minD1, minD2, maxD1, maxD2 data
	
	// Initialize the minimum and maximum values for the first two arrays
	if arr0[0] < arr1[0] {
		// If the first element of arr0 is smaller, set it as minD1 and the first element of arr1 as minD2
		minD1 = data{arr0[0], 0}
		minD2 = data{arr1[0], 1}
	} else {
		// Otherwise, set the first element of arr1 as minD1 and the first element of arr0 as minD2
		minD1 = data{arr1[0], 1}
		minD2 = data{arr0[0], 0}
	}
	
	if arr0[len(arr0)-1] > arr1[len(arr1)-1] {
		// If the last element of arr0 is larger, set it as maxD1 and the last element of arr1 as maxD2
		maxD1 = data{arr0[len(arr0)-1], 0}
		maxD2 = data{arr1[len(arr1)-1], 1}
	} else {
		// Otherwise, set the last element of arr1 as maxD1 and the last element of arr0 as maxD2
		maxD1 = data{arr1[len(arr1)-1], 1}
		maxD2 = data{arr0[len(arr0)-1], 0}
	}
	
	// Iterate through the rest of the arrays
	for i := 2; i < m; i++ {
		arr := arrs[i]
		
		// Update the minimum values
		if arr[0] < minD1.val {
			// If the first element of the current array is smaller than minD1, update minD1 and minD2
			minD2 = minD1
			minD1 = data{arr[0], i}
		} else if arr[0] < minD2.val {
			// If the first element of the current array is smaller than minD2, update minD2
			minD2 = data{arr[0], i}
		}
		
		// Update the maximum values
		if arr[len(arr)-1] > maxD1.val {
			// If the last element of the current array is larger than maxD1, update maxD1 and maxD2
			maxD2 = maxD1
			maxD1 = data{arr[len(arr)-1], i}
		} else if arr[len(arr)-1] > maxD2.val {
			// If the last element of the current array is larger than maxD2, update maxD2
			maxD2 = data{arr[len(arr)-1], i}
		}
	}
	
	// Calculate the maximum distance
	if minD1.idx != maxD1.idx {
		// If the indices of minD1 and maxD1 are different, return the difference between their values
		return maxD1.val - minD1.val
	} else {
		// Otherwise, return the maximum of the absolute differences between maxD1 and minD2, and maxD2 and minD1
		return max(abs(maxD1.val-minD2.val), abs(maxD2.val-minD1.val))
	}
}

// Function to calculate the absolute value of a number
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}