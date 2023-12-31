package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Definisi data buku
type Book struct {
	Title  string
	Author string
	Year   int
	Stock  int
}

// Definisi data peminjaman buku
type Borrowing struct {
	Book       Book
	Borrower   string
	BorrowDate time.Time
	ReturnDate time.Time
}

// Slice untuk menyimpan data buku dan peminjaman
var books []Book
var borrowings []Borrowing

// Fungsi untuk menampilkan menu
func displayMenu() {
	fmt.Println("┌────────────────────────────────────────────────┐")
	fmt.Println("|              Program Perpustakaan              |")
	fmt.Println("|================================================|")
	fmt.Println("| 1. Tambah Buku                                 |")
	fmt.Println("| 2. Cari Buku                                   |")
	fmt.Println("| 3. Hapus Buku                                  |")
	fmt.Println("| 4. Tambah Peminjaman                           |")
	fmt.Println("| 5. Cari Peminjaman                             |")
	fmt.Println("| 6. Hapus Peminjaman                            |")
	fmt.Println("| 7. Tutorial menjalankan program dengan argumen |")
	fmt.Println("| 8. Keluar                                      |")
	fmt.Println("└────────────────────────────────────────────────┘")
	fmt.Print("Pilih Menu [1-8]: ")

}

// Fungsi untuk menambahkan buku baru
func addBook() {
	var newBook Book

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("┌────────────────────────────────────────────────┐")
	fmt.Println("|           Masukkan informasi buku:             |")
	fmt.Println("└────────────────────────────────────────────────┘")
	fmt.Print("Judul: ")
	scanner.Scan()
	newBook.Title = scanner.Text()
	fmt.Print("Pengarang: ")
	scanner.Scan()
	newBook.Author = scanner.Text()
	fmt.Print("Tahun Terbit: ")
	fmt.Scanln(&newBook.Year)
	fmt.Print("Stok: ")
	fmt.Scanln(&newBook.Stock)

	books = append(books, newBook)
	saveData()
	fmt.Println("┌────────────────────────────────────────────────┐")
	fmt.Println("|          Buku berhasil ditambahkan.            |")
	fmt.Println("└────────────────────────────────────────────────┘")
}

// Fungsi untuk mencari buku
func searchBook() {
	var bookTitle string

	fmt.Print("Masukkan Judul Buku: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	bookTitle = scanner.Text()

	matchingBooks := []Book{}
	// Looping data buku sampai mendapat buku yang diinginkan
	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(bookTitle)) {
			matchingBooks = append(matchingBooks, book)
		}
	}
	// Cek jika ada buku yang sesuai dengan judul yang diinput/ mengandung keyword dari judul yang diinput
	if len(matchingBooks) > 0 {

		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|                Buku ditemukan!                 |")
		fmt.Println("└────────────────────────────────────────────────┘")
		for _, book := range matchingBooks {
			fmt.Println("")
			fmt.Println("------------------------------")
			fmt.Println("")
			fmt.Printf("Judul Buku: %s\n", book.Title)
			fmt.Printf("Pengarang: %s\n", book.Author)
			fmt.Printf("Tahun Terbit: %d\n", book.Year)
			fmt.Printf("Stok: %d\n", book.Stock)
			fmt.Println("")
			fmt.Println("------------------------------")
		}
	} else {
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|              Buku tidak ditemukan.             |")
		fmt.Println("└────────────────────────────────────────────────┘")
	}
}

// Fungsi untuk menghapus buku berdasarkan pilihan pengguna
func deleteBook() {

	fmt.Println("Daftar Buku:")
	for i, book := range books {
		fmt.Printf("%d. %s\n", i+1, book.Title)
	}

	// Menginput pilihan buku yang ingin dihapus
	var indexToDelete int
	fmt.Print("Pilih Buku yang akan dihapus [1-", len(books), "]: ")
	fmt.Scanln(&indexToDelete)

	if indexToDelete < 1 || indexToDelete > len(books) {
		fmt.Println("Pilihan tidak valid. Operasi dibatalkan.")
		return
	}

	// Menyesuaikan indeks agar sesuai dengan indeks slice (perhitungan komputer mulai dari 0)
	indexToDelete--

	// Menghapus buku yang diinginkan dari slice
	books = append(books[:indexToDelete], books[indexToDelete+1:]...)

	saveData()
	fmt.Println("Buku berhasil dihapus!")
}

// Fungsi untuk menambahkan peminjaman buku
func addBorrowing() {
	fmt.Println("Daftar Buku:")
	for i, book := range books {
		fmt.Printf("%d. %s\n", i+1, book.Title)
	}

	var bookIndex int
	fmt.Print("Pilih buku yang akan dipinjam [1-", len(books), "]: ")
	fmt.Scanln(&bookIndex)

	if bookIndex < 1 || bookIndex > len(books) {
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|    Pilihan tidak valid. Operasi dibatalkan.    |")
		fmt.Println("└────────────────────────────────────────────────┘")
		return
	}

	// Menyesuaikan indeks agar sesuai dengan indeks slice (perhitungan komputer dimulai dari 0)
	bookIndex--

	selectedBook := books[bookIndex]

	fmt.Println("┌────────────────────────────────────────────────┐")
	fmt.Printf("  Buku yang dipinjam adalah %s\n", selectedBook.Title)
	fmt.Println("└────────────────────────────────────────────────┘")

	var borrowerName string
	fmt.Print("Nama Peminjam: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	borrowerName = scanner.Text()

	var borrowDuration int
	fmt.Print("Masukkan jumlah hari peminjaman (1-30): ")
	fmt.Scanln(&borrowDuration)

	if borrowDuration < 1 || borrowDuration > 30 {
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|   Jumlah durasi peminjaman tidak valid.        |")
		fmt.Println("|   Maksimal peminjaman 30 hari.                 |")
		fmt.Println("└────────────────────────────────────────────────┘")
		return
	}

	if selectedBook.Stock > 0 {
		// Jika stok tersedia, lanjutkan proses peminjaman
		borrowDate := time.Now()
		returnDate := borrowDate.AddDate(0, 0, borrowDuration)

		newBorrowing := Borrowing{
			Book:       selectedBook,
			Borrower:   borrowerName,
			BorrowDate: borrowDate,
			ReturnDate: returnDate,
		}

		borrowings = append(borrowings, newBorrowing)

		// Mengurangi jumlah stok buku yang dipinjam
		books[bookIndex].Stock--

		saveData()
		fmt.Printf("Tanggal Pinjam: %s\n", borrowDate.Format("2006-01-02"))
		fmt.Printf("Tanggal Kembali: %s\n", returnDate.Format("2006-01-02"))
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|              Peminjaman berhasil.              |")
		fmt.Println("└────────────────────────────────────────────────┘")
	} else {
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|    Buku tidak tersedia atau tidak ditemukan.   |")
		fmt.Println("└────────────────────────────────────────────────┘")
	}
}

// Fungsi untuk mencari peminjaman
func searchBorrowing() {
	var bookTitle string

	fmt.Print("Masukkan Judul Buku: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	bookTitle = scanner.Text()

	// deklarasi slice matchingBorrowings dengan isi dari slice nya adalah struct dari Borrowing
	var matchingBorrowings []Borrowing

	// Looping data borrowings sampai mendapat data judul buku yang diinginkan lalu masukkan data ke slice matchingBorrowings
	for _, borrowing := range borrowings {
		if borrowing.Book.Title == bookTitle {
			matchingBorrowings = append(matchingBorrowings, borrowing)
		}
	}

	// mengecek apakah ada peminjaman pada buku tersebut
	if len(matchingBorrowings) > 0 {
		fmt.Println("Peminjaman ditemukan!")

		// Menampilkan semua user yang meminjam buku yang disebutkan
		fmt.Printf("Judul Buku: %s\n", matchingBorrowings[0].Book.Title)
		fmt.Println("Daftar Peminjam:")
		for i, borrowing := range matchingBorrowings {
			fmt.Printf("%d. %s\n", i+1, borrowing.Borrower)
		}

		var borrowingIndex int
		fmt.Print("Pilih peminjam yang akan ditampilkan [1-", len(matchingBorrowings), "]: ")
		fmt.Scanln(&borrowingIndex)

		if borrowingIndex < 1 || borrowingIndex > len(matchingBorrowings) {
			fmt.Println("┌────────────────────────────────────────────────┐")
			fmt.Println("|    Pilihan tidak valid. Operasi dibatalkan.    |")
			fmt.Println("└────────────────────────────────────────────────┘")
			return
		}

		// Menyesuaikan indeks agar sesuai dengan indeks slice (perhitungan komputer mulai dari 0)
		borrowingIndex--

		selectedBorrowing := matchingBorrowings[borrowingIndex]

		fmt.Printf("Nama Peminjam: %s\n", selectedBorrowing.Borrower)
		fmt.Printf("Tanggal Pinjam: %s\n", selectedBorrowing.BorrowDate.Format("2006-01-02"))
		fmt.Printf("Tanggal Kembali: %s\n", selectedBorrowing.ReturnDate.Format("2006-01-02"))
	} else {
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|           Peminjaman tidak ditemukan.          |")
		fmt.Println("└────────────────────────────────────────────────┘")
	}
}

// Fungsi untuk menghapus peminjaman
func deleteBorrowing() {
	displayBorrowings()

	if len(borrowings) == 0 {
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|           Tidak ada peminjaman tersedia.       |")
		fmt.Println("└────────────────────────────────────────────────┘")
		return
	}

	var borrowingIndex int
	fmt.Print("Pilih Peminjaman yang akan dihapus [1-", len(borrowings), "]: ")
	fmt.Scanln(&borrowingIndex)

	if borrowingIndex < 1 || borrowingIndex > len(borrowings) {
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|    Pilihan tidak valid. Operasi dibatalkan.    |")
		fmt.Println("└────────────────────────────────────────────────┘")
		return
	}

	// Menyesuaikan indeks agar sesuai dengan indeks slice (perhitungan komputer mulai dari 0)
	borrowingIndex--

	// Menyimpan data buku yang akan dihapus peminjaman nya
	deletedBorrowing := borrowings[borrowingIndex]

	// Menambah stok buku dari buku yang sudah dikembalikan (sebelumnya dipinjam)
	for i, book := range books {
		if book.Title == deletedBorrowing.Book.Title {
			books[i].Stock++
			break
		}
	}

	// Menghapus index borrowing yang dipilih dari slice
	borrowings = append(borrowings[:borrowingIndex], borrowings[borrowingIndex+1:]...)

	fmt.Printf("Peminjaman untuk buku %s oleh %s berhasil dihapus!\n", deletedBorrowing.Book.Title, deletedBorrowing.Borrower)
	saveData()
}

// Fungsi untuk menampilkan peminjaman
func displayBorrowings() {
	fmt.Println("Daftar Peminjaman:")
	for i, borrowing := range borrowings {
		fmt.Printf("%d. %s - %s\n", i+1, borrowing.Book.Title, borrowing.Borrower)
	}
}

// Fungsi untuk menyimpan data buku dan peminjaman ke file JSON
func saveData() {
	// Simpan data buku
	bookData, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|         Error saat menyimpan data buku:        |", err)
		fmt.Println("└────────────────────────────────────────────────┘")
		return
	}
	err = os.WriteFile("books.json", bookData, 0644)
	if err != nil {
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|         Error saat menyimpan data buku:        |", err)
		fmt.Println("└────────────────────────────────────────────────┘")
		return
	}

	// Simpan data peminjaman
	borrowData, err := json.MarshalIndent(borrowings, "", "  ")
	if err != nil {
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|     Error saat menyimpan data peminjaman:      |", err)
		fmt.Println("└────────────────────────────────────────────────┘")
		return
	}
	err = os.WriteFile("borrowings.json", borrowData, 0644)
	if err != nil {
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|     Error saat menyimpan data peminjaman:      |", err)
		fmt.Println("└────────────────────────────────────────────────┘")
		return
	}
}

const booksFilePath = "books.json"
const borrowingsFilePath = "borrowings.json"

func main() {
	// Load data dari file json
	loadData()

	if len(os.Args) > 1 {
		// Jika user menggunakan argumen, maka proses data jika user menginput argumen sesuai format dibawah
		command := flag.String("command", "", "Menu command")

		title := flag.String("judul", "", "Judul buku")
		author := flag.String("pengarang", "", "Penulis buku")
		year := flag.Int("tahun", 0, "Tahun terbit buku")
		stock := flag.Int("stok", 0, "Stok buku")

		borrower := flag.String("peminjam", "", "Nama peminjam")
		borrowDateStr := flag.String("tgl_pinjam", "", "Tanggal peminjaman (format: YYYY-MM-DD)")
		returnDateStr := flag.String("tgl_kembali", "", "Tanggal pengembalian (format: YYYY-MM-DD)")

		flag.Parse()

		// Memproses argumen yang digunakan oleh user
		switch *command {
		case "tambah_buku":
			addBookFromArgs(*title, *author, *year, *stock)
		case "cari_buku":
			searchBookFromArgs(*title)
		case "hapus_buku":
			deleteBookFromArgs(*title)
		case "tambah_pinjam":
			addBorrowingFromArgs(*title, *borrower, *borrowDateStr, *returnDateStr)
		case "cari_pinjam":
			searchBorrowingFromArgs(*title)
		case "hapus_pinjam":
			deleteBorrowingFromArgs(*title, *borrower)
		default:
			displayTutorialArgs()
		}

		// Save data argumen ke files
		saveData()
	} else {
		// Jika user tidak memberikan perintah argumen, maka masuk ke menu perpustakaan
		for {
			displayMenu()

			var choice int
			fmt.Scan(&choice)

			switch choice {
			case 1:
				addBook()
			case 2:
				searchBook()
			case 3:
				deleteBook()
			case 4:
				addBorrowing()
			case 5:
				searchBorrowing()
			case 6:
				deleteBorrowing()
			case 7:
				displayTutorialArgs()
			case 8:
				// Save data sebelum keluar dari program
				saveData()
				os.Exit(0)
			default:
				fmt.Println("┌─────────────────────────────────────────────────────────────────┐")
				fmt.Println("|           Pilihan tidak valid. Silahkan pilih kembali.          |")
				fmt.Println("└─────────────────────────────────────────────────────────────────┘")
			}
		}
	}
}

// fungsi untuk load data dari file json
func loadData() {
	// Load data buku
	booksData, err := os.ReadFile(booksFilePath)
	if err == nil {
		err = json.Unmarshal(booksData, &books)
		if err != nil {
			fmt.Println("┌────────────────────────────────────────────────┐")
			fmt.Println("|             Gagal memuat data buku:            |", err)
			fmt.Println("└────────────────────────────────────────────────┘")
		}
	}

	// Load data peminjaman
	borrowingsData, err := os.ReadFile(borrowingsFilePath)
	if err == nil {
		err = json.Unmarshal(borrowingsData, &borrowings)
		if err != nil {
			fmt.Println("┌────────────────────────────────────────────────┐")
			fmt.Println("|          Gagal memuat data peminjaman:         |", err)
			fmt.Println("└────────────────────────────────────────────────┘")
		}
	}
}

// fungsi menampilkan tutorial menjalankan program menggunakan argumen
func displayTutorialArgs() {
	fmt.Println("┌───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐")
	fmt.Println("| Tutorial menjalankan program dengan argumen:                                                                                                      |")
	fmt.Println("| Pilih antara tambah_buku, cari_buku, hapus_buku, tambah_pinjam, cari_pinjam, atau hapus_pinjam                                                    |")
	fmt.Println("|                                                                                                                                                   |")
	fmt.Println("| Contoh penggunaan:                                                                                                                                |")
	fmt.Println("| go run perpustakaan.go -command=\"tambah_buku\" -judul=\"Belajar Go\" -pengarang=\"Agus Kopling\" tahun=2020 stok=10                                    |")
	fmt.Println("| go run perpustakaan.go -command=\"cari_buku\" -judul=\"Belajar Go\"                                                                                   |")
	fmt.Println("| go run perpustakaan.go -command=\"hapus_buku\" -judul=\"Belajar Go\"                                                                                  |")
	fmt.Println("| go run perpustakaan.go -command=\"tambah_pinjam\" -judul=\"Belajar Go\" -peminjam=\"Agus Kopling\" -tgl_pinjam=2023-12-31 -tgl_kembali=2024-01-30       |")
	fmt.Println("| go run perpustakaan.go -command=\"cari_pinjam\" -judul=\"Belajar Go\"                                                                                 |")
	fmt.Println("| go run perpustakaan.go -command=\"hapus_pinjam\" -judul=\"Belajar Go\" -peminjam=\"Agus Kopling\"                                                       |")
	fmt.Println("└───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘")
}

// fungsi untuk tambah buku 1
func addBookFromArgs(title string, author string, year int, stock int) {
	if title == "" || author == "" || year == 0 || stock == 0 {
		fmt.Println("┌───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐")
		fmt.Println("|  Perintah tidak valid. contoh penggunaan:                                                                                 |")
		fmt.Println("|  go run perpustakaan.go -command=\"tambah_buku\" -judul=\"Belajar Go\" -pengarang=\"Agus Kopling\" tahun=2020 stok=10           |")
		fmt.Println("└───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘")
	} else {
		newBook := Book{
			Title:  title,
			Author: author,
			Year:   year,
			Stock:  stock,
		}

		books = append(books, newBook)
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|          Buku berhasil ditambahkan!            |")
		fmt.Println("└────────────────────────────────────────────────┘")
	}
}

// fungsi untuk cari buku 2
func searchBookFromArgs(title string) {
	foundBooks := searchBooks(title)

	if len(foundBooks) == 0 {
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|             Buku tidak ditemukan.              |")
		fmt.Println("└────────────────────────────────────────────────┘")
	} else {
		fmt.Println("Buku yang ditemukan:")
		for _, book := range foundBooks {
			fmt.Println("┌────────────────────────────────────────────────┐")
			fmt.Printf(" Judul: %s\n Penulis: %s\n Tahun: %d\n Stok: %d\n", book.Title, book.Author, book.Year, book.Stock)
			fmt.Println("└────────────────────────────────────────────────┘")
		}
	}
}

func searchBooks(title string) []Book {
	var foundBooks []Book

	for _, book := range books {
		// Periksa apakah judul cocok (case-insensitive)
		if title == "" || strings.Contains(strings.ToLower(book.Title), strings.ToLower(title)) {
			foundBooks = append(foundBooks, book)
		}
	}

	return foundBooks
}

// fungsi untuk hapus buku 3
func deleteBookFromArgs(title string) {
	// Mencari buku sesuai judul yang diinput user
	book, bookIndex := findBookByTitle(title)
	if book == nil {
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|             Buku tidak ditemukan.              |")
		fmt.Println("└────────────────────────────────────────────────┘")
		return
	}

	// Menghapus buku yang diinginkan dari slice
	books = append(books[:bookIndex], books[bookIndex+1:]...)

	saveData()
	fmt.Printf("Buku %s berhasil dihapus!\n", title)
}

// fungsi untuk tambah peminjam 4
func addBorrowingFromArgs(title string, borrower string, borrowDateStr string, returnDateStr string) {
	// Mengubah format date/tanggal dari string menjadi time.time
	borrowDate, err := time.Parse("2006-01-02", borrowDateStr)
	if err != nil {
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|     Format tanggal peminjaman tidak valid.     |")
		fmt.Println("└────────────────────────────────────────────────┘")
		return
	}

	returnDate, err := time.Parse("2006-01-02", returnDateStr)
	if err != nil {
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|    Format tanggal pengembalian tidak valid.    |")
		fmt.Println("└────────────────────────────────────────────────┘")
		return
	}

	// Mencari buku sesuai judul yang diinput user
	book, bookIndex := findBookByTitle(title)
	if book == nil {
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|             Buku tidak ditemukan.              |")
		fmt.Println("└────────────────────────────────────────────────┘")
		return
	}

	// Mengecek stok ketersediaan buku untuk dipinjam
	if book.Stock <= 0 {
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|               Stok buku habis.                 |")
		fmt.Println("└────────────────────────────────────────────────┘")
		return
	}

	// Menambahkan data peminjaman yang baru
	newBorrowing := Borrowing{
		Book:       *book,
		Borrower:   borrower,
		BorrowDate: borrowDate,
		ReturnDate: returnDate,
	}

	// memperbarui jumlah stok buku dan data peminjaman
	books[bookIndex].Stock--
	borrowings = append(borrowings, newBorrowing)

	fmt.Println("┌────────────────────────────────────────────────┐")
	fmt.Println("|        Peminjaman berhasil ditambahkan!        |")
	fmt.Println("└────────────────────────────────────────────────┘")
}

// fungsi untuk mencari buku sesuai judul yang diinput user
func findBookByTitle(title string) (*Book, int) {
	for i, book := range books {
		if strings.EqualFold(book.Title, title) {
			return &book, i
		}
	}
	return nil, -1
}

// fungsi untuk Cari peminjaman 5
func searchBorrowingFromArgs(title string) {
	//deklarasi slice matchingBorrowings dengan isi dari slice nya adalah struct dari Borrowing
	var matchingBorrowings []Borrowing

	// Looping data borrowings hingga mendapatkan judul buku yang ingin diketahui informasi peminjamannya
	for _, borrowing := range borrowings {
		if borrowing.Book.Title == title {
			matchingBorrowings = append(matchingBorrowings, borrowing)
		}
	}

	// Mengecek apakah ada peminjaman untuk buku tersebut
	if len(matchingBorrowings) > 0 {
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|              Peminjaman ditemukan.             |")
		fmt.Println("└────────────────────────────────────────────────┘")

		// Menampilkan semua peminjam buku yang diinginkan
		fmt.Printf("Judul Buku: %s\n", matchingBorrowings[0].Book.Title)
		fmt.Println("Daftar Peminjam:")
		for i, borrowing := range matchingBorrowings {
			fmt.Printf("%d. %s\n", i+1, borrowing.Borrower)
		}

		var borrowingIndex int
		fmt.Print("Pilih peminjam yang akan ditampilkan [1-", len(matchingBorrowings), "]: ")
		fmt.Scanln(&borrowingIndex)

		if borrowingIndex < 1 || borrowingIndex > len(matchingBorrowings) {
			fmt.Println("┌────────────────────────────────────────────────┐")
			fmt.Println("|    Pilihan tidak valid. Operasi dibatalkan.    |")
			fmt.Println("└────────────────────────────────────────────────┘")
			return
		}

		// Menyesuaikan indeks agar sesuai dengan indeks slice (Perhitungan komputer mulai dari 0)
		borrowingIndex--

		selectedBorrowing := matchingBorrowings[borrowingIndex]

		fmt.Printf("Nama Peminjam: %s\n", selectedBorrowing.Borrower)
		fmt.Printf("Tanggal Pinjam: %s\n", selectedBorrowing.BorrowDate.Format("2006-01-02"))
		fmt.Printf("Tanggal Kembali: %s\n", selectedBorrowing.ReturnDate.Format("2006-01-02"))
	} else {
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|           Peminjaman tidak ditemukan.          |")
		fmt.Println("└────────────────────────────────────────────────┘")
	}
}

// fungsi untuk hapus peminjam 6
func deleteBorrowingFromArgs(title string, borrower string) {
	// Temukan peminjaman berdasarkan judul buku dan nama peminjam
	borrowing, borrowingIndex := findBorrowing(title, borrower)
	if borrowing == nil {
		fmt.Println("┌────────────────────────────────────────────────┐")
		fmt.Println("|           Peminjaman tidak ditemukan.          |")
		fmt.Println("└────────────────────────────────────────────────┘")
		return
	}

	// Tambahkan stok untuk buku yang sesuai
	for i, book := range books {
		if book.Title == borrowing.Book.Title {
			books[i].Stock++
			break
		}
	}

	// Hapus peminjaman yang dipilih dari slice
	borrowings = append(borrowings[:borrowingIndex], borrowings[borrowingIndex+1:]...)

	fmt.Printf("Peminjaman untuk buku %s oleh %s berhasil dihapus!\n", borrowing.Book.Title, borrowing.Borrower)
	saveData()
}

func findBorrowing(title string, borrower string) (*Borrowing, int) {
	for i, borrowing := range borrowings {
		if borrowing.Book.Title == title && borrowing.Borrower == borrower {
			return &borrowing, i
		}
	}
	return nil, -1
}
