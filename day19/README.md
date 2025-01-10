# Day 19: Linen Layout

I feel like a complete and total loser. I have spent an embarrassing amount of time troubleshooting why part one would _never_ equal the right answer. AOC continued to tell me that my number was too high. Too high! As you can see, I wrote and re-wrote every possible means I could to attempt to find the "proper" answer here. Every single time the sample input would come through without any issues. It would always print out the right answer. But when I ran it against the challenge input...always wrong.

I have such a headache from banging my head against my desk. Only after I increased verbosity to an almost obnoxious degree did I see the issue:

```terminal
Found 0 possible combinations for wwrrrbbwugrurubuwwwbwwuggwubuuwggrwubbgbrw
Found 1322464992653 possible combinations for bgwurbwgrbgrbubrbuuugrwugurwgbrugbbggbwubguuurggwrg
Found 64600933497 possible combinations for wgbuggurgurguruwugbuwurwuugbwgwbbrubwgburuwurgwb
Found 1292897016866 possible combinations for rbgbuuwbgwwurrbgbuubbggruuwggbrbuwrrruururrwbwbrwuuwgb
Found 0 possible combinations for rbbburrbuggbwrrgrrrgbbwrwbwbruwbbrguwbwrburbbrw
Found 1349330002009 possible combinations for gwwgggugurubguuwurwugbbrrurgrwbwgbgubbrwrugbbbwgrgwbbgwgww
Found 27569910465 possible combinations for gwbbbuuwgbbwrbwurwrgbgbbbwugbwwrwuwwrwwubuwbuggwbburg
Found 29669252603886 possible combinations for bwrwrgrurwbrrbggbgggburggggwwburuggbuububugrrggwwuuurbbbu
Found 0 possible combinations for brgbgrguuwrrwuwrubugbbggwgburuwubuggrrwubrrgubwgrbrgbww
Found 0 possible combinations for wuguruwbugwbubbgggbwwgbbruwwggrwwwrgubbrwwguwrrrbguburbww
Found 156053877382 possible combinations for wwgugwuugbuugrurwwuwbuwwuuuuwwbruwbuwwbrbbbuuwwwbbgrburgwr
Found 1 possible combinations for
Total towel combinations possible for part 1: 227
Total time elapsed: 54.1243ms
```

Do you see it? I'll give you a minute.

...How about the last line before it displays the total? "Found 1 possible combinations for " ...and nothing. Why?

Because there was an additional newline at the bottom of the input file. And because of that, the logic in my code stated that if the length of the challenge design _is zero_, then technically YES this potential towel is possible. It's easily possible because you don't need any towels to display absolutely nothing!

I removed the newline and it ran...and gave me 226 total towel combinations. The correct answer.

I also am happy to say that I guessed part 2's challenge correctly. It wanted the total amount of combinations. I already wrote the function for part 1! I simply modified the code a little bit to accommodate the total amount of combinations and it ripped through it in about 53ms. I kept all my previous attempts at determining the values in as terrible functions for all of you to witness my ultimate shame.

## Takeways

I added memoization here to increase the speed of the code a bit. Took a while to truly understand how to do that but I figured it out. Also I think this is the first challenge since day 1 in which I _didn't_ use a struct.