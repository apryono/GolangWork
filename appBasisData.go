package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var isInt = regexp.MustCompile("^[0-9]+$")

func main() {
	var inputMenu, age, getMenu int
	var nama, jurusan string

	symbol := strings.Repeat("=", 40)

	// fungsi untuk pilihan
	getMenu = chooseMenu(inputMenu)

	if getMenu == 1 {

		fmt.Println(symbol)
		fmt.Println("Add Mahasiswa")
		fmt.Println(symbol)

		// fungsi untuk input nama
		addName(nama)
		// fungsi untuk input umur
		addAge(age)
		// fungsi untuk input jurusan
		addMajor(jurusan)

	} else if getMenu == 2 {

		fmt.Println(symbol)
		fmt.Println("Delete Mahasiswa")
		fmt.Println(symbol)

	} else if getMenu == 3 {

		fmt.Println(symbol)
		fmt.Println("View Mahasiswa")
		fmt.Println(symbol)

	} else {
		return
	}

}

// fungsi untuk menu pilihan
func chooseMenu(input int) int {

	for {
		symbol := strings.Repeat("=", 40)
		fmt.Println(symbol)
		fmt.Println("Main Menu")
		fmt.Println(symbol)
		fmt.Println(`
	Silahkan Pilih Menu 
	1. Add Mahasiswa
	2. Delete Mahasiswa
	3. View Mahasiswa
	4. Exit
	`)
		fmt.Println(symbol)
		fmt.Print("Masukkan Menu yang dipilih : ")
		fmt.Scanln(&input)

		if input > 0 && input < 5 {
			break
		}
		fmt.Println("Error : Pilihan Tidak Ada")
	}

	return input
}

//fungsi untuk input nama
func addName(nama string) string {
	for {
		fmt.Print("Masukkan Nama ( Karakter > 2 dan < 21 ): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		nama = scanner.Text()

		if len(nama) > 2 && len(nama) < 21 {
			break
		}
		fmt.Println("ERROR : Karakter harus lebih dari 2 dan kurang dari 21")
	}
	fmt.Println("Nama Anda : ", nama)
	return nama
}

// fungsi untuk input umur
func addAge(age int) int {
	for {
		fmt.Print("Umur Anda ( Umur > 17 ): ")
		fmt.Scan(&age)

		if age >= 17 {
			break
		}
		fmt.Println("ERROR : Umur harus lebih dari 17 tahun")
	}
	return age
}

// fungsi untuk input jurusan
func addMajor(major string) string {
	for {
		fmt.Print("Jurusan (Karakter <= 10): ")
		fmt.Scan(&major)

		if len(major) <= 10 {
			break
		}
		fmt.Println("ERROR :Karakter harus Di bawah atau sama dengan 10")
	}
	return major
}
