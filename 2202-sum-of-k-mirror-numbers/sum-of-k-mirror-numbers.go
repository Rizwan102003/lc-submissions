func kMirror(k int, n int) int64 {
    dig:=make([]int,64)
    l:=1
    cnt:=0
    var sum int64=0
    for cnt<n{
        r:=l*10
        for i:=0;i<2&&cnt<n;i++{
            for j:=l;j<r&&cnt<n;j++{
                pali:=makePali(i==0,j)
                if palicheck(pali,k,dig){
                    sum+=pali
                    cnt++
                }
            }
        }
        l=r
    }
    return sum
}

func palicheck(m int64,k int,dig []int)bool{
    len:=-1
    for m>0{
        len++
        dig[len]=int(m%int64(k))
        m/=int64(k)
    }
    for i,j:=0,len;i<j;i,j=i+1,j-1{
        if dig[i]!=dig[j]{
            return false
        }
    }
    return true
}
func makePali(odd bool,base int) int64 {
    res := int64(base)
    if odd {
        base /= 10
    }
    for base > 0 {
        res = res*10 + int64(base%10)
        base /= 10
    }
    return res
}
