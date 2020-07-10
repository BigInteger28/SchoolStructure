// main.go
package main

import (
	"fmt"
)

func main() {
	for {
		var choice int
		var slaapdiscipline float64
		var slaaphorloge float64
		var hardlopen float64
		var pompen float64
		var hygiëne float64
		var productief float64

		fmt.Println("1. Slaap discipline")
		fmt.Println("2. Slaap met horloge")
		fmt.Println("3. Hardlopen")
		fmt.Println("4. Pompen")
		fmt.Println("5. Hygiëne")
		fmt.Println("6. Productief zijn")
		fmt.Println("\nChoice: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			fmt.Println("Gaan slapen om uur (zonder minuten): ")
			var uur float64
			fmt.Scanln(&uur)
			fmt.Println("Gaan slapen om minuten (zonder uren): ")
			var minuten float64
			fmt.Scanln(&minuten)
			minuten = (uur * 60) + minuten
			if uur < 18 {
				if uur > 1 && minuten > 30 {
					slaapdiscipline = 0
				} else {
					slaapdiscipline = 30 - (minuten / 3)
				}
			} else {
				slaapdiscipline = 30 + (1440-minuten)/3
			}

			fmt.Println("\nScore slaap discipline): ", slaapdiscipline)
		}

		if choice == 2 {
			fmt.Println("Aantal uur geslapen (zonder minuten): ")
			var uren float64
			fmt.Scanln(&uren)
			fmt.Println("Aantal minuten geslapen (zonder uren): ")
			var minuten float64
			fmt.Scanln(&minuten)
			minuten = (uren * 60) + minuten
			if minuten < 480 {
				slaaphorloge = (minuten - 240) / 3.75
			} else {
				slaaphorloge = 64 + (minuten-480)/6
			}
			fmt.Println("\nScore Slaap (horloge): ", slaaphorloge)
		}

		if choice == 3 {
			fmt.Print("\nMinuten per km (zonder seconden): ")
			var minuten float64
			fmt.Scanln(&minuten)
			fmt.Print("\nSeconden per km (zonder minuten): ")
			var seconden float64
			fmt.Scanln(&seconden)
			fmt.Print("\nAfstand in km: ")
			var afstand float64
			fmt.Scanln(&afstand)
			hardlopen = (0 + ((650-((minuten*60)+seconden))/5)*afstand) / 4
			fmt.Println("\nScore hardlopen: ", hardlopen)
		}

		if choice == 4 {
			fmt.Print("\nLevel: ")
			var level float64
			fmt.Scanln(&level)
			fmt.Print("\nAantal keer gepompt: ")
			var pomp float64
			fmt.Scanln(&pomp)
			pompen = pomp / (level / 5)
			fmt.Println("Score pompen: ", pompen)
		}

		if choice == 5 {
			fmt.Print("\nHoeveel keer tanden gepoetst: ")
			var tanden float64
			fmt.Scanln(&tanden)
			fmt.Print("\nAantal keer gewassen: ")
			var gewassen float64
			fmt.Scanln(&gewassen)
			fmt.Print("\nAantal dagen: ")
			var dagen float64
			fmt.Scanln(&dagen)
			hygiëne = ((gewassen / dagen) * 50) + ((tanden / dagen) * 25)
			fmt.Println("Score hygiëne: ", hygiëne)
		}

		if choice == 6 {
			fmt.Println("Level: ")
			var level float64
			fmt.Scanln(&level)
			fmt.Println("Uren productief geweest (zonder minuten): ")
			var uren float64
			fmt.Scanln(&uren)
			fmt.Println("Minuten productief geweest (zonder uren): ")
			var minuten float64
			fmt.Scanln(&minuten)
			productief = ((uren * 60) + minuten) / level
			fmt.Println("Score productief: ", productief)
		}
		fmt.Println("\n\n")
	}
}
