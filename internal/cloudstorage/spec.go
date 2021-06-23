package cloudstorage

// Provider todo
type Provider interface {
	UploadFile(bucketName, objectKey, localFilePath,public string) error
	GetFile(bucketName, objectKey, localFilePath string) error
}
