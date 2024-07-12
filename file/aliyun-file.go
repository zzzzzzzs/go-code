package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func main() {
	// 阿里云OSS的Access Key ID和Access Key Secret
	accessKeyId := "your_access_key_id"
	accessKeySecret := "your_access_key_secret"

	// OSS的Endpoint和Bucket名称
	endpoint := "oss-cn-hangzhou.aliyuncs.com"
	bucketName := "your_bucket_name"

	// 创建OSS客户端实例
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 获取Bucket对象
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 本地文件路径
	localFilePath := "path/to/local/file"

	// 上传文件到OSS
	err = bucket.PutObjectFromFile("remote/file/path", localFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Upload success!")
}
