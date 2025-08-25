func findDiagonalOrder(mat [][]int) []int {
    m, n := len(mat), len(mat[0])
    result := make([]int, 0, m*n)
    
    dirs := [2][2]int{{-1, 1}, {1, -1}}
    row, col, d := 0, 0, 0
    
    for i := 0; i < m*n; i++ {
        result = append(result, mat[row][col])
        row += dirs[d][0]
        col += dirs[d][1]
        
        if row >= m {
            row = m - 1
            col += 2
            d ^= 1
        } else if col >= n {
            col = n - 1
            row += 2
            d ^= 1
        } else if row < 0 {
            row = 0
            d ^= 1
        } else if col < 0 {
            col = 0
            d ^= 1
        }
    }
    return result
}