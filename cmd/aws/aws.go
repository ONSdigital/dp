package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

func getAWSSession() *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
}
