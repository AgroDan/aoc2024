# Day 17: Chronospatial Computer

A reverse engineering challenge! Definitely one of those "think outside the box" ones because I really had to scratch my head quite a bit to come up with this solution. Eventually I had to work with some fuzzy logic but it got the job done I suppose.

Part one was extremely straightforward. After reviewing other peoples' code I noticed that the weirdo `x / 2^y` is actually just a bit shift right (`>>`) of `3`, something I don't think I ever would have come up with myself. At first I did some ridiculous Go casting logic of something like `int(float64(operand) / math.Pow(2, float64(B)))`, but since that was just painful to look at (as well as painful to compute in hindsight), but using bitwise math not only simplified it quite a bit, but also explained better how to reverse engineer it for part 2. That said though, Part 1 was easy. Basically just copy down the formula and execute it. Don't have to worry too much about computational expense, since I'm literally doing what it told me to do.

Part two on the other hand was much different. Find out the first number it would take to initialize to register A to print out an exact copy of itself. This is evidentally called a [Quine](https://en.wikipedia.org/wiki/Quine_(computing)), a computer program that takes no input and produces a copy of its own source code as its only output. And in advent-of-code style, you can imagine that in order to determine what number to initialize the A register to, brute forcing would be virtually impossible. The number would be so incredibly high that even a processor doing nothing but incrementing itself would take an unreasonable amount of time to even count to it. So how do you determine this?

Well, you have to reverse engineer it. By looking at the code for my particular challenge, I'm able to determine _exactly_ what the program will do given each runtime execution. And with my challenge being `Program: 2,4,1,1,7,5,1,5,4,3,0,3,5,5,3,0`, I can break it down like so:

```
b = a % 8     // bst  (2,4)
b = b ^ 1     // bxl  (1,1)
c = a >> b    // cdv  (7,5)
b = b ^ 5     // bxl  (1,5)
b = b ^ c     // bxc  (4,3)
a = a >> 3    // adv  (0,3)
out += b % 8  // out  (5,5)
loop          // jnz  (3,0)
```

So with an initialized value in register A, I know that running through the program it will eventually output a number between `0` and `7` for each item in the output. So I can just run the program from `0` to `10` and see what it outputs. For me, I noticed that if I initialize the value of A to `4`, the resulting output is `0`, which is the last number in the program. So I can do the same search by bit-shift-left'ing by 3 and check again. `4 << 3 == 32`, so I can search again starting at `32`, and I find that if I initialize A to the number `37`, the value it prints is `[3, 0]`, which is the last two digits of the program! I repeat this process until eventually I find a number that works for the entirety of the program, and since this is a day of firsts, I can tell you that this is the first time I used a `panic()` to determine if I was successful.

Hey, don't judge me. I see you judging me. Stop it. Not cool.