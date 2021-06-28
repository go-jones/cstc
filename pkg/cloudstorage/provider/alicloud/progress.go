package alicloud

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

// NewProvider todo
func NewProvider(endpoint, accessID, accessKey string) (*Provider, error) {
	p := &Provider{
		Endpoint:  endpoint,
		AccessID:  accessID,
		AccessKey: accessKey,
	}

	if err := p.Validate(); err != nil {
		return nil, err
	}

	return p, nil
}

// Provider todo
type Provider struct {
	Endpoint  string `validate:"required"`
	AccessID  string `validate:"required"`
	AccessKey string `validate:"required"`
}

func (p *Provider) Validate() error {
	return validate.Struct(p)
}

// GetBucket todo
func (p *Provider) GetBucket(bucketName string) (*oss.Bucket, error) {
	if bucketName == "" {
		return nil, fmt.Errorf("upload bucket name required")
	}

	// New client
	client, err := oss.New(p.Endpoint, p.AccessID, p.AccessKey)
	if err != nil {
		return nil, err
	}
	// Get bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

type OssProgressListener struct {
}

func (listener *OssProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		fmt.Printf("Transfer Started, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferDataEvent:
		if event.TotalBytes != 0 {
			fmt.Printf("\rTransfer Data, ConsumedBytes: %d, TotalBytes %d, %d%%.",
				event.ConsumedBytes, event.TotalBytes, event.ConsumedBytes*100/event.TotalBytes)
		}
		fmt.Printf("\rTransfer Data, ConsumedBytes: %d, TotalBytes %d, %d%%.",
			event.ConsumedBytes, event.TotalBytes, event.ConsumedBytes*100/event.TotalBytes)
	case oss.TransferCompletedEvent:
		fmt.Printf("\nTransfer Completed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferFailedEvent:
		fmt.Printf("\nTransfer Failed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	default:
	}
}

func (p *Provider) UploadFile(bucketName, objectKey, localFilePath, public string) error {
	bucket, err := p.GetBucket(bucketName)
	if err != nil {
		return err
	}
	fmt.Printf("上传位置: bucket名称: %s bucket路径: %s\n", bucketName, objectKey)
	err = bucket.PutObjectFromFile(objectKey, localFilePath, oss.Progress(&OssProgressListener{}))
	if err != nil {
		return fmt.Errorf("upload file to bucket: %s error, %s", bucketName, err)
	}
	if public == "true" {
		signedURL, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24*3)
		if err != nil {
			return fmt.Errorf("SignURL error, %s", err)
		}
		fmt.Printf("下载链接: %s\n", signedURL)
		fmt.Println("\n注意: 文件下载链接有效期为3天,请及时下载")
	}
	return nil
}

func (p *Provider) GetFile(bucketName, objectKey, localFilePath string) error {
	bucket, err := p.GetBucket(bucketName)
	if err != nil {
		return err
	}
	fmt.Printf("下载位置: bucket名称: %s bucket路径: %s\n", bucketName, objectKey)
	err = bucket.GetObjectToFile(objectKey, localFilePath, oss.Progress(&OssProgressListener{}))
	if err != nil {
		return fmt.Errorf("download cloud storge file to local: %s error, %s", bucketName, err)
	}
	return nil
}

func (p *Provider) WatchFile(bucketName, watchFileDir, watchFilePrefix string, maxkeys int) error {
	bucket, err := p.GetBucket(bucketName)
	if err != nil {
		return err
	}
	continueToken := ""
	marker := oss.Marker(watchFileDir)
	watchFilePrefix = fmt.Sprintf("%s/%s", watchFileDir, watchFilePrefix)
	prefix := oss.Prefix(watchFilePrefix)
	for {
		lsRes, err := bucket.ListObjectsV2(marker, prefix, oss.MaxKeys(maxkeys), oss.ContinuationToken(continueToken))
		if err != nil {
			return fmt.Errorf("watch cloud storge file : %s error, %s", watchFileDir, err)
		}
		fmt.Println("对象名称      文件大小(b)      最后修改时间      存储类型")
		for _, object := range lsRes.Objects {
			// fmt.Println(object.Key, object.Type, object.Size, object.ETag, object.LastModified, object.StorageClass)
			fmt.Println(object.Key,  object.Size, object.LastModified, object.StorageClass)
		}

		if lsRes.IsTruncated {
			continueToken = lsRes.NextContinuationToken
		} else {
			break
		}
	}
	return nil
}
