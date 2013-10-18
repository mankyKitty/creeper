package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
	"strconv"
	"strings"
)

// Some basic and kinda stupid protection against fail-biscuit arguments.
func checkInputArgs(f, c string, t int) bool {
	return (f == "" || c == "" || t <= 0)	
}

// Display a basic output to the user of what we're planning to do
// with the chosen file that we're creeping on..
func displayIntent(f, c string, a []string, t int) {
	fmt.Printf("Creeping on: '%v'\n", f)
	fmt.Printf("When modified, Creeper will...\nExecute: '%v',\nWith Arguments: '%v'\n", c, a)
}

func main() {
	// Hide the starting message of intent.
	quiet := flag.Bool("q", false, "Intent message won't be displayed")

	// Try to hide any output of anything except Creeper crashes.
	superQuiet := flag.Bool("shuttup", false, "Creeper will make a best error to not make any std/out/err noise at all.")

	// Time in milliseconds between creepy drive bys on the file.
	tickerDuration := flag.Int("wait", 1000, "Time in Milliseconds between creeping on the file")
	
	// This is what we're actually going to creep on.
	file := flag.String("file", "", "This is the file to watch!")
	
	// This is what we will do if there is a successful creep event.
	cmd := flag.String("cmd", "", "This is the cmd to execute")
	
	// Nom the args into useful values..
	flag.Parse()

	// Assume (ho ho ho) that everything trailing is an argument to be parsed to *cmd
	cmdArgs := flag.Args()

	// Examine our arguments to make sure we're lined up..
	if checkInputArgs(*file, *cmd, *tickerDuration) || !flag.Parsed() {
		flag.Usage()
		os.Exit(1)
	}

	// If we make it this far print some message about what we're configured to do
	// so the user knows what they're in for..
	if !*quiet && !*superQuiet {
		displayIntent(*file, *cmd, cmdArgs, *tickerDuration)
	}

	// Get the initial file information.
	f, err := os.Stat(*file)

	if err != nil {
		// File error!
		fmt.Println("File error! Check the path and try again.")
		os.Exit(1)
	}

	// Get an initial modification time for comparison.
	lastModTime := f.ModTime()

	duration := strconv.Itoa(*tickerDuration) + "ms"

	// Configure our tick duration.. 
	sec, err := time.ParseDuration(duration)

	if err != nil {
		// BE BROKE! :(
		fmt.Println("Timer Creation Failed")
		os.Exit(1)
	}

	// Ticker is just a channel that puts a value on the channel on each tick of
	// of our configured duration.
	ticker := time.Tick(sec)

	// Start monitoring and tickening !
	for _ = range ticker {

		// Check the file so we have fresh information.
		fileInfo, err := os.Stat(*file)

		// Make sure the file hasn't disappeared for some reason.
		if err != nil {
			fmt.Println("File update error! Check file hasn't done a Batman and restart.")
			os.Exit(1)
		}

		// Compare the modification times of the files as that is our core
		// trigger to take action.
		if fileInfo.ModTime() != lastModTime {

			// Create our command thingy
			stuffDoer := exec.Command(*cmd, strings.Join(cmdArgs, " "))

			// Run the command
			out, cmdErr := stuffDoer.CombinedOutput()

			if cmdErr != nil {
				fmt.Println("Command Execution Failed. Output below...")
				fmt.Println(cmdErr.Error())
				os.Exit(1)
			}

			if !*superQuiet {
				fmt.Printf("%s", out)
			}

			lastModTime = fileInfo.ModTime()
		}
	}
}