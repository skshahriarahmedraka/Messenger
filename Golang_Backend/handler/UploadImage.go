package handler

import (
	"app/logerror"
	"app/model"
	"context"
	"encoding/json"
	"os"
	"strconv"


	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func (H *DatabaseCollections)UploadImg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

		// c.Request.Header.Set()
		// 	c.Request.Header.Set("Access-Control-Allow-Headers", "Content-Type")
		// c.Request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		// c.Request.Header.Set("Access-Control-Allow-Origin", "*")

		var l model.ImgLinkModel
		// err := json.NewDecoder(r.Body).Decode(&l)
		var resError model.ErrorRes
		err := json.NewDecoder(r.Body).Decode(&l)
        fmt.Println("🚀 ~ file: imglink.go ~ line 32 ~ func ~ REquested data : ", l)
		if err != nil {
        	logerror.ERROR("🚀 ~ file: imglink.go ~ line 37 ~ func ~ err : ", err)
			w.WriteHeader(http.StatusBadRequest)
			resError.ErrorRes=`{"error": "request is not valid"}`
			json.NewEncoder(w).Encode(&resError)
			// c.JSON(http.StatusBadRequest, gin.H{"error": "bind error"})
			return
		}

		endpoint := os.Getenv("MINIO_ENDPOINT")
		accessKeyID := os.Getenv("MINIO_ACCESS_KEY_ID")
		secretAccessKey := os.Getenv("MINIO_ACCESS_KEY")

		useSSL , _ := strconv.ParseBool(os.Getenv("MINIO_USE_SSL"))

		minioClient, err := minio.New(endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: useSSL,
		})
		if err != nil {
			logerror.ERROR("🚀 ~ file: imglink.go ~ line 56 ~ func ~ err : ", err)

			w.WriteHeader(http.StatusInternalServerError)
			resError.ErrorRes=`{"error": " minio connection error"}`
			json.NewEncoder(w).Encode(&resError)
			// c.JSON(http.StatusInternalServerError, gin.H{"error": "minio connection error"})
			return
		}
		fmt.Println("minio connected successfully", minioClient)

		// Set request parameters for content-disposition.
		reqParams := make(url.Values)
		// reqParams.Set("response-content-disposition", "attachment; filename=\"p.png\"")

		presignedURL, err := minioClient.PresignedGetObject(context.Background(), l.BucketName, l.Name, time.Second*24*60*60, reqParams)
		if err != nil {
			logerror.ERROR("🚀 ~ file: imglink.go ~ line 72 ~ func ~ err : ", err)
			w.WriteHeader(http.StatusInternalServerError)
			resError.ErrorRes=`{"error": " minio url generation error"}`
			json.NewEncoder(w).Encode(&resError)
			// c.JSON(http.StatusInternalServerError, gin.H{"error": "minio getting url error"})

			return
		}
		l.Link = (*presignedURL).String()
		fmt.Println("🚀 ~ file: imglink.go ~ line 59 ~ returnfunc ~ l.Link : ", l.Link)

		json.NewEncoder(w).Encode(&l)
		// c.JSON(http.StatusOK, l)

	}