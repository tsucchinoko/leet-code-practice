function twoSum(nums: number[], target: number): number[] {
  const map = new Map<number, number>();

  for (let i = 0; i < nums.length; i++) {
    const diff = target - nums[i];
    if (map.has(diff)) {
      return [map.get(diff)!, i];
    }
    map.set(nums[i], i);
  }

  throw new Error("solution was not found!");
}

// Main execution
if (import.meta.main) {
  const nums = [2, 7, 11, 15];
  const target = 9;
  const result = twoSum(nums, target);
  console.log("TwoSum result:", result);
}

export { twoSum };
