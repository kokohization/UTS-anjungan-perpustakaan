package main

import (
	"fmt"
	"strings"
)

var (
	riwayat []string

	// var buat validasi akun
	inputusername, inputpassword string

	// var informasi akun
	username  = "fazle"
	password  = "2406424335"
	pekerjaan = "Mahasiswa"
	institusi = "Universitas Indonesia"

	pilihan int
)

// var buat buku
type buku struct {
	namabuku   string
	jumlahbuku int
}

var daftarbuku []buku

func main() {

	// isi dari variable buku saat awal
	daftarbuku = []buku{
		{namabuku: "Pemrograman", jumlahbuku: 10},
		{namabuku: "Film", jumlahbuku: 5},
		{namabuku: "Printing", jumlahbuku: 20},
	}

	fmt.Println("=========================================================")
	fmt.Println("=       Selamat Datang di Anjungan Perpustakaan!        =")
	fmt.Println("=========================================================")

	// validasi username & password
	fmt.Println("Masukkan Username & Password Anda!")

	fmt.Print("Username : ")
	fmt.Scan(&inputusername)

	fmt.Print("Password : ")
	fmt.Scan(&inputpassword)

	if !validasiakun() {
		fmt.Println("Username atau Password salah! Program akan dihentikan.")
		fmt.Println("=========================================================")

		return
	}

	fmt.Println("     Username & Password Benar, Selamat Datang!     ")

	// menu
	for {

		fmt.Println("=========================================================")
		fmt.Println("=                       MENU                            =")
		fmt.Println("=========================================================")

		fmt.Println("1. Lihat Informasi Pengguna Program")
		fmt.Println("2. Lihat Daftar Buku")
		fmt.Println("3. Tambah Daftar Buku")
		fmt.Println("4. Tambah Peminjaman Buku")
		fmt.Println("5. Riwayat Peminjaman Buku")
		fmt.Println("6. Keluar dari Program")

		fmt.Println("Masukkan Pilihan Anda : ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			lihatinformasipengguna()
		case 2:
			lihatdaftarbuku()
		case 3:
			tambahdaftarbuku()
		case 4:
			pinjambuku()
		case 5:
			riwayatpeminjaman()
		case 6:
			fmt.Println("=========================================================")
			fmt.Println("=           Terimakasih Telah Menggunakan Kami!         =")
			fmt.Println("-------------------------xoxo----------------------------")
			return
		default:
			fmt.Println("------Input Salah!, Tolong Masukkan Input yang Benar!----")
			continue
		}
	}
}

// func buat validasi password dan username
func validasiakun() bool {
	return inputusername == username && inputpassword == password
}

// func buat lihat informasi pengguna
func lihatinformasipengguna() {
	fmt.Println("=========================================================")
	fmt.Println("=                 INFORMASI PENGGUNA                    =")
	fmt.Println("=========================================================")

	fmt.Println("Username  : ", username)
	fmt.Println("Password  : ", password)
	fmt.Println("Pekerjaan : ", pekerjaan)
	fmt.Println("Institusi : ", institusi)

}

// func untuk melihat daftar buku-buku yang ada
func lihatdaftarbuku() {
	fmt.Println("=========================================================")
	fmt.Println("=                    DAFTAR BUKU                        =")
	fmt.Println("=========================================================")

	for _, buku := range daftarbuku {
		fmt.Println("Nama Buku:", buku.namabuku)
		fmt.Println("Jumlah Buku:", buku.jumlahbuku)
		fmt.Println("----------------------------------")
	}
}

// func buat user nambahin buku mereka
func tambahdaftarbuku() {
	var (
		namabuku   string
		jumlahbuku int
	)

	fmt.Println("=========================================================")
	fmt.Println("=                 TAMBAHKAN BUKU BARU                   =")
	fmt.Println("=========================================================")

	for {
		fmt.Print("Masukkan Nama Buku Baru : ")
		fmt.Scan(&namabuku)

		fmt.Print("Masukkan Jumlah Buku : ")
		fmt.Scan(&jumlahbuku)

		if jumlahbuku <= 0 {
			fmt.Println("Input Jumlah Tidak Valid!")
			fmt.Println("----------------------------------")

			continue
		}

		//buat ngehapus spasi & jadiin semua inputnya lowercase
		namabuku = strings.TrimSpace(namabuku)
		//buat push input user ke dalam array daftar buku
		daftarbuku = append(daftarbuku, buku{namabuku: namabuku, jumlahbuku: jumlahbuku})
		break
	}

	fmt.Println("Buku Berhasil Ditambahkan!")
	fmt.Println("----------------------------------")
}

// func buat user meminjam buku
func pinjambuku() {
	var (
		namabuku      string
		jumlahpinjam  int
		bukuditemukan bool
	)

	fmt.Println("=========================================================")
	fmt.Println("=                   PEMINJAMAN BUKU                     =")
	fmt.Println("=========================================================")

	fmt.Println("----BUKU YANG TERSEDIA----")

	//buat memanggil ulang array daftar buku kepada user
	for _, buku := range daftarbuku {
		fmt.Println("Nama Buku   :", buku.namabuku)
		fmt.Println("Jumlah Buku :", buku.jumlahbuku)
		fmt.Println("----------------------------------")
	}

	// loop untuk input dan validasi input user
	for {
		fmt.Print("Masukkan Nama Buku : ")
		fmt.Scan(&namabuku)
		namabuku = strings.TrimSpace(namabuku)

		// bool untuk validasi apakah input user tersedia dengan yang ada di array
		bukuditemukan = false

		// mencari buku dalam daftar
		for i := range daftarbuku {
			if strings.EqualFold(daftarbuku[i].namabuku, namabuku) {
				bukuditemukan = true

				fmt.Print("Masukkan Jumlah Buku : ")
				fmt.Scan(&jumlahpinjam)

				// memeriksa apakah stok cukup
				if daftarbuku[i].jumlahbuku >= jumlahpinjam {
					daftarbuku[i].jumlahbuku -= jumlahpinjam

					//push input user kedalam array riwayat
					riwayat = append(riwayat, fmt.Sprintf("Peminjaman Buku %s, Sebanyak %d Buah", namabuku, jumlahpinjam))

					fmt.Printf("Berhasil Meminjam Buku %s, Sebanyak %d\n", namabuku, jumlahpinjam)
					fmt.Println("----------------------------------")
					return
				} else {
					fmt.Println("======================================================================")
					fmt.Println("=      Jumlah yang ingin dipinjam melebihi stok yang tersedia.       =")
					fmt.Println("======================================================================")
					break
				}
			}
		}

		// jika buku tidak ditemukan setelah loop pencarian
		if !bukuditemukan {
			fmt.Println("======================================================================")
			fmt.Printf("=     Buku %s tidak ditemukan dalam sistem. Silakan coba lagi!   =\n", namabuku)
			fmt.Println("======================================================================")
		}
	}
}

// dunc buat lihat riwayat peminjaman yang dilakukan sebelumnya
func riwayatpeminjaman() {
	fmt.Println("=========================================================")
	fmt.Println("=                  RIWAYAT PEMINJAMAN                   =")
	fmt.Println("=========================================================")

	for _, riwayat := range riwayat {
		fmt.Println(" -", riwayat)
		fmt.Println("----------------------------------")
	}
}
