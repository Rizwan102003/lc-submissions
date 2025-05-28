func join(a,b int,mp map[int][]int){
    mp[a]=append(mp[a],b)
    mp[b]=append(mp[b],a)
}
var visited map[int]bool
var cnt int
func dfs(src,k int, mp map[int][]int){
    if k<0{
        return
    }
    cnt++
    visited[src]=true;
    for _,v:=range mp[src]{
        if !visited[v]{
            dfs(v,k-1,mp)
        }
    }
    return
}
func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
    visited=make(map[int]bool)
    mp1,mp2:=make(map[int][]int),make(map[int][]int)
    for _,v:=range edges1{
        join(v[0],v[1],mp1)
    }
    for _,v:=range edges2{
        join(v[0],v[1],mp2)
    }
    n:=len(edges2)
    val:=0
    for i:=0;i<=n;i++{
        for k:=0;k<=n;k++{
            visited[k]=false
        }
        cnt=0
        dfs(i,k-1,mp2)
        
        val=max(val,cnt)
    }
    ans:=[]int{}
    n=len(edges1)
    for i:=0;i<=n;i++{
        for k:=0;k<=n;k++{
            visited[k]=false
        }
        cnt=0
        dfs(i,k,mp1)
        ans=append(ans,cnt+val)
    }
    return ans

    
}