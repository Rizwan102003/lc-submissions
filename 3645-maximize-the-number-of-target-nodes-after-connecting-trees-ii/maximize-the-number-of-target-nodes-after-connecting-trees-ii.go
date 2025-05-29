func maxTargetNodes(edges1 [][]int, edges2 [][]int) []int {
    n, m := len(edges1)+1, len(edges2)+1
    g1 := make([][]int, n)
    g2 := make([][]int, m)
    for i := 0; i < n; i++ {
        g1[i] = make([]int, 0, 2)
    }
    for i := 0; i < m; i++ {
        g2[i] = make([]int, 0, 2)
    }
    for _, e := range edges1 {
        u, v := e[0], e[1]
        g1[u] = append(g1[u], v)
        g1[v] = append(g1[v], u)
    }
    for _, e := range edges2 {
        u, v := e[0], e[1]
        g2[u] = append(g2[u], v)
        g2[v] = append(g2[v], u)
    }
    p1, c1 := bfs(n, g1)
    _, c2 := bfs(m, g2)
    best2 := c2[0]
    if c2[1] > best2 {
        best2 = c2[1]
    }
    ans := make([]int, n)
    for i := 0; i < n; i++ {
        ans[i] = c1[p1[i]] + best2
    }
    return ans
}

func bfs(N int, g [][]int) ([]int, [2]int) {
    parity := make([]int, N)
    cnt := [2]int{}
    vis := make([]bool, N)
    q := make([]int, N)
    head, tail := 0, 0
    q[tail] = 0
    tail++
    vis[0] = true
    cnt[0] = 1
    for head < tail {
        u := q[head]
        head++
        pu := parity[u]
        for _, v := range g[u] {
            if !vis[v] {
                vis[v] = true
                pv := pu ^ 1
                parity[v] = pv
                cnt[pv]++
                q[tail] = v
                tail++
            }
        }
    }
    return parity, cnt
}