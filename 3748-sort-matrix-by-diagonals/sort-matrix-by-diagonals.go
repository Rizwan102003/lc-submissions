func sortMatrix(grid [][]int) [][]int {
  n := len(grid)
  for i := range n {
    var s []int
    for j := range n-i {
      s = append(s, grid[i+j][j])
    }
    sort.Slice(s, func(x, y int) bool { return s[x] > s[y] })
    for j := range n-i {
      grid[i+j][j] = s[j]
    }
  }

  for j:=1; j<n; j++ {
    var s []int
    for i := range n-j {
      s = append(s, grid[i][i+j])
    }
    sort.Ints(s)
    for i := range n-j {
      grid[i][i+j] = s[i]
    }
  }

  return grid
}
