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

    details, totals = defaultdict(list), defaultdict(int)
    current_guard, fell_asleep = None, None
    for timestamp, text  in entries:
        if text == 'falls asleep':
            fell_asleep = timestamp
        elif text == 'wakes up':
            details[current_guard].append((fell_asleep, timestamp))
            totals[current_guard] += (timestamp - fell_asleep).total_seconds() / 60
        else:
            match = re.match('Guard #(\d+)', text)
            current_guard = int(match.group(1))
            fell_asleep = None
    sleepiest_guard = sorted((v, k) for k, v in totals.items())[-1][1]

    minutes = defaultdict(int)
    for minute in range(60):
        for begin, end in details[sleepiest_guard]:
            if minute >= begin.minute and minute < end.minute:
                minutes[minute] += 1
    sleepiest_minute = sorted((v, k) for k, v in minutes.items())[-1][1]

    print(sleepiest_guard * sleepiest_minute)
    return 0


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('input', type=argparse.FileType())
    return parser.parse_args()


if __name__ == '__main__':
    raise SystemExit(main())
