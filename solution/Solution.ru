# Solution in RUST
# @param {Integer[][]} arrays
# @return {Integer}
def max_distance(arrays)
  # If the input array is empty, return 0
  return 0 if arrays.empty?

  # Initialize the minimum value to the first element of the first array
  min_val = arrays[0][0]
  
  # Initialize the maximum value to the last element of the first array
  max_val = arrays[0][-1]
  
  # Initialize the maximum distance to 0
  max_distance = 0

  # Iterate through the remaining arrays (starting from the second array)
  (1...arrays.length).each do |i|
    # Get the current array
    current_array = arrays[i]
    
    # Get the smallest element in the current array
    current_min = current_array[0]
    
    # Get the largest element in the current array
    current_max = current_array[-1]

    # Update the maximum distance by considering two possibilities:
    # 1. The absolute difference between the largest element in the current array and the minimum value
    # 2. The absolute difference between the maximum value and the smallest element in the current array
    max_distance = [max_distance, (current_max - min_val).abs, (max_val - current_min).abs].max

    # Update the minimum value to be the minimum of the current minimum value and the smallest element in the current array
    min_val = [min_val, current_min].min
    
    # Update the maximum value to be the maximum of the current maximum value and the largest element in the current array
    max_val = [max_val, current_max].max
  end

  # Return the maximum distance found
  max_distance
end