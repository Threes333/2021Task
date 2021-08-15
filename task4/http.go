package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello,world"))
	})
	http.HandleFunc("/html", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("./template/home.html"))
		t.Execute(w, nil)
	})
	http.HandleFunc("/upload", Upload())
	http.HandleFunc("/download", Download())
	http.ListenAndServe(":8080", nil)
}

func Download() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PostFormValue("filename")
		src := path.Join("./", name)
		f1, err := os.Open(src)
		if err != nil {
			w.Write([]byte("file no exist"))
			log.Println(err)
			return
		}
		defer f1.Close()
		w.Header().Add("Content-Disposition", "attachment; filename="+f1.Name())
		for {
			dst := make([]byte, 1024)
			_, err := f1.Read(dst)
			w.Write(dst)
			if err == io.EOF {
				break
			}
		}
	}
}

func Upload() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		src, fh, err := r.FormFile("f1")
		if err != nil {
			w.Write([]byte("fail"))
			log.Fatalln(err)
		}
		defer src.Close()
		dst := path.Join("./", fh.Filename)
		out, err := os.Create(dst)
		if err != nil {
			w.Write([]byte("fail"))
			log.Fatalln(err)
		}
		defer out.Close()
		_, err = io.Copy(out, src)
		if err != nil {
			w.Write([]byte("fail"))
			log.Fatalln(err)
		}
		w.Write([]byte("success"))
	}
}
