package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/fatih/color"
)

var (
	delayMs   = flag.Int("delay-ms", 50, "delay in printing each char, in ms")
	clrScreen = flag.Bool("clr-screen", true, "whether to clear the screen before printing")
	col       = flag.String("color", "green", "the color to use for printing; possible values: yellow, blue, red")
	sleepMs   = flag.Int("sleep-ms", 0, "ms to sleep for at the end")
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ezPrint(str string, delay time.Duration, c *color.Color, clrScreen bool, sleepMs time.Duration) {
	if clrScreen {
		clearScreen()
	}
	for _, char := range str {
		c.Printf("%c", char)
		time.Sleep(delay * time.Millisecond)
	}
	if sleepMs > 0 {
		time.Sleep(sleepMs * time.Millisecond)
	}
}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "write stdin to stdout, but slowly\n\nFlags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	time.Sleep(50 * time.Millisecond)

	fi, err := os.Stdin.Stat()
	if err != nil {
		os.Exit(1)
	}
	size := fi.Size()
	if size == 0 {
		os.Exit(0)
	}

	scanner := bufio.NewScanner(os.Stdin)

	var input string
	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}

	var ca color.Attribute

	switch *col {
	case "green":
		ca = color.FgGreen
	case "yellow":
		ca = color.FgYellow
	case "blue":
		ca = color.FgHiBlue
	case "red":
		ca = color.FgHiRed
	default:
		ca = color.FgGreen
	}

	var c = color.New(ca, color.Bold)

	if scanner.Err() != nil {
		fmt.Fprintf(os.Stdout, "Error reading input: %s", scanner.Err())
		os.Exit(1)
	}

	ezPrint(input, time.Duration(*delayMs), c, *clrScreen, time.Duration(*sleepMs))
}
