/*
NIM		: 1301190324
Nama	: Kaenova Mahendra Auditama
Kelas	: IF-43-02
Topik	: Penjadwalan Kegiatan 2
*/

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//ClearScreen Method (source : https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go)

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

//Global Var :

var (
	clear  map[string]func()
	user   DataUser
	jadwal [1000]Event
)

type DataUser struct {
	nama     string
	password string
}

type EventTime struct {
	tanggal     int
	bulan       int
	tahun       int
	jam_mulai   int
	menit_mulai int
	jam_akhir   int
	menit_akhir int
}

type Event struct {
	time      EventTime
	tempat    string
	eventname string
	id        string
}

//UI

func Header() {
	fmt.Println("==============================================")
	fmt.Println("MyPlanner by. Kaenova Mahendra Auditama")
	fmt.Println("==============================================")

}

func AccountDetails() {
	fmt.Println("Account Details")
	fmt.Println("==================")
	fmt.Println("Username :", user.nama)
	fmt.Println("Password :", user.password)
	fmt.Println("==================")
}

func menuawal(terpilih *int) {
	var pilihan int
	CallClear()
	Header()
	fmt.Println("Menu :")
	fmt.Println("1. Create Akun")
	fmt.Println("2. Login Akun")
	fmt.Println("3. Exit")
	fmt.Print("Pilihan : ")
	fmt.Scanln(&pilihan)
	for (pilihan < 1) || (pilihan > 3) {
		CallClear()
		Header()
		fmt.Println("Pilihan tidak valid")
		fmt.Println("Menu :")
		fmt.Println("1. Create Akun")
		fmt.Println("2. Login Akun")
		fmt.Println("3. Exit")
		fmt.Print("Pilihan : ")
		fmt.Scanln(&pilihan)
	}

	*terpilih = pilihan
}

func menu(terpilih *int, memori_n int) {
	var pilihan int
	CallClear()
	Header()
	AccountDetails()
	fmt.Println("Banyaknya jadwal yang dibuat :", memori_n)
	fmt.Println("Menu :")
	fmt.Println("1. Create Jadwal")
	fmt.Println("2. Hapus Jadwal")
	fmt.Println("3. Cek Jadwal")
	fmt.Println("4. Exit")
	fmt.Print("Pilihan : ")
	fmt.Scanln(&pilihan)
	for (pilihan < 1) || (pilihan > 4) {
		CallClear()
		Header()
		AccountDetails()
		fmt.Println("Banyaknya jadwal yang dibuat :", memori_n)
		fmt.Println("Menu :")
		fmt.Println("1. Create Jadwal")
		fmt.Println("2. Hapus Jadwal")
		fmt.Println("3. Cek Jadwal")
		fmt.Println("4. Exit")
		fmt.Println("Pilihan tidak valid")
		fmt.Print("Pilihan : ")
		fmt.Scanln(&pilihan)
	}

	*terpilih = pilihan

}

func menulihatjadwal(terpilih *int) {
	var pilihan int
	fmt.Println("Menu")
	fmt.Println("1. Tampilkan jadwal dengan tanggal")
	fmt.Println("2. Tampilkan jadwal dengan bulan")
	fmt.Println("3. Tampilkan jadwal berdasarkan waktu mulai")
	fmt.Println("4. Tampilkan jadwal yang sudah selesai")
	fmt.Println("5. Tampilkan jadwal yang ada")
	fmt.Println("6. Kembali ke menu")
	fmt.Print("Pilihan : ")
	fmt.Scanln(&pilihan)

	for (pilihan < 1) || (pilihan > 6) {
		CallClear()
		Header()
		AccountDetails()
		fmt.Println("Menu")
		fmt.Println("1. Tampilkan jadwal dengan tanggal")
		fmt.Println("2. Tampilkan jadwal dengan bulan")
		fmt.Println("3. Tampilkan jadwal berdasarkan waktu mulai")
		fmt.Println("4. Tampilkan jadwal yang sudah selesai")
		fmt.Println("5. Tampilkan jadwal yang ada")
		fmt.Println("6. Kembali ke menu")
		fmt.Println("Pilihan tidak valid")
		fmt.Print("Pilihan : ")
		fmt.Scanln(&pilihan)
	}

	*terpilih = pilihan
}

func cleartime(sec int) {
	fmt.Println("Loading...")
	time.Sleep(time.Duration(sec) * time.Second)
	CallClear()
}

//Create Acc
func Username(nama_out *string) {
	var nama string
	fmt.Scanln(&nama)
	for len(nama) == 0 {
		CallClear()
		fmt.Println("==============================================")
		fmt.Println("MyPlanner by. Kaenova Mahendra Auditama")
		fmt.Println("==============================================")
		fmt.Println("Create Account")
		fmt.Println("Username tidak boleh kosong. Harap masukkan username lagi")
		fmt.Println("Username :")
		fmt.Scanln(&nama)
	}
	*nama_out = nama

}

func Password(password_out *string, nama_in string) { //Parameter minimal 2
	var (
		inpassword                                string
		parameter_huruf, parameter_angka, counter int
		slice                                     []string
		passlen                                   int
	)
	parameter_huruf = 0
	parameter_angka = 0
	passlen = 0
	for ((parameter_huruf == 0) && (parameter_angka == 0)) || ((parameter_huruf >= 1) && (parameter_angka == 0)) || ((parameter_huruf == 0) && (parameter_angka >= 1)) || passlen < 8 {
		fmt.Scanln(&inpassword)
		slice = strings.Split(inpassword, "")
		passlen = len(slice)
		counter = 0
		parameter_huruf = 0
		parameter_angka = 0
		for counter < len(slice) {
			if (slice[counter] == "A") || (slice[counter] == "B") || (slice[counter] == "C") || (slice[counter] == "D") || (slice[counter] == "E") || (slice[counter] == "F") || (slice[counter] == "G") || (slice[counter] == "H") || (slice[counter] == "I") || (slice[counter] == "J") || (slice[counter] == "K") || (slice[counter] == "L") || (slice[counter] == "M") || (slice[counter] == "N") || (slice[counter] == "O") || (slice[counter] == "P") || (slice[counter] == "Q") || (slice[counter] == "R") || (slice[counter] == "S") || (slice[counter] == "T") || (slice[counter] == "U") || (slice[counter] == "V") || (slice[counter] == "W") || (slice[counter] == "X") || (slice[counter] == "Y") || (slice[counter] == "Z") {
				parameter_huruf = parameter_huruf + 1
			}
			if (slice[counter] == "0") || (slice[counter] == "1") || (slice[counter] == "2") || (slice[counter] == "3") || (slice[counter] == "4") || (slice[counter] == "5") || (slice[counter] == "6") || (slice[counter] == "7") || (slice[counter] == "8") || (slice[counter] == "9") || (slice[counter] == "0") {
				parameter_angka = parameter_angka + 1
			}
			counter++
		}

		//Hanya untuk tampilan aja ini.
		if ((parameter_huruf <= 0) || (parameter_angka <= 0)) || (passlen < 8) {
			CallClear()
			fmt.Println("==============================================")
			fmt.Println("MyPlanner by. Kaenova Mahendra Auditama")
			fmt.Println("==============================================")
			fmt.Println("Create Account")
			fmt.Println("Username :")
			fmt.Println(nama_in)
			fmt.Println("Password setidaknya terdiri dari 1 angka dan 1 huruf kapital. Password setidaknya sepanjang 8 karakter.")
			fmt.Println("Password :")

		}
	}
	*password_out = inpassword

}

func CreateAcc(nama_out *string, password_out *string) {
	/*
		Username(input : user string) prosedur membuat username yang disimpan di "user"
		Password(input/output : pass string input : user) prosedur yang membuat password
	*/

	fmt.Println("==============================================")
	fmt.Println("MyPlanner by. Kaenova Mahendra Auditama")
	fmt.Println("==============================================")
	fmt.Println("Create Account")
	fmt.Println("Username :")
	Username(&*nama_out)
	nama_temp := nama_out
	fmt.Println("Password :")
	Password(&(*password_out), *nama_temp)

}

//Create Jadwal

func convMonthString(month string) string {

	if month == "January" {
		return "01"
	} else if month == "February" {
		return "02"
	} else if month == "March" {
		return "03"
	} else if month == "April" {
		return "04"
	} else if month == "May" {
		return "05"
	} else if month == "June" {
		return "06"
	} else if month == "July" {
		return "07"
	} else if month == "August" {
		return "08"
	} else if month == "September" {
		return "09"
	} else if month == "October" {
		return "10"
	} else if month == "November" {
		return "11"
	} else if month == "December" {
		return "12"
	} else {
		return "Salah"
	}
}

func convMonthInt(month string) int {

	if month == "January" {
		return 1
	} else if month == "February" {
		return 2
	} else if month == "March" {
		return 3
	} else if month == "April" {
		return 4
	} else if month == "May" {
		return 5
	} else if month == "June" {
		return 6
	} else if month == "July" {
		return 7
	} else if month == "August" {
		return 8
	} else if month == "September" {
		return 9
	} else if month == "October" {
		return 10
	} else if month == "November" {
		return 11
	} else if month == "December" {
		return 12
	} else {
		return -1
	}
}

func InTahun(tahun_in int, tahun_out *int) {
	var temp int
	fmt.Println("Masukkan Tahun : ")
	fmt.Scanln(&temp)
	for (temp < 0) || (temp < tahun_in) {
		fmt.Println("Tahun tidak valid.")
		fmt.Println("Masukkan Tahun :")
		fmt.Scanln(&temp)
	}
	*tahun_out = temp
}

func InBulan(bulan_out *int, tahunawal, tahunakhir, bulan_in int) {
	var temp int
	if tahunawal == tahunakhir {
		fmt.Println("Masukkan Bulan (1-12) :")
		fmt.Scanln(&temp)
		for (temp < bulan_in) || (temp < 1) || (temp > 12) {
			fmt.Println("Bulan tidak valid")
			fmt.Println("Masukkan Bulan (1-12) :")
			fmt.Scanln(&temp)
		}
		*bulan_out = temp
	} else {
		fmt.Println("Masukkan Bulan (1-12) :")
		fmt.Scanln(&temp)
		for (temp < 1) || (temp > 12) {
			fmt.Println("Bulan tidak valid")
			fmt.Println("Masukkan Bulan (1-12) :")
			fmt.Scanln(&temp)
		}
		*bulan_out = temp
	}
}

func kabisat(tahun int) bool {
	// Pengecekkan habis dibagi 400 (x)
	x1 := tahun % 400
	x2 := x1 == 0

	// Pengecekkan habis dibagi 4 (y)
	y1 := tahun % 4
	y2 := y1 == 0

	// Pengecekkan habis dibagi 100 (z)
	z1 := tahun % 100
	z2 := z1 != 0

	//Return
	return (x2 || (y2 && z2))

}

func InTanggalAwal(tanggal *int, bulan, tahun int) {
	var (
		temp            int
		valid, cekbulan bool
	)
	cekbulan = bulan == convMonthInt(time.Now().Month().String())
	fmt.Println("Masukan Tanggal : ")
	valid = false
	fmt.Scanln(&temp) //Input Tanggal
	//Cek Tanggal
	for valid == false {
		if time.Now().Year() == tahun {
			if bulan == 2 && cekbulan {
				if kabisat(tahun) {
					if (temp > 0) && (temp <= 29) && (temp >= time.Now().Day()) {
						*tanggal = temp
						valid = true
					} else {
						fmt.Println("Tanggal tidak valid")
						fmt.Scanln(&temp)
					}
				} else {
					if (temp > 0) && (temp <= 28) && (temp >= time.Now().Day()) {
						*tanggal = temp
						valid = true
					} else {
						fmt.Println("Tanggal tidak valid")
						fmt.Scanln(&temp)
					}
				}
			} else if bulan%2 != 0 && cekbulan {
				if (temp > 0) && (temp <= 31) && (temp >= time.Now().Day()) {
					*tanggal = temp
					valid = true
				} else {
					fmt.Println("Tanggal tidak valid")
					fmt.Scanln(&temp)
				}
			} else if bulan%2 == 0 && cekbulan {
				if (temp > 0) && (temp <= 30) && (temp >= time.Now().Day()) {
					*tanggal = temp
					valid = true
				} else {
					fmt.Println("Tanggal tidak valid")
					fmt.Scanln(&temp)
				}
			} else if bulan == 2 {
				if kabisat(tahun) {
					if (temp > 0) && (temp <= 29) {
						*tanggal = temp
						valid = true
					} else {
						fmt.Println("Tanggal tidak valid")
						fmt.Scanln(&temp)
					}
				} else {
					if (temp > 0) && (temp <= 28) {
						*tanggal = temp
						valid = true
					} else {
						fmt.Println("Tanggal tidak valid")
						fmt.Scanln(&temp)
					}
				}
			} else if bulan%2 != 0 {
				if (temp > 0) && (temp <= 31) {
					*tanggal = temp
					valid = true
				} else {
					fmt.Println("Tanggal tidak valid")
					fmt.Scanln(&temp)
				}
			} else if bulan%2 == 0 {
				if (temp > 0) && (temp <= 30) {
					*tanggal = temp
					valid = true
				} else {
					fmt.Println("Tanggal tidak valid")
					fmt.Scanln(&temp)
				}
			}
		} else if time.Now().Year() < tahun {
			if bulan == 2 {
				if kabisat(tahun) {
					if (temp > 0) && (temp <= 29) {
						*tanggal = temp
						valid = true
					} else {
						fmt.Println("Tanggal tidak valid")
						fmt.Scanln(&temp)
					}
				} else {
					if (temp > 0) && (temp <= 28) {
						*tanggal = temp
						valid = true
					} else {
						fmt.Println("Tanggal tidak valid")
						fmt.Scanln(&temp)
					}
				}
			} else if bulan%2 != 0 {
				if (temp > 0) && (temp <= 31) {
					*tanggal = temp
					valid = true
				} else {
					fmt.Println("Tanggal tidak valid")
					fmt.Scanln(&temp)
				}
			} else if bulan%2 == 0 {
				if (temp > 0) && (temp <= 30) {
					*tanggal = temp
					valid = true
				} else {
					fmt.Println("Tanggal tidak valid")
					fmt.Scanln(&temp)
				}
			}
		}
	}
}

func InJamAwal(jam_out, menit_out *int, bulan, tanggal, tahun int) {
	var (
		tempjam, tempmenit, jamnow, menitnow int
		cekbulan, cektanggal, cektahun       bool //22nya harus true
	)
	cekbulan = bulan == convMonthInt(time.Now().Month().String())
	cektanggal = tanggal == time.Now().Day()
	cektahun = tahun == time.Now().Year()
	jamnow = time.Now().Hour()
	menitnow = time.Now().Minute()
	fmt.Println("Masukkan jam dengan format 24 jam")
	fmt.Scanln(&tempjam)
	if cekbulan && cektanggal && cektahun {
		for tempjam >= 24 || tempjam < 00 || tempjam < jamnow {
			fmt.Println("Masukan jam tidak valid")
			fmt.Println("Masukkan jam dengan format 24 jam")
			fmt.Scanln(&tempjam)
		}

		if tempjam > jamnow {
			fmt.Println("Masukkan menit")
			fmt.Scanln(&tempmenit)
			for tempmenit >= 60 || tempmenit < 0 {
				fmt.Println("Masukan menit tidak valid")
				fmt.Println("Masukkan menit")
				fmt.Scanln(&tempmenit)
			}
		} else if tempjam == jamnow {
			fmt.Println("Masukkan menit")
			fmt.Scanln(&tempmenit)
			for tempmenit >= 60 || tempmenit < 0 || tempmenit < menitnow {
				fmt.Println("Masukan menit tidak valid")
				fmt.Println("Masukkan menit")
				fmt.Scanln(&tempmenit)

			}
		}
	} else {
		for tempjam >= 24 || tempjam < 00 {
			fmt.Println("Masukan jam tidak valid")
			fmt.Println("Masukkan jam dengan format 24 jam")
			fmt.Scanln(&tempjam)
		}

		if tempjam > jamnow {
			fmt.Println("Masukkan menit")
			fmt.Scanln(&tempmenit)
			for tempmenit >= 60 || tempmenit < 0 {
				fmt.Println("Masukan menit tidak valid")
				fmt.Println("Masukkan menit")
				fmt.Scanln(&tempmenit)
			}
		} else {
			fmt.Println("Masukkan menit")
			fmt.Scanln(&tempmenit)
			for tempmenit >= 60 || tempmenit < 0 {
				fmt.Println("Masukan menit tidak valid")
				fmt.Println("Masukkan menit")
				fmt.Scanln(&tempmenit)

			}
		}
	}
	*jam_out = tempjam
	*menit_out = tempmenit
}

func InJamAkhir(jam_out, menit_out *int, tahunawal, bulanawal, jamawal, menitawal, tanggalawal int) {
	var (
		tempjam, tempmenit int
	)

	fmt.Println("Masukkan waktu jam :")
	fmt.Scanln(&tempjam)
	for (tempjam >= 24 || tempjam < 00) || (tempjam < jamawal) {
		fmt.Println("Masukkan jam tidak valid")
		fmt.Println("Masukkan waktu jam :")
		fmt.Scanln(&tempjam)
	}
	*jam_out = tempjam

	fmt.Println("Masukkan waktu menit :")
	fmt.Scanln(&tempmenit)
	if tempjam > jamawal {
		for tempmenit >= 60 || tempmenit < 0 {
			fmt.Println("Masukkan menit tidak valid")
			fmt.Println("Masukkan waktu menit :")
			fmt.Scanln(&tempmenit)
		}
		*menit_out = tempmenit
	} else {
		for (tempmenit >= 60 || tempmenit < 0) || (tempmenit <= menitawal) {
			fmt.Println("Masukkan menit tidak valid")
			fmt.Println("Masukkan waktu menit :")
			fmt.Scanln(&tempmenit)
		}
		*menit_out = tempmenit
	}

}

func namatempat(tempat_out, nama_out *string) {
	var temp1, temp2 string
	fmt.Println("Masukkan nama kegiatan :")
	fmt.Scanln(&temp1)
	for len(temp1) == 0 {
		CallClear()
		Header()
		AccountDetails()
		fmt.Println("Nama jadwal tidak boleh kosong")
		fmt.Println("Masukkan nama kegiatan :")
		fmt.Scanln(&temp1)
	}
	*nama_out = temp1
	fmt.Println("Masukkan tempat kegiatan :")
	fmt.Scanln(&temp2)
	for len(temp2) == 0 {
		CallClear()
		Header()
		AccountDetails()
		fmt.Println("Nama Kegiatan :", temp1)
		fmt.Println("Nama jadwal tidak boleh kosong")
		fmt.Println("Masukkan nama kegiatan :")
		fmt.Scanln(&temp2)
	}
	*tempat_out = temp2
}

func makeID(memori int, idout *string) {
	var (
		tahun, memtemp         int
		bulan, tahunstr, nomor string
		ID                     string
	)
	memtemp = (memori % 1000) + 1
	tahun = (time.Now().Year()) % 100
	tahunstr = strconv.Itoa(tahun)
	bulan = convMonthString(time.Now().Month().String())
	nomor = strconv.Itoa(memtemp)
	if memtemp >= 100 {
		ID = bulan + tahunstr + "-" + nomor
	} else {
		if memtemp < 10 {
			ID = bulan + tahunstr + "-00" + nomor
		} else {
			ID = bulan + tahunstr + "-0" + nomor
		}
	}
	*idout = ID
}

func JadwalCreated(memori int, nama, tempat, id string, tanggalmulai, bulanmulai, tahunmulai, jammulai, jamakhir, menitmulai, menitakhir int) {
	fmt.Println("================================")
	fmt.Println("Jadwal ke-", memori+1, " sudah terbuat")
	fmt.Println("================================")
	fmt.Println("Nama Jadwal :", nama)
	fmt.Println("Tempat :", tempat)
	fmt.Println("Waktu Mulai :")
	fmt.Println("Tanggal [dd/mm/yyyy] :", tanggalmulai, bulanmulai, tahunmulai)
	fmt.Println("jam [hh/mm] :", jammulai, ":", menitmulai)
	fmt.Println("Waktu Selesai :")
	fmt.Println("jam [hh/mm] :", jamakhir, ":", menitakhir)
	fmt.Println("ID :", id)
	fmt.Println("================================")
	cleartime(5)
}

func beririsan(p1, p2 Event) bool {
	//p1 = yang dicek p2 = yang ngecek atau parameternya
	//p1 tidak mungkin lebih kecil karena sudah ada pembatasnya
	//Batas tahun
	var (
		p1mulai, p1selesai, p2mulai, p2selesai int
	)
	p1mulai = (p1.time.jam_mulai * 60) + p1.time.menit_mulai
	p2mulai = (p2.time.jam_mulai * 60) + p2.time.menit_mulai
	p1selesai = (p1.time.jam_akhir * 60) + p1.time.menit_akhir
	p2selesai = (p2.time.jam_akhir * 60) + p2.time.menit_akhir
	if (p1.time.tanggal == p2.time.tanggal) && (p1.time.bulan == p2.time.bulan) && (p1.time.tahun == p2.time.tahun) {
		if ((p1mulai >= p2mulai) && (p1mulai <= p2selesai)) || ((p1selesai >= p2mulai) && (p1selesai <= p2selesai)) || ((p1mulai <= p2mulai) && (p1selesai >= p2selesai)) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

//Cek Jadwal
func tampilkanada(p1 Event) { //pilihan5
	fmt.Println("Nama Jadwal :", p1.eventname)
	fmt.Println("Tempat :", p1.tempat)
	fmt.Println("Waktu Mulai :")
	fmt.Println("Tanggal [dd/mm/yyyy] :", p1.time.tanggal, p1.time.bulan, p1.time.tahun)
	fmt.Println("jam [hh/mm] :", p1.time.jam_mulai, ":", p1.time.menit_mulai)
	fmt.Println("Waktu Selesai :")
	fmt.Println("jam [hh/mm] :", p1.time.jam_akhir, ":", p1.time.menit_akhir)
	fmt.Println("ID :", p1.id)
}

func tampilkantanggal(banyakjadwal int) { //Pilihan1
	var (
		tanggaltemp, bulantemp, tahuntemp int
		countertemparr                    int
		temparr                           [1000]Event
	)
	CallClear()
	Header()
	AccountDetails()
	fmt.Print("Masukkan tanggal : ") //Ini dibuat tidak terlalu detail
	fmt.Scanln(&tanggaltemp)
	fmt.Print("Masukkan bulan (1-12): ")
	fmt.Scanln(&bulantemp)
	fmt.Print("Masukkan tahun : ")
	fmt.Scanln(&tahuntemp)
	cleartime(2)
	Header()
	AccountDetails()
	fmt.Println("Jadwal yang sudah dibuat")
	fmt.Println("=========================")
	countertemparr = 0
	for counter := 0; counter < banyakjadwal; counter++ {
		if (jadwal[counter].time.tanggal == tanggaltemp) && (jadwal[counter].time.bulan == bulantemp) && (jadwal[counter].time.tahun == tahuntemp) {
			temparr[countertemparr] = jadwal[counter]
			countertemparr++
		}
	}

	tampilkanmulai(countertemparr, temparr)
}

func convjammenitmulai(p Event) int {
	return ((p.time.jam_mulai * 60) + p.time.menit_mulai)
}

func tampilkanmulai(banyakjadwal int, jadwal [1000]Event) { //Menggunakan Selection Sort
	var (
		temparr [1000]Event
	)
	temparr = jadwal

	//Tahun
	for i := 0; i < banyakjadwal; i++ {
		min := i
		for j := i + 1; j < banyakjadwal; j++ { //j := i + 1 -_-
			if temparr[min].time.tahun > temparr[j].time.tahun {
				min = j
			}
		}

		temp := temparr[i]
		temparr[i] = temparr[min]
		temparr[min] = temp
	}

	//Bulan
	k := 0
	totaltahun := 0 //Untuk menentukan posisi tahun yang berbeda
	for k < banyakjadwal {
		banyaktahun := 0
		for l := k; l < banyakjadwal; l++ {
			if temparr[k].time.tahun == temparr[l].time.tahun {
				banyaktahun++
			}
		}

		totaltahun = totaltahun + banyaktahun

		for i := k; i < totaltahun-1; i++ { // Harusnya pakai total tahun
			min := i
			for j := i + 1; j < totaltahun; j++ { //j := i + 1 -_-
				if temparr[min].time.bulan > temparr[j].time.bulan {
					min = j
				}
			}

			temp := temparr[i]
			temparr[i] = temparr[min]
			temparr[min] = temp
		}

		k = totaltahun
	}

	//Tanggal
	k = 0
	totaltahun = 0 //Untuk menentukan posisi tahun dan bulan yang berbeda
	for k < banyakjadwal {
		banyaktahun := 0
		for l := k; l < banyakjadwal; l++ {
			if (temparr[k].time.bulan == temparr[l].time.bulan) && (temparr[k].time.tahun == temparr[l].time.tahun) {
				banyaktahun++
			}
		}

		totaltahun = totaltahun + banyaktahun

		for i := k; i < totaltahun-1; i++ {
			min := i
			for j := i + 1; j < totaltahun; j++ { //j := i + 1 -_-
				if temparr[min].time.tanggal > temparr[j].time.tanggal {
					min = j
				}
			}

			temp := temparr[i]
			temparr[i] = temparr[min]
			temparr[min] = temp
		}

		k = totaltahun
	}

	//Menit
	k = 0
	totaltahun = 0 //Untuk menentukan posisi tahun yang berbeda
	for k < banyakjadwal {
		banyaktahun := 0
		for l := k; l < banyakjadwal; l++ {
			if (temparr[k].time.tanggal == temparr[l].time.tanggal) && (temparr[k].time.bulan == temparr[l].time.bulan) && (temparr[k].time.tahun == temparr[l].time.tahun) {
				banyaktahun++
			}
		}

		totaltahun = totaltahun + banyaktahun

		for i := k; i < totaltahun-1; i++ {
			min := i
			for j := i; j < totaltahun; j++ {
				if convjammenitmulai(temparr[min]) > convjammenitmulai(temparr[j]) {
					min = j
				}
			}

			temp := temparr[i]
			temparr[i] = temparr[min]
			temparr[min] = temp
		}

		k = totaltahun
	}

	fmt.Println("Jadwal yang sudah dibuat terurut")
	fmt.Println("Jadwal yang ada :", banyakjadwal)
	for x := 0; x < banyakjadwal; x++ {
		fmt.Println("================================")
		tampilkanada(temparr[x])
	}

}

func tampilkanbulantahun(tempbulan, temptahun, memori int) { // Menggunakan Insertion Sort
	var (
		temparr       [1000]Event
		k, totaltahun int
	)
	y := 0
	//Masukkan temparr
	for x := 0; x < memori; x++ {
		if tempbulan == jadwal[x].time.bulan && temptahun == jadwal[x].time.tahun {
			temparr[y] = jadwal[x]
			y++
		}
	}

	banyakjadwal := memori

	//Pengolahan + Output
	//Tanggal	//Hanya disini saya pakai Insertion soalnya untuk menitnya angkanya terlalu besar
	k = 0
	totaltahun = 0 //Untuk menentukan posisi tahun dan bulan yang berbeda
	for k < banyakjadwal {
		banyaktahun := 0
		for l := k; l < banyakjadwal; l++ {
			if (temparr[k].time.bulan == temparr[l].time.bulan) && (temparr[k].time.tahun == temparr[l].time.tahun) {
				banyaktahun++
			}
		}

		totaltahun = totaltahun + banyaktahun

		for i := k; i < totaltahun-1; i++ {
			j := i + 1
			for j > 0 && temparr[j].time.tanggal < temparr[j-1].time.tanggal {
				temp := temparr[j]
				temparr[j] = temparr[j-1]
				temparr[j-1] = temp
				j--
			}
		}

		k = totaltahun
	}

	//Menit	//Kembali menggunakan Selection
	k = 0
	totaltahun = 0 //Untuk menentukan posisi tahun yang berbeda
	for k < banyakjadwal {
		banyaktahun := 0
		for l := k; l < banyakjadwal; l++ {
			if (temparr[k].time.tanggal == temparr[l].time.tanggal) && (temparr[k].time.bulan == temparr[l].time.bulan) && (temparr[k].time.tahun == temparr[l].time.tahun) {
				banyaktahun++
			}
		}

		totaltahun = totaltahun + banyaktahun

		for i := k; i < totaltahun-1; i++ {
			min := i
			for j := i; j < totaltahun; j++ {
				if convjammenitmulai(temparr[min]) > convjammenitmulai(temparr[j]) {
					min = j
				}
			}

			temp := temparr[i]
			temparr[i] = temparr[min]
			temparr[min] = temp
		}

		k = totaltahun
	}

	fmt.Println("Jadwal yang sudah dibuat :", y)
	fmt.Println("Jadwal yang sudah dibuat terurut")
	for x := 0; x < y; x++ {
		fmt.Println("================================")
		tampilkanada(temparr[x])
	}

}

func tampilkansudah(n int) {
	var (
		temparr                                  [1000]Event
		tahunnow, bulannow, tanggalnow, menitnow int
		counterbanyak                            int
	)

	cleartime(1)
	Header()
	AccountDetails()

	tahunnow = time.Now().Year()
	bulannow = convMonthInt(time.Now().Month().String())
	tanggalnow = time.Now().Day()
	menitnow = (time.Now().Hour() * 60) + (time.Now().Minute())
	counterbanyak = 0

	//Masukkan ke temparr[]
	for i := 0; i < n; i++ {
		//Tahun
		if jadwal[i].time.tahun < tahunnow {
			temparr[counterbanyak] = jadwal[i]
			counterbanyak++
		} else if jadwal[i].time.tahun == tahunnow {
			//Bulan
			if jadwal[i].time.bulan < bulannow {
				temparr[counterbanyak] = jadwal[i]
				counterbanyak++
			} else if jadwal[i].time.bulan == bulannow {
				//Tanggal
				if jadwal[i].time.tanggal < tanggalnow {
					temparr[counterbanyak] = jadwal[i]
					counterbanyak++
				} else if jadwal[i].time.tanggal == tanggalnow {
					//Menit
					if convjammenitselesai(jadwal[i]) < menitnow {
						temparr[counterbanyak] = jadwal[i]
						counterbanyak++
					}
				}
			}
		}
	}

	tampilkanmulai(counterbanyak, temparr)
}

//Hapus Jadwal
func convjammenitselesai(p Event) int {
	return ((p.time.jam_akhir * 60) + p.time.menit_akhir)
}

func hapusjadwal(n *int) {
	var (
		banyakjadwal, index int
		id                  string
	)
	banyakjadwal = *n
	tahunnow := time.Now().Year()
	bulannow := convMonthInt(time.Now().Month().String())
	tanggalnow := time.Now().Day()
	menitnow := (time.Now().Hour() * 60) + (time.Now().Minute())

	cleartime(1)
	Header()
	AccountDetails()
	fmt.Print("Masukkan ID yang ingin dihapus (ketik -1 untuk kembali) : ")
	fmt.Scanln(&id)

	//Cari Index dan ID
	index = -1
	if id != "-1" {
		for x := 0; x < banyakjadwal; x++ {
			if id == jadwal[x].id {
				index = x
			}
		}
	} else {
		goto akhir
	}

	for index == -1 {
		CallClear()
		Header()
		AccountDetails()
		fmt.Println("ID yang dicari tidak ada")
		fmt.Print("Masukkan ID yang ingin dihapus  (ketik -1 untuk kembali) : ")
		fmt.Scanln(&id)

		index = -1
		if id != "-1" {
			for x := 0; x < banyakjadwal; x++ {
				if id == jadwal[x].id {
					index = x
				}
			}
		} else {
			goto akhir
		}
	}

	//Menghapus Index dan ID
	if jadwal[index].time.tahun > tahunnow {
		fmt.Println("Jadwal belum terlewat")
		cleartime(3)
	} else {
		if jadwal[index].time.bulan > bulannow {
			fmt.Println("Jadwal belum terlewat")
			cleartime(3)
		} else {
			if jadwal[index].time.tanggal > tanggalnow {
				fmt.Println("Jadwal belum terlewat")
				cleartime(3)
			} else {
				if convjammenitselesai(jadwal[index]) > menitnow {
					fmt.Println("Jadwal belum terlewat")
					cleartime(3)
				} else {
					for y := index + 1; y < banyakjadwal; y++ {
						jadwal[y-1] = jadwal[y]
					}

					*n = banyakjadwal - 1
				}
			}
		}
	}

akhir: //Pointer Jumper
}

//Main
func main() {
	/*
		Header() menampilan header
		AccountDetails() memnampilkan account details
		CallClear() clearscreen
		CreateAcc(string, string) membuat akun
		cleartime(x int) pause screen selama x detik
	*/

	var (
		pilihan, tempbulan, temptahun int
		n                             int //memori_create_jadwal ini untuk menunjukkan sudah berapa banyak jadwal yang dibuat
		idcounter                     int
		cekusername, cekpassword      string
		usernamelama                  string
	)
	n = 0
	idcounter = 0
	CallClear()
	//Buat Akun
menuawal:
	cekusername = ""
	cekpassword = ""
	menuawal(&pilihan)
	if pilihan == 1 {
		CallClear()
		CreateAcc(&user.nama, &user.password)
		cleartime(1)
		Header() // Menampilkan header
		fmt.Println("Account Created")
		fmt.Println("==================")
		fmt.Println("Username :", user.nama)
		fmt.Println("Password :", user.password)
		cleartime(2)
		goto menuawal
	} else if pilihan == 2 {
		CallClear()
		Header()
		fmt.Println("Login : ")
		fmt.Print("Username : ")
		fmt.Scanln(&cekusername)
		for cekusername == "" {
			CallClear()
			Header()
			fmt.Println("Login : ")
			fmt.Println("Username tidak boleh kosong")
			fmt.Print("Username : ")
			fmt.Scanln(&cekusername)
		}
		fmt.Print("Password : ")
		fmt.Scanln(&cekpassword)
		for cekpassword == "" {
			CallClear()
			Header()
			fmt.Println("Login : ")
			fmt.Println("Username : ", cekusername)
			fmt.Println("Password tidak boleh kosong")
			fmt.Print("Password : ")
			fmt.Scanln(&cekpassword)
		}

		if (usernamelama != user.nama) && (cekusername == user.nama) && (cekpassword == user.password) {
			n = 0
			idcounter = 0
		}

		if (cekusername == user.nama) && (cekpassword == user.password) {
			CallClear()
			Header()
			fmt.Println("Account Logged In")
			fmt.Println("==================")
			fmt.Println("Username :", user.nama)
			fmt.Println("Password :", user.password)
			cleartime(1)
		menu: //Pointer jumper
			CallClear()
			menu(&pilihan, n)

			if pilihan == 1 {
				CallClear()
				irisan := 0
				Header()
				AccountDetails()
				fmt.Println("Nama & Tempat Jadwal : ")
				namatempat(&jadwal[n].tempat, &jadwal[n].eventname)
				CallClear()
				Header()
				AccountDetails()

				//Input Mulai
				fmt.Println("Waktu mulai : ")
				fmt.Println("=====================")
				InTahun(time.Now().Year(), &jadwal[n].time.tahun)
				InBulan(&jadwal[n].time.bulan, jadwal[n].time.tahun, time.Now().Year(), convMonthInt(time.Now().Month().String()))
				InTanggalAwal(&jadwal[n].time.tanggal, jadwal[n].time.bulan, jadwal[n].time.tahun)
				InJamAwal(&jadwal[n].time.jam_mulai, &jadwal[n].time.menit_mulai, jadwal[n].time.bulan, jadwal[n].time.tanggal, jadwal[n].time.tahun)
				CallClear()
				Header()
				AccountDetails()

				//Input Selesai
				fmt.Println("Waktu mulai : ", jadwal[n].time.tanggal, jadwal[n].time.bulan, jadwal[n].time.tahun, "    jam [hh/mm] :", jadwal[n].time.jam_mulai, ":", jadwal[n].time.menit_mulai)
				fmt.Println("Waktu selesai : ")
				fmt.Println("=====================")
				InJamAkhir(&jadwal[n].time.jam_akhir, &jadwal[n].time.menit_akhir, jadwal[n].time.tahun, jadwal[n].time.bulan, jadwal[n].time.jam_mulai, jadwal[n].time.menit_mulai, jadwal[n].time.tanggal)
				cleartime(2)

				//Cek Beririsan
				for j := 0; j < n; j++ {
					if beririsan(jadwal[n], jadwal[j]) {
						irisan = irisan + 1
					}
				}

				for irisan > 2 {
					irisan = 0
					CallClear()
					Header()
					AccountDetails()
					fmt.Println("Buat Jadwal : ")
					fmt.Println("=====================")
					fmt.Println("Jadwal tidak valid, hanya boleh 3 jadwal yang beririsan")
					fmt.Println("Nama & Tempat Jadwal : ")
					namatempat(&jadwal[n].tempat, &jadwal[n].eventname)
					CallClear()
					Header()
					AccountDetails()

					//Input Mulai
					fmt.Println("Waktu mulai : ")
					fmt.Println("=====================")
					InTahun(time.Now().Year(), &jadwal[n].time.tahun)
					InBulan(&jadwal[n].time.bulan, jadwal[n].time.tahun, time.Now().Year(), convMonthInt(time.Now().Month().String()))
					InTanggalAwal(&jadwal[n].time.tanggal, jadwal[n].time.bulan, jadwal[n].time.tahun)
					InJamAwal(&jadwal[n].time.jam_mulai, &jadwal[n].time.menit_mulai, jadwal[n].time.bulan, jadwal[n].time.tanggal, jadwal[n].time.tahun)
					CallClear()
					Header()
					AccountDetails()

					//Input Selesai
					fmt.Println("Waktu mulai : ", jadwal[n].time.tanggal, jadwal[n].time.bulan, jadwal[n].time.tahun, "    jam [hh/mm] :", jadwal[n].time.jam_mulai, jadwal[n].time.menit_mulai)
					fmt.Println("Waktu selesai : ")
					fmt.Println("=====================")
					InJamAkhir(&jadwal[n].time.jam_akhir, &jadwal[n].time.menit_akhir, jadwal[n].time.tahun, jadwal[n].time.bulan, jadwal[n].time.jam_mulai, jadwal[n].time.menit_mulai, jadwal[n].time.tanggal)
					cleartime(2)

					//Cek Beririsan
					for j := 0; j < n; j++ {
						if beririsan(jadwal[n], jadwal[j]) {
							irisan = irisan + 1
						}
					}

				}
				Header()
				AccountDetails()
				makeID(idcounter, &jadwal[n].id)
				JadwalCreated(n, jadwal[n].eventname, jadwal[n].tempat, jadwal[n].id, jadwal[n].time.tanggal, jadwal[n].time.bulan, jadwal[n].time.tahun, jadwal[n].time.jam_mulai, jadwal[n].time.jam_akhir, jadwal[n].time.menit_mulai, jadwal[n].time.menit_akhir)

				n++
				idcounter++
				goto menu //Jumper
			} else if pilihan == 2 {
				if n > 0 {
					hapusjadwal(&n)
				} else {
					CallClear()
					Header()
					AccountDetails()
					fmt.Println("Jadwal belum dibuat")
					fmt.Println("Akan kembali ke menu dalam 2 detik")
					cleartime(2)
				}
				// Berdasarkan ID
				goto menu //Jumper
			} else if pilihan == 3 {
				counter := 0
				CallClear()
				Header()
				AccountDetails()
				menulihatjadwal(&pilihan)

				if pilihan == 1 { //Tangagal
					tampilkantanggal(n)
					fmt.Print("Untuk kembali ke menu tekan (1) : ")
					fmt.Scanln(&pilihan)
					for pilihan != 1 {
						fmt.Println("Pilihan tidak valid")
						fmt.Print("Untuk kembali ke menu tekan (1) : ")
						fmt.Scanln(&pilihan)
					}
				} else if pilihan == 2 { //Bulan
					CallClear()
					Header()
					AccountDetails()
					fmt.Print("Masukkan bulan yang anda inginkan (1-12) : ")
					fmt.Scanln(&tempbulan)
					for tempbulan < 1 || tempbulan > 12 {
						fmt.Println("Masukkan bulan tidak valid")
						fmt.Print("Masukkan bulan yang anda inginkan (1-12) : ")
						fmt.Scanln(&tempbulan)

					}
					fmt.Print("Masukkan tahun yang anda inginkan : ")
					fmt.Scanln(&temptahun)
					CallClear()
					Header()
					AccountDetails()
					fmt.Println("Jadwal pada bulan ke-", tempbulan, " tahun ", temptahun)
					fmt.Println("===============================")
					tampilkanbulantahun(tempbulan, temptahun, n)

					fmt.Print("Untuk kembali ke menu tekan (1) : ")
					fmt.Scanln(&pilihan)
					for pilihan != 1 {
						fmt.Println("Pilihan tidak valid")
						fmt.Print("Untuk kembali ke menu tekan (1) : ")
						fmt.Scanln(&pilihan)
					}

				} else if pilihan == 3 { //Menampilkan semua berdasarkan waktu mulai
					CallClear()
					cleartime(3)
					Header()
					AccountDetails()
					tampilkanmulai(n, jadwal)
					fmt.Print("Untuk kembali ke menu tekan (1) : ")
					fmt.Scanln(&pilihan)
					for pilihan != 1 {
						fmt.Println("Pilihan tidak valid")
						fmt.Print("Untuk kembali ke menu tekan (1) : ")
						fmt.Scanln(&pilihan)
					}
				} else if pilihan == 4 { //Urutan yang sudah selesai
					tampilkansudah(n)
					fmt.Print("Untuk kembali ke menu tekan (1) : ")
					fmt.Scanln(&pilihan)
					for pilihan != 1 {
						fmt.Println("Pilihan tidak valid")
						fmt.Print("Untuk kembali ke menu tekan (1) : ")
						fmt.Scanln(&pilihan)
					}
				} else if pilihan == 5 { // Yang ada
					fmt.Println()
					fmt.Println("Jadwal yang sudah dibuat")
					fmt.Println("=========================")
					for counter < n {
						fmt.Println("================================")
						fmt.Println("Jadwal ke-", counter+1)
						fmt.Println("================================")
						tampilkanada(jadwal[counter])
						counter++
					}
					fmt.Print("Untuk kembali ke menu tekan (1) : ")
					fmt.Scanln(&pilihan)
					for pilihan != 1 {
						fmt.Println("Pilihan tidak valid")
						fmt.Print("Untuk kembali ke menu tekan (1) : ")
						fmt.Scanln(&pilihan)
					}
				} else if pilihan == 6 { //Goto menu
					goto menu
				}
				goto menu //Jumper
			} else if pilihan == 4 {
				fmt.Println("Apakah anda yakin ingin keluar? tekan 0 untuk iya tekan 1 untuk tidak")
				fmt.Scanln(&pilihan)
				if pilihan == 0 {
					usernamelama = user.nama
					fmt.Println("Program akan keluar dalam 3 deitk")
					cleartime(3)
					goto menuawal
				} else {
					goto menu
				}
			}
		} else {
			CallClear()
			Header()
			fmt.Println("Akun salah, akan kembali ke menu awal dalam 2 detik")
			cleartime(2)
			goto menuawal
		}
	} else if pilihan == 3 {
		fmt.Println("Apakah anda yakin ingin keluar? tekan 0 untuk iya tekan 1 untuk tidak")
		fmt.Scanln(&pilihan)
		if pilihan == 0 {
			fmt.Println("Program akan keluar dalam 3 deitk")
			cleartime(3)
		} else {
			goto menuawal
		}
	}

}
