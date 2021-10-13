package handler

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-redis/redis/v8"
)

func AddPicture(ctx context.Context, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(30 << 40)

		fmt.Println(r.Body)
		fmt.Println(r.ContentLength)

		// file, handler, err := r.FormFile("myFile")
		// if err != nil {
		// 	fmt.Println("Error Retrieving the File")
		// 	fmt.Println(err)
		// 	return
		// }
		// defer file.Close()

		// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		// fmt.Printf("File Size: %+v\n", handler.Size)
		// fmt.Printf("MIME Header: %+v\n", handler.Header)

		tempFile, err := ioutil.TempFile("./yanto", "upload-*.png")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		tempFile.Write(fileBytes)
	}
}
