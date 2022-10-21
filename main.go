package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
)

/**
/* http Handle ile route/port tanımlamalarını yaptık ve server ayaklandırdık.
*/
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", imageUpload)
	err := http.ListenAndServe(":2222", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server kapatıldı\n")
	} else if err != nil {
		fmt.Printf("Bir hata oluştu: %s\n", err)
		os.Exit(1)
	}
}

/**
/* Açılış sayfamızda index olarak çekmek istediğimiz html sayfasının yolunu verdik.
*/
func index(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("html/index.html")
	view.Execute(w, nil)
}

/**
/* file ile gelen request değerlerini aldık. newfile ile oluşturacağımız dosyayı belirledik 
/* gelen dosyayı oluşturduğumuz newfile üzerine yazdık 
*/
func imageUpload(w http.ResponseWriter, r *http.Request) {
	file, header, _ := r.FormFile("file")
	newFile, _ := os.OpenFile("storage/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0755)
	io.Copy(newFile, file)
}
