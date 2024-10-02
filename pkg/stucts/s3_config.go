package stucts

type S3Config struct {
	AccessKeyID     string   `name:"access-key-id"`
	SecretAccessKey string   `name:"secret-access-key"`
	SessionToken    string   `name:"session-token"`
	Region          string   `name:"region"`
	FileLocations   []string `name:"object-urls"`
}
