# Day 11

My method for this was to create a linked list believe it or not...I felt that since the author said that order actually mattered that a linked list would preserve that the best, and I could insert things into the list without really doing some weird slice/array merge/unmerge nonsense so there you go.

...until part 2 happened. I don't think I've ever seen a segfault happen so quickly after I told it to up the game to 75 iterations. Clearly a linked list was too crazy for my meager VM and I needed a better way to manage memory before it expanded so much it blew up in catastrophic fashion. I did some research and found a neat trick: memoization.

Instead of representing an object of an array/linked list/whatever for every single stone in this challenge, I simply created a finite set. Essentially a map of every single stone and kept a record of how many stones I had of that type. Since, despite what the author said, order didn't actually matter here I could get away with just counting the different types. After I figured out the proper way of doing it I was able to complete both part 1 AND part 2 (mind you, part 1 used linked lists) in under 65 milliseconds.

Happy christmas eve everyone! This challenge certainly tells me exactly how much spare time I _don't_ have compared to all the other participants.