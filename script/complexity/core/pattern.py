#! /usr/bin/env python
# -*- coding: utf-8 -*-
#
# Copyright (C) distroy
#


import os
import sys
import re


class Pattern(object):
    def __init__(self, excludes: list[str] = [], includes: list[str] = []):
        self.__includes: list[re.Pattern] = []
        self.__excludes: list[re.Pattern] = []

        self.add_includes(includes)
        self.add_excludes(excludes)
        return

    def add_excludes(self, patterns: list[str]):
        for pattern in patterns:
            self.__excludes.append(self.__parse_pattern(pattern))
        return

    def add_includes(self, patterns: list[str]):
        for pattern in patterns:
            self.__includes.append(self.__parse_pattern(pattern))
        return

    def __parse_pattern(self, pattern: str) -> re.Pattern:
        return re.compile(pattern)

    def check_file(self, path: str) -> bool:
        for rexpr in self.__includes:
            res = rexpr.search(path)
            if res:
                # print('=====', 'include', path, str(rexpr))
                return True

        for rexpr in self.__excludes:
            res = rexpr.search(path)
            if res:
                # print('*****', 'exclude', path, str(rexpr))
                return False

        return True
