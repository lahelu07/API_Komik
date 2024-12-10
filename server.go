package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Komik struct {
	ID          int    `json:"id"`
	Nama        string `json:"nama"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	TahunTerbit int    `json:"tahun_terbit"`
	Publisher   string `json:"publisher"`
}

var komikData []Komik
var idCounter = 1 // Counter untuk ID komik

func main() {
	// Buat router
	router := mux.NewRouter()

	// Daftarkan endpoint
	router.HandleFunc("/komik", getKomik).Methods("GET")
	router.HandleFunc("/komik", createKomik).Methods("POST")
	router.HandleFunc("/komik/{id}", getKomikByID).Methods("GET")
	router.HandleFunc("/komik/{id}", updateKomik).Methods("PUT")
	router.HandleFunc("/komik/{id}", deleteKomik).Methods("DELETE")

	// Jalankan server
	fmt.Println("Terhubung dengan http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Ambil semua data komik
func getKomik(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(komikData)
}

// Tambah data komik baru
func createKomik(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var komik Komik
	json.NewDecoder(r.Body).Decode(&komik)

	// Tambahkan ID ke komik baru
	komik.ID = idCounter
	idCounter++

	komikData = append(komikData, komik)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Data berhasil ditambah"))
}

// Ambil data komik berdasarkan ID
func getKomikByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, item := range komikData {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Data tidak ditemukan"))
}

// Perbarui data komik (partial update)
func updateKomik(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, item := range komikData {
		if item.ID == id {
			var updatedKomik Komik
			json.NewDecoder(r.Body).Decode(&updatedKomik)

			// Lakukan pembaruan hanya pada field yang disediakan
			if updatedKomik.Nama != "" {
				item.Nama = updatedKomik.Nama
			}
			if updatedKomik.Author != "" {
				item.Author = updatedKomik.Author
			}
			if updatedKomik.Genre != "" {
				item.Genre = updatedKomik.Genre
			}
			if updatedKomik.TahunTerbit != 0 {
				item.TahunTerbit = updatedKomik.TahunTerbit
			}
			if updatedKomik.Publisher != "" {
				item.Publisher = updatedKomik.Publisher
			}

			komikData[index] = item
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Data berhasil diupdate"))
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Data tidak ditemukan"))
}

// Hapus data komik berdasarkan ID
func deleteKomik(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, item := range komikData {
		if item.ID == id {
			komikData = append(komikData[:index], komikData[index+1:]...)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Data berhasil dihapus"))
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Data tidak ditemukan"))
}
