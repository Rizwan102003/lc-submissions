class Solution:
  def distributeCandies(self, n: int, limit: int) -> int:
    res = 0
    for i in range(min(n, limit) + 1):
      v = n - i
      res += max(0, v + 1 - 2 * max(0, v - limit))
    return res 