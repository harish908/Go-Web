package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
)

func putFile(sess *session.Session, bucketName string) error{
	file, err := os.Open("C:\\Users\\HARISH\\Desktop\\which_fingers.png")
	if err != nil{
		return err
	}
	defer file.Close()

	fileName := "which_fingers.png"
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: &bucketName,
		Key: &fileName,
		Body: file,
	})

	if err != nil{
		return err
	}
	return nil
}

func main()  {
	// These access keys are incorrect
	accessId := "AKIT74N4"
	accessKey := "/5gy+2CNYiJE1lr3FEgL"
	bucketName := "harish-bucket-2021"
	awsRegion := "ap-south-1"
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(accessId, accessKey, ""),
		},
	)

	if err != nil{
		log.Panic(err)
	}
	err  = putFile(sess, bucketName)
	if err != nil{
		fmt.Println(err)
	}
}
