import { assertEquals, assertThrows } from "https://deno.land/std@0.224.0/assert/mod.ts";
import { twoSum } from "./main.ts";

Deno.test("twoSum - basic case", () => {
  const nums = [2, 7, 11, 15];
  const target = 9;
  const result = twoSum(nums, target);
  assertEquals(result, [0, 1]);
});

Deno.test("twoSum - different indices", () => {
  const nums = [3, 2, 4];
  const target = 6;
  const result = twoSum(nums, target);
  assertEquals(result, [1, 2]);
});

Deno.test("twoSum - same number twice", () => {
  const nums = [3, 3];
  const target = 6;
  const result = twoSum(nums, target);
  assertEquals(result, [0, 1]);
});

Deno.test("twoSum - no solution throws error", () => {
  const nums = [1, 2, 3];
  const target = 10;
  assertThrows(
    () => twoSum(nums, target),
    Error,
    "solution was not found!"
  );
});
