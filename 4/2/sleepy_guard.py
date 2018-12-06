#!/usr/bin/env python
# -*- coding: utf-8 -*-

import argparse
import datetime
import enum
import re
from collections import defaultdict


def main():
    args = parse_args()

    entries = []
    for line in args.input:
        match = re.match('\[(\d{4}-\d{2}-\d{2} \d{2}:\d{2})\] (.*)\\n', line)
        timestamp = datetime.datetime.strptime(match.group(1),
                                               '%Y-%m-%d %H:%M')
        entries.append((timestamp, match.group(2)))
    entries.sort()

    naps = defaultdict(list)
    current_guard, fell_asleep = None, None
    for timestamp, text  in entries:
        if text == 'falls asleep':
            fell_asleep = timestamp
        elif text == 'wakes up':
            naps[current_guard].append((fell_asleep, timestamp))
        else:
            match = re.match('Guard #(\d+)', text)
            current_guard = int(match.group(1))
            fell_asleep = None

    sleepiest_minutes = []
    for guard in naps:
        minutes = defaultdict(int)
        for minute in range(60):
            for begin, end in naps[guard]:
                if minute >= begin.minute and minute < end.minute:
                    minutes[minute] += 1
        count, sleepiest_minute = sorted((v, k) for k, v in minutes.items())[-1]
        sleepiest_minutes.append((count, sleepiest_minute, guard))
    sleepiest_minutes.sort()

    count, sleepiest_minute, sleepiest_guard = sleepiest_minutes[-1]
    print(sleepiest_minute * sleepiest_guard)

    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
