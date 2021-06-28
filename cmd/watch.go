package cmd

import (
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "watch cloud storage file",
	Long:  `watch cloud storage file`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := getProvider()
		if err != nil {
			return err
		}
		// if watchFilePath == "" {
		// 	return fmt.Errorf("upload file path is missing")
		// }

		// fn := path.Base(watchFilePath)
		// ok := fmt.Sprintf("%s/%s", day, fn)
		err = p.WatchFile(buckName, watchFilePrefix, watchFileMaxkeys)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	watchCmd.PersistentFlags().StringVarP(&watchFilePrefix, "watch_prefix", "x", "", "watch cloud storage file prefix ")
	watchCmd.PersistentFlags().IntVarP(&watchFileMaxkeys, "watch_Max", "m", 100, "watch cloud storage file max")
	watchCmd.PersistentFlags().StringVarP(&buckName, "bucket_name", "b", defaultBuckName, "cloud storage bucket name")
	watchCmd.PersistentFlags().StringVarP(&bucketEndpoint, "bucket_endpoint", "e", defaultEndpoint, "cloud storage endpoint")
	RootCmd.AddCommand(watchCmd)
}
