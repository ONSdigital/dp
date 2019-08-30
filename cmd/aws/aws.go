package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

func getAWSSession(environment, profile string) *session.Session {
	opts := session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}
	if profile != "" {
		opts.Profile = profile
	}
	return session.Must(session.NewSessionWithOptions(opts))
}
