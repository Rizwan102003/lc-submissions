func maxDiff(num int) int {
    s:=strconv.Itoa(num)
    maxstr:=s
    for _,ch:=range s{
        if ch!='9'{
            maxstr = strings.ReplaceAll(s, string(ch), "9")
            break
        }
    }
    minstr:=s
    if s[0]!='1'{
        minstr=strings.ReplaceAll(s,string(s[0]),"1")
    }else{
        for _,ch:=range s[1:]{
            if ch!='0' && ch!='1'{
                minstr=strings.ReplaceAll(s,string(ch),"0")
                break
            }
        }
    }
    n1,_:=strconv.Atoi(maxstr)
    n2,_:=strconv.Atoi(minstr)
    return n1-n2
}