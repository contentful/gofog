package autoscaling_test

import (
	"github.com/hailocab/goamz/autoscaling"
	"github.com/hailocab/goamz/aws"
	"github.com/hailocab/goamz/testutil"
	"launchpad.net/gocheck"
	"testing"
	"time"
)

func Test(t *testing.T) {
	gocheck.TestingT(t)
}

var _ = gocheck.Suite(&S{})

type S struct {
	as *autoscaling.AutoScaling
}

var testServer = testutil.NewHTTPServer()

func (s *S) SetUpSuite(c *gocheck.C) {
	testServer.Start()
	auth := aws.Auth{AccessKey: "abc", SecretKey: "123"}
	s.as = autoscaling.New(auth, aws.Region{AutoScalingEndpoint: testServer.URL})
}

func (s *S) TearDownTest(c *gocheck.C) {
	testServer.Flush()
}

func (s *S) TestCreateLaunchConfiguration(c *gocheck.C) {
	testServer.Response(200, nil, CreateLaunchConfiguration)
	testServer.Response(200, nil, DeleteLaunchConfiguration)

	launchConfig := &autoscaling.CreateLaunchConfiguration{
		LaunchConfigurationName:  "my-test-lc",
		AssociatePublicIpAddress: true,
		EbsOptimized:             true,
		SecurityGroups:           []string{"sec-grp1", "sec-grp2"},
		UserData:                 "1234",
		KeyName:                  "secretKeyPair",
		ImageId:                  "ami-0078da69",
		InstanceType:             "m1.small",
		SpotPrice:                "0.03",
		BlockDeviceMappings: []autoscaling.BlockDeviceMapping{
			{
				DeviceName:  "/dev/sda1",
				VirtualName: "ephemeral0",
			},
			{
				DeviceName:  "/dev/sdb",
				VirtualName: "ephemeral1",
			},
			{
				DeviceName: "/dev/sdf",
				Ebs: autoscaling.EBS{
					DeleteOnTermination: true,
					SnapshotId:          "snap-2a2b3c4d",
					VolumeSize:          100,
				},
			},
		},
		InstanceMonitoring: autoscaling.InstanceMonitoring{
			Enabled: true,
		},
	}
	resp, err := s.as.CreateLaunchConfiguration(launchConfig)
	c.Assert(err, gocheck.IsNil)
	defer s.as.DeleteLaunchConfiguration(launchConfig.LaunchConfigurationName)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "CreateLaunchConfiguration")
	c.Assert(values.Get("LaunchConfigurationName"), gocheck.Equals, "my-test-lc")
	c.Assert(values.Get("AssociatePublicIpAddress"), gocheck.Equals, "true")
	c.Assert(values.Get("EbsOptimized"), gocheck.Equals, "true")
	c.Assert(values.Get("SecurityGroups.member.1"), gocheck.Equals, "sec-grp1")
	c.Assert(values.Get("SecurityGroups.member.2"), gocheck.Equals, "sec-grp2")
	c.Assert(values.Get("UserData"), gocheck.Equals, "MTIzNA==")
	c.Assert(values.Get("KeyName"), gocheck.Equals, "secretKeyPair")
	c.Assert(values.Get("ImageId"), gocheck.Equals, "ami-0078da69")
	c.Assert(values.Get("InstanceType"), gocheck.Equals, "m1.small")
	c.Assert(values.Get("SpotPrice"), gocheck.Equals, "0.03")
	c.Assert(values.Get("BlockDeviceMappings.member.1.DeviceName"), gocheck.Equals, "/dev/sda1")
	c.Assert(values.Get("BlockDeviceMappings.member.1.VirtualName"), gocheck.Equals, "ephemeral0")
	c.Assert(values.Get("BlockDeviceMappings.member.2.DeviceName"), gocheck.Equals, "/dev/sdb")
	c.Assert(values.Get("BlockDeviceMappings.member.2.VirtualName"), gocheck.Equals, "ephemeral1")
	c.Assert(values.Get("BlockDeviceMappings.member.3.DeviceName"), gocheck.Equals, "/dev/sdf")
	c.Assert(values.Get("BlockDeviceMappings.member.3.Ebs.DeleteOnTermination"), gocheck.Equals, "true")
	c.Assert(values.Get("BlockDeviceMappings.member.3.Ebs.SnapshotId"), gocheck.Equals, "snap-2a2b3c4d")
	c.Assert(values.Get("BlockDeviceMappings.member.3.Ebs.VolumeSize"), gocheck.Equals, "100")
	c.Assert(values.Get("InstanceMonitoring.Enabled"), gocheck.Equals, "true")
	c.Assert(resp.RequestId, gocheck.Equals, "7c6e177f-f082-11e1-ac58-3714bEXAMPLE")
}

func (s *S) TestDeleteLaunchConfiguration(c *gocheck.C) {
	testServer.Response(200, nil, DeleteLaunchConfiguration)
	resp, err := s.as.DeleteLaunchConfiguration("my-test-lc")
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "DeleteLaunchConfiguration")
	c.Assert(values.Get("LaunchConfigurationName"), gocheck.Equals, "my-test-lc")
	c.Assert(resp.RequestId, gocheck.Equals, "7347261f-97df-11e2-8756-35eEXAMPLE")
}

