const std = @import("std");
const roman_to_integer = @import("roman_to_integer");

fn value(c: u8) i32 {
    const result: i32 = switch (c) {
        'I' => 1,
        'V' => 5,
        'X' => 10,
        'L' => 50,
        'C' => 100,
        'D' => 500,
        'M' => 1000,
        else => 0,
    };

    return result;
}

pub fn romanToInt(s: []const u8) i32 {
    var i: usize = 0;
    var sum: i32 = 0;
    const n = s.len;

    while (i < n) {
        const cur = value(s[i]);
        const next = if (i + 1 < n) value(s[i + 1]) else 0;

        if (cur < next) {
            sum += next - cur;
            i += 2;
        } else {
            sum += cur;
            i += 1;
        }
    }

    return sum;
}

pub fn main() !void {
    var stdout_buffer: [1024]u8 = undefined;
    var stdout_writer = std.fs.File.stdout().writer(&stdout_buffer);
    const stdout = &stdout_writer.interface;

    const examples = [_][]const u8{ "III", "LVIII", "MCMXCIV" };
    for (examples) |ex| {
        const v = romanToInt(ex);
        try stdout.print("{s} -> {d}\n", .{ ex, v });
    }
    try stdout.flush();
}
