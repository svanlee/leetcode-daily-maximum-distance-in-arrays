// Solution in PHP
class Solution {

/**
* @param Integer[][] $arrays
* @return Integer
*/
function maxDistance($arrays) {
// Initialize the maximum distance found so far
$res = 0;

// Initialize the minimum and maximum values seen so far
$minVal = $arrays[0][0];
$maxVal = $arrays[0][count($arrays[0]) - 1];

// Iterate through each array starting from the second array
for ($i = 1; $i < count($arrays); $i++) { // Calculate the maximum distance between the current array's last element and
    the minimum value seen so far // and between the current array's first element and the maximum value seen so far
    $res=max( $res, max( abs($arrays[$i][count($arrays[$i]) - 1] - $minVal), abs($maxVal - $arrays[$i][0]) ) ); //
    Update the minimum and maximum values seen so far $minVal=min($minVal, $arrays[$i][0]); $maxVal=max($maxVal,
    $arrays[$i][count($arrays[$i]) - 1]); } // Return the maximum distance found return $res; } }