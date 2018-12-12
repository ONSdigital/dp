package aws

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var resultCache = make(map[string][]EC2Result)

// EC2Result is the information returned for an individual EC2 instance
type EC2Result struct {
	Name          string
	Environment   string
	IPAddress     string
	AnsibleGroups []string
}

func ListEC2ByAnsibleGroup(environment string, ansibleGroup string) ([]EC2Result, error) {
	r, err := ListEC2(environment)
	if err != nil {
		return r, err
	}

	var res []EC2Result
	for _, i := range r {
		for _, t := range i.AnsibleGroups {
			if t == ansibleGroup {
				res = append(res, i)
				break
			}
		}
	}

	return res, nil
}

// ListEC2 returns a list of EC2 instances which match the environment name
func ListEC2(environment string) ([]EC2Result, error) {
	if r, ok := resultCache[environment]; ok {
		return r, nil
	}
	resultCache[environment] = make([]EC2Result, 0)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create new EC2 client
	ec2Svc := ec2.New(sess)

	var result *ec2.DescribeInstancesOutput
	var err error
	request := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("tag:Environment"),
				Values: []*string{aws.String(environment)},
			},
		},
	}

	for {
		if result != nil {
			if result.NextToken == nil {
				break
			}
			request.SetNextToken(*result.NextToken)
		}
		result, err = ec2Svc.DescribeInstances(request)

		if err != nil {
			return nil, err
		}

		for _, r := range result.Reservations {
			for _, i := range r.Instances {
				var name, ansibleGroup string
				for _, tag := range i.Tags {
					if tag.Key != nil && *tag.Key == "Name" && tag.Value != nil {
						name = *tag.Value
						continue
					} else if tag.Key != nil && *tag.Key == "AnsibleGroup" && tag.Value != nil {
						ansibleGroup = *tag.Value
						continue
					}
				}
				var ipAddr string
				if len(i.NetworkInterfaces) > 0 && len(i.NetworkInterfaces[0].PrivateIpAddresses) > 0 {
					if i.NetworkInterfaces[0].PrivateIpAddresses[0].PrivateIpAddress != nil {
						ipAddr = *i.NetworkInterfaces[0].PrivateIpAddresses[0].PrivateIpAddress
					}
				}
				resultCache[environment] = append(resultCache[environment], EC2Result{
					Name:          name,
					IPAddress:     ipAddr,
					Environment:   environment,
					AnsibleGroups: strings.Split(ansibleGroup, ","),
				})
			}
		}
	}

	return resultCache[environment], nil
}
