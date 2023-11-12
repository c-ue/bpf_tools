#!/bin/bash

# find dependencies header files
# clang -MD -MF mybpf.d -target bpf -I /lib/linux/ -c mybpf.c

# compile
clang -target bpf -Wall -O2 -emit-llvm -g -S -Iinclude -c mybpf.c -o mybpf.bc
llc -march=bpf -mcpu=probe -filetype=obj -o mybpf.o mybpf.bc