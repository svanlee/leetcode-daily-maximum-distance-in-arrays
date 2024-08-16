// Solution in RUST
impl Solution {
    pub fn max_distance(arrays: Vec<Vec<i32>>) -> i32 {
        // Get the number of arrays
        let m = arrays.len();
        
        // Initialize the minimum and maximum values with their corresponding indices
        let mut min_d1 = (arrays[0][0], 0);
        let mut min_d2 = (arrays[1][0], 1);
        let mut max_d1 = (arrays[0][arrays[0].len()-1], 0);
        let mut max_d2 = (arrays[1][arrays[1].len()-1], 1);
        
        // Initialize the minimum and maximum values for the first two arrays
        if arrays[0][0] < arrays[1][0] {
            min_d1 = (arrays[0][0], 0);
            min_d2 = (arrays[1][0], 1);
        } else {
            min_d1 = (arrays[1][0], 1);
            min_d2 = (arrays[0][0], 0);
        }
        
        if arrays[0][arrays[0].len()-1] > arrays[1][arrays[1].len()-1] {
            max_d1 = (arrays[0][arrays[0].len()-1], 0);
            max_d2 = (arrays[1][arrays[1].len()-1], 1);
        } else {
            max_d1 = (arrays[1][arrays[1].len()-1], 1);
            max_d2 = (arrays[0][arrays[0].len()-1], 0);
        }
        
        // Iterate through the rest of the arrays
        for i in 2..m {
            let arr = &arrays[i];
            if arr[0] < min_d1.0 {
                min_d2 = min_d1;
                min_d1 = (arr[0], i as i32);
            } else if arr[0] < min_d2.0 {
                min_d2 = (arr[0], i as i32);
            }
            
            if arr[arr.len()-1] > max_d1.0 {
                max_d2 = max_d1;
                max_d1 = (arr[arr.len()-1], i as i32);
            } else if arr[arr.len()-1] > max_d2.0 {
                max_d2 = (arr[arr.len()-1], i as i32);
            }
        }
        
        // Calculate the maximum distance
        if min_d1.1 != max_d1.1 {
            return max_d1.0 - min_d1.0;
        } else {
            return i32::max((max_d1.0 - min_d2.0).abs(), (max_d2.0 - min_d1.0).abs());
        }
    }
}