#! /usr/bin/env python
# -*- coding: utf-8 -*-
#
# Copyright (C) distroy
#


import sys
import os
import threading


def exec(args: list[str]) -> tuple[int, str]:
    status, output, _ = __exec(args)
    return status, output


def __exec(args: list[str]) -> tuple[int, str, str]:
    out_read, out_write = os.pipe()
    err_read, err_write = os.pipe()

    pid = os.fork()
    if pid == 0:  # child
        os.dup2(out_write, sys.stdout.fileno())
        os.dup2(err_write, sys.stderr.fileno())
        os.execvp(args[0], args)

    # parent
    def thread_handler(f: int, buffer: list[bytes]):
        size = 1024
        while True:
            byts = os.read(f, size)
            if len(byts) == 0:
                break
            buffer.append(byts)

    os.close(out_write)
    os.close(err_write)

    out_buf: list[bytes] = []
    out_thread = threading.Thread(target=thread_handler, args=(out_read, out_buf))
    out_thread.daemon = True
    out_thread.start()

    err_buf: list[bytes] = []
    err_thread = threading.Thread(target=thread_handler, args=(err_read, err_buf))
    err_thread.daemon = True
    err_thread.start()

    _, status = os.waitpid(pid, 0)
    out_thread.join()
    err_thread.join()

    os.close(out_read)
    os.close(err_read)

    out_text = str(bytes(b'').join(out_buf).decode('UTF-8'))
    if out_text[-1:] == '\n':
        out_text = out_text[:-1]

    err_text = str(bytes(b'').join(err_buf).decode('UTF-8'))
    if err_text[-1:] == '\n':
        err_text = err_text[:-1]

    # print(args, status, type(output), out_text, err_text)
    return status, out_text, err_text

