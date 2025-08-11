export function twoSum(nums: number[], target: number): number[] {
  const indexMap = new Map<number, number>();

  for (let i = 0; i < nums.length; i++) {
    const diff = target - nums[i];
    const complementIndex = indexMap.get(diff);
    if (complementIndex !== undefined) {
      return [complementIndex, i];
    }
    indexMap.set(nums[i], i);
  }

  throw new Error("No two sum solution found");
}
