package s3connManager

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type connManager struct {
	connCache map[string]*s3.S3
	err       error
}

func (c *connManager) defaultNewS3(host, accessKey, secretKey string) (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:         aws.String(host),
		Region:           aws.String("Region"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
		//MaxRetries:                        nil,
		//Retryer:                           nil,
		//DisableParamValidation:            nil,
		//DisableComputeChecksums:           nil,
		//S3Disable100Continue:              nil,
		//S3UseAccelerate:                   nil,
		//S3DisableContentMD5Validation:     nil,
		//S3UseARNRegion:                    nil,
		//LowerCaseHeaderMaps:               nil,
		//EC2MetadataDisableTimeoutOverride: nil,
		//EC2MetadataEnableFallback:         nil,
		//UseDualStack:                      nil,
		//UseDualStackEndpoint:              0,
		//UseFIPSEndpoint:                   0,
		//SleepDelay:                        nil,
		//DisableRestProtocolURICleaning:    nil,
		//EnableEndpointDiscovery:           nil,
		//DisableEndpointHostPrefix:         nil,
		//STSRegionalEndpoint:               0,
		//S3UsEast1RegionalEndpoint:         0,
	})
	if err != nil {
		return nil, err
	}

	s3 := s3.New(sess)
	return s3, nil
}

func (c *connManager) GetS3Client(host, accessKey, secretKey string) (*s3.S3, error) {
	s, ok := c.connCache[host]
	if ok {
		return s, nil
	}
	s, err := c.GetS3ClientForm(host, accessKey, secretKey, c.defaultNewS3)
	if err != nil {
		return nil, err
	}
	c.connCache[host] = s
	return s, nil
}

func (c *connManager) GetS3ClientForm(host, accessKey, secretKey string, fn func(host, accessKey, secretKey string) (*s3.S3, error)) (*s3.S3, error) {
	return fn(host, accessKey, secretKey)
}

func NewConnManager() *connManager {
	return &connManager{
		connCache: make(map[string]*s3.S3),
	}
}
