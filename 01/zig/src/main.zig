const std = @import("std");

pub fn main() !void {
    var file = try std.fs.cwd().openFile("input-real.txt", .{});
    defer file.close();

    var buf: [1024]u8 = undefined;
    var file_reader = file.reader(&buf);

    var dial2: i32 = 50;
    var zeros2: i32 = 0;

    var dial1: i32 = 50;
    var zeros1: i32 = 0;

    while (try file_reader.interface.takeDelimiter('\n')) |line| {
        const o = try step_part2(dial2, line);
        dial2 = o.@"0";
        zeros2 += o.@"1";

        dial1 = try step_part1(dial1, line);
        zeros1 += @intFromBool(dial1 == 0);
    }

    std.debug.print("found {} zeros in password 1.\n", .{zeros1});
    std.debug.print("found {} zeros in password 2.\n", .{zeros2});
}

pub fn step_part1(start: i32, input: []const u8) !i32 {
    const inputNum = try std.fmt.parseInt(i32, input[1..], 10);
    const move = @mod(inputNum, 100);
    const end = start + (@as(i32, @intFromBool(input[0] == 'L')) * -1 + @as(i32, @intFromBool(input[0] == 'R'))) * move;

    return @mod(end,100);
}

pub fn step_part2(start: i32, input: []const u8) !struct { i32, i32 } {
    const inputNum = try std.fmt.parseInt(i32, input[1..], 10);
    const move = @mod(inputNum, 100);
    const end = start + (@as(i32, @intFromBool(input[0] == 'L')) * -1 + @as(i32, @intFromBool(input[0] == 'R'))) * move;
    const zeros = @divTrunc(inputNum, 100) + @intFromBool(end >= 100 or (start != 0 and end <= 0));

    return struct { i32, i32 }{ @intCast(@abs(@mod(end,100))), zeros };
}

test "simple test" {
    // sample
    try std.testing.expectEqual(@as(struct { i32, i32 }, struct { i32, i32 }{ 82, 1 }), step_part2(50, "L68"));
    try std.testing.expectEqual(@as(struct { i32, i32 }, struct { i32, i32 }{ 52, 0 }), step_part2(82, "L30"));
    try std.testing.expectEqual(@as(struct { i32, i32 }, struct { i32, i32 }{ 0, 1 }), step_part2(52, "R48"));
    try std.testing.expectEqual(@as(struct { i32, i32 }, struct { i32, i32 }{ 95, 0 }), step_part2(0, "L5"));
    try std.testing.expectEqual(@as(struct { i32, i32 }, struct { i32, i32 }{ 55, 1 }), step_part2(95, "R60"));
    try std.testing.expectEqual(@as(struct { i32, i32 }, struct { i32, i32 }{ 0, 1 }), step_part2(55, "L55"));
    try std.testing.expectEqual(@as(struct { i32, i32 }, struct { i32, i32 }{ 99, 0 }), step_part2(0, "L1"));
    try std.testing.expectEqual(@as(struct { i32, i32 }, struct { i32, i32 }{ 0, 1 }), step_part2(99, "L99"));
    try std.testing.expectEqual(@as(struct { i32, i32 }, struct { i32, i32 }{ 14, 0 }), step_part2(0, "R14"));
    try std.testing.expectEqual(@as(struct { i32, i32 }, struct { i32, i32 }{ 32, 1 }), step_part2(14, "L82"));

    //edge cases
    try std.testing.expectEqual(@as(struct { i32, i32 }, struct { i32, i32 }{ 58, 0 }), step_part2(0, "L42"));
    try std.testing.expectEqual(@as(struct { i32, i32 }, struct { i32, i32 }{ 30, 4 }), step_part2(10, "R420"));
    try std.testing.expectEqual(@as(struct { i32, i32 }, struct { i32, i32 }{ 0, 5 }), step_part2(10, "R490"));
    try std.testing.expectEqual(@as(struct { i32, i32 }, struct { i32, i32 }{ 0, 6 }), step_part2(0, "L600"));
    try std.testing.expectEqual(@as(struct { i32, i32 }, struct { i32, i32 }{ 10, 6 }), step_part2(10, "L600"));
    try std.testing.expectEqual(@as(struct { i32, i32 }, struct { i32, i32 }{ 87, 4 }), step_part2(20, "L333"));
}
