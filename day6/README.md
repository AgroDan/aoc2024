# Day 6: Guard Gallivant

Day 6 was my downfall. I struggled so hard to get a working version of this. Part one would complete fine, but all my algorithms that I came up with for finding out part 2 were messy. I kept getting incorrect guesses so many times after tweaking slight variations that I gave up and re-wrote the whole thing from scratch again **twice**. Eventually after trying to put obstacles only on places where the guard was going to walk, I ultimately decided to attempt to put obstacles on _every_ tile and just re-run the guard algorithm path and see if we hit an infinite loop. It wasn't efficient, nor was it elegant, but it finally gave me the right answer. I am still ashamed of this, but I did learn one thing: So early into the Advent of Code challenges, sometimes you shouldn't worry about efficiency and just let it run.

This challenge was the first one to set me off the schedule. This took me a couple of days to solve if only due to prior engagements, and I would never catch up. Oh well. Temper your expectations for things that take up a large portion of your time around the holidays.

## Utils

Incidentally, this is the first day I created the `utils/` directory containing objects I would use throughout the challenges.