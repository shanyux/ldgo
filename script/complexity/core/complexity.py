#! /usr/bin/env python
# -*- coding: utf-8 -*-
#
# Copyright (C) distroy
#


import os

from . import exec
from . import pattern


class Complexity(object):
    def __init__(self, file: str = '', complexity: int = 0, package: str = '',
                 function: str = '', pos: int = 0, end: int = 0):

        self.complexity: int = complexity
        self.package: str = package
        self.function: str = function
        self.file: str = file
        self.pos: int = pos  # start with 1
        self.end: int = end

    def __str__(self) -> str:
        file_text = '%s:%d,%d' % (self.file, self.pos, self.end)
        return '%d %s %s %s' % (self.complexity, self.package, self.function, file_text)


def install_gocognit():
    cmd = ['type', 'gocognit']
    status, _ = exec.exec(cmd)
    if status == 0:
        return

    cmd = ['go', 'install', 'github.com/distroy/gocognit/cmd/gocognit']
    status, _ = exec.exec(cmd)
    if status != 0:
        raise Exception('intall gocognit fail. status:%d, cmd:%s\n' % (status, ' '.join(cmd)))


def get_cogntive(path: str, threshold: int = 15, excludes: list[str] = [], includes: list[str] = []) -> list[Complexity]:
    install_gocognit()

    cmd = ['gocognit', path]
    status, output = exec.exec(cmd)
    if status != 0:
        raise Exception('exec gocognit fail. status:%d, cmd:%s\n' % (status, ' '.join(cmd)))

    if not output:
        return []

    patterns = pattern.Pattern(excludes=excludes, includes=includes)

    lines: list[str] = output.split('\n')
    # print(lines)
    buffer: list[Complexity] = []

    for line in lines:
        items = line.split(' ')
        if len(items) < 4:
            raise Exception('invalid complexity line. line: %s\n' % line)

        complexity = int(items[0])
        package = items[1]
        function = items[2]
        if complexity <= threshold:
            continue

        # print(line)
        items = items[3].split(':')
        file = items[0]
        file = os.path.relpath(file, path)

        items = items[1].split(',')
        pos = int(items[0])
        end = int(items[1])

        if not patterns.check_file(file):
            continue

        o = Complexity(file, complexity=complexity, package=package,
                       function=function, pos=pos, end=end)
        # print([line, str(o)])
        buffer.append(o)

    return buffer
