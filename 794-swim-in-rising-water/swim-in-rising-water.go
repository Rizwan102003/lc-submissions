const MAXN = 50
const LOWERMASK = 0xFFFF
const SHIFT = 16

var UF, INVIDX [MAXN * MAXN]int

func pack(i, j int) int {
	return i << SHIFT | j
}

func unpack(idx int) (i, j int) {
    return idx >> SHIFT, idx & LOWERMASK
}

func find(x int) int {
	if UF[x] == x {
		return x
	}
	UF[x] = find(UF[x])
	return UF[x]
}

func union(x, y int) {
	UF[find(x)] = find(y)
}

func swimInWater(grid [][]int) int {
	N := len(grid)
	if N == 1 {
		return 0
	}
	// Make the inverted mapping of value -> grid index
	// and initialize the union find (each value is connected to itself)
	for i := range N {
		for j := range N {
			level := grid[i][j]
            INVIDX[level] = pack(i, j)
			UF[level] = level
		}
	}
	// Swimmer reaches end when start and end connects
	start, end := grid[0][0], grid[N-1][N-1]
	reachedEnd := func() bool {
		return find(start) == find(end)
	}
	// Tries to swim to adjacent lower-level cell
	swimToAdjCell := func(level, i, j int) {
		adjlevel := grid[i][j]
		if level > adjlevel {
			union(level, adjlevel)
		}
	}
	level := 1
	for ; !reachedEnd(); level++ {
        i, j := unpack(INVIDX[level])
		if ii := i - 1; ii >= 0 {
			swimToAdjCell(level, ii, j)
		}
		if ii := i + 1; ii < N {
			swimToAdjCell(level, ii, j)
		}
		if jj := j - 1; jj >= 0 {
			swimToAdjCell(level, i, jj)
		}
		if jj := j + 1; jj < N {
			swimToAdjCell(level, i, jj)
		}
	}
	return level - 1
}