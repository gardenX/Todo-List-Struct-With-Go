package main

import (
	"fmt"
	"os"
)

func actList() {
	fmt.Println("=== To-DO List ====")
	fmt.Println("1. Tambah Tugas")
	fmt.Println("2. Lihat Daftar Tugas")
	fmt.Println("3. Tandai Tugas Selesai")
	fmt.Println("4. Filter Tugas")
	fmt.Println("5. Hapus Tugas")
	fmt.Println("6. Keluar")
}

func AddList(ToDoList []ToDo) []ToDo {
	var listNew string
	fmt.Print("Masukan Nama Tugas :")
	fmt.Scanln(&listNew)

	var todoAdd ToDo
	todoAdd.ListData = listNew
	todoAdd.Status = false

	ToDoList = append(ToDoList, todoAdd)

	return ToDoList
}

func ViewList(ToDoList []ToDo) {
	fmt.Println("Daftar Tugas:")
	for i := 0; i < len(ToDoList); i++ {
		if ToDoList[i].Status == true {
			fmt.Println(i+1, ToDoList[i].ListData, "- Selesai")
		} else {
			fmt.Println(i+1, ToDoList[i].ListData, "- Belum Selesai")
		}
	}
}

func MarkList(ToDoList []ToDo) {
	fmt.Print("Pilih Daftar Tugas :")
	var NumList int
	fmt.Scanln(&NumList)
	NumList = NumList - 1

	ToDoList[NumList].Status = true

	ViewList(ToDoList)

}

func filterList(ToDoList []ToDo) {
	fmt.Println("1. Tugas Selesai")
	fmt.Println("2. Tugas Belum Selesai")
	fmt.Print("Pilih Filter :")

	var numList int
	fmt.Scan(&numList)

	for i := 0; i < len(ToDoList); i++ {
		if ToDoList[i].Status == (numList == 1) {
			fmt.Println(i+1, ToDoList[i].ListData)
		}
	}
}

func delList(ToDolist []ToDo) []ToDo {
	fmt.Print("Pilih Daftar Tugas :")
	var NumList int
	fmt.Scanln(&NumList)

	var indDel int = NumList - 1

	ToDolist = append(ToDolist[:indDel], ToDolist[indDel+1:]...)

	return ToDolist
}

type ToDo struct {
	ListData string
	Status   bool
}

func main() {
	var ToDoList []ToDo

	for {
		actList()

		var act int
		fmt.Print("Pilih Operasi (1/2/3/4/5) :")
		fmt.Scanln(&act)

		switch act {
		case 1:
			ToDoList = AddList(ToDoList)
			fmt.Println("Tugas Di Tambahkan !")
			fmt.Println("")
			fmt.Println("")
		case 2:
			ViewList(ToDoList)
			fmt.Println("")
			fmt.Println("")
		case 3:
			MarkList(ToDoList)
			fmt.Println("")
			fmt.Println("")
		case 4:
			filterList(ToDoList)
			fmt.Println("")
			fmt.Println("")
		case 5:
			ToDoList = delList(ToDoList)
			//ViewList(slc1, status)
			fmt.Println("")
			fmt.Println("")
		case 6:
			os.Exit(0)
		}
	}
}
