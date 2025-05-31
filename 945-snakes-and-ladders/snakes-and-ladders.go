func snakesAndLadders(board [][]int) int {
    n:=len(board)
    Cord:=func(k int) (int,int){
        q,r:=(k-1)/n,(k-1)%n//getting coordinates following zigzag pattern(boustrophedon)
        row:=n-1-q      
        col:=r//if the rows even
        if q%2==1{  
            col=n-1-r//if the rows odd
        }
        return row,col
    }
    vis:=make([]bool,n*n+1)
    queue:=[]int{1}//initialize queue at 1 as the start point of the board
    moves:=0//record number of steps taken
    for len(queue)>0{
        next:=[]int{}//record next dice rolls(moves)
        for _,curr:=range queue{
            if curr==n*n{//if we reach the destination return the moves
                return moves
            }
            for i:=1;i<=6 && curr+i<=n*n;i++{//now compute for all possibilities up until n*n
                k:=curr+i
                r,c:=Cord(k)//get coordinates
                if board[r][c]!=-1{
                    k=board[r][c]//get position
                }
                if !vis[k]{//visit the unvisited
                    vis[k]=true
                    next=append(next,k)
                }
            }
        }
        queue=next//move to the next layer
        moves++//increment the steps taken
    }
    return -1//if the nodes are over before n*n return -1
}