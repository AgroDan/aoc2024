# This is the prep.ps1 script rewritten to work with PowerShell.
# Export your AOC_SESSION variable using:
# $Env:AOC_SESSION="<session key>"
# .\prep.ps1 <day without a leading zero>
#
# "key" is the session key from your logged-in cookie, and "day" is
# the day you want to prep. Avoid leading zeros in the day input, as
# it will fail.
#
# Credit where due: the concept of this script was inspired heavily by 0xdf's version.

# Set base directory and URL
$AOC_DIR = Get-Location
$AOC_URL_BASE = "https://adventofcode.com/2024/day"

# Check if the directory is a Git repository
if (-Not (Test-Path "$AOC_DIR\.git")) {
    Write-Error "Please run this within the root of the AOC directory"
    exit 1
}

# Check if AOC_SESSION is set
if (-Not $Env:AOC_SESSION) {
    Write-Error "AOC_SESSION is not set!"
    Write-Host "Please set it using: `$Env:AOC_SESSION='<session key>'"
    exit 2
}

# Check if a day argument is provided
if (-Not $args[0]) {
    Write-Error "Need the day we are working with. Please enter the advent day."
    Write-Host "Example: .\prep.ps1 5"
    exit 3
}

# Prepare the working directory
$Day = $args[0]
$WORKDIR = Join-Path $AOC_DIR "day$Day"

New-Item -ItemType Directory -Path $WORKDIR -Force | Out-Null
Set-Location $WORKDIR

# Set up the cookie _the microsoft way_
$cookieContainer = New-Object System.Net.CookieContainer
$cookie = New-Object System.Net.Cookie("session", $Env:AOC_SESSION, "/", "adventofcode.com")
$cookieContainer.Add($cookie)

# Web session
$webSession = New-Object Microsoft.PowerShell.Commands.WebRequestSession
$webSession.Cookies.Add($cookie)

# Fetch the input file
$AOC_FULL = "$AOC_URL_BASE/$Day"
Invoke-WebRequest -Uri "$AOC_FULL/input" -WebSession $webSession -OutFile "input"

# Create the Go stub file
$GOFILE = "day$Day.go"
$GoStub = @"
package main

import (
    "os"
    "bufio"
    "fmt"
    "flag"
    "time"
    "utils"
)

func main() {
    t := time.Now()
    filePtr := flag.String("f", "input", "Input file if not 'input'")
    // any additional flags add here

	flag.Parse()

    // Choose based on the challenge...
    // individual lines:
    // lines, err := utils.GetTextBlob(*filePtr)
    // if err != nil {
    //     fmt.Println("Fatal:", err)
    // }
    
    // giant text blob:
    // challengeText, err := utils.GetTextBlob(*filePtr)
    // if err != nil {
    //     fmt.Println("Fatal:", err)
    // }

    // Insert code here

    fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
"@

Set-Content -Path (Join-Path $WORKDIR $GOFILE) -Value $GoStub

# Initialize Go module
Set-Location $WORKDIR
& go mod init "day$Day"
Set-Location $AOC_DIR
& go work use "day$Day"

Write-Host "Preparation for day $Day completed successfully!"
