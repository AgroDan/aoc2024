# Day 7: Bridge Repair

A challenge appears with the concept of many permutations of a result. I took it upon myself to create a recursive function that simply generated an amount of permutations based on the type of equations that the challenge called for. In part one, it would either be multiply (`*`), or add (`+`). But for part two, we had to add another permutation altogether: concatenation.

So not only did I need to find a way to concatenate two numbers (That is, if `12 || 345` appeared, it would create a new number: `12345`), but to also generate enough permutations for three possible operations. The first bit was easier than I expected, albeit a little cheat-y. I just converted the numbers to a string, combined them, then converted them back to a number. That's the way I've found to be the least complex, because apparently concatenating two numbers _mathematically_ tends to get pretty silly.

Anyways, once I got the logic for that it was just a matter of letting it crunch away. Eventually it gave me what I needed after about `6` seconds of processing time. 