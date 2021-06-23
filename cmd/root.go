package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const defaultdir  = "20210623"
var (
	storageProvider string
	EndpointType    string
	AccessID        string
	AccessKey       string
)

var (
	putfilename		string
	getfilename     string
	buckName       string
	uploadFilePath string
	bucketEndpoint string
	public         string
)


var RootCmd = &cobra.Command{
	Use:   "cstcli",
	Short: "cloud-storage-transfer-cli",
	Long:  `cloud-storage-transfer-cli ...`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("no flags find")
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&storageProvider, "oss_provider", "p", "ali", "cloud storage Provider [ali/tx]")
	RootCmd.PersistentFlags().StringVarP(&AccessID, "access_id", "i", "", "the cloud storage access id")
	RootCmd.PersistentFlags().StringVarP(&AccessKey, "access_key", "k", "", "the cloud storage access key")
}
