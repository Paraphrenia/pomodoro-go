package main

import (
	"fmt"
	"os"
	"time"
)

type Stats struct {
	breakTime     int
	longBreakTime int
	workTime      int
	totalPomos    int
	totalBreaks   int
}

func newStats(breakTime int, longBreakTime int, workTime int, totalPomos int, totalBreaks int) *Stats {
	return &Stats{
		breakTime:     breakTime,
		longBreakTime: longBreakTime,
		workTime:      workTime,
		totalPomos:    totalPomos,
		totalBreaks:   totalBreaks,
	}
}

func printLineBreak() {
	fmt.Println("----------------")
}

func ASCIIArt() {
	fmt.Println(`
__________                         .___                    ___________.__                      
\______   \____   _____   ____   __| _/___________  ____   \__    ___/|__| _____   ___________ 
 |     ___/  _ \ /     \ /  _ \ / __ |/  _ \_  __ \/  _ \    |    |   |  |/     \_/ __ \_  __ \
 |    |  (  <_> )  Y Y  (  <_> ) /_/ (  <_> )  | \(  <_> )   |    |   |  |  Y Y  \  ___/|  | \/
 |____|   \____/|__|_|  /\____/\____ |\____/|__|   \____/    |____|   |__|__|_|  /\___  >__|   
                      \/            \/                                         \/     \/       
`)

}

func introMenuDisplay(stats *Stats) {
	fmt.Println("MAIN MENU")
	printLineBreak()
	fmt.Println("1. Start")
	fmt.Println("2. Stats")
	fmt.Println("3. Settings")
	fmt.Println("4. Exit")
	printLineBreak()

	for {
		var x int
		fmt.Scanln(&x)

		switch x {
		case 1:
			printLineBreak()
			fmt.Println("Starting Pomodoro")
			printLineBreak()
			startPomo(stats)
		case 2:
			printLineBreak()
			fmt.Println("Stats page")
			statsDisplay(stats)
		case 3:
			printLineBreak()
			fmt.Println("Settings page")
			settingsDisplay(stats)
		case 4:
			printLineBreak()
			fmt.Println("Exiting program")
			printLineBreak()
			exitProgram()
		default:
			printLineBreak()
			fmt.Println("Invalid choice. Please try again.")
			printLineBreak()
			introMenuDisplay(stats)
		}
	}
}

func startPomo(stats *Stats) {
	fmt.Println("Pomodoro Menu")
	fmt.Println()
	fmt.Println("1. Start pomodoro")
	fmt.Println("2. Exit")
	fmt.Println()

	var x = -1
	fmt.Println("Enter choice:")
	printLineBreak()
	fmt.Scanln(&x)

	switch x {
	case 1:
		var work = stats.workTime
		var breaks = stats.breakTime
		var longBreak = stats.longBreakTime

		printLineBreak()

		if stats.totalPomos%4 == 0 && stats.totalPomos != 0 {
			breaks = longBreak
		}

		for work > 0 {
			fmt.Println(work, "minutes of work")
			time.Sleep(1 * time.Minute)
			work--
		}
		fmt.Print("\a")

		for breaks > 0 {
			fmt.Println(breaks, "minutes of break")
			time.Sleep(1 * time.Minute)
			breaks--
		}
		fmt.Print("\a")

		stats.totalPomos++
		stats.totalBreaks++

		startPomo(stats)
	case 2:
		printLineBreak()
		introMenuDisplay(stats)
		printLineBreak()
	default:
		printLineBreak()
		fmt.Println("Invalid choice. Please try again.")
		printLineBreak()
		startPomo(stats)
	}

}

func statsDisplay(stats *Stats) {
	fmt.Println("Total Pomodoros:", stats.totalPomos)
	fmt.Println("Total Breaks:", stats.totalBreaks)

	fmt.Println()
	fmt.Println("1. Reset total pomodoros")
	fmt.Println("2. Reset total breaks")
	fmt.Println("3. Exit")

	fmt.Println()
	var x = -1
	fmt.Println("Enter choice: ")
	fmt.Scanln(&x)

	for {
		switch x {
		case 1:
			printLineBreak()
			stats.totalPomos = 0
			fmt.Println()
			statsDisplay(stats)
		case 2:
			printLineBreak()
			stats.totalBreaks = 0
			fmt.Println()
			statsDisplay(stats)
		case 3:
			printLineBreak()
			introMenuDisplay(stats)
		default:
			printLineBreak()
			fmt.Println("Invalid choice. Please try again.")
			printLineBreak()
			fmt.Println()
			statsDisplay(stats)
		}
	}
}

func settingsDisplay(stats *Stats) {
	printLineBreak()
	fmt.Println("Break time:", stats.breakTime)
	fmt.Println("Long break time:", stats.longBreakTime)
	fmt.Println("Work time:", stats.workTime)

	fmt.Println()
	fmt.Println("1. Set break time")
	fmt.Println("2. Set long break time")
	fmt.Println("3. Set work time")
	fmt.Println("4. Exit ")
	fmt.Println()

	var x = -1
	fmt.Println("Enter choice: ")
	fmt.Scanln(&x)

	switch x {
	case 1:
		printLineBreak()
		fmt.Println("Enter break time")
		fmt.Scanln(&stats.breakTime)
		settingsDisplay(stats)
	case 2:
		printLineBreak()
		fmt.Println("Enter break time")
		fmt.Scanln(&stats.longBreakTime)
		settingsDisplay(stats)
	case 3:
		printLineBreak()
		fmt.Println("Enter work time")
		fmt.Scanln(&stats.workTime)
		settingsDisplay(stats)
	case 4:
		printLineBreak()
		introMenuDisplay(stats)
	default:
		printLineBreak()
		fmt.Println("Invalid choice. Please try again.")
		printLineBreak()
		fmt.Println()
		settingsDisplay(stats)
	}
}

func exitProgram() {
	os.Exit(0)
}

func main() {
	breakTime := 5
	longBreakTime := 15
	workTime := 25
	totalPomos := 0
	totalBreaks := 0
	stats := newStats(breakTime, longBreakTime, workTime, totalPomos, totalBreaks)

	ASCIIArt()
	introMenuDisplay(stats)
}
