// Solution in CS
public class Solution
{
    public int MaxDistance(IList<IList<int>> arrays)
    {
        // Initialize the result variable to store the maximum distance
        int result = 0;

        // Initialize the maximum and minimum values with the first array
        int max = arrays[0][arrays[0].Count - 1]; // Maximum value in the first array
        int min = arrays[0][0]; // Minimum value in the first array

        // Iterate through the rest of the arrays
        for (int i = 1; i < arrays.Count; i++)
        {
            // Get the last and first elements of the current array
            int last = arrays[i][arrays[i].Count - 1]; // Last element of the current array
            int first = arrays[i][0]; // First element of the current array

            // Calculate the maximum distance between the current array and the previous arrays
            result = Math.Max(result, Math.Abs(last - min)); // Distance between the last element of the current array and the minimum value so far
            result = Math.Max(result, Math.Abs(max - first)); // Distance between the maximum value so far and the first element of the current array

            // Update the minimum and maximum values
            min = Math.Min(min, first); // Update the minimum value
            max = Math.Max(max, last); // Update the maximum value
        }

        // Return the maximum distance
        return result;
    }
}