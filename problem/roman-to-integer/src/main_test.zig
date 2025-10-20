const std = @import("std");
const testing = std.testing;
const romanToInt = @import("main.zig").romanToInt;

test "romanToInt examples and edge cases" {
    try testing.expect(romanToInt("III") == 3);
    try testing.expect(romanToInt("IV") == 4);
    try testing.expect(romanToInt("IX") == 9);
    try testing.expect(romanToInt("LVIII") == 58);
    try testing.expect(romanToInt("MCMXCIV") == 1994);
    try testing.expect(romanToInt("XL") == 40);
    try testing.expect(romanToInt("XC") == 90);
    try testing.expect(romanToInt("CD") == 400);
    try testing.expect(romanToInt("CM") == 900);
    // minimal and maximal valid values
    try testing.expect(romanToInt("I") == 1);
    try testing.expect(romanToInt("MMMCMXCIX") == 3999); // 3000 + 900 + 90 + 9
}
