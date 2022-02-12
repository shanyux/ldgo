#! /usr/bin/env python
# -*- coding: utf-8 -*-
#
# Copyright (C) distroy
#


import sys

import exec


class Complexity(object):
    def __init__(self, file: str = '', complexity: int = 0, pos: int = 0, end: int = 0):
        self.file: str = file
        self.pos: int = pos  # start with 1
        self.end: int = end
        self.complexity: int = complexity

    def __str__(self) -> str:
        return '%d %s:%d,%d' % (self.complexity, self.file, self.pos, self.end)


def install_gocognit():
    cmd = ['go', 'install', 'github.com/distroy/gocognit/cmd/gocognit@v1.0.5.2']
    status, _ = exec.exec(cmd)
    if status != 0:
        sys.stderr.write('intall gocognit fail. cmd:%s\n' % ''.join(cmd))
        sys.exit(status)


def get_cogntive(path: str, exclude: list[str] = [], exclude_dirs: list[str] = []) -> list[Complexity]:
    cmd = ['gocognit', path]
    status, output = exec.exec(cmd)
    if status != 0:
        sys.stderr.write('exec gocognit fail. cmd:%s\n' % ''.join(cmd))
        sys.exit(status)

    lines: list[str] = output.split('\n')
    # print(lines)
    buffer: list[Complexity] = []

    for line in lines:
        items = line.split(' ')
        if len(items) < 4:
            sys.stderr.write('invalid complexity line. line: %s\n' % line)
            sys.exit(-1)

        complexity = int(items[0])

        items = items[3].split(':')
        file = items[0]
        items = items[1].split(',')
        pos = int(items[0])
        end = int(items[1])

        o = Complexity(file, complexity=complexity, pos=pos, end=end)
        print(line, str(o))
        buffer.append(o)

    return buffer
