func atmost(nums *[]int , k int) int {
    ans := 0 ;
    var m = make(map[int]int);
    start := 0 ;
    for i:=0 ; i<len(*nums) ; i++ {
        m[(*nums)[i]]++;
        for len(m) > k {
            if m[(*nums)[start]] == 1 {
                delete(m , (*nums)[start]);
            }else{
                m[(*nums)[start]]--;
            }

            start++;
        }

        ans += i - start + 1 ;
    }

    return ans ;
}
func countCompleteSubarrays(nums []int) int {
    var m = make(map[int]int);

    for _ , num := range nums {
        m[num]++;
    }

    uni := len(m);

    return atmost(&nums , uni) - atmost(&nums , uni - 1) ;
}