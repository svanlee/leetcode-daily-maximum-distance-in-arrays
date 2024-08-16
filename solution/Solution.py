# Solution in PYTHON3
from typing import List


class Solution:
    def maxDistance(self, arrays: List[List[int]]) -> int:
        # Initialize the minimum and second minimum values to positive infinity
        mn1, mn2 = float('inf'), float('inf')
        
        # Initialize the maximum and second maximum values to negative infinity
        mx1, mx2 = float('-inf'), float('-inf')
        
        # Iterate over each array in the input
        for i in arrays:
            # Get the minimum and maximum values in the current array
            mn, mx = i[0], i[-1]
            
            # If the current minimum is less than the overall minimum, update the overall minimum and second minimum
            if mn < mn1:
                mn2, mn1 = mn1, mn
            # If the current minimum is less than the second minimum but not less than the overall minimum, update the second minimum
            elif mn < mn2:
                mn2 = mn
            
            # If the current maximum is greater than the overall maximum, update the overall maximum and second maximum
            if mx > mx1:
                mx2, mx1 = mx1, mx
            # If the current maximum is greater than the second maximum but not greater than the overall maximum, update the second maximum
            elif mx > mx2:
                mx2 = mx
        
        # If the overall maximum and minimum are in the same array, or the second maximum and minimum are in the same array, return the difference
        if mx1 == mx2 or mn1 == mn2:
            return mx1 - mn1
        
        # If the overall maximum and minimum are not in the same array, iterate over the arrays again to find the array that contains the overall maximum and minimum
        for i in arrays:
            # If the current array contains the overall maximum and minimum, return the maximum distance
            if (i[0], i[-1]) == (mn1, mx1):
                return max(abs(mx2 - mn1), abs(mx1 - mn2))
        
        # If the overall maximum and minimum are not in the same array, return the difference
        return mx1 - mn1