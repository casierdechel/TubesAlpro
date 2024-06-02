package main

import (
	"fmt"
	"os"
)

type Parkir struct {
	jenis     string
	plat      string
	jamIn     int
	menitIn   int
	detikIn   int
	jamOut    int
	menitOut  int
	detikOut  int
	costParkir int
}

type Pengguna struct {
	uName string
	pw    string
}

var masuk int
var jumlahTP int

const N int = 1000

type TabPark [N]Parkir
type TabPengguna [N]Pengguna

var TP TabPengguna
var P TabPark

func main() {
	SelamatDatang()
}

func SelamatDatang() {
	var yesno int
	for {
		fmt.Println("Welcome to ParKing")
		fmt.Println("1. Log in")
		fmt.Println("2. Register")
		fmt.Println("3. Exit")
		fmt.Scan(&yesno)
		fmt.Println()
		if yesno == 1 {
			LogIn()
		} else if yesno == 2 {
			Register()
		} else if yesno == 3 {
			Exit()
		}
	}
}

func Register() {
	fmt.Println("Sign up your username.")
	fmt.Scan(&TP[jumlahTP].uName)
	fmt.Println("Make your password.")
	fmt.Scan(&TP[jumlahTP].pw)
	jumlahTP++
	fmt.Println()
	Menu()
}

func LogIn() {
	if jumlahTP == 0 {
		fmt.Println("No registered users. Please register first.")
		fmt.Println()
		Register()
		return
	}

	var lastUname string
	var lastpw string
	var tAttempt int = 3
	fmt.Println("Log into your account.")
	fmt.Println("Username :")
	fmt.Scan(&lastUname)
	fmt.Println("Password :")
	fmt.Scan(&lastpw)
	for i := 0; i < jumlahTP; i++ {
		for j := 0; j < tAttempt; j++ {
			if lastUname == TP[i].uName && lastpw == TP[i].pw {
				fmt.Println("Success")
				fmt.Println()
				Menu()
				return
			} else if j < tAttempt-1 {
				fmt.Println("Login Failed")
				fmt.Printf("You have %d attempt left\n", tAttempt-(j+1))
				fmt.Println("Username :")
				fmt.Scan(&lastUname)
				fmt.Println("Password :")
				fmt.Scan(&lastpw)
				fmt.Println()
			} else {
				fmt.Println("Login Failed")
				fmt.Println("You have no attempts left. Exiting program.")
				Exit()
			}
		}
	}
}

func Menu() {
	var choose int
	for {
		fmt.Println("What can ParKing help you with today?")
		fmt.Println("1. Add Parking Data")
		fmt.Println("2. Change Parking Data")
		fmt.Println("3. Erase Parking Data")
		fmt.Println("4. Vehicle Out")
		fmt.Println("5. Print Data")
		fmt.Println("6. Log Out")
		fmt.Println("7. Exit")
		fmt.Scan(&choose)
		fmt.Println()
		if choose == 1 {
			BacaData()
		} else if choose == 2 {
			change()
		} else if choose == 3 {
			erase()
		} else if choose == 4 {
			KeluarParkir()
		} else if choose == 5 {
			CetakData()
		} else if choose == 6 {
			fmt.Println()
			return
		} else if choose == 7 {
			fmt.Println("Thank You for using ParKing!")
			Exit()
		}
		fmt.Println()
	}
}

func change() {
	var tempChange string
	fmt.Println("Please input the vehicle plate data that you want to change.")
	fmt.Scan(&tempChange)
	for i := 0; i < masuk; i++ {
		if tempChange == P[i].plat {
			fmt.Println("Please input all the data here (including the data change).")
			fmt.Println("Type :")
			fmt.Scan(&P[i].jenis)
			fmt.Println("Vehicle Plate :")
			fmt.Scan(&P[i].plat)
			fmt.Println("Entry Time (hour, minute, second) :")
			fmt.Scan(&P[i].jamIn, &P[i].menitIn, &P[i].detikIn)
			fmt.Println("Exit Time (hour, minute, second) :")
			fmt.Scan(&P[i].jamOut, &P[i].menitOut, &P[i].detikOut)
			fmt.Println("Parking Cost :")
			fmt.Scan(&P[i].costParkir)
			fmt.Println("Data has been changed.")
			return
		}
	}
	fmt.Println("Vehicle with the specified plate not found.")
}

