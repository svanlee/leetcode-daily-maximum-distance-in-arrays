// Solution in JS
/**
 * @param {number[][]} arrays
 * @return {number}
 */
var maxDistance = function(arrays) {
    // If the input array is empty, return 0
    if (arrays.length === 0) return 0;

    // Initialize the minimum value to the first element of the first array
    let minVal = arrays[0][0];
    
    // Initialize the maximum value to the last element of the first array
    let maxVal = arrays[0][arrays[0].length - 1];
    
    // Initialize the maximum distance to 0
    let maxDistance = 0;

    // Iterate through the remaining arrays (starting from the second array)
    for (let i = 1; i < arrays.length; i++) {
        // Get the current array
        let currentArray = arrays[i];
        
        // Get the smallest element in the current array
        let currentMin = currentArray[0];
        
        // Get the largest element in the current array
        let currentMax = currentArray[currentArray.length - 1];

        // Update the maximum distance by considering two possibilities:
        // 1. The absolute difference between the largest element in the current array and the minimum value
        // 2. The absolute difference between the maximum value and the smallest element in the current array
        maxDistance = Math.max(maxDistance, Math.abs(currentMax - minVal), Math.abs(maxVal - currentMin));

        // Update the minimum value to be the minimum of the current minimum value and the smallest element in the current array
        minVal = Math.min(minVal, currentMin);
        
        // Update the maximum value to be the maximum of the current maximum value and the largest element in the current array
        maxVal = Math.max(maxVal, currentMax);
    }

    // Return the maximum distance found
    return maxDistance;
};