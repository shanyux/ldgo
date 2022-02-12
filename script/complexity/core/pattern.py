#! /usr/bin/env python
# -*- coding: utf-8 -*-
#
# Copyright (C) distroy
#


import os
import sys
import re


class Pattern(object):
    def __init__(self, file_patterns: list[str] = [], dir_patterns: list[str] = []):
        self.__dir_patterns: list[re.Pattern] = []
        self.__file_patterns: list[re.Pattern] = []

        self.add_dir_patterns(dir_patterns)
        self.add_file_patterns(file_patterns)
        return

    def add_file_patterns(self, patterns: list[str]):
        for pattern in patterns:
            self.__file_patterns.append(self.__parse_pattern(pattern))
        return

    def add_dir_patterns(self, patterns: list[str]):
        for pattern in patterns:
            self.__dir_patterns.append(self.__parse_pattern(pattern))
        return

    def __parse_pattern(self, pattern: str) -> re.Pattern:
        return re.compile(pattern)

    def check_file(self, path: str) -> bool:
        path = os.path.normpath(path)
        dirs = path.split(os.sep)
        file_name = dirs[-1]
        dirs =  dirs[:-1]
        for rexpr in self.__file_patterns:
            res = rexpr.search(file_name)
            if res:
                # print('=====', path, file_name, str(rexpr))
                return True

        for name in dirs:
            for rexpr in self.__dir_patterns:
                res = rexpr.search(name)
                if res:
                    # print('*****', path, name, str(rexpr))
                    return True

        return False
