// Solution in JAVA
class Solution {
    public int maxDistance(List<List<Integer>> arrays) {
        // Initialize the maximum distance to 0
        int res = 0;
        
        // Initialize the index to 1, starting from the second array
        int i = 1;

        // Initialize the minimum value to the first element of the first array
        int min = arrays.get(0).get(0);
        
        // Initialize the maximum value to the last element of the first array
        int max = arrays.get(0).get(arrays.get(0).size() - 1);

        // Iterate through the remaining arrays (starting from the second array)
        while (i < arrays.size()) {
            // Get the current array
            List<Integer> m = arrays.get(i);
            
            // Get the smallest element in the current array
            int currentMin = m.get(0);
            
            // Get the largest element in the current array
            int currentMax = m.get(m.size() - 1);

            // Update the maximum distance by considering two possibilities:
            // 1. The absolute difference between the largest element in the current array and the minimum value
            // 2. The absolute difference between the maximum value and the smallest element in the current array
            res = Math.max(res, Math.abs(currentMax - min));
            res = Math.max(res, Math.abs(max - currentMin));

            // Update the minimum value to be the minimum of the current minimum value and the smallest element in the current array
            min = Math.min(currentMin, min);
            
            // Update the maximum value to be the maximum of the current maximum value and the largest element in the current array
            max = Math.max(currentMax, max);

            // Move to the next array
            i++;
        } 

        // Return the maximum distance found
        return res;
    }
}