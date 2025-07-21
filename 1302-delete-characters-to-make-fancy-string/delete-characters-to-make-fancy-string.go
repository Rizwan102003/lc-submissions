func makeFancyString(s string) string {     
    ans := []byte{s[0]}
    
    count := 1
    
    for i := 1; i < len(s); i++ {   
        if s[i] == s[i - 1] {
            count++
            
            if count < 3 {
                ans = append(ans, s[i])
            }
        } else{
            count =  1
             ans = append(ans, s[i])
        }
    }
        
    return string(ans)   
} 