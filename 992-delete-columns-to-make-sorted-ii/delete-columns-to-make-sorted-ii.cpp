class Solution {
public:
    using  uint128_t=unsigned __int128;
    int minDeletionSize(vector<string>& strs) {
        int n=strs.size(), m=strs[0].size();
        uint128_t isSorted=0, remove=0;
        int cnt=0;
        for (int j=0; j<m; j++) {
            for (int i=0; i<n-1; i++) {
                if ((((isSorted>>i)&1)==0) && strs[i][j]>strs[i+1][j] ) {
                    remove|=(((uint128_t)1)<<j);
                    break;
                }
            }

            if ((remove>>j)&1) continue;
            
            for (int i=0; i<n-1; i++) 
                isSorted|=(((uint128_t)(strs[i][j]<strs[i+1][j])<<i));
        }
        return __builtin_popcountll((unsigned long long)(remove& ULLONG_MAX))
     +__builtin_popcountll((unsigned long long)(remove>>64));
    }
};