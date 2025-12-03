const std = @import("std");

const nums = "0123456789";
const size = 12;

pub fn main() !void {
    var file = try std.fs.cwd().openFile("input-real.txt", .{});
    defer file.close();

    var buf: [1024]u8 = undefined;
    var file_reader = file.reader(&buf);

    var result1: i32 = 0;
    var result2: i64 = 0;

    while (try file_reader.interface.takeDelimiter('\n')) |line| {
        result1 = result1 + try step_part1(line);
        result2 = result2 + try step_part2(line, size);
    }

    std.debug.print("result1 {}\n", .{result1});
    std.debug.print("result2 {}\n", .{result2});
}

pub fn step_part1(input: []const u8) !i32 {
    var i: usize = 9;
    var j: usize = 9;

    while (i > 0) : (i -= 1) {
        const char = nums[i];
        const index = std.mem.indexOfScalar(u8, input, char);
        if (index == null) {
            continue;
        }
        if (index == input.len - 1) {
            continue;
        }
        while (j > 0) : (j -= 1) {
            const char2 = nums[j];
            const index2 = std.mem.indexOfScalarPos(u8, input, index.? + 1, char2);
            if (index2 == null) {
                continue;
            }
            const s = [_]u8{ input[index.?], input[index2.?] };
            return try std.fmt.parseInt(u8, &s, 10);
        }
    }
    return 0;
}

test "simple test part1" {
    // sample
    try std.testing.expectEqual(98, try step_part1("987654321111111"));
    try std.testing.expectEqual(89, try step_part1("811111111111119"));
    try std.testing.expectEqual(78, try step_part1("234234234234278"));
    try std.testing.expectEqual(92, try step_part1("818181911112111"));
}

pub fn step_part2(input: []const u8, digits: i32) !i64 {
    if (input.len < digits) {
        return 0;
    }
    if (digits == 0) {
        return 0;
    }
    var searchCharIndex: usize = 9;

    while (searchCharIndex > 0) : (searchCharIndex -= 1) {
        const index = std.mem.indexOfScalar(u8, input, nums[searchCharIndex]);
        if (index == null) {
            continue;
        }
        if (index.? >= input.len) {
            continue;
        }
        if (digits == 1) {
            return @as(i64, @intCast(searchCharIndex));
        }
        const o = try step_part2(input[(index.? + 1)..], digits - 1);
        if (o == 0) {
            continue;
        }

        return @as(i64, @intCast(searchCharIndex)) * std.math.pow(i64, 10, digits-1) + o;
    }
    return 0;
}

test "simple test part2" {
    // edge cases
    try std.testing.expectEqual(92, try step_part2("911112", 2));
    try std.testing.expectEqual(9, try step_part2("911112", 1));

    // sample
    try std.testing.expectEqual(987654321111, try step_part2("987654321111111", 12));
    try std.testing.expectEqual(811111111119, try step_part2("811111111111119", 12));
    try std.testing.expectEqual(434234234278, try step_part2("234234234234278", 12));
    try std.testing.expectEqual(888911112111, try step_part2("818181911112111", 12));
}
