class Solution {
    public boolean uniqueOccurrences(int[] arr) {
        Map<Integer, Integer> store = new HashMap<>();
        for (int i : arr)
            store.put(i, store.getOrDefault(i, 0) + 1);
        Set<Integer> seen = new HashSet<>();
        for (int i : store.values())
            if (!seen.add(i)) 
                return false;
        return true;
    }
}