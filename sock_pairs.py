#!/usr/bin/python3
# https://www.hackerrank.com/challenges/sock-merchant/problem


def sockMerchant(n, ar):
    pairs = 0
    colors_seen = dict()
    for sock_color_number in ar:
        if colors_seen.get(sock_color_number) is None:
            colors_seen.update({sock_color_number: 1})
        else:
            colors_seen[sock_color_number] += 1
    for k, v in colors_seen.items():
        pairs += int(v / 2)
    print(pairs)
    return pairs


sockMerchant(9, [10, 20, 20, 10, 10, 30, 50, 10, 20])
sockMerchant(7, [1, 2, 1, 2, 1, 3, 2])