func (s *S) TestDeleteLaunchConfigurationInUse(c *gocheck.C) {
	testServer.Response(400, nil, DeleteLaunchConfigurationInUse)
	resp, err := s.as.DeleteLaunchConfiguration("my-test-lc")
	testServer.WaitRequest()
	c.Assert(resp, gocheck.IsNil)
	c.Assert(err, gocheck.NotNil)
	e, ok := err.(*autoscaling.Error)
	c.Assert(ok, gocheck.Equals, true)
	c.Assert(e.Message, gocheck.Equals, "Cannot delete launch configuration my-test-lc because it is attached to AutoScalingGroup test")
	c.Assert(e.Code, gocheck.Equals, "ResourceInUse")
	c.Assert(e.StatusCode, gocheck.Equals, 400)
	c.Assert(e.RequestId, gocheck.Equals, "7347261f-97df-11e2-8756-35eEXAMPLE")
}

func (s *S) TestCreateAutoScalingGroup(c *gocheck.C) {
	testServer.Response(200, nil, CreateAutoScalingGroup)
	testServer.Response(200, nil, DeleteAutoScalingGroup)

	createAS := &autoscaling.CreateAutoScalingGroup{
		AutoScalingGroupName:    "my-test-asg",
		AvailabilityZones:       []string{"us-east-1a", "us-east-1b"},
		MinSize:                 3,
		MaxSize:                 3,
		DefaultCooldown:         600,
		DesiredCapacity:         3,
		LaunchConfigurationName: "my-test-lc",
		LoadBalancerNames:       []string{"elb-1", "elb-2"},
		Tags: []autoscaling.Tag{
			{
				Key:   "foo",
				Value: "bar",
			},
			{
				Key:   "baz",
				Value: "qux",
			},
		},
		VPCZoneIdentifier: "subnet-610acd08,subnet-530fc83a",
	}
	resp, err := s.as.CreateAutoScalingGroup(createAS)
	c.Assert(err, gocheck.IsNil)
	defer s.as.DeleteAutoScalingGroup(createAS.AutoScalingGroupName, true)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "CreateAutoScalingGroup")
	c.Assert(values.Get("AutoScalingGroupName"), gocheck.Equals, "my-test-asg")
	c.Assert(values.Get("AvailabilityZones.member.1"), gocheck.Equals, "us-east-1a")
	c.Assert(values.Get("AvailabilityZones.member.2"), gocheck.Equals, "us-east-1b")
	c.Assert(values.Get("MinSize"), gocheck.Equals, "3")
	c.Assert(values.Get("MaxSize"), gocheck.Equals, "3")
	c.Assert(values.Get("DefaultCooldown"), gocheck.Equals, "600")
	c.Assert(values.Get("DesiredCapacity"), gocheck.Equals, "3")
	c.Assert(values.Get("LaunchConfigurationName"), gocheck.Equals, "my-test-lc")
	c.Assert(values.Get("LoadBalancerNames.member.1"), gocheck.Equals, "elb-1")
	c.Assert(values.Get("LoadBalancerNames.member.2"), gocheck.Equals, "elb-2")
	c.Assert(values.Get("Tags.member.1.Key"), gocheck.Equals, "foo")
	c.Assert(values.Get("Tags.member.1.Value"), gocheck.Equals, "bar")
	c.Assert(values.Get("Tags.member.2.Key"), gocheck.Equals, "baz")
	c.Assert(values.Get("Tags.member.2.Value"), gocheck.Equals, "qux")
	c.Assert(values.Get("VPCZoneIdentifier"), gocheck.Equals, "subnet-610acd08,subnet-530fc83a")
	c.Assert(resp.RequestId, gocheck.Equals, "8d798a29-f083-11e1-bdfb-cb223EXAMPLE")
}

func (s *S) TestDeleteAutoScalingGroup(c *gocheck.C) {
	testServer.Response(200, nil, DeleteAutoScalingGroup)
	resp, err := s.as.DeleteAutoScalingGroup("my-test-asg", true)
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "DeleteAutoScalingGroup")
	c.Assert(values.Get("AutoScalingGroupName"), gocheck.Equals, "my-test-asg")
	c.Assert(resp.RequestId, gocheck.Equals, "70a76d42-9665-11e2-9fdf-211deEXAMPLE")
}

func (s *S) TestDeleteAutoScalingGroupWithExistingInstances(c *gocheck.C) {
	testServer.Response(400, nil, DeleteAutoScalingGroupError)
	resp, err := s.as.DeleteAutoScalingGroup("my-test-asg", false)
	testServer.WaitRequest()
	c.Assert(resp, gocheck.IsNil)
	c.Assert(err, gocheck.NotNil)
	e, ok := err.(*autoscaling.Error)
	c.Assert(ok, gocheck.Equals, true)
	c.Assert(e.Message, gocheck.Equals, "You cannot delete an AutoScalingGroup while there are instances or pending Spot instance request(s) still in the group.")
	c.Assert(e.Code, gocheck.Equals, "ResourceInUse")
	c.Assert(e.StatusCode, gocheck.Equals, 400)
	c.Assert(e.RequestId, gocheck.Equals, "70a76d42-9665-11e2-9fdf-211deEXAMPLE")
}

func (s *S) TestAttachInstances(c *gocheck.C) {
	testServer.Response(200, nil, AttachInstances)
	resp, err := s.as.AttachInstances("my-test-asg", []string{"i-21321afs", "i-baaffg23"})
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "AttachInstances")
	c.Assert(values.Get("AutoScalingGroupName"), gocheck.Equals, "my-test-asg")
	c.Assert(values.Get("InstanceIds.member.1"), gocheck.Equals, "i-21321afs")
	c.Assert(values.Get("InstanceIds.member.2"), gocheck.Equals, "i-baaffg23")
	c.Assert(resp.RequestId, gocheck.Equals, "8d798a29-f083-11e1-bdfb-cb223EXAMPLE")
}

func (s *S) TestCreateOrUpdateTags(c *gocheck.C) {
	testServer.Response(200, nil, CreateOrUpdateTags)
	tags := []autoscaling.Tag{
		{
			Key:        "foo",
			Value:      "bar",
			ResourceId: "my-test-asg",
		},
		{
			Key:               "baz",
			Value:             "qux",
			ResourceId:        "my-test-asg",
			PropagateAtLaunch: true,
		},
	}
	resp, err := s.as.CreateOrUpdateTags(tags)
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "CreateOrUpdateTags")
	c.Assert(values.Get("Tags.member.1.Key"), gocheck.Equals, "foo")
	c.Assert(values.Get("Tags.member.1.Value"), gocheck.Equals, "bar")
	c.Assert(values.Get("Tags.member.1.ResourceId"), gocheck.Equals, "my-test-asg")
	c.Assert(values.Get("Tags.member.2.Key"), gocheck.Equals, "baz")
	c.Assert(values.Get("Tags.member.2.Value"), gocheck.Equals, "qux")
	c.Assert(values.Get("Tags.member.2.ResourceId"), gocheck.Equals, "my-test-asg")
	c.Assert(values.Get("Tags.member.2.PropagateAtLaunch"), gocheck.Equals, "true")
	c.Assert(resp.RequestId, gocheck.Equals, "b0203919-bf1b-11e2-8a01-13263EXAMPLE")
}

func (s *S) TestDeleteTags(c *gocheck.C) {
	testServer.Response(200, nil, DeleteTags)
	tags := []autoscaling.Tag{
		{
			Key:        "foo",
			Value:      "bar",
			ResourceId: "my-test-asg",
		},
		{
			Key:               "baz",
			Value:             "qux",
			ResourceId:        "my-test-asg",
			PropagateAtLaunch: true,
		},
	}
	resp, err := s.as.DeleteTags(tags)
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "DeleteTags")
	c.Assert(values.Get("Tags.member.1.Key"), gocheck.Equals, "foo")
	c.Assert(values.Get("Tags.member.1.Value"), gocheck.Equals, "bar")
	c.Assert(values.Get("Tags.member.1.ResourceId"), gocheck.Equals, "my-test-asg")
	c.Assert(values.Get("Tags.member.2.Key"), gocheck.Equals, "baz")
	c.Assert(values.Get("Tags.member.2.Value"), gocheck.Equals, "qux")
	c.Assert(values.Get("Tags.member.2.ResourceId"), gocheck.Equals, "my-test-asg")
	c.Assert(values.Get("Tags.member.2.PropagateAtLaunch"), gocheck.Equals, "true")
	c.Assert(resp.RequestId, gocheck.Equals, "b0203919-bf1b-11e2-8a01-13263EXAMPLE")
}

func (s *S) TestDescribeAccountLimits(c *gocheck.C) {
	testServer.Response(200, nil, DescribeAccountLimits)

	resp, err := s.as.DescribeAccountLimits()
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "DescribeAccountLimits")
	c.Assert(resp.RequestId, gocheck.Equals, "a32bd184-519d-11e3-a8a4-c1c467cbcc3b")
	c.Assert(resp.MaxNumberOfAutoScalingGroups, gocheck.Equals, 20)
	c.Assert(resp.MaxNumberOfLaunchConfigurations, gocheck.Equals, 100)

}

func (s *S) TestDescribeAdjustmentTypes(c *gocheck.C) {
	testServer.Response(200, nil, DescribeAdjustmentTypes)
	resp, err := s.as.DescribeAdjustmentTypes()
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "DescribeAdjustmentTypes")
	c.Assert(resp.RequestId, gocheck.Equals, "cc5f0337-b694-11e2-afc0-6544dEXAMPLE")
	c.Assert(resp.AdjustmentTypes, gocheck.DeepEquals, []autoscaling.AdjustmentType{{"ChangeInCapacity"}, {"ExactCapacity"}, {"PercentChangeInCapacity"}})
}

func (s *S) TestDescribeAutoScalingGroups(c *gocheck.C) {
	testServer.Response(200, nil, DescribeAutoScalingGroups)
	resp, err := s.as.DescribeAutoScalingGroups([]string{"my-test-asg-lbs"}, 0, "")
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	t, _ := time.Parse(time.RFC3339, "2013-05-06T17:47:15.107Z")
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "DescribeAutoScalingGroups")
	c.Assert(values.Get("AutoScalingGroupNames.member.1"), gocheck.Equals, "my-test-asg-lbs")

	expected := &autoscaling.DescribeAutoScalingGroupsResp{
		AutoScalingGroups: []autoscaling.AutoScalingGroup{
			{
				AutoScalingGroupName: "my-test-asg-lbs",
				Tags: []autoscaling.Tag{
					{
						Key:               "foo",
						Value:             "bar",
						ResourceId:        "my-test-asg-lbs",
						PropagateAtLaunch: true,
						ResourceType:      "auto-scaling-group",
					},
					{
						Key:               "baz",
						Value:             "qux",
						ResourceId:        "my-test-asg-lbs",
						PropagateAtLaunch: true,
						ResourceType:      "auto-scaling-group",
					},
				},
				Instances: []autoscaling.Instance{
					{
						AvailabilityZone:        "us-east-1b",
						HealthStatus:            "Healthy",
						InstanceId:              "i-zb1f313",
						LaunchConfigurationName: "my-test-lc",
						LifecycleState:          "InService",
					},
					{
						AvailabilityZone:        "us-east-1a",
						HealthStatus:            "Healthy",
						InstanceId:              "i-90123adv",
						LaunchConfigurationName: "my-test-lc",
						LifecycleState:          "InService",
					},
				},
				HealthCheckType:         "ELB",
				CreatedTime:             t,
				LaunchConfigurationName: "my-test-lc",
				DesiredCapacity:         2,
				AvailabilityZones:       []string{"us-east-1b", "us-east-1a"},
				LoadBalancerNames:       []string{"my-test-asg-loadbalancer"},
				MinSize:                 2,
				MaxSize:                 10,
				VPCZoneIdentifier:       "subnet-32131da1,subnet-1312dad2",
				HealthCheckGracePeriod:  120,
				DefaultCooldown:         300,
				AutoScalingGroupARN:     "arn:aws:autoscaling:us-east-1:803981987763:autoScalingGroup:ca861182-c8f9-4ca7-b1eb-cd35505f5ebb:autoScalingGroupName/my-test-asg-lbs",
				TerminationPolicies:     []string{"Default"},
			},
		},
		RequestId: "0f02a07d-b677-11e2-9eb0-dd50EXAMPLE",
	}
	c.Assert(resp, gocheck.DeepEquals, expected)
}

func (s *S) TestDescribeAutoScalingInstances(c *gocheck.C) {
	testServer.Response(200, nil, DescribeAutoScalingInstances)
	resp, err := s.as.DescribeAutoScalingInstances([]string{"i-78e0d40b"}, 0, "")
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "DescribeAutoScalingInstances")
	c.Assert(resp.RequestId, gocheck.Equals, "df992dc3-b72f-11e2-81e1-750aa6EXAMPLE")
	c.Assert(resp.AutoScalingInstances, gocheck.DeepEquals, []autoscaling.Instance{
		{
			AutoScalingGroupName:    "my-test-asg",
			AvailabilityZone:        "us-east-1a",
			HealthStatus:            "Healthy",
			InstanceId:              "i-78e0d40b",
			LaunchConfigurationName: "my-test-lc",
			LifecycleState:          "InService",
		},
	})
}

func (s *S) TestDescribeLaunchConfigurations(c *gocheck.C) {
	testServer.Response(200, nil, DescribeLaunchConfigurations)
	resp, err := s.as.DescribeLaunchConfigurations([]string{"my-test-lc"}, 0, "")
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	t, _ := time.Parse(time.RFC3339, "2013-01-21T23:04:42.200Z")
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "DescribeLaunchConfigurations")
	c.Assert(values.Get("LaunchConfigurationNames.member.1"), gocheck.Equals, "my-test-lc")
	expected := &autoscaling.DescribeLaunchConfigurationsResp{
		LaunchConfigurations: []autoscaling.LaunchConfiguration{
			{
				AssociatePublicIpAddress: true,
				BlockDeviceMappings: []autoscaling.BlockDeviceMapping{
					{
						DeviceName:  "/dev/sdb",
						VirtualName: "ephemeral0",
					},
					{
						DeviceName: "/dev/sdf",
						Ebs: autoscaling.EBS{
							SnapshotId: "snap-XXXXYYY",
							VolumeSize: 100,
						},
					},
				},
				EbsOptimized:            false,
				CreatedTime:             t,
				LaunchConfigurationName: "my-test-lc",
				InstanceType:            "m1.small",
				ImageId:                 "ami-514ac838",
				InstanceMonitoring:      autoscaling.InstanceMonitoring{Enabled: true},
				LaunchConfigurationARN:  "arn:aws:autoscaling:us-east-1:803981987763:launchConfiguration:9dbbbf87-6141-428a-a409-0752edbe6cad:launchConfigurationName/my-test-lc",
			},
		},
		RequestId: "d05a22f8-b690-11e2-bf8e-2113fEXAMPLE",
	}
	c.Assert(resp, gocheck.DeepEquals, expected)
}

func (s *S) TestDescribeMetricCollectionTypes(c *gocheck.C) {
	testServer.Response(200, nil, DescribeMetricCollectionTypes)
	resp, err := s.as.DescribeMetricCollectionTypes()
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "DescribeMetricCollectionTypes")
	c.Assert(resp.RequestId, gocheck.Equals, "07f3fea2-bf3c-11e2-9b6f-f3cdbb80c073")
	c.Assert(resp.Metrics, gocheck.DeepEquals, []autoscaling.MetricCollection{
		{
			Metric: "GroupMinSize",
		},
		{
			Metric: "GroupMaxSize",
		},
		{
			Metric: "GroupDesiredCapacity",
		},
		{
			Metric: "GroupInServiceInstances",
		},
		{
			Metric: "GroupPendingInstances",
		},
		{
			Metric: "GroupTerminatingInstances",
		},
		{
			Metric: "GroupTotalInstances",
		},
	})
	c.Assert(resp.Granularities, gocheck.DeepEquals, []autoscaling.MetricGranularity{
		{
			Granularity: "1Minute",
		},
	})
}

func (s *S) TestDescribeNotificationConfigurations(c *gocheck.C) {
	testServer.Response(200, nil, DescribeNotificationConfigurations)
	resp, err := s.as.DescribeNotificationConfigurations([]string{"i-78e0d40b"}, 0, "")
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "DescribeNotificationConfigurations")
	c.Assert(resp.RequestId, gocheck.Equals, "07f3fea2-bf3c-11e2-9b6f-f3cdbb80c073")
	c.Assert(resp.NotificationConfigurations, gocheck.DeepEquals, []autoscaling.NotificationConfiguration{
		{
			AutoScalingGroupName: "my-test-asg",
			NotificationType:     "autoscaling: EC2_INSTANCE_LAUNCH",
			TopicARN:             "vajdoafj231j41231/topic",
		},
	})
}

func (s *S) TestDescribePolicies(c *gocheck.C) {
	testServer.Response(200, nil, DescribePolicies)
	resp, err := s.as.DescribePolicies("my-test-asg", []string{}, 2, "")
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "DescribePolicies")
	c.Assert(values.Get("MaxRecords"), gocheck.Equals, "2")
	expected := &autoscaling.DescribePoliciesResp{
		RequestId: "ec3bffad-b739-11e2-b38d-15fbEXAMPLE",
		NextToken: "3ef417fe-9202-12-8ddd-d13e1313413",
		ScalingPolicies: []autoscaling.ScalingPolicy{
			{
				PolicyARN:            "arn:aws:autoscaling:us-east-1:803981987763:scalingPolicy:c322761b-3172-4d56-9a21-0ed9d6161d67:autoScalingGroupName/my-test-asg:policyName/MyScaleDownPolicy",
				AdjustmentType:       "ChangeInCapacity",
				ScalingAdjustment:    -1,
				PolicyName:           "MyScaleDownPolicy",
				AutoScalingGroupName: "my-test-asg",
				Cooldown:             60,
				Alarms: []autoscaling.Alarm{
					{
						AlarmName: "TestQueue",
						AlarmARN:  "arn:aws:cloudwatch:us-east-1:803981987763:alarm:TestQueue",
					},
				},
			},
			{
				PolicyARN:            "arn:aws:autoscaling:us-east-1:803981987763:scalingPolicy:c55a5cdd-9be0-435b-b60b-a8dd313159f5:autoScalingGroupName/my-test-asg:policyName/MyScaleUpPolicy",
				AdjustmentType:       "ChangeInCapacity",
				ScalingAdjustment:    1,
				PolicyName:           "MyScaleUpPolicy",
				AutoScalingGroupName: "my-test-asg",
				Cooldown:             60,
				Alarms: []autoscaling.Alarm{
					{
						AlarmName: "TestQueue",
						AlarmARN:  "arn:aws:cloudwatch:us-east-1:803981987763:alarm:TestQueue",
					},
				},
			},
		},
	}
	c.Assert(resp, gocheck.DeepEquals, expected)
}

func (s *S) TestDescribeScalingActivities(c *gocheck.C) {
	testServer.Response(200, nil, DescribeScalingActivities)
	resp, err := s.as.DescribeScalingActivities("my-test-asg", []string{}, 1, "")
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "DescribeScalingActivities")
	c.Assert(values.Get("MaxRecords"), gocheck.Equals, "1")
	c.Assert(values.Get("AutoScalingGroupName"), gocheck.Equals, "my-test-asg")
	st, _ := time.Parse(time.RFC3339, "2012-04-12T17:32:07.882Z")
	et, _ := time.Parse(time.RFC3339, "2012-04-12T17:32:08Z")
	expected := &autoscaling.DescribeScalingActivitiesResp{
		RequestId: "7a641adc-84c5-11e1-a8a5-217ebEXAMPLE",
		NextToken: "3ef417fe-9202-12-8ddd-d13e1313413",
		Activities: []autoscaling.Activity{
			{
				StatusCode:           "Failed",
				Progress:             0,
				ActivityId:           "063308ae-aa22-4a9b-94f4-9faeEXAMPLE",
				StartTime:            st,
				AutoScalingGroupName: "my-test-asg",
				Details:              "{}",
				Cause:                "At 2012-04-12T17:31:30Z a user request created an AutoScalingGroup changing the desired capacity from 0 to 1.  At 2012-04-12T17:32:07Z an instance was started in response to a difference between desired and actual capacity, increasing the capacity from 0 to 1.",
				Description:          "Launching a new EC2 instance.  Status Reason: The image id 'ami-4edb0327' does not exist. Launching EC2 instance failed.",
				EndTime:              et,
				StatusMessage:        "The image id 'ami-4edb0327' does not exist. Launching EC2 instance failed.",
			},
		},
	}
	c.Assert(resp, gocheck.DeepEquals, expected)
}

func (s *S) TestDescribeScalingProcessTypes(c *gocheck.C) {
	testServer.Response(200, nil, DescribeScalingProcessTypes)
	resp, err := s.as.DescribeScalingProcessTypes()
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "DescribeScalingProcessTypes")
	c.Assert(resp.RequestId, gocheck.Equals, "27f2eacc-b73f-11e2-ad99-c7aba3a9c963")
	c.Assert(resp.Processes, gocheck.DeepEquals, []autoscaling.ProcessType{
		{"AZRebalance"},
		{"AddToLoadBalancer"},
		{"AlarmNotification"},
		{"HealthCheck"},
		{"Launch"},
		{"ReplaceUnhealthy"},
		{"ScheduledActions"},
		{"Terminate"},
	})
}

func (s *S) TestDescribeTerminationPolicyTypes(c *gocheck.C) {
	testServer.Response(200, nil, DescribeTerminationPolicyTypes)
	resp, err := s.as.DescribeTerminationPolicyTypes()
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "DescribeTerminationPolicyTypes")
	c.Assert(resp.RequestId, gocheck.Equals, "d9a05827-b735-11e2-a40c-c79a5EXAMPLE")
	c.Assert(resp.TerminationPolicyTypes, gocheck.DeepEquals, []string{"ClosestToNextInstanceHour", "Default", "NewestInstance", "OldestInstance", "OldestLaunchConfiguration"})
}

func (s *S) TestDescribeTags(c *gocheck.C) {
	testServer.Response(200, nil, DescribeTags)
	filter := autoscaling.NewFilter()
	filter.Add("auto-scaling-group", "my-test-asg")
	resp, err := s.as.DescribeTags(filter, 1, "")
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "DescribeTags")
	c.Assert(values.Get("MaxRecords"), gocheck.Equals, "1")
	c.Assert(values.Get("Filters.member.1.Name"), gocheck.Equals, "auto-scaling-group")
	c.Assert(values.Get("Filters.member.1.Values.member.1"), gocheck.Equals, "my-test-asg")
	c.Assert(resp.RequestId, gocheck.Equals, "086265fd-bf3e-11e2-85fc-fbb1EXAMPLE")
	c.Assert(resp.Tags, gocheck.DeepEquals, []autoscaling.Tag{
		{
			Key:               "version",
			Value:             "1.0",
			ResourceId:        "my-test-asg",
			PropagateAtLaunch: true,
			ResourceType:      "auto-scaling-group",
		},
	})
}

func (s *S) TestDisableMetricsCollection(c *gocheck.C) {
	testServer.Response(200, nil, DisableMetricsCollection)
	resp, err := s.as.DisableMetricsCollection("my-test-asg", []string{"GroupMinSize"})
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "DisableMetricsCollection")
	c.Assert(values.Get("AutoScalingGroupName"), gocheck.Equals, "my-test-asg")
	c.Assert(values.Get("Metrics.member.1"), gocheck.Equals, "GroupMinSize")
	c.Assert(resp.RequestId, gocheck.Equals, "8d798a29-f083-11e1-bdfb-cb223EXAMPLE")
}

func (s *S) TestEnableMetricsCollection(c *gocheck.C) {
	testServer.Response(200, nil, DisableMetricsCollection)
	resp, err := s.as.EnableMetricsCollection("my-test-asg", []string{"GroupMinSize", "GroupMaxSize"}, "1Minute")
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "EnableMetricsCollection")
	c.Assert(values.Get("AutoScalingGroupName"), gocheck.Equals, "my-test-asg")
	c.Assert(values.Get("Granularity"), gocheck.Equals, "1Minute")
	c.Assert(values.Get("Metrics.member.1"), gocheck.Equals, "GroupMinSize")
	c.Assert(values.Get("Metrics.member.2"), gocheck.Equals, "GroupMaxSize")
	c.Assert(resp.RequestId, gocheck.Equals, "8d798a29-f083-11e1-bdfb-cb223EXAMPLE")
}

func (s *S) TestExecutePolicy(c *gocheck.C) {
	testServer.Response(200, nil, ExecutePolicy)
	resp, err := s.as.ExecutePolicy("my-scaleout-policy", "my-test-asg",true)
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "ExecutePolicy")
	c.Assert(values.Get("AutoScalingGroupName"), gocheck.Equals, "my-test-asg")
	c.Assert(values.Get("PolicyName"), gocheck.Equals, "my-scaleout-policy")
	c.Assert(values.Get("HonorCooldown"), gocheck.Equals, "true")
	c.Assert(resp.RequestId, gocheck.Equals, "8d798a29-f083-11e1-bdfb-cb223EXAMPLE")
}

func (s *S) TestPutNotificationConfiguration(c *gocheck.C) {
	testServer.Response(200, nil, PutNotificationConfiguration)
	resp, err := s.as.PutNotificationConfiguration("my-test-asg", []string{"autoscaling:EC2_INSTANCE_LAUNCH", "autoscaling:EC2_INSTANCE_LAUNCH_ERROR"}, "myTopicARN")
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "PutNotificationConfiguration")
	c.Assert(values.Get("AutoScalingGroupName"), gocheck.Equals, "my-test-asg")
	c.Assert(values.Get("TopicARN"), gocheck.Equals, "myTopicARN")
	c.Assert(values.Get("NotificationTypes.member.1"), gocheck.Equals, "autoscaling:EC2_INSTANCE_LAUNCH")
	c.Assert(values.Get("NotificationTypes.member.2"), gocheck.Equals, "autoscaling:EC2_INSTANCE_LAUNCH_ERROR")
	c.Assert(resp.RequestId, gocheck.Equals, "8d798a29-f083-11e1-bdfb-cb223EXAMPLE")
}

func (s *S) TestPutScalingPolicy(c *gocheck.C) {
	testServer.Response(200, nil, PutScalingPolicy)
	resp, err := s.as.PutScalingPolicy("my-test-asg","my-scaleout-policy" ,30, "PercentChangeInCapacity", 0, 0)
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "PutScalingPolicy")
	c.Assert(values.Get("AutoScalingGroupName"), gocheck.Equals, "my-test-asg")
	c.Assert(values.Get("PolicyName"), gocheck.Equals, "my-scaleout-policy")
	c.Assert(values.Get("AdjustmentType"), gocheck.Equals, "PercentChangeInCapacity")
	c.Assert(values.Get("ScalingAdjustment"), gocheck.Equals, "30")
	c.Assert(resp.RequestId, gocheck.Equals, "3cfc6fef-c08b-11e2-a697-2922EXAMPLE")
	c.Assert(resp.PolicyARN, gocheck.Equals, "arn:aws:autoscaling:us-east-1:803981987763:scalingPolicy:b0dcf5e8-02e6-4e31-9719-0675d0dc31ae:autoScalingGroupName/my-test-asg:policyName/my-scaleout-policy")
}

func (s *S) TestPutScheduledUpdateGroupAction(c *gocheck.C) {
	testServer.Response(200, nil, PutScheduledUpdateGroupAction)
	st, _ := time.Parse(time.RFC3339, "2013-05-25T08:00:00Z")
	request := &autoscaling.PutScheduledUpdateGroupAction{
		AutoScalingGroupName: "my-test-asg",
		DesiredCapacity:      3,
		ScheduledActionName:  "ScaleUp",
		StartTime:            st,
	}
	resp, err := s.as.PutScheduledUpdateGroupAction(request)
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "PutScheduledUpdateGroupAction")
	c.Assert(values.Get("AutoScalingGroupName"), gocheck.Equals, "my-test-asg")
	c.Assert(values.Get("ScheduledActionName"), gocheck.Equals, "ScaleUp")
	c.Assert(values.Get("DesiredCapacity"), gocheck.Equals, "3")
	c.Assert(values.Get("StartTime"), gocheck.Equals, "2013-05-25T08:00:00Z")
	c.Assert(resp.RequestId, gocheck.Equals, "3bc8c9bc-6a62-11e2-8a51-4b8a1EXAMPLE")
}

func (s *S) TestPutScheduledUpdateGroupActionCron(c *gocheck.C) {
	testServer.Response(200, nil, PutScheduledUpdateGroupAction)
	st, _ := time.Parse(time.RFC3339, "2013-05-25T08:00:00Z")
	request := &autoscaling.PutScheduledUpdateGroupAction{
		AutoScalingGroupName: "my-test-asg",
		DesiredCapacity:      3,
		ScheduledActionName:  "scaleup-schedule-year",
		StartTime:            st,
		Recurrence:           "30 0 1 1,6,12 *",
	}
	resp, err := s.as.PutScheduledUpdateGroupAction(request)
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "PutScheduledUpdateGroupAction")
	c.Assert(values.Get("AutoScalingGroupName"), gocheck.Equals, "my-test-asg")
	c.Assert(values.Get("ScheduledActionName"), gocheck.Equals, "scaleup-schedule-year")
	c.Assert(values.Get("DesiredCapacity"), gocheck.Equals, "3")
	c.Assert(values.Get("Recurrence"), gocheck.Equals, "30 0 1 1,6,12 *")
	c.Assert(resp.RequestId, gocheck.Equals, "3bc8c9bc-6a62-11e2-8a51-4b8a1EXAMPLE")

}

func (s *S) TestResumeProcesses(c *gocheck.C) {
	testServer.Response(200, nil, ResumeProcesses)
	resp, err := s.as.ResumeProcesses("my-test-asg", []string{"Launch", "Terminate"})
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "ResumeProcesses")
	c.Assert(values.Get("AutoScalingGroupName"), gocheck.Equals, "my-test-asg")
	c.Assert(values.Get("ScalingProcesses.member.1"), gocheck.Equals, "Launch")
	c.Assert(values.Get("ScalingProcesses.member.2"), gocheck.Equals, "Terminate")
	c.Assert(resp.RequestId, gocheck.Equals, "8d798a29-f083-11e1-bdfb-cb223EXAMPLE")

}

func (s *S) TestSetDesiredCapacity(c *gocheck.C) {
	testServer.Response(200, nil, SetDesiredCapacity)
	resp, err := s.as.SetDesiredCapacity("my-test-asg", 3, true)
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "SetDesiredCapacity")
	c.Assert(values.Get("AutoScalingGroupName"), gocheck.Equals, "my-test-asg")
	c.Assert(values.Get("HonorCooldown"), gocheck.Equals, "true")
	c.Assert(values.Get("DesiredCapacity"), gocheck.Equals, "3")
	c.Assert(resp.RequestId, gocheck.Equals, "9fb7e2db-6998-11e2-a985-57c82EXAMPLE")
}

func (s *S) TestSetInstanceHealth(c *gocheck.C) {
	testServer.Response(200, nil, SetInstanceHealth)
	resp, err := s.as.SetInstanceHealth("i-baha3121", "Unhealthy", false)
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "SetInstanceHealth")
	c.Assert(values.Get("HealthStatus"), gocheck.Equals, "Unhealthy")
	c.Assert(values.Get("InstanceId"), gocheck.Equals, "i-baha3121")
	c.Assert(values.Get("ShouldRespectGracePeriod"), gocheck.Equals, "false")
	c.Assert(resp.RequestId, gocheck.Equals, "9fb7e2db-6998-11e2-a985-57c82EXAMPLE")
}

func (s *S) TestSuspendProcesses(c *gocheck.C) {
	testServer.Response(200, nil, SuspendProcesses)
	resp, err := s.as.SuspendProcesses("my-test-asg", []string{"Launch", "Terminate"})
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "SuspendProcesses")
	c.Assert(values.Get("AutoScalingGroupName"), gocheck.Equals, "my-test-asg")
	c.Assert(values.Get("ScalingProcesses.member.1"), gocheck.Equals, "Launch")
	c.Assert(values.Get("ScalingProcesses.member.2"), gocheck.Equals, "Terminate")
	c.Assert(resp.RequestId, gocheck.Equals, "8d798a29-f083-11e1-bdfb-cb223EXAMPLE")
}

func (s *S) TestTerminateInstanceInAutoScalingGroup(c *gocheck.C) {
	testServer.Response(200, nil, TerminateInstanceInAutoScalingGroup)
	st, _ := time.Parse(time.RFC3339, "2014-01-26T14:08:30.560Z")
	resp, err := s.as.TerminateInstanceInAutoScalingGroup("i-br234123", false)
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "TerminateInstanceInAutoScalingGroup")
	c.Assert(values.Get("InstanceId"), gocheck.Equals, "i-br234123")
	c.Assert(values.Get("ShouldDecrementDesiredCapacity"), gocheck.Equals, "false")
	expected := &autoscaling.TerminateInstanceInAutoScalingGroupResp{
		Activity: autoscaling.Activity{
			ActivityId:  "cczc44a87-7d04-dsa15-31-d27c219864c5",
			Cause:       "At 2014-01-26T14:08:30Z instance i-br234123 was taken out of service in response to a user request.",
			Description: "Terminating EC2 instance: i-br234123",
			Details:     "{\"Availability Zone\":\"us-east-1b\"}",
			Progress:    0,
			StartTime:   st,
			StatusCode:  "InProgress",
		},
		RequestId: "8d798a29-f083-11e1-bdfb-cb223EXAMPLE",
	}
	c.Assert(resp, gocheck.DeepEquals, expected)
}

func (s *S) TestUpdateAutoScalingGroup(c *gocheck.C) {
	testServer.Response(200, nil, UpdateAutoScalingGroup)

	asg := &autoscaling.UpdateAutoScalingGroup{
		AutoScalingGroupName:    "my-test-asg",
		AvailabilityZones:       []string{"us-east-1a", "us-east-1b"},
		MinSize:                 3,
		MaxSize:                 3,
		DefaultCooldown:         600,
		DesiredCapacity:         3,
		LaunchConfigurationName: "my-test-lc",
		VPCZoneIdentifier:       "subnet-610acd08,subnet-530fc83a",
	}
	resp, err := s.as.UpdateAutoScalingGroup(asg)
	c.Assert(err, gocheck.IsNil)
	values := testServer.WaitRequest().PostForm
	c.Assert(values.Get("Version"), gocheck.Equals, "2011-01-01")
	c.Assert(values.Get("Action"), gocheck.Equals, "UpdateAutoScalingGroup")
	c.Assert(values.Get("AutoScalingGroupName"), gocheck.Equals, "my-test-asg")
	c.Assert(values.Get("AvailabilityZones.member.1"), gocheck.Equals, "us-east-1a")
	c.Assert(values.Get("AvailabilityZones.member.2"), gocheck.Equals, "us-east-1b")
	c.Assert(values.Get("MinSize"), gocheck.Equals, "3")
	c.Assert(values.Get("MaxSize"), gocheck.Equals, "3")
	c.Assert(values.Get("DefaultCooldown"), gocheck.Equals, "600")
	c.Assert(values.Get("DesiredCapacity"), gocheck.Equals, "3")
	c.Assert(values.Get("LaunchConfigurationName"), gocheck.Equals, "my-test-lc")
	c.Assert(values.Get("VPCZoneIdentifier"), gocheck.Equals, "subnet-610acd08,subnet-530fc83a")
	c.Assert(resp.RequestId, gocheck.Equals, "8d798a29-f083-11e1-bdfb-cb223EXAMPLE")
}
