package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Student struct {
	Nama      string
	Alamat    string
	Pekerjaan string
}

func (s Student) GetStudentInfo() {
	fmt.Printf("Nama : %v, Alamat : %v, Pekerjaan : %v\n", s.Nama, s.Alamat, s.Pekerjaan)
}

func main() {
	f, err := excelize.OpenFile("hacktiv8_golang.xlsx")
	if err != nil {
		fmt.Println("File excel absensi tidak ada")
		return
	}
	defer f.Close()

	if len(os.Args) < 1 {
		fmt.Println("Not enough arguments")
	} else {
		nomorAbsen, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		if nomorAbsen > 25 {
			fmt.Println("Kelas hanya berisi 24 orang!")
			return
		}
		rows, err := f.GetRows("Sheet1")
		if err != nil {
			panic(err)
		}
		siswa := Student{
			Nama:      rows[nomorAbsen][0],
			Alamat:    rows[nomorAbsen][1],
			Pekerjaan: rows[nomorAbsen][2],
		}
		siswa.GetStudentInfo()
	}

}
