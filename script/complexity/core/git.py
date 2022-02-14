#! /usr/bin/env python
# -*- coding: utf-8 -*-
#
# Copyright (C) distroy
#


import traceback

from . import exec


class Diff(object):
    def __init__(self, file: str = '', pos: int = 0, end: int = 0):
        self.file = file
        self.pos = pos  # start with 1
        self.end = end

    def __str__(self) -> str:
        return '%s:%d,%d' % (self.file, self.pos, self.end)


def get_diffs(branch: str) -> list[Diff]:
    cmd = ['git', 'diff', '--unified=0', branch]
    status, output = exec.exec(cmd)
    if status != 0:
        raise Exception('get git diff fail. status:%d, cmd:%s\n' % (status, ' '.join(cmd)))

    # print(output)
    lines: list[str] = output.split('\n')
    # print(lines)

    buffer: list[Diff] = []
    i, l = 0, len(lines)
    file = ''
    while i < l:
        line = lines[i]
        i += 1

        prefix = '+++ b/'
        if line.startswith(prefix):
            file = line[len(prefix):]

        if line.startswith('@@'):
            items = line.split(' ')
            pos = items[2].split(',')
            if len(pos) == 1:
                diff = Diff(file, int(pos[0]), int(pos[0]))
            else:
                diff = Diff(file, int(pos[0]), int(pos[0]) + int(pos[1]) - 1)
            # print(line, pos, str(diff))
            buffer.append(diff)

    return buffer


def repo_root() -> str:
    cmd = ['git', 'rev-parse', '--show-toplevel']
    status, output = exec.exec(cmd)
    if status != 0:
        raise Exception('get repo root fail. status:%d, cmd:%s\n' % (status, ' '.join(cmd)))
    return output


def get_branch() -> str:
    cmd = ['git', 'rev-parse', '--verify', 'HEAD']
    status, _ = exec.exec(cmd)
    if status == 0:
        return 'HEAD'

    cmd = ['git', 'hash-object', '-t', 'tree', '/dev/null']
    status, output = exec.exec(cmd)
    if status != 0:
        raise Exception('get hash-object fail. status:%d, cmd:%s\n' % (status, ' '.join(cmd)))
    return output
