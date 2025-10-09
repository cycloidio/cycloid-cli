package externalbackends

import "github.com/spf13/cobra"

var (
	regionAwsFlag       string
	regionFlag          string
	awsBucketName       string
	awsBucketPath       string
	awsS3BucketEndpoint string
	urls                []string
	prefilters          map[string]string
	skipVerifySSL       bool
	awsS3ForcePathStyle bool
	def                 bool
)

func WithFlagAwsRegion(cmd *cobra.Command) string {
	flagName := "region"
	cmd.Flags().StringVar(&regionAwsFlag, flagName, "eu-west-1", "region")
	return flagName
}

func WithFlagRegion(cmd *cobra.Command) string {
	flagName := "region"
	cmd.Flags().StringVar(&regionFlag, flagName, "", "region")
	return flagName
}

func WithFlagBucketName(cmd *cobra.Command) string {
	flagName := "bucket-name"
	cmd.Flags().StringVar(&awsBucketName, flagName, "", "bucket name")
	return flagName
}
func WithFlagBucketPath(cmd *cobra.Command) string {
	flagName := "bucket-path"
	cmd.Flags().StringVar(&awsBucketPath, flagName, "", "bucket path")
	return flagName
}

func WithFlagS3BucketEndpoint(cmd *cobra.Command) string {
	flagName := "endpoint"
	cmd.Flags().StringVar(&awsS3BucketEndpoint, flagName, "", "Aws S3 endpoint")
	return flagName
}
func WithFlagS3ForcePathStyle(cmd *cobra.Command) string {
	flagName := "s3-force-path-style"
	cmd.Flags().BoolVar(&awsS3ForcePathStyle, flagName, true, "")
	return flagName
}
func WithFlagSkipVerifySSL(cmd *cobra.Command) string {
	flagName := "skip-verify-ssl"
	cmd.Flags().BoolVar(&skipVerifySSL, flagName, false, "")
	return flagName
}
func WithFlagDefault(cmd *cobra.Command) string {
	flagName := "default"
	cmd.Flags().BoolVar(&def, flagName, false, "")
	return flagName
}

func WithFlagURL(cmd *cobra.Command) string {
	flagName := "url"
	cmd.Flags().StringSliceVar(&urls, flagName, nil, "urls")
	return flagName
}
func WithFlagPrefilter(cmd *cobra.Command) string {
	flagName := "prefilter"
	cmd.Flags().StringToStringVar(&prefilters, flagName, nil, "key=value")
	return flagName
}
