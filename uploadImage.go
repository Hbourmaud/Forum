package forum

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const Max_upload_size = 20480 * 20480 //20MB

func uploadImage(w http.ResponseWriter, r *http.Request) (string, string, string, string) {
	r.Body = http.MaxBytesReader(w, r.Body, Max_upload_size)
	err := r.ParseMultipartForm(Max_upload_size)
	if err != nil {
		fmt.Println("0", err)
		return "", "", "", "tooBig"
	}

	file, fileHeader, err := r.FormFile("picture")
	title := r.FormValue("title")
	content := r.FormValue("content")
	category := r.FormValue("category")
	if file == nil {
		return title, content, category, ""
	}
	if err != nil {
		fmt.Println("1", err)
		return "", "", "", ""
	}
	defer file.Close()
	buffMime := make([]byte, 512)
	_, err = file.Read(buffMime)
	if err != nil {
		fmt.Println("2", err)
		return "", "", "", ""
	}

	filetype := http.DetectContentType(buffMime)
	if filetype != "image/jpeg" && filetype != "image/png" && filetype != "image/gif" {
		return "", "", "", "wrongType"
	}
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println("3,5", err)
		return "", "", "", ""
	}
	err = os.MkdirAll("./img", os.ModePerm)
	if err != nil {
		fmt.Println("3", err)
		return "", "", "", ""
	}
	save_name := fmt.Sprintf("./img/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))
	dst, err := os.Create(save_name)
	if err != nil {
		fmt.Println("4", err)
		return "", "", "", ""
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		fmt.Println("5", err)
		return "", "", "", ""
	}
	fmt.Println("upload success")
	return title, content, category, save_name
}
