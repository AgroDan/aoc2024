# Day 23: LAN Party

Ahhhh, this is one of those challenges where like it or not, you're going to learn something about computer science. For me, it was learning about the [Clique Problem](https://en.wikipedia.org/wiki/Clique_problem). This is one of those things that you never really sit down and think about in real life because the concept of adjacent and similar vertices within a set tend to just present themselves in some sort of tangible form that your brain can just deduce just by observing. But when you break it all down into granular operations, it quickly becomes one of those P vs NP questions, where people far smarter than I could ever hope to be have found ways of finding the solutions to these questions, such as my discovery of the [Bron-Kerbosch Algorithm](https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm), which is ultimately what I used to find the solution here.

The TL;DR of the challenge is you're given a gigantic list of two-letter computer names, a hyphen, and another two-letter computer name, so `aa-bb`, and this signifies that computer `aa` is connected to computer `bb` and vice versa. The problem comes in finding the largest lan party among this list of connections.

First, part one was straightforward, albeit a bit daunting. What I did was create a two-dimensional _map_, which came in the form of `map[string]map[string]bool`, which pretty tangibly laid out the framework for me to determine if `X` is connected to `Y`. In this case, I used something like this:

```go

party := make(map[string]map[string]bool)

// generate the gigantic two-dimensional map using my parser function...

if party["aa"]["bb"] {
    fmt.Printf("aa and bb are connected!\n")
}
```

So I basically just followed the members they were connected to and looked three layers deep, then looked for members of this group with a name starting with the letter `t`. It was simply enough.

But now finding the largest lan party was an exercise in futility, as I tried to work through the logic of doing this myself, only to find out after searching that there existed an algorithm that found exactly that. But my first attempt had me using a bitmask to make it easier to detect membership of cliques, a technique I saw through my searches. This technique worked pretty well...for the example input. Once I ran it for the challenge input, it returned a zero-length "maximum lan party" slice, which was just freakin' impossible. But I'm mentioning it because of _why_ it returned no results, and I think it was a pretty cool finding. First, this is the function in question:

```go
// Generate all subsets of nodes
func generateSubsets(nodes []string) [][]string {
	subsets := [][]string{}
	n := len(nodes)
	for i := 0; i < (1 << n); i++ {
		subset := []string{}
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				subset = append(subset, nodes[j])
			}
		}
		subsets = append(subsets, subset)
	}
	return subsets
}
```

The above basically generated a two-dimensional array of strings, all containing a list of all the nodes in the challenge and every permutation consisting of all nodes _except_ the node in question. To accomplish this, it uses a bitmask to determine if the node we're iterating around should be added to this particular permutation. Well this is fine on smaller datasets, like the example data which only had about 32 different connections. Well with the challenge dataset, there were `3380` different connections, and the amount of permutations would have been `520`, which means the very first `for` loop up there would have set `i` to `0`, then checked to see if `i` was less than `(1 << n)`, which is actually `3432398830065304857490950399540696608634717650071652704697231729592771591698828026061279820330727277488648155695740429018560993999858321906287014145557528576`! So as you can see, this is much larger than a standard integer datatype can handle, so it returned zero (I'm assuming it was truncated, because left-bit-shifting pads it with zeroes to the right side of the number), so it just...never looped. That stumped me for quite a bit.

Once I discovered a better algorithm for handling this kind of data with an unreasonable amount of permutations to loop through, I simply just used that. Lo and behold, it...just worked.

So I guess thank you Coenraad Bron and Joep Kerbosch for doing all the work for me for the Advent of Code Day 23. Fun fact, Coenraad Bron worked with [Edsger W. Dijkstra](https://en.wikipedia.org/wiki/Edsger_W._Dijkstra), the creator of [Dijkstra's algorithm](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm), something I've definitely used in past advent of code challenges. Small world!

By the way, I'm leaving most of my failed code in this challenge as a warning to others who may think I'm actually a good developer or something. Trust me, it's mostly blind luck and standing on the shoulders of giants.