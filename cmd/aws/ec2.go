package aws

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	cli "gopkg.in/urfave/cli.v1"

	"github.com/ONSdigital/dp/cmd/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func portHash(username string) int64 {
	var hash int64
	for s := range username {
		hash += int64(s)
	}
	// return numbers that (in all expected cases) are between 10000 and 12000
	port := (hash * 20) + 10000

	// extremely long username strings will exceed allowable port ranges
	if port > 32766 {
		panic("Are you sure DP_SSH_USER is correct?")
	}

	return port
}

func getEC2Service(environment, profile string) *ec2.EC2 {
	// Create new EC2 client
	return ec2.New(getAWSSession(environment, profile))
}

var resultCache = make(map[string][]EC2Result)

// EC2Result is the information returned for an individual EC2 instance
type EC2Result struct {
	Name          string
	Environment   string
	IPAddress     string
	AnsibleGroups []string
}

func GetNamedSG(name, environment, profile string) (string, error) {
	ec2Svc := getEC2Service(environment, profile)
	filters := []*ec2.Filter{
		{
			Name:   aws.String("tag:Name"),
			Values: []*string{aws.String(name)},
		},
	}
	if len(environment) > 0 {
		filters = append(filters, &ec2.Filter{
			Name:   aws.String("tag:Environment"),
			Values: []*string{aws.String(environment)},
		})
	}

	res, err := ec2Svc.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{
		Filters: filters,
	})
	if err != nil {
		return "", err
	}

	if len(res.SecurityGroups) < 1 {
		return "", fmt.Errorf("no security groups matching environment: %q with name %q", environment, name)
	}
	if len(res.SecurityGroups) > 1 {
		return "", fmt.Errorf("too many security groups matching environment: %s name: %q", environment, name)
	}
	if res.SecurityGroups[0].GroupId == nil {
		return "", fmt.Errorf("no groupId found for security group on environment: %q name: %q", environment, name)
	}

	return *res.SecurityGroups[0].GroupId, nil
}

func GetBastionSGForEnvironment(environment, profile string) (string, error) {
	return GetNamedSG(environment+" - bastion", environment, profile)
}

func GetELBPublishingSGForEnvironment(environment, profile string) (string, error) {
	return GetNamedSG(environment+" - publishing elb", environment, profile)
}

func GetELBWebSGForEnvironment(environment, profile string) (string, error) {
	return GetNamedSG(environment+" - web elb", environment, profile)
}

func GetConcourseWebSG() (string, error) {
	return GetNamedSG("concourse-ci-web", "", "")
}

func GetManagementACLForEnvironment(environment, profile string) (string, error) {
	ec2Svc := getEC2Service(environment, profile)

	res, err := ec2Svc.DescribeNetworkAcls(&ec2.DescribeNetworkAclsInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("tag:Environment"),
				Values: []*string{aws.String(environment)},
			},
			{
				Name:   aws.String("tag:Name"),
				Values: []*string{aws.String(environment + " - management")},
			},
		},
	})
	if err != nil {
		return "", err
	}

	if len(res.NetworkAcls) < 1 {
		return "", fmt.Errorf("no network acls matching environment: %s", environment)
	}
	if len(res.NetworkAcls) > 1 {
		return "", fmt.Errorf("too many network acls matching environment: %s", environment)
	}
	if res.NetworkAcls[0].NetworkAclId == nil {
		return "", fmt.Errorf("no networkAclId found for network acl: %s", environment)
	}

	return *res.NetworkAcls[0].NetworkAclId, nil
}

func AllowIPForConcourse(cfg config.Config) error {
	ec2Svc := getEC2Service("", "")

	sg, err := GetConcourseWebSG()
	if err != nil {
		return err
	}

	myIP, err := config.GetMyIP()
	if err != nil {
		return err
	}

	var errs []error

	_, err = ec2Svc.AuthorizeSecurityGroupIngress(&ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: aws.String(sg),
		IpPermissions: []*ec2.IpPermission{
			{
				IpProtocol: aws.String("tcp"),
				FromPort:   aws.Int64(int64(443)),
				ToPort:     aws.Int64(int64(443)),
				IpRanges:   getIpRangesFor(myIP, cfg.SSHUser),
			},
		},
	})
	if err != nil {
		errs = append(errs, fmt.Errorf("error adding rules to sg: %s", err))
	}

	if len(errs) > 0 {
		return cli.NewMultiError(errs...)
	}

	return nil
}

func DenyIPForConcourse(cfg config.Config) error {
	ec2Svc := getEC2Service("", "")

	sg, err := GetConcourseWebSG()
	if err != nil {
		return err
	}

	myIP, err := config.GetMyIP()
	if err != nil {
		return err
	}

	var errs []error

	_, err = ec2Svc.RevokeSecurityGroupIngress(&ec2.RevokeSecurityGroupIngressInput{
		GroupId: aws.String(sg),
		IpPermissions: []*ec2.IpPermission{
			{
				IpProtocol: aws.String("tcp"),
				FromPort:   aws.Int64(int64(443)),
				ToPort:     aws.Int64(int64(443)),
				IpRanges:   getIpRangesFor(myIP, cfg.SSHUser),
			},
		},
	})
	if err != nil {
		errs = append(errs, fmt.Errorf("error removing rules from sg: %s", err))
	}

	if len(errs) > 0 {
		return cli.NewMultiError(errs...)
	}

	return nil
}

func DenyIPForEnvironment(cfg config.Config, environment, profile string) error {
	return ChangeIPForEnvironment(false, cfg, environment, profile)
}
func AllowIPForEnvironment(cfg config.Config, environment, profile string) error {
	return ChangeIPForEnvironment(true, cfg, environment, profile)
}

func ChangeIPForEnvironment(isAllow bool, cfg config.Config, environment, profile string) error {
	ec2Svc := getEC2Service(environment, profile)

	if len(cfg.SSHUser) == 0 {
		return errors.New("please set DP_SSH_USER to change remote access")
	}
	ruleBase := portHash(cfg.SSHUser)

	bastionSG, err := GetBastionSGForEnvironment(environment, profile)
	if err != nil {
		return err
	}

	envIsProduction := environment == "production"
	var pubSG string
	if !envIsProduction {
		var err error
		pubSG, err = GetELBPublishingSGForEnvironment(environment, profile)
		if err != nil {
			return err
		}
	}

	webSG, err := GetELBWebSGForEnvironment(environment, profile)
	if err != nil {
		return err
	}

	acl, err := GetManagementACLForEnvironment(environment, profile)
	if err != nil {
		return err
	}

	myIP, err := config.GetMyIP()
	if err != nil {
		return err
	}

	var (
		errs        []error
		ipPermHTTPS = &ec2.IpPermission{
			IpProtocol: aws.String("tcp"),
			FromPort:   aws.Int64(int64(443)),
			ToPort:     aws.Int64(int64(443)),
			IpRanges:   getIpRangesFor(myIP, cfg.SSHUser),
		}
		ipPermSSH = &ec2.IpPermission{
			IpProtocol: aws.String("tcp"),
			FromPort:   aws.Int64(int64(22)),
			ToPort:     aws.Int64(int64(22)),
			IpRanges:   getIpRangesFor(myIP, cfg.SSHUser),
		}
		ipPermHTTP = &ec2.IpPermission{
			IpProtocol: aws.String("tcp"),
			FromPort:   aws.Int64(int64(80)),
			ToPort:     aws.Int64(int64(80)),
			IpRanges:   getIpRangesFor(myIP, cfg.SSHUser),
		}
		ipPermsAllHTTP = []*ec2.IpPermission{
			ipPermHTTP,
			ipPermHTTPS,
		}
		ipPermsAllSecure = []*ec2.IpPermission{
			ipPermSSH,
			ipPermHTTPS,
		}
	)

	if isAllow {

		_, err = ec2Svc.AuthorizeSecurityGroupIngress(&ec2.AuthorizeSecurityGroupIngressInput{
			GroupId:       aws.String(bastionSG),
			IpPermissions: ipPermsAllSecure,
		})
		if err != nil {
			errs = append(errs, fmt.Errorf("error adding rules to bastionSG: %s", err))
		}

		_, err = ec2Svc.AuthorizeSecurityGroupIngress(&ec2.AuthorizeSecurityGroupIngressInput{
			GroupId:       aws.String(webSG),
			IpPermissions: ipPermsAllHTTP,
		})
		if err != nil {
			errs = append(errs, fmt.Errorf("error removing rules from webSG: %s", err))
		}

		if !envIsProduction {
			_, err = ec2Svc.AuthorizeSecurityGroupIngress(&ec2.AuthorizeSecurityGroupIngressInput{
				GroupId:       aws.String(pubSG),
				IpPermissions: ipPermsAllHTTP,
			})
			if err != nil {
				errs = append(errs, fmt.Errorf("error removing rules from pubSG: %s", err))
			}
		}

		_, err = ec2Svc.CreateNetworkAclEntry(&ec2.CreateNetworkAclEntryInput{
			Egress:       aws.Bool(false),
			NetworkAclId: aws.String(acl),
			RuleNumber:   aws.Int64(ruleBase), // 1 to 32766
			CidrBlock:    aws.String(myIP + "/32"),
			Protocol:     aws.String("6"),
			RuleAction:   aws.String("allow"),
			PortRange: &ec2.PortRange{
				From: aws.Int64(int64(22)),
				To:   aws.Int64(int64(22)),
			},
		})
		if err != nil {
			errs = append(errs, fmt.Errorf("error adding rules to acl: %s", err))
		}

		_, err = ec2Svc.CreateNetworkAclEntry(&ec2.CreateNetworkAclEntryInput{
			Egress:       aws.Bool(false),
			NetworkAclId: aws.String(acl),
			RuleNumber:   aws.Int64(ruleBase + 1), // 1 to 32766
			CidrBlock:    aws.String(myIP + "/32"),
			Protocol:     aws.String("6"),
			RuleAction:   aws.String("allow"),
			PortRange: &ec2.PortRange{
				From: aws.Int64(int64(443)),
				To:   aws.Int64(int64(443)),
			},
		})
		if err != nil {
			errs = append(errs, fmt.Errorf("error adding rules to acl: %s", err))
		}

		_, err = ec2Svc.CreateNetworkAclEntry(&ec2.CreateNetworkAclEntryInput{
			Egress:       aws.Bool(true),
			NetworkAclId: aws.String(acl),
			RuleNumber:   aws.Int64(ruleBase + 2), // 1 to 32766
			CidrBlock:    aws.String(myIP + "/32"),
			Protocol:     aws.String("6"),
			RuleAction:   aws.String("allow"),
			PortRange: &ec2.PortRange{
				From: aws.Int64(int64(32768)),
				To:   aws.Int64(int64(61000)),
			},
		})
		if err != nil {
			errs = append(errs, fmt.Errorf("error adding rules to acl: %s", err))
		}

	} else {

		_, err = ec2Svc.RevokeSecurityGroupIngress(&ec2.RevokeSecurityGroupIngressInput{
			GroupId:       aws.String(bastionSG),
			IpPermissions: ipPermsAllSecure,
		})
		if err != nil {
			errs = append(errs, fmt.Errorf("error removing rules from bastionSG: %s", err))
		}

		_, err = ec2Svc.RevokeSecurityGroupIngress(&ec2.RevokeSecurityGroupIngressInput{
			GroupId:       aws.String(webSG),
			IpPermissions: ipPermsAllHTTP,
		})
		if err != nil {
			errs = append(errs, fmt.Errorf("error removing rules from webSG: %s", err))
		}

		if !envIsProduction {
			_, err = ec2Svc.RevokeSecurityGroupIngress(&ec2.RevokeSecurityGroupIngressInput{
				GroupId:       aws.String(pubSG),
				IpPermissions: ipPermsAllHTTP,
			})
			if err != nil {
				errs = append(errs, fmt.Errorf("error removing rules from pubSG: %s", err))
			}
		}

		_, err = ec2Svc.DeleteNetworkAclEntry(&ec2.DeleteNetworkAclEntryInput{
			Egress:       aws.Bool(false),
			NetworkAclId: aws.String(acl),
			RuleNumber:   aws.Int64(ruleBase), // 1 to 32766
		})
		if err != nil {
			errs = append(errs, fmt.Errorf("error removing rules from acl: %s", err))
		}

		_, err = ec2Svc.DeleteNetworkAclEntry(&ec2.DeleteNetworkAclEntryInput{
			Egress:       aws.Bool(false),
			NetworkAclId: aws.String(acl),
			RuleNumber:   aws.Int64(ruleBase + 1), // 1 to 32766
		})
		if err != nil {
			errs = append(errs, fmt.Errorf("error removing rules from acl: %s", err))
		}

		_, err = ec2Svc.DeleteNetworkAclEntry(&ec2.DeleteNetworkAclEntryInput{
			Egress:       aws.Bool(true),
			NetworkAclId: aws.String(acl),
			RuleNumber:   aws.Int64(ruleBase + 2), // 1 to 32766
		})
		if err != nil {
			errs = append(errs, fmt.Errorf("error removing rules from acl: %s", err))
		}
	}

	if len(errs) > 0 {
		return cli.NewMultiError(errs...)
	}

	return nil
}

func ListEC2ByAnsibleGroup(environment, profile string, ansibleGroup string) ([]EC2Result, error) {
	r, err := ListEC2(environment, profile)
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
func ListEC2(environment, profile string) ([]EC2Result, error) {
	if r, ok := resultCache[environment]; ok {
		return r, nil
	}
	resultCache[environment] = make([]EC2Result, 0)

	ec2Svc := getEC2Service(environment, profile)

	var result *ec2.DescribeInstancesOutput
	var err error
	request := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("tag:Environment"),
				Values: []*string{aws.String(environment)},
			},
			{
				Name:   aws.String("instance-state-name"),
				Values: []*string{aws.String(ec2.InstanceStateNameRunning)},
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

	sort.Slice(resultCache[environment], func(i, j int) bool { return resultCache[environment][i].Name < resultCache[environment][j].Name })
	return resultCache[environment], nil
}

func getIpRangesFor(myIP, sshUser string) []*ec2.IpRange {
	return []*ec2.IpRange{
		{
			CidrIp:      aws.String(myIP + "/32"),
			Description: &sshUser,
		},
	}
}
