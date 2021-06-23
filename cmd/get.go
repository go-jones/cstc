package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "下载文件到本地",
	Long:  `下载云存储文件到本地`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := getProvider()
		if err != nil {
			return err
		}
		if putfilename == "" {
			return fmt.Errorf("Download file path is missing")
		}
		day := defaultdir
		fn := putfilename
		ok := fmt.Sprintf("%s/%s", day, fn)
		if getfilename == "" {
			getfilename = putfilename
		}
		err = p.GetFile(buckName, ok, getfilename)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	getCmd.PersistentFlags().StringVarP(&putfilename, "file_path", "f", "", "download file path")
	getCmd.PersistentFlags().StringVarP(&getfilename, "out_path_filename", "o", "", "save file path and filename")
	getCmd.PersistentFlags().StringVarP(&buckName, "bucket_name", "b", defaultBuckName, "cloud storage bucket name")
	getCmd.PersistentFlags().StringVarP(&bucketEndpoint, "bucket_endpoint", "e", defaultEndpoint, "cloud storage endpoint")
	RootCmd.AddCommand(getCmd)
}
