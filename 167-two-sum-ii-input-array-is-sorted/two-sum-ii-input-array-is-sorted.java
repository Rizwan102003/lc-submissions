class Solution {
    public int[] twoSum(int[] numbers, int target) {
        HashMap<Integer,Integer> index= new HashMap<>();
        int n=numbers.length;
        for (int i = 0; i < n; i++) {
            int current = numbers[i]; 
            int complement = target - current;
            if (index.containsKey(complement))
                return new int[] {index.get(complement)+1, i+1};
            index.put(current, i);
        }
        return new int[] {0,0};
    }
}