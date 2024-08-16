// Solution in TS
function maxDistance(arrays: number[][]): number {
    // Initialize the maximum distance found so far
    let res = 0;

    // Initialize the minimum and maximum values seen so far
    let minVal = arrays[0][0];
    let maxVal = arrays[0][arrays[0].length - 1];

    // Iterate through each array starting from the second array
    for (let i = 1; i < arrays.length; i++) {
        // Calculate the maximum distance between the current array's last element and the minimum value seen so far
        // and between the current array's first element and the maximum value seen so far
        res = Math.max(
            res,
            Math.max(
                Math.abs(arrays[i][arrays[i].length - 1] - minVal),
                Math.abs(maxVal - arrays[i][0])
            )
        );

        // Update the minimum and maximum values seen so far
        minVal = Math.min(minVal, arrays[i][0]);
        maxVal = Math.max(maxVal, arrays[i][arrays[i].length - 1]);
    }

    // Return the maximum distance found
    return res;
}