function minSubarray(nums: number[], p: number): number {
    const n = nums.length;
    let totalSum = 0;

    totalSum = nums.reduce((acc, cur) => (acc + cur) % p, totalSum);
    const target = totalSum % p;
    if(!target) return 0;

    const map = new Map();
    map.set(0, -1);
    let curSum = 0, minLen = n;

    nums.forEach((n, i) => {
        curSum = (curSum + n) % p;
        const needed = (curSum - target + p) % p;
        if(map.has(needed)) minLen = Math.min(minLen, i - map.get(needed));
        map.set(curSum, i);
    })

    return minLen === n ? -1 : minLen;
};