import { assertEquals } from "@std/assert";
import { isValidParentheses } from "./main.ts";

Deno.test("isValidParentheses", () => {
  assertEquals(isValidParentheses("()"), true);
  assertEquals(isValidParentheses("()[]{}"), true);
  assertEquals(isValidParentheses("(]"), false);
  assertEquals(isValidParentheses("([)]"), false);
  assertEquals(isValidParentheses("{[]}"), true);
});
