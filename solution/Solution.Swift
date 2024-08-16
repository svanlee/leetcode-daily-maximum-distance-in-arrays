// Solution in SWIFT
class Solution {
    func maxDistance(_ arrays: [[Int]]) -> Int {
        // Initialize the maximum distance found so far
        var res = 0
        
        // Initialize the minimum and maximum values seen so far
        var minVal = arrays[0].min()!
        var maxVal = arrays[0].max()!
        
        // Iterate through each array starting from the second array
        for i in 1..<arrays.count {
            // Calculate the maximum distance between the current array's last element and the minimum value seen so far
            // and between the current array's first element and the maximum value seen so far
            res = max(
                res,
                max(
                    abs(arrays[i].max()! - minVal),
                    abs(maxVal - arrays[i].min()!)
                )
            )
            
            // Update the minimum and maximum values seen so far
            minVal = min(minVal, arrays[i].min()!)
            maxVal = max(maxVal, arrays[i].max()!)
        }
        
        // Return the maximum distance found
        return res
    }
}