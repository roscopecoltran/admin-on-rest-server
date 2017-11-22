package aws

type S3Config struct {
	AccessKeyID     string `env:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
	Region          string `env:"AWS_Region"`
	S3Bucket        string `env:"AWS_Bucket"`
}
