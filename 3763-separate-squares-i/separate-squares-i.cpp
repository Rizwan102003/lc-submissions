const int N=1e5;
uint64_t events[N];
class Solution {
public:
    static inline uint64_t pack(uint64_t y, uint64_t L, bool isStart){
        return (y<<32)|(L<<1)|isStart;
    }
    static double separateSquares(vector<vector<int>>& squares) {
        const int n=squares.size(), n2=n*2;
        double A0=0;
        for(int i=0; i<n; i++){
            auto& square=squares[i];
            uint64_t y=square[1], L=square[2];
            A0+=L*L;
            events[2*i]=pack(y, L, 1);
            events[2*i+1]=pack(y+L, L, 0);
        }
        sort(events, events+n2);
        double sumWidth=0, area=0, prevY=0, A2=A0/2.0;
        for(int i=0; i<n2; i++){
            uint64_t event=events[i];
            int y=event>>32, L=(event&UINT_MAX)>>1, dy=y-prevY;
            double curA=dy*sumWidth;
            if(area+curA>=A2){// sumWidth is never 0  
                return prevY+(A2-area)/sumWidth;
            }
            bool isStart=event&1;
            sumWidth+=(isStart*2-1)*L;
            area+=curA;
            prevY=y;
        }
        return 0;
    }
};
auto init = []()
{ 
    ios::sync_with_stdio(0);
    cin.tie(0);
    cout.tie(0);
    return 'c';
}();