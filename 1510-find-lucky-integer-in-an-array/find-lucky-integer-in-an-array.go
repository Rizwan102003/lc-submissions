func findLucky(arr []int) int {
  sort.Ints(arr)
  count := 1
  for i := len(arr)-2; i >= 0; i-- {
    if arr[i] != arr[i+1] {
      if count == arr[i+1] {
        return arr[i+1]
      } else {
        count = 1
      }
    } else {
      count++
    }
  }
  
  if arr[0] == count {
    return arr[0]
  }

  return -1
}