#!/usr/bin/env bash

# cat input1_1.prod | paste -sd+ | sed 's/++/\n/g' | bc | sort -n | tail -1 


cat input1_1.prod | paste -sd+ | sed 's/++/\n/g' | bc | sort -n | tail -3 | paste -sd+ | bc

