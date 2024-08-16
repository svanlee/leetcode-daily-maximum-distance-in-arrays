// Solution in CPP
class Solution
{
public:
    int maxDistance(vector<vector<int>> &arrays)
    {
        // Disable synchronization between C++ streams and C streams for performance optimization
        ios_base::sync_with_stdio(false);
        cin.tie(NULL);
        cout.tie(NULL);

        // Initialize the maximum distance to 0
        int ans(0);

        // Initialize the minimum value to the first element of the first array
        int mn(arrays[0][0]);

        // Initialize the maximum value to the last element of the first array
        int mx(arrays[0][arrays[0].size() - 1]);

        // Iterate through the remaining arrays (starting from the second array)
        for (int i(1); i < arrays.size(); ++i)
        {
            // Update the maximum distance by considering two possibilities:
            // 1. The absolute difference between the largest element in the current array and the minimum value
            // 2. The absolute difference between the maximum value and the smallest element in the current array
            ans = max(ans, max(abs(arrays[i][arrays[i].size() - 1] - mn), abs(mx - arrays[i][0])));

            // Update the minimum value to be the minimum of the current minimum value and the smallest element in the current array
            mn = min(mn, arrays[i][0]);

            // Update the maximum value to be the maximum of the current maximum value and the largest element in the current array
            mx = max(mx, arrays[i][arrays[i].size() - 1]);
        }

        // Return the maximum distance found
        return ans;
    }
};