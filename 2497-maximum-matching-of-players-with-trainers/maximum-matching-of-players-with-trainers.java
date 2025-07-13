// self
class Solution {
    public int matchPlayersAndTrainers(int[] players, int[] trainers) {
        Arrays.sort(players);
        Arrays.sort(trainers);
        int i = 0, j = 0, count = 0;
        int m=players.length,n=trainers.length;
        while (i<m && j<n) {
            if (players[i]<=trainers[j]) {
                i++; 
                count++;
            }
            j++;
        }
        return count;
    }
}