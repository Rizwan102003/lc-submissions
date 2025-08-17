class Solution {
    public int[] asteroidCollision(int[] asteroids) {
        Stack <Integer> boom = new Stack <Integer> ();
        for(int i:asteroids){
            boolean destroyed = false;
            while (!boom.empty() && boom.peek() > 0 && i < 0) { // [-ve,+ve] no collision
                int top = boom.peek();
                if (Math.abs(top) < Math.abs(i)) {
                    boom.pop();
                    continue;   // keep checking collisions
                } else if (Math.abs(top) == Math.abs(i))
                    boom.pop();
                destroyed = true;
                break;
            }
            if (!destroyed) boom.push(i);
        }
        return boom.stream().mapToInt(Integer::intValue).toArray();
    }
}