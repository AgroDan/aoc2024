#!/usr/bin/env fish
#
# This is a fish script to pull the input file of the specific day.
#
# In order to execute, run:
#
# set -x AOC_SESSION "<session key>"
# . ./prep.fish <day>
#
# where "key" is the session key in your logged-in cookie, and "day" is
# the day you want to prep.
#
# Credit where it is due, the concept of this script was inspired heavily by 0xdf's version.

set AOC_DIR "/home/dan/Documents/aoc2024"
set AOC_URL_BASE "https://adventofcode.com/2024/day"

# AOC Session is not set
if test -z "$AOC_SESSION"
    echo "AOC Session is not set!"
    echo "Please enter 'set -x AOC_SESSION <session key>"
    exit 1
end

if test -z "$argv[1]"
    echo "Need the day we are working with, please enter the advent day"
    echo "Example: ./prep.fish 5"
    exit 2
end

set WORKDIR "$AOC_DIR/day$argv[1]"

mkdir -p $WORKDIR
cd $WORKDIR

set AOC_FULL "$AOC_URL_BASE/$argv[1]"

curl -s --cookie "session=$AOC_SESSION" "$AOC_FULL/input" > input

# Now let's set up the go stub
set GOFILE "day$argv[1].go"

echo > $WORKDIR/$GOFILE "\
package main

import (
    \"os\"
    \"bufio\"
    \"fmt\"
    \"flag\"
    \"time\"
    \"utils\"
)

func main() {
    t := time.Now()
    filePtr := flag.String(\"f\", \"input\", \"Input file if not 'input'\")
    // any additional flags add here

	flag.Parse()

    // Choose based on the challenge...
    // individual lines:
    // lines, err := utils.GetTextBlob(*filePtr)
    // if err != nil {
    //     fmt.Println(\"Fatal:\", err)
    // }

    // giant text blob:
    // challengeText, err := utils.GetTextBlob(*filePtr)
    // if err != nil {
    //     fmt.Println(\"Fatal:\", err)
    // }

    // Insert code here

    fmt.Printf(\"Total time elapsed: %s\\n\", time.Since(t))
}
"

# Now handle the go module stuff
pushd $AOC_DIR
pushd day$argv[1]
go mod init day$argv[1]
popd
go work use day$argv[1]
popd
