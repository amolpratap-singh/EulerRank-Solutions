#!/bin/python3

import sys

t = int(input().strip())
for a0 in range(t):
    n = int(input().strip())
    sum_natural_sqr = ((n * (n + 1)) // 2).__pow__(2)
    sum_sqr_natural = (n * (n + 1) * (2*n + 1)) // 6
    print(sum_natural_sqr - sum_sqr_natural)