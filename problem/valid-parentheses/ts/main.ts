export function isValidParentheses(text: string): boolean {
  const pairs: Record<string, string> = {
    ")": "(",
    "]": "[",
    "}": "{",
    ">": "<",
  } as const;

  const stack: string[] = [];

  for (const char of text) {
    if (char in pairs) {
      // 閉じ括弧の場合はスタックのトップと比較
      if (stack.length === 0) return false;

      const expectedOpen = pairs[char];
      const lastOpen = stack[stack.length - 1];
      if (lastOpen !== expectedOpen) return false;

      stack.pop();
    } else {
      // 開き括弧はスタックに積む
      stack.push(char);
    }
  }

  // スタックが空ならすべて対応した括弧ペア
  return stack.length === 0;
}
