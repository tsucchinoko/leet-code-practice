const std = @import("std");
const testing = std.testing;
const main = @import("main.zig");

fn sliceEquals(a: []const i32, b: []const i32) bool {
    if (a.len != b.len) return false;
    for (a, 0..) |v, i| {
        if (v != b[i]) return false;
    }
    return true;
}

test "productExceptSelf example 1" {
    var allocator = std.testing.allocator;
    const nums: [4]i32 = .{ 1, 2, 3, 4 };
    const expected: [4]i32 = .{ 24, 12, 8, 6 };

    const out = try main.productExceptSelf(&allocator, nums[0..]);
    defer allocator.free(out);

    try testing.expect(sliceEquals(out, expected[0..]));
}

test "productExceptSelf example 2" {
    var allocator = std.testing.allocator;
    const nums: [5]i32 = .{ -1, 1, 0, -3, 3 };
    const expected: [5]i32 = .{ 0, 0, 9, 0, 0 };

    const out = try main.productExceptSelf(&allocator, nums[0..]);
    defer allocator.free(out);

    try testing.expect(sliceEquals(out, expected[0..]));
}
