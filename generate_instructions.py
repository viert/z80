#!/usr/bin/env python

#
#  Go port of Gabriel Gambetta's brilliant libz80 library.
#  Original C version can be found at https://github.com/ggambetta/libz80
#

import os
import re
import argparse

EXPAND_RE = re.compile(r'\[([^\]]+)\]')
NUM_RANGE_RE = re.compile(r'^(\d)\-(\d)$')
COMMENT_RE = re.compile(r'^\s*#')
WHITESPACESTART_RE = re.compile(r'^\s+')

def parse_pattern(group):
    nr = NUM_RANGE_RE.match(group)
    if nr:
        first = int(nr.groups()[0])
        second = int(nr.groups()[1])
        return [str(x) for x in range(first,second+1)]

    spl = group.split(',')
    if len(spl) == 1:
        return list(group)
    else:
        return spl


def expand(line, vars=[]):
    match = EXPAND_RE.search(line)
    if match is None:
        yield (line, vars)
    else:
        variants = parse_pattern(match.groups()[0])
        for variant in variants:
            new_line = EXPAND_RE.sub(variant, line, 1)
            for out in expand(new_line, vars + [variant]):
                yield out

def generate_funcname(func_name):
    func_name = func_name.split()
    func_name[0] = func_name[0].lower()
    func_name = '_'.join(func_name)
    func_name = func_name.replace(',','_').replace('+','_').replace("'",'_')
    func_name = re.sub('\(([^\)]+)\)', 'off_\\1', func_name)
    return func_name

def render(declaration, implementation):
    result = ""
    if not declaration is None:
        for variant, vars in expand(declaration):
            func_name = generate_funcname(variant)

            result += "func (c *Context) %s() {\n" % func_name
            for line in implementation:
                for i, var in enumerate(vars):
                    line = line.replace("$%d"%(i+1), var)
                result += "%s\n" % line
            result += "}\n\n"
    return result

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--specfile', '-s', dest='specfile', type=str, required=True, help='specifications metafile')
    parser.add_argument('--output', '-o', dest='output', type=str, default=None, help='output file. stdout is used when omitted')
    args = parser.parse_args()

    f = open(args.specfile)

    result = "package z80\n\n\n"
    declaration = None
    implementation = []

    for line in f:
        line = line.rstrip()
        if line == "" or COMMENT_RE.match(line): continue
        if WHITESPACESTART_RE.match(line):
            implementation.append(line)
        else:
            result += render(declaration, implementation)

            declaration = line
            implementation = []

    result += render(declaration, implementation)

    if args.output:
        f = open(args.output, "w")
        f.write(result)
        f.close()
    else:
        print result
