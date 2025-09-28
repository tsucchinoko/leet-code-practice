const std = @import("std");
const expect = std.testing.expect;
const main = @import("main.zig");

test "canCompleteCircuit basic cases" {
    // Test case 1: should return 3
    const gas1 = [_]i32{ 1, 2, 3, 4, 5 };
    const cost1 = [_]i32{ 3, 4, 5, 1, 2 };
    const res1 = main.canCompleteCircuit(gas1[0..], cost1[0..]);
    try expect(res1 == 3);

    // Test case 2: should return -1
    const gas2 = [_]i32{ 2, 3, 4 };
    const cost2 = [_]i32{ 3, 4, 3 };
    const res2 = main.canCompleteCircuit(gas2[0..], cost2[0..]);
    try expect(res2 == -1);
}

test "canCompleteCircuit edge cases" {
    // Single station with zero gas and cost
    const gas3 = [_]i32{0};
    const cost3 = [_]i32{0};
    try expect(main.canCompleteCircuit(gas3[0..], cost3[0..]) == 0);

    // Case where solution exists from index 4
    const gas4 = [_]i32{ 5, 1, 2, 3, 4 };
    const cost4 = [_]i32{ 4, 4, 1, 5, 1 };
    // Total gas = total cost = 15, starting from index 4 works
    try expect(main.canCompleteCircuit(gas4[0..], cost4[0..]) == 4);

    // Truly impossible case: total gas < total cost
    const gas5 = [_]i32{ 1, 2 };
    const cost5 = [_]i32{ 3, 4 };
    // Total: (1+2) - (3+4) = -4 < 0, impossible
    try expect(main.canCompleteCircuit(gas5[0..], cost5[0..]) == -1);
}

test "canCompleteCircuit single element" {
    // Single station, enough gas
    const gas = [_]i32{10};
    const cost = [_]i32{5};
    try expect(main.canCompleteCircuit(gas[0..], cost[0..]) == 0);

    // Single station, not enough gas
    const gas2 = [_]i32{2};
    const cost2 = [_]i32{5};
    try expect(main.canCompleteCircuit(gas2[0..], cost2[0..]) == -1);
}
