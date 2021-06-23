package cmd

import (
	"fmt"
	"path"

	"github.com/go-jones/cstc/internal/cloudstorage"
	"github.com/go-jones/cstc/internal/cloudstorage/provider/alicloud"
	"github.com/spf13/cobra"
)


var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload file to cloud storage",
	Long:  `upload file to cloud storage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := getProvider()
		if err != nil {
			return err
		}
		if uploadFilePath == "" {
			return fmt.Errorf("upload file path is missing")
		}
		// day := time.Now().Format("20060102")
		// 上传到指定存储目录下方便文件生命周期管理同时也方便下载
		day := defaultdir
		if public == "true" {
			day = fmt.Sprintf("%s_tmp",day)
		}
		fn := path.Base(uploadFilePath)
		ok := fmt.Sprintf("%s/%s", day, fn)
		err = p.UploadFile(buckName, ok, uploadFilePath, public)
		if err != nil {
			return err
		}
		return nil
	},
}

func getProvider() (p cloudstorage.Provider, err error) {
	switch storageProvider {
	case "ali":
		fmt.Printf("云商: 阿里云[%s]\n", bucketEndpoint)
		if AccessID == "" {
			AccessID = defaultALIAK
		}
		if AccessKey == "" {
			AccessKey = defaultALISK
		}
		fmt.Printf("用户: %s\n", AccessID)
		p, err = alicloud.NewProvider(bucketEndpoint, AccessID, AccessKey)
		return
	case "tx":
		return nil, fmt.Errorf("not impl")
	default:
		return nil, fmt.Errorf("unknown oss privier options [ali/tx]")
	}
}

func init() {
	uploadCmd.PersistentFlags().StringVarP(&uploadFilePath, "file_path", "f", "", "upload file path")
	uploadCmd.PersistentFlags().StringVarP(&buckName, "bucket_name", "b", defaultBuckName, "cloud storage bucket name")
	uploadCmd.PersistentFlags().StringVarP(&bucketEndpoint, "bucket_endpoint", "e", defaultEndpoint, "cloud storage endpoint")
	uploadCmd.PersistentFlags().StringVarP(&public, "public", "u", "false", "provide public download link")
	RootCmd.AddCommand(uploadCmd)
}
