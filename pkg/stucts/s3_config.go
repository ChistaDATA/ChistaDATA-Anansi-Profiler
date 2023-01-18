package stucts

type S3Config struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
	Region          string
	FileLocations   []string
}
