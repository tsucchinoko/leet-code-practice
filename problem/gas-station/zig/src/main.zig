const std = @import("std");

pub fn canCompleteCircuit(gas: []const i32, cost: []const i32) i32 {
    const n = gas.len;
    if (n == 0 or n != cost.len) return -1;

    var total: i64 = 0;
    var tank: i64 = 0;
    var start: i32 = 0;

    var i: usize = 0;
    while (i < n) : (i += 1) {
        const diff: i64 = @as(i64, @intCast(gas[i])) - @as(i64, @intCast(cost[i]));
        total += diff;
        tank += diff;
        if (tank < 0) {
            // can't start from current start, move start to i+1
            start = @as(i32, @intCast(i)) + 1;
            tank = 0;
        }
    }

    if (total < 0) return -1;
    return start;
}

pub fn main() !void {
    var stdout_buffer: [1024]u8 = undefined;
    var stdout_writer = std.fs.File.stdout().writer(&stdout_buffer);
    const stdout = &stdout_writer.interface;
    // Example usage: read nothing, show examples
    const gas1 = [_]i32{ 1, 2, 3, 4, 5 };
    const cost1 = [_]i32{ 3, 4, 5, 1, 2 };
    const gas2 = [_]i32{ 2, 3, 4 };
    const cost2 = [_]i32{ 3, 4, 3 };

    try stdout.print("Example1 -> start = {d}\n", .{canCompleteCircuit(gas1[0..], cost1[0..])});
    try stdout.print("Example2 -> start = {d}\n", .{canCompleteCircuit(gas2[0..], cost2[0..])});
    try stdout.flush();
}
