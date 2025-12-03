const std = @import("std");

pub fn main() !void {
    var file = try std.fs.cwd().openFile("input-real.txt", .{});
    defer file.close();

    var buf: [1024]u8 = undefined;
    var file_reader = file.reader(&buf);

    var result1: f64 = 0;
    var result2: f64 = 0;

    while (try file_reader.interface.takeDelimiter(',')) |line| {
        // This is needed to make sure we don't
        const trimLine = std.mem.trimRight(u8, line, "\n");
        result1 = result1 + try step_part1(trimLine);
        result2 = result2 + try step_part2(trimLine);
    }

    std.debug.print("result1 {}\n", .{result1});
    std.debug.print("result2 {}\n", .{result2});
}

pub fn step_part1(input: []const u8) !f64 {
    var split = std.mem.splitAny(u8, input, "-");
    const startString = split.next();
    const endString = split.next();
    const start = @as(f64, @floatFromInt(try std.fmt.parseInt(u64, startString.?, 10)));
    const end = @as(f64, @floatFromInt(try std.fmt.parseInt(u64, endString.?, 10)));
    var i = start;
    var result: f64 = 0;

    std.debug.print("start {}, end {}\n", .{ start, end });

    while (i <= end and i >= start) {
        // length calculation
        var log = @floor(std.math.log10(i)) + 1;
        if (@mod(log, 2) == 1) {
            i = std.math.pow(f64, 10, log);
            continue;
        }

        if (log > 1) {
            log = @divFloor(log, 2);
        }

        const logSplit = std.math.pow(f64, 10, log);

        const second = @mod(i, logSplit);
        const first = @divFloor(i, logSplit);

        if (first == second) {
            result += i;
        }

        // find next possible match
        const next = first * logSplit + first;
        if (next > i and next > start) {
            i = next;
        } else {
            i = (first + 1) * logSplit + first + 1;
        }
    }
    return result;
}

test "simple test part1" {
    // sample
    try std.testing.expectEqual(33, try step_part1("11-22"));
    try std.testing.expectEqual(99, try step_part1("95-115"));
    try std.testing.expectEqual(1010, try step_part1("998-1012"));
    try std.testing.expectEqual(1188511885, try step_part1("1188511880-1188511890"));
    try std.testing.expectEqual(222222, try step_part1("222220-222224"));
    try std.testing.expectEqual(0, try step_part1("1698522-1698528"));
    try std.testing.expectEqual(446446, try step_part1("446443-446449"));
    try std.testing.expectEqual(38593859, try step_part1("38593856-38593862"));
    try std.testing.expectEqual(0, try step_part1("2121212118-2121212124"));
    try std.testing.expectEqual(0, try step_part1("824824821-824824827"));
    try std.testing.expectEqual(0, try step_part1("565653-565659"));
}

pub fn step_part2(input: []const u8) !f64 {
    var split = std.mem.splitAny(u8, input, "-");
    const startString = split.next();
    const endString = split.next();
    const start = @as(f64, @floatFromInt(try std.fmt.parseInt(u64, startString.?, 10)));
    const end = @as(f64, @floatFromInt(try std.fmt.parseInt(u64, endString.?, 10)));
    var i = start;
    var result: f64 = 0;

    while (i <= end and i >= start) : (i += 1) {
        const s = std.fmt.allocPrint(std.heap.page_allocator, "{}", .{i}) catch unreachable;
        if (try test_part2(s)) {
            result += i;
        }
    }
    return result;
}

test "simple test part2" {
    try std.testing.expectEqual(33, try step_part2("11-22"));
    try std.testing.expectEqual(210, try step_part2("95-115"));
    try std.testing.expectEqual(2109, try step_part1("998-1012"));
    try std.testing.expectEqual(1188511885, try step_part1("1188511880-1188511890"));
    try std.testing.expectEqual(222222, try step_part1("222220-222224"));
    try std.testing.expectEqual(0, try step_part1("1698522-1698528"));
    try std.testing.expectEqual(446446, try step_part1("446443-446449"));
    try std.testing.expectEqual(38593859, try step_part1("38593856-38593862"));
    try std.testing.expectEqual(0, try step_part1("2121212118-2121212124"));
    try std.testing.expectEqual(0, try step_part1("824824821-824824827"));
    try std.testing.expectEqual(0, try step_part1("565653-565659"));
}

pub fn test_part2(input: []const u8) !bool {
    var i = @divFloor(input.len, 2);

    while (i > 0) : (i -= 1) {
        var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
        defer arena.deinit();
        const allocator = arena.allocator();

        const output = try std.mem.replaceOwned(u8, allocator, input, input[0..i], "");
        defer allocator.free(output);
        if (output.len == 0) {
            return true;
        }
    }
    return false;
}

test "simple test filer2" {
    try std.testing.expectEqual(true, try test_part2("11"));
    try std.testing.expectEqual(false, try test_part2("12"));
}