func erase() {
	var tempPlat string
	fmt.Print("Input the Vehicle Plate to find the data that you want to erase: ")
	fmt.Scan(&tempPlat)

	for i := 0; i < masuk; i++ {
		if tempPlat == P[i].plat {
			for j := i; j < masuk-1; j++ {
				P[j] = P[j+1]
			}
			masuk--
			P[masuk] = Parkir{}
			fmt.Println("Data has been deleted.")
			return
		}
	}
	fmt.Println("Vehicle with the specified plate not found.")
}

func BacaData() {
	var tempMasuk int
	fmt.Println("How many data do you want to add?")
	fmt.Scan(&tempMasuk)
	fmt.Println("Please Input the data here.")
	for i := masuk; i < masuk+tempMasuk; i++ {
		fmt.Println("Type :")
		fmt.Scan(&P[i].jenis)
		fmt.Println("Vehicle Plate :")
		fmt.Scan(&P[i].plat)
		fmt.Println("Entry Time (hour, minute, second) :")
		fmt.Scan(&P[i].jamIn, &P[i].menitIn, &P[i].detikIn)
	}
	masuk += tempMasuk
}

func KeluarParkir() {
	var tempPlat string
	fmt.Println("Please input the vehicle plate.")
	fmt.Scan(&tempPlat)
	for i := 0; i < masuk; i++ {
		if P[i].plat == tempPlat {
			fmt.Println("Please input the exit hour, exit minute, and exit second")
			fmt.Scan(&P[i].jamOut, &P[i].menitOut, &P[i].detikOut)
			ParkCost(i)
			return
		}
	}
	fmt.Println("Vehicle with the specified plate not found.")
}

func ParkCost(index int) {
	totalDurationInSeconds := ((P[index].jamOut * 3600) + (P[index].menitOut * 60) + P[index].detikOut) - ((P[index].jamIn * 3600) + (P[index].menitIn * 60) + P[index].detikIn)
	totalHours := totalDurationInSeconds / 3600
	if P[index].jenis == "Car" || P[index].jenis == "car" {
		P[index].costParkir = (totalHours * 3000) + 5000
	} else if P[index].jenis == "Motorcycle" || P[index].jenis == "motorcycle" {
		P[index].costParkir = (totalHours * 1000) + 2000
	}
	fmt.Println("Total Parking Cost:", P[index].costParkir)
}

func CetakData() {
	fmt.Println("ParKing Data History")
	fmt.Printf("%-15s %-20s %-15s %-15s %-15s %-15s %-15s %-15s %-15s\n", "Type", "Vehicle Plate", "Entry Hour", "Entry Minute", "Entry Second", "Exit Hour", "Exit Minute", "Exit Second", "Parking Cost")
	for i := 0; i < masuk; i++ {
		fmt.Printf("%-15s %-20s %-15d %-15d %-15d %-15d %-15d %-15d %-15d\n", P[i].jenis, P[i].plat, P[i].jamIn, P[i].menitIn, P[i].detikIn, P[i].jamOut, P[i].menitOut, P[i].detikOut, P[i].costParkir)
		fmt.Println()
	}
	Total()
}

func Total() {
	var totalP int = 0
	for i := 0; i < masuk; i++ {
		totalP += P[i].costParkir
	}
	fmt.Println("Total Parking Cost:", totalP)
	fmt.Println()
}

func Exit() {
	fmt.Println()
	os.Exit(0)
}