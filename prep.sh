#!/bin/bash
# This is the prep.fish script re-written to work with bash. I'm coding on a
# new system and I don't want to install fish, so I'm going to do the sensible
# thing and rewriting it in bash instead of some esoteric shell script
#
# export AOC_SESSION="<session key>"
# . ./prep.sh <day without a pre-pended zero>
#
# where "key" is the session key in your logged-in cookie, and "day" is
# the day you want to prep, just don't prepend a zero to the day because
# it will just fail
#
# Credit where it is due, the concept of this script was inspired heavily by 0xdf's version.

AOC_DIR=$(pwd)
AOC_URL_BASE="https://adventofcode.com/2024/day"

if [ ! -d ${AOC_DIR}/.git ]
then
    echo "Please run this within the root of the aoc directory"
    exit 1
fi

# AOC Session is not set
if [ -z "$AOC_SESSION" ]
then
    echo "AOC Session is not set!"
    echo "Please enter 'export AOC_SESSION=<session key>"
    exit 2
fi

if [ -z "$1" ]
then
    echo "Need the day we are working with, please enter the advent day"
    echo "Example: ./prep.sh 5"
    exit 3
fi

WORKDIR="$AOC_DIR/day$1"

mkdir -p $WORKDIR
pushd $WORKDIR

AOC_FULL="$AOC_URL_BASE/$1"

curl -s --cookie "session=$AOC_SESSION" "$AOC_FULL/input" > input

# Now let's set up the go stub
GOFILE="day$1.go"

cat > $WORKDIR/$GOFILE <<EOF
package main

import (
    "os"
    "bufio"
    "fmt"
    "flag"
    "time"
)

func main() {
    t := time.Now()
    filePtr := flag.String("f", "input", "Input file if not 'input'")

	flag.Parse()
	readFile, err := os.Open(*filePtr)

	if err != nil {
		fmt.Println("Fatal:", err)
	}
	defer readFile.Close()

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    var lines []string

    for fileScanner.Scan() {
        lines = append(lines, fileScanner.Text())
    }

    // Insert code here

    fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
EOF

# Now handle the go module stuff
go mod init day$1
popd
go work use day$1

