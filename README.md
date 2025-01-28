# Advent of Code 2024

Once more into the breach, dear friends.

I type this now having just completed the Advent of Code 2024. I'm a little late to the end goal here, but I can finally turn around and look behind me at the long journey it's been. The crowd has long since dissipated (if there even really was a crowd to begin with), and many of the tents have been collapsed and trucked out of here. But here I am at the finish line and I couldn't be happier. The advent of code is a series of challenges that are all more difficult than the last, and if you've never tried it then the author has a really great keynote speech about it that he made which you can see [here](https://www.youtube.com/watch?v=uZ8DcbhojOw), and I strongly advise you watch it because it's really cool and 100% true to form.

I'm a cyber security guy through and through, don't get me wrong. The AOC has absolutely nothing to do with information security, but everything to do with computer science. If you take it, I guarantee you will learn things. The sheer amount of new concepts that I learn every year is staggering, and given that I have a focus in application security, seeing these concepts play out in a hands-on way helps me appreciate the developers that work on things like this every day. As someone who gets his feet wet in this sort of thing all the time, it's hard to find someone who can truly appreciate the elegance of simplicity, the toil of complexity, and the success of code that completes correctly within a reasonable amount of time. Developers, wherever and whoever you are out there, _you have my utmost respect._

Here I went to hone my grasping of [Go](https://go.dev/). I typically work with Python, but Go has always caught my eye. It has an incredible standard library and is one of those languages that just...clicks when you get it. And I think I finally have a good handle on it. Now I think I'm going to work on something worthwhile with it. We'll see about that.

Anyway, this year's AOC has brought with it a few more tricks up my sleeve since the last time I completed it. I have a better set of "prep files" that pull the code from the remote server and set up my working directory. I have built a `utils/` directory full of utilities that I have used to some degree. I'm sure some of them have been already replicated in other libraries, but my sorta-unwritten rule for the AOC challenge is to stick with _only_ the standard library, or at the very least libraries that I created _using_ the standard library. As a result I've learned a bit about [Go Generics](https://go.dev/doc/tutorial/generics) and how I can use it to create general utilities, such as if I want to create a [Queue](https://en.wikipedia.org/wiki/Queue_(abstract_data_type)) or a [Stack](https://en.wikipedia.org/wiki/Stack_(abstract_data_type)), I can set it up so that I can add _whatever I want_ to the queue or stack, be it pointers to custom structs, integers, maps, anything. If I write the generics for these abstract data types, I can just stick whatever I want in there and it won't complain.

I've learned how to create an [A* Search Algorithm](https://en.wikipedia.org/wiki/A*_search_algorithm) which helped me out of quite a few binds this time. It turns out that compared to something like [Breadth-First Search](https://en.wikipedia.org/wiki/Breadth-first_search) or [Depth-First Search](https://en.wikipedia.org/wiki/Depth-first_search), it is incredibly efficient since it offers a weighted calculation on where to go, rather than try all potential points and report back the most efficient method. A* cuts through the cruft and just makes informed decisions from its neighboring starting points to get from point A to point B. It's so efficient in getting there, [I actually used it alongside BFS to solve a challenge](day16/maze/astar.go), using the result of A* as the cutoff limit to speed things along with BFS, meaning if a path took longer than it was to get there with A*, then ignore the results of that path. A* also wound up being the clear winner in pathfinding algorithms for challenges like [day 18: RAM Run](day18/).

The single most difficult challenge this year was easily [Day 24: Crossed Wires](day24/), followed closely by [Day 21: Keypad Conundrum](day21/). The former involving a lot of reading, and a lot of trial and error to even find out what was even broken. The latter though was just extremely difficult to conceptualize, since it involved many many layers upon layers of control and finding the most efficient method between those controls, and of course working on those layers caused a tremendous amount of load on my CPU until I discovered how to [cache my function calls](utils/cache.go) similar to Python's `[lru_cache()](https://docs.python.org/3/library/functools.html)` functions in their standard library. Still staying true to my concept of sticking only to the standard Go libraries, I made my own little function wrapper which would cache the results rather than re-calculate the results of a function, allowing it to just return the results of a function if it's been pre-computed rather than force the function to recalculate itself, saving myself an incredible amount of computation time in solving Day 21.

My favorite challenge was probably [Day 14: Restroom Redoubt](day14/). Really any of the challenges where you have to rely on the output of a _picture_ is cool in my book. And the way it was presented was so ambiguous I almost raged over it. How do you determine if a tree design appears? Well apparently there is a way, even if you aren't given specifics, and I was pretty happy with how I figured it out.

Ultimately, I think the theme of this year's AOC is [Memoization](https://en.wikipedia.org/wiki/Memoization). Finding ways to store results rather than re-compute things. No need to keep recalculating if you can infer the result of a previously-determined state. There were a ton of ways to accomplish one particular result, but most of them involved just repeating calculations. This is a technique I can see myself using quite a bit.

So with that, I hang my christmas hat on the hook and go back to my usual day-to-day. The Advent of Code puts me in a completely different mental state for as long as I'm participating in it, and now I can finally exit the developer frame of mind and move back into the infosec frame. I'm back baby. But what a ride, huh?

Some code highlights I can refer back to:

- [Go Regular Expressions/Regex](day3/)
- [Slice Manipulation](day5/manuals/fixer.go)
- [Runemap Utility](utils/runemap.go)
- [Coordinates Object](utils/coords.go)
- [Queue Object w/ Generics](utils/gQueue.go)
- [Stack Object w/ Generics](utils/gStack.go)
- [Cartesian Product (all permutations)](utils/heap.go)
- [File Parsing](utils/fileparse.go)
- [Instant STDOUT Flush for Debugging](utils/debug.go)
- [Breadcrumbs object for pathfinding](utils/breadcrumbs.go)
- [Caching function calls](utils/cache.go)
- [Linked Lists](day11/stones/stoneset.go)
- [Memoization Example 1](day11/stones/memoization.go)
- [Memoization Example 2](day19/towels/towels.go)
- [Memoization Example 3](day22/pseudorandom/monkeys.go)
- [Counting Complex Sides](day12/gardens/sides.go)
- [Using Algebra to "cheat"](day13/arcade/math.go)
- [Tree-finding Algorithm](day14/robotmap/tree.go)
- [Sweet, sweet recursion](day15/warehouse/widewarehouse.go)
- [A* + BFS = Optimized AF](day16/maze/astar.go)
- [Reverse Engineering a Quine](day17/)
- [Best use of A* by far](day18/)
- [Taxicab Geometry/Manhattan Distance and its radius](day20/)
- [Caching a recursive function](day21/robots/robots.go)
- [Bron Kerbosch Algorithm](day23/lanparty/bronkerbosch.go)
- [Cliques](day23/lanparty/cliques.go)
- [Fixing a logical adder](day24/)
- [Dealing with trailing newlines](day25/locksandkeys/parser.go)