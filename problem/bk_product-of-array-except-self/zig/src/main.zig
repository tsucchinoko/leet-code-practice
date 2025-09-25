const std = @import("std");

pub fn productExceptSelf(allocator: *std.mem.Allocator, nums: []const i32) ![]i32 {
    const n = nums.len;
    var out = try allocator.alloc(i32, n);
    // left products into out
    var left: i32 = 1;
    for (nums, 0..) |v, i| {
        out[i] = left;
        left = left * v;
    }
    // multiply by right products on the fly
    var right: i32 = 1;
    var i = n;
    while (i > 0) {
        i -= 1;
        out[i] = out[i] * right;
        right = right * nums[i];
    }
    return out;
}

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();
    var allocator = std.heap.page_allocator;

    // example 1
    const nums1: [4]i32 = .{ 1, 2, 3, 4 };
    const out1 = try productExceptSelf(allocator, nums1[0..]);
    try stdout.print("input: {any}\n", .{nums1});
    try stdout.print("output: ");
    for (out1, 0..) |v, i| {
        if (i != 0) try stdout.print(", ", .{});
        try stdout.print("{d}", .{v});
    }
    try stdout.print("\n", .{});
    allocator.free(out1);

    // example 2
    const nums2: [5]i32 = .{ -1, 1, 0, -3, 3 };
    const out2 = try productExceptSelf(allocator, nums2[0..]);
    try stdout.print("input: {any}\n", .{nums2});
    try stdout.print("output: ");
    for (out2, 0..) |v, i| {
        if (i != 0) try stdout.print(", ", .{});
        try stdout.print("{d}", .{v});
    }
    try stdout.print("\n", .{});
    allocator.free(out2);
}
