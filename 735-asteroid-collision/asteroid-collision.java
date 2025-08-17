class Solution {
    public int[] asteroidCollision(int[] asteroids) {
        Stack <Integer> boom = new Stack <Integer> ();
        for(int i:asteroids){
            boolean destroyed = false;
            while (!boom.empty() && boom.peek() > 0 && i < 0) {
                int top = boom.peek();
                if (Math.abs(top) < Math.abs(i)) {
                    boom.pop(); // top asteroid destroyed
                    continue;   // keep checking collisions
                } else if (Math.abs(top) == Math.abs(i)) {
                    boom.pop(); // both destroyed
                }
                destroyed = true;
                break;
            }
            if (!destroyed) {
                boom.push(i);
            }
        }
        return boom.stream().mapToInt(Integer::intValue).toArray();
    }
}