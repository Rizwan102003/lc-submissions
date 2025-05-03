class Solution:
    def minDominoRotations(self, A: List[int], B: List[int]) -> int:
        
        d = { i: set() for i in range(1, 7) }
        for i in range(len(A)):
            d[A[i]].add(i)
            d[B[i]].add(i)
                
        res = float('inf')
        flag = True
        
        for k in d:
            if len(d[k]) == len(A):
                flag = False
                top, bottom = 0, 0
                for i in range(len(A)):
                    if A[i] != k: top += 1
                    if B[i] != k: bottom += 1
                res = min([res, top, bottom])
        
        if flag: return -1
        return res