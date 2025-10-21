const std = @import("std");

fn digitToRoman(allocator: std.mem.Allocator, d: u8, one: []const u8, five: []const u8, ten: []const u8) ![]u8 {
    if (d == 0) {
        return allocator.dupe(u8, "");
    } else if (d <= 3) {
        var result = try allocator.alloc(u8, d * one.len);
        var i: u8 = 0;
        while (i < d) : (i += 1) {
            @memcpy(result[i * one.len .. (i + 1) * one.len], one);
        }
        return result;
    } else if (d == 4) {
        var result = try allocator.alloc(u8, one.len + five.len);
        @memcpy(result[0..one.len], one);
        @memcpy(result[one.len..], five);
        return result;
    } else if (d == 5) {
        return allocator.dupe(u8, five);
    } else if (d <= 8) {
        const extra_ones = d - 5;
        var result = try allocator.alloc(u8, five.len + extra_ones * one.len);
        @memcpy(result[0..five.len], five);
        var i: u8 = 0;
        while (i < extra_ones) : (i += 1) {
            @memcpy(result[five.len + i * one.len .. five.len + (i + 1) * one.len], one);
        }
        return result;
    } else { // d == 9
        var result = try allocator.alloc(u8, one.len + ten.len);
        @memcpy(result[0..one.len], one);
        @memcpy(result[one.len..], ten);
        return result;
    }
}

pub fn toRoman(allocator: std.mem.Allocator, num: u16) ![]u8 {
    const thousands = @as(u8, @intCast(num / 1000));
    const hundreds = @as(u8, @intCast((num % 1000) / 100));
    const tens = @as(u8, @intCast((num % 100) / 10));
    const ones = @as(u8, @intCast(num % 10));

    // Calculate total length needed
    var total_len: usize = 0;

    // thousands: only 'M' repeated
    total_len += thousands;

    // hundreds
    const hundreds_part = try digitToRoman(allocator, hundreds, "C", "D", "M");
    defer allocator.free(hundreds_part);
    total_len += hundreds_part.len;

    // tens
    const tens_part = try digitToRoman(allocator, tens, "X", "L", "C");
    defer allocator.free(tens_part);
    total_len += tens_part.len;

    // ones
    const ones_part = try digitToRoman(allocator, ones, "I", "V", "X");
    defer allocator.free(ones_part);
    total_len += ones_part.len;

    // Allocate result buffer
    var result = try allocator.alloc(u8, total_len);
    var pos: usize = 0;

    // Add thousands
    var i: u8 = 0;
    while (i < thousands) : (i += 1) {
        result[pos] = 'M';
        pos += 1;
    }

    // Add hundreds
    if (hundreds_part.len > 0) {
        @memcpy(result[pos .. pos + hundreds_part.len], hundreds_part);
        pos += hundreds_part.len;
    }

    // Add tens
    if (tens_part.len > 0) {
        @memcpy(result[pos .. pos + tens_part.len], tens_part);
        pos += tens_part.len;
    }

    // Add ones
    if (ones_part.len > 0) {
        @memcpy(result[pos .. pos + ones_part.len], ones_part);
        pos += ones_part.len;
    }

    return result;
}

pub fn main() !void {
    const allocator = std.heap.page_allocator;

    const nums = [_]u16{ 3749, 58, 1994, 4, 9, 40, 944 };

    for (nums) |n| {
        const r = try toRoman(allocator, n);
        defer allocator.free(r);
        std.debug.print("{d} -> {s}\n", .{ n, r });
    }
}

test "roman conversion basic cases" {
    const allocator = std.testing.allocator;

    const cases = &[_]struct {
        num: u16,
        want: []const u8,
    }{
        .{ .num = 1, .want = "I" },
        .{ .num = 3, .want = "III" },
        .{ .num = 4, .want = "IV" },
        .{ .num = 9, .want = "IX" },
        .{ .num = 58, .want = "LVIII" },
        .{ .num = 1994, .want = "MCMXCIV" },
        .{ .num = 3749, .want = "MMMDCCXLIX" },
        .{ .num = 40, .want = "XL" },
        .{ .num = 90, .want = "XC" },
        .{ .num = 400, .want = "CD" },
        .{ .num = 900, .want = "CM" },
        .{ .num = 3999, .want = "MMMCMXCIX" },
    };

    for (cases) |c| {
        const got = try toRoman(allocator, c.num);
        try std.testing.expect(std.mem.eql(u8, got, c.want));
        allocator.free(got);
    }
}
