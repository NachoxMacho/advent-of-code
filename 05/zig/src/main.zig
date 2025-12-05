const std = @import("std");

const Range = struct {
    low: i64,
    high: i64,
};

pub fn main() !void {
    var file = try std.fs.cwd().openFile("input-real.txt", .{});
    defer file.close();

    var buf: [1024]u8 = undefined;
    var file_reader = file.reader(&buf);

    // var result1: i32 = 0;
    // var result2: i64 = 0;

    // var ranges: []Range = &[_]Range{};
    var ranges = std.ArrayList(Range).initCapacity(std.heap.page_allocator, 16) catch unreachable;
    defer ranges.deinit(std.heap.page_allocator);

    while (try file_reader.interface.takeDelimiter('\n')) |line| {
        if (std.mem.eql(u8, line, "")) {
            break;
        }
        var iter = std.mem.splitScalar(u8, line, '-');
        const low = try std.fmt.parseInt(i64, iter.next().?, 10);
        const high = try std.fmt.parseInt(i64, iter.next().?, 10);
        try ranges.append(std.heap.page_allocator, Range{ .low = low, .high = high });
    }
    const rangeArray = ranges.items;

    std.mem.sort(Range, rangeArray, {}, cmpRange);

    const start = std.time.microTimestamp();
    const result2 = step_part2(rangeArray);
    const end = std.time.microTimestamp();

    std.debug.print("result2 {}, diff {}\n", .{result2, (end - start)});

    // std.debug.print("result1 {}\n", .{result1});
    // std.debug.print("result2 {}\n", .{result2});
}

pub fn cmpRange(_: void, a: Range, b: Range) bool {
    return (a.low < b.low);
}

pub fn step_part2(input: []Range) i64 {
    // const copy: []Range = &[_]Range{};
    // std.mem.copyForwards(Range, copy, input);
    var res: i64 = 0;
    var comparisons: i32 = 0;

    for (0..input.len) |index| {
        if (input[index].low == 0 and input[index].high == 0) {
            continue;
        }
        inner: for ((index+1)..input.len) |index2| {
            if (input[index2].low == 0 and input[index2].high == 0) {
                continue;
            }
            if (input[index].high < input[index2].low) {
                continue;
            }
            if (input[index].low > input[index2].high) {
                break :inner;
            }
            comparisons += 1;
            if (input[index].high >= input[index2].high and input[index].low <= input[index2].low) {
                input[index2].low = 0;
                input[index2].high = 0;
                continue;
            }
            if (input[index].low <= input[index2].low) {
                input[index2].low = input[index].high + 1;
                continue;
            }
            if (input[index].high >= input[index2].high) {
                input[index2].high = input[index].low - 1;
                continue;
            }
        }
    }

    for (input) |r| {
        if (r.low == 0 and r.high == 0) {
            continue;
        }
        res += r.high - r.low + 1;
    }
    std.debug.print("comparisons {}, size {}\n", .{comparisons, input.len});

    return res;
}

test "simple test part2" {
    // edge cases
    // try std.testing.expectEqual(92, try step_part2("911112", 2));

    // sample
}
