from collections import deque

class Solution(object):
    def findLexSmallestString(self, s, a, b):
        """
        :type s: str
        :type a: int
        :type b: int
        :rtype: str
        """
        l = len(s)
        q = deque([s])
        seen = set([s])
        res = s
        
        while q:
            curr = q.pop()
            if curr < res:
                res = curr
            
            temp = list(curr)
            for i in range(1, l, 2):
                temp[i] = str((int(temp[i]) + a) % 10)
            temp_s = ''.join(temp)
            if temp_s not in seen:
                seen.add(temp_s)
                q.append(temp_s)
            
            rotated = curr[-b:] + curr[:-b]
            if rotated not in seen:
                seen.add(rotated)
                q.append(rotated)
        
        return res