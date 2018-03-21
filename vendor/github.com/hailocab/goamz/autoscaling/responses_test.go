package autoscaling_test

var AttachInstances = `
<AttachInstancesResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
	<ResponseMetadata>
		<RequestId>8d798a29-f083-11e1-bdfb-cb223EXAMPLE</RequestId>
	</ResponseMetadata>
</AttachInstancesResponse> 
`

var CreateAutoScalingGroup = `
<CreateAutoScalingGroupResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
	<ResponseMetadata>
		<RequestId>8d798a29-f083-11e1-bdfb-cb223EXAMPLE</RequestId>
	</ResponseMetadata>
</CreateAutoScalingGroupResponse> 
`

var CreateLaunchConfiguration = `
<CreateLaunchConfigurationResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
<ResponseMetadata>
   <RequestId>7c6e177f-f082-11e1-ac58-3714bEXAMPLE</RequestId>
</ResponseMetadata>
</CreateLaunchConfigurationResponse> 
`
var DeleteAutoScalingGroup = `
 <DeleteAutoScalingGroupResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
  <ResponseMetadata>
    <RequestId>70a76d42-9665-11e2-9fdf-211deEXAMPLE</RequestId>
  </ResponseMetadata>
 </DeleteAutoScalingGroupResponse> 
`
var DeleteAutoScalingGroupError = `
<ErrorResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
  <Error>
    <Type>Sender</Type>
    <Code>ResourceInUse</Code>
    <Message>You cannot delete an AutoScalingGroup while there are instances or pending Spot instance request(s) still in the group.</Message>
  </Error>
  <RequestId>70a76d42-9665-11e2-9fdf-211deEXAMPLE</RequestId>
</ErrorResponse>
`
var DeleteLaunchConfiguration = `
<DeleteLaunchConfigurationResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
  <ResponseMetadata>
    <RequestId>7347261f-97df-11e2-8756-35eEXAMPLE</RequestId>
  </ResponseMetadata>
</DeleteLaunchConfigurationResponse> 
`
var DeleteLaunchConfigurationInUse = `
<ErrorResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
  <Error>
    <Type>Sender</Type>
    <Code>ResourceInUse</Code>
    <Message>Cannot delete launch configuration my-test-lc because it is attached to AutoScalingGroup test</Message>
  </Error>
  <RequestId>7347261f-97df-11e2-8756-35eEXAMPLE</RequestId>
</ErrorResponse>
`
var CreateOrUpdateTags = `
<CreateOrUpdateTagsResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
  <ResponseMetadata>
    <RequestId>b0203919-bf1b-11e2-8a01-13263EXAMPLE</RequestId>
  </ResponseMetadata>
</CreateOrUpdateTagsResponse>
`
var DeleteTags = `
<CreateOrUpdateTagsResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
  <ResponseMetadata>
    <RequestId>b0203919-bf1b-11e2-8a01-13263EXAMPLE</RequestId>
  </ResponseMetadata>
</CreateOrUpdateTagsResponse>
`
var DescribeAccountLimits = `
<DescribeAccountLimitsResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
  <DescribeAccountLimitsResult>
    <MaxNumberOfLaunchConfigurations>100</MaxNumberOfLaunchConfigurations>
    <MaxNumberOfAutoScalingGroups>20</MaxNumberOfAutoScalingGroups>
  </DescribeAccountLimitsResult>
  <ResponseMetadata>
    <RequestId>a32bd184-519d-11e3-a8a4-c1c467cbcc3b</RequestId>
  </ResponseMetadata>
</DescribeAccountLimitsResponse> 
`
var DescribeAdjustmentTypes = `
<DescribeAdjustmentTypesResponse xmlns="http://autoscaling.amazonaws.com/doc/201-01-01/">
  <DescribeAdjustmentTypesResult>
    <AdjustmentTypes>
      <member>
        <AdjustmentType>ChangeInCapacity</AdjustmentType>
      </member>
      <member>
        <AdjustmentType>ExactCapacity</AdjustmentType>
      </member>
      <member>
        <AdjustmentType>PercentChangeInCapacity</AdjustmentType>
      </member>
    </AdjustmentTypes>
  </DescribeAdjustmentTypesResult>
  <ResponseMetadata>
    <RequestId>cc5f0337-b694-11e2-afc0-6544dEXAMPLE</RequestId>
  </ResponseMetadata>
</DescribeAdjustmentTypesResponse> 
`
var DescribeAutoScalingGroups = `
<DescribeAutoScalingGroupsResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
<DescribeAutoScalingGroupsResult>
    <AutoScalingGroups>
      <member>
         <Tags>
           <member>
              <ResourceId>my-test-asg-lbs</ResourceId>
              <PropagateAtLaunch>true</PropagateAtLaunch>
              <Value>bar</Value>
              <Key>foo</Key>
              <ResourceType>auto-scaling-group</ResourceType>
            </member>
            <member>
              <ResourceId>my-test-asg-lbs</ResourceId>
              <PropagateAtLaunch>true</PropagateAtLaunch>
              <Value>qux</Value>
              <Key>baz</Key>
              <ResourceType>auto-scaling-group</ResourceType>
            </member>
        </Tags>
        <SuspendedProcesses/>
        <AutoScalingGroupName>my-test-asg-lbs</AutoScalingGroupName>
        <HealthCheckType>ELB</HealthCheckType>
        <CreatedTime>2013-05-06T17:47:15.107Z</CreatedTime>
        <EnabledMetrics/>
        <LaunchConfigurationName>my-test-lc</LaunchConfigurationName>
        <Instances>
          <member>
            <HealthStatus>Healthy</HealthStatus>
            <AvailabilityZone>us-east-1b</AvailabilityZone>
            <InstanceId>i-zb1f313</InstanceId>
            <LaunchConfigurationName>my-test-lc</LaunchConfigurationName>
            <LifecycleState>InService</LifecycleState>
          </member>
          <member>
            <HealthStatus>Healthy</HealthStatus>
            <AvailabilityZone>us-east-1a</AvailabilityZone>
            <InstanceId>i-90123adv</InstanceId>
            <LaunchConfigurationName>my-test-lc</LaunchConfigurationName>
            <LifecycleState>InService</LifecycleState>
          </member>
        </Instances>
        <DesiredCapacity>2</DesiredCapacity>
        <AvailabilityZones>
          <member>us-east-1b</member>
          <member>us-east-1a</member>
        </AvailabilityZones>
        <LoadBalancerNames>
          <member>my-test-asg-loadbalancer</member>
        </LoadBalancerNames>
        <MinSize>2</MinSize>
        <VPCZoneIdentifier>subnet-32131da1,subnet-1312dad2</VPCZoneIdentifier>
        <HealthCheckGracePeriod>120</HealthCheckGracePeriod>
        <DefaultCooldown>300</DefaultCooldown>
        <AutoScalingGroupARN>arn:aws:autoscaling:us-east-1:803981987763:autoScalingGroup:ca861182-c8f9-4ca7-b1eb-cd35505f5ebb:autoScalingGroupName/my-test-asg-lbs</AutoScalingGroupARN>
        <TerminationPolicies>
          <member>Default</member>
        </TerminationPolicies>
        <MaxSize>10</MaxSize>
      </member>
    </AutoScalingGroups>
  </DescribeAutoScalingGroupsResult>
  <ResponseMetadata>
    <RequestId>0f02a07d-b677-11e2-9eb0-dd50EXAMPLE</RequestId>
  </ResponseMetadata>
</DescribeAutoScalingGroupsResponse>
`
var DescribeAutoScalingInstances = `
<DescribeAutoScalingInstancesResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
  <DescribeAutoScalingInstancesResult>
    <AutoScalingInstances>
      <member>
        <HealthStatus>Healthy</HealthStatus>
        <AutoScalingGroupName>my-test-asg</AutoScalingGroupName>
        <AvailabilityZone>us-east-1a</AvailabilityZone>
        <InstanceId>i-78e0d40b</InstanceId>
        <LaunchConfigurationName>my-test-lc</LaunchConfigurationName>
        <LifecycleState>InService</LifecycleState>
      </member>
    </AutoScalingInstances>
  </DescribeAutoScalingInstancesResult>
  <ResponseMetadata>
    <RequestId>df992dc3-b72f-11e2-81e1-750aa6EXAMPLE</RequestId>
  </ResponseMetadata>
</DescribeAutoScalingInstancesResponse>
`
var DescribeLaunchConfigurations = `
<DescribeLaunchConfigurationsResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
  <DescribeLaunchConfigurationsResult>
    <LaunchConfigurations>
      <member>
        <AssociatePublicIpAddress>true</AssociatePublicIpAddress>
        <SecurityGroups/>
        <CreatedTime>2013-01-21T23:04:42.200Z</CreatedTime>
        <KernelId/>
        <LaunchConfigurationName>my-test-lc</LaunchConfigurationName>
        <UserData/>
        <InstanceType>m1.small</InstanceType>
        <LaunchConfigurationARN>arn:aws:autoscaling:us-east-1:803981987763:launchConfiguration:9dbbbf87-6141-428a-a409-0752edbe6cad:launchConfigurationName/my-test-lc</LaunchConfigurationARN>
        <BlockDeviceMappings>
          <member>
            <VirtualName>ephemeral0</VirtualName>
            <DeviceName>/dev/sdb</DeviceName>
          </member>
          <member>
            <Ebs> 
               <SnapshotId>snap-XXXXYYY</SnapshotId>
               <VolumeSize>100</VolumeSize>
            </Ebs>
            <DeviceName>/dev/sdf</DeviceName>
          </member>
        </BlockDeviceMappings>
        <ImageId>ami-514ac838</ImageId>
        <KeyName/>
        <RamdiskId/>
        <InstanceMonitoring>
          <Enabled>true</Enabled>
        </InstanceMonitoring>
        <EbsOptimized>false</EbsOptimized>
      </member>
    </LaunchConfigurations>
  </DescribeLaunchConfigurationsResult>
  <ResponseMetadata>
    <RequestId>d05a22f8-b690-11e2-bf8e-2113fEXAMPLE</RequestId>
  </ResponseMetadata>
</DescribeLaunchConfigurationsResponse> 
`
var DescribeMetricCollectionTypes = `
<DescribeMetricCollectionTypesResponse xmlns="http://autoscaling.amazonaws.co
oc/2011-01-01/">
  <DescribeMetricCollectionTypesResult>
    <Metrics>
      <member>
        <Metric>GroupMinSize</Metric>
      </member>
      <member>
        <Metric>GroupMaxSize</Metric>
      </member>
      <member>
        <Metric>GroupDesiredCapacity</Metric>
      </member>
      <member>
        <Metric>GroupInServiceInstances</Metric>
      </member>
      <member>
        <Metric>GroupPendingInstances</Metric>
      </member>
      <member>
        <Metric>GroupTerminatingInstances</Metric>
      </member>
      <member>
        <Metric>GroupTotalInstances</Metric>
      </member>
    </Metrics>
    <Granularities>
      <member>
        <Granularity>1Minute</Granularity>
      </member>
    </Granularities>
  </DescribeMetricCollectionTypesResult>
  <ResponseMetadata>
    <RequestId>07f3fea2-bf3c-11e2-9b6f-f3cdbb80c073</RequestId>
  </ResponseMetadata>
</DescribeMetricCollectionTypesResponse> 
`
var DescribeNotificationConfigurations = `
<DescribeNotificationConfigurationsResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
  <DescribeNotificationConfigurationsResult>
    <NotificationConfigurations>
      <member>
        <AutoScalingGroupName>my-test-asg</AutoScalingGroupName>
        <NotificationType>autoscaling: EC2_INSTANCE_LAUNCH</NotificationType>
        <TopicARN>vajdoafj231j41231/topic</TopicARN>
      </member>
    </NotificationConfigurations>
  </DescribeNotificationConfigurationsResult>
  <ResponseMetadata>
    <RequestId>07f3fea2-bf3c-11e2-9b6f-f3cdbb80c073</RequestId>
  </ResponseMetadata>
</DescribeNotificationConfigurationsResponse>
`
var DescribePolicies = `
<DescribePoliciesResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
  <DescribePoliciesResult>
   <NextToken>3ef417fe-9202-12-8ddd-d13e1313413</NextToken>
    <ScalingPolicies>
      <member>
        <PolicyARN>arn:aws:autoscaling:us-east-1:803981987763:scalingPolicy:c322761b-3172-4d56-9a21-0ed9d6161d67:autoScalingGroupName/my-test-asg:policyName/MyScaleDownPolicy</PolicyARN>
        <AdjustmentType>ChangeInCapacity</AdjustmentType>
        <ScalingAdjustment>-1</ScalingAdjustment>
        <PolicyName>MyScaleDownPolicy</PolicyName>
        <AutoScalingGroupName>my-test-asg</AutoScalingGroupName>
        <Cooldown>60</Cooldown>
        <Alarms>
          <member>
            <AlarmName>TestQueue</AlarmName>
            <AlarmARN>arn:aws:cloudwatch:us-east-1:803981987763:alarm:TestQueue</AlarmARN>
          </member>
        </Alarms>
      </member>
      <member>
        <PolicyARN>arn:aws:autoscaling:us-east-1:803981987763:scalingPolicy:c55a5cdd-9be0-435b-b60b-a8dd313159f5:autoScalingGroupName/my-test-asg:policyName/MyScaleUpPolicy</PolicyARN>
        <AdjustmentType>ChangeInCapacity</AdjustmentType>
        <ScalingAdjustment>1</ScalingAdjustment>
        <PolicyName>MyScaleUpPolicy</PolicyName>
        <AutoScalingGroupName>my-test-asg</AutoScalingGroupName>
        <Cooldown>60</Cooldown>
        <Alarms>
          <member>
            <AlarmName>TestQueue</AlarmName>
            <AlarmARN>arn:aws:cloudwatch:us-east-1:803981987763:alarm:TestQueue</AlarmARN>
          </member>
        </Alarms>
      </member>
    </ScalingPolicies>
  </DescribePoliciesResult>
  <ResponseMetadata>
    <RequestId>ec3bffad-b739-11e2-b38d-15fbEXAMPLE</RequestId>
  </ResponseMetadata>
</DescribePoliciesResponse> 
`
var DescribeScalingActivities = `
<DescribeScalingActivitiesResponse xmlns="http://ec2.amazonaws.com/doc/2011-01-01/">
<DescribeScalingActivitiesResult>
 <NextToken>3ef417fe-9202-12-8ddd-d13e1313413</NextToken>
<Activities>
   <member>
     <StatusCode>Failed</StatusCode>
     <Progress>0</Progress>
     <ActivityId>063308ae-aa22-4a9b-94f4-9faeEXAMPLE</ActivityId>
     <StartTime>2012-04-12T17:32:07.882Z</StartTime>
     <AutoScalingGroupName>my-test-asg</AutoScalingGroupName>
     <Cause>At 2012-04-12T17:31:30Z a user request created an AutoScalingGroup changing the desired capacity from 0 to 1.  At 2012-04-12T17:32:07Z an instance was started in response to a difference between desired and actual capacity, increasing the capacity from 0 to 1.</Cause>
     <Details>{}</Details>
     <Description>Launching a new EC2 instance.  Status Reason: The image id 'ami-4edb0327' does not exist. Launching EC2 instance failed.</Description>
     <EndTime>2012-04-12T17:32:08Z</EndTime>
     <StatusMessage>The image id 'ami-4edb0327' does not exist. Launching EC2 instance failed.</StatusMessage>
   </member>
</Activities>
  </DescribeScalingActivitiesResult>
  <ResponseMetadata>
   <RequestId>7a641adc-84c5-11e1-a8a5-217ebEXAMPLE</RequestId>
  </ResponseMetadata>
</DescribeScalingActivitiesResponse> 
`
var DescribeScalingProcessTypes = `
<DescribeScalingProcessTypesResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
  <DescribeScalingProcessTypesResult>
    <Processes>
      <member>
        <ProcessName>AZRebalance</ProcessName>
      </member>
      <member>
        <ProcessName>AddToLoadBalancer</ProcessName>
      </member>
      <member>
        <ProcessName>AlarmNotification</ProcessName>
      </member>
      <member>
        <ProcessName>HealthCheck</ProcessName>
      </member>
      <member>
        <ProcessName>Launch</ProcessName>
      </member>
      <member>
        <ProcessName>ReplaceUnhealthy</ProcessName>
      </member>
      <member>
        <ProcessName>ScheduledActions</ProcessName>
      </member>
      <member>
        <ProcessName>Terminate</ProcessName>
      </member>
    </Processes>
  </DescribeScalingProcessTypesResult>
  <ResponseMetadata>
    <RequestId>27f2eacc-b73f-11e2-ad99-c7aba3a9c963</RequestId>
  </ResponseMetadata>
</DescribeScalingProcessTypesResponse> 
`
var DescribeTerminationPolicyTypes = `
<DescribeTerminationPolicyTypesResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
  <DescribeTerminationPolicyTypesResult>
    <TerminationPolicyTypes>
      <member>ClosestToNextInstanceHour</member>
      <member>Default</member>
      <member>NewestInstance</member>
      <member>OldestInstance</member>
      <member>OldestLaunchConfiguration</member>
    </TerminationPolicyTypes>
  </DescribeTerminationPolicyTypesResult>
  <ResponseMetadata>
    <RequestId>d9a05827-b735-11e2-a40c-c79a5EXAMPLE</RequestId>
  </ResponseMetadata>
</DescribeTerminationPolicyTypesResponse> 
`
var DescribeTags = `
<DescribeTagsResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
  <DescribeTagsResult>
    <Tags>      
      <member>
        <ResourceId>my-test-asg</ResourceId>
        <PropagateAtLaunch>true</PropagateAtLaunch>
        <Value>1.0</Value>
        <Key>version</Key>
        <ResourceType>auto-scaling-group</ResourceType>
      </member>
    </Tags>
  </DescribeTagsResult>
  <ResponseMetadata>
    <RequestId>086265fd-bf3e-11e2-85fc-fbb1EXAMPLE</RequestId>
  </ResponseMetadata>
</DescribeTagsResponse> 
`
var DisableMetricsCollection = `
<DisableMetricsCollectionResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
	<ResponseMetadata>
		<RequestId>8d798a29-f083-11e1-bdfb-cb223EXAMPLE</RequestId>
	</ResponseMetadata>
</DisableMetricsCollectionResponse> 
`
var EnableMetricsCollection = `
<EnableMetricsCollectionResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
	<ResponseMetadata>
		<RequestId>8d798a29-f083-11e1-bdfb-cb223EXAMPLE</RequestId>
	</ResponseMetadata>
</EnableMetricsCollectionResponse> 
`
var ExecutePolicy = `
<EnableMetricsCollectionResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
	<ResponseMetadata>
		<RequestId>8d798a29-f083-11e1-bdfb-cb223EXAMPLE</RequestId>
	</ResponseMetadata>
</EnableMetricsCollectionResponse> 
`
var PutNotificationConfiguration = `
<EnableMetricsCollectionResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
	<ResponseMetadata>
		<RequestId>8d798a29-f083-11e1-bdfb-cb223EXAMPLE</RequestId>
	</ResponseMetadata>
</EnableMetricsCollectionResponse> 
`
var PutScalingPolicy = `
<PutScalingPolicyResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
  <PutScalingPolicyResult>
    <PolicyARN>arn:aws:autoscaling:us-east-1:803981987763:scalingPolicy:b0dcf5e8-02e6-4e31-9719-0675d0dc31ae:autoScalingGroupName/my-test-asg:policyName/my-scaleout-policy</PolicyARN>
  </PutScalingPolicyResult>
  <ResponseMetadata>
    <RequestId>3cfc6fef-c08b-11e2-a697-2922EXAMPLE</RequestId>
  </ResponseMetadata>
</PutScalingPolicyResponse> 
`
var PutScheduledUpdateGroupAction = `
<PutScheduledUpdateGroupActionResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
  <ResponseMetadata>
    <RequestId>3bc8c9bc-6a62-11e2-8a51-4b8a1EXAMPLE</RequestId>
  </ResponseMetadata>
  </PutScheduledUpdateGroupActionResponse>
  `
var ResumeProcesses = `
<ResumeProcessesResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
	<ResponseMetadata>
		<RequestId>8d798a29-f083-11e1-bdfb-cb223EXAMPLE</RequestId>
	</ResponseMetadata>
</ResumeProcessesResponse> 
`
var SetDesiredCapacity = `
<SetDesiredCapacityResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
  <ResponseMetadata>
    <RequestId>9fb7e2db-6998-11e2-a985-57c82EXAMPLE</RequestId>
  </ResponseMetadata>
</SetDesiredCapacityResponse>
`
var SetInstanceHealth = `
<SetInstanceHealthResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
  <ResponseMetadata>
    <RequestId>9fb7e2db-6998-11e2-a985-57c82EXAMPLE</RequestId>
  </ResponseMetadata>
</SetInstanceHealthResponse>
`
var SuspendProcesses = `
<SuspendProcessesResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
	<ResponseMetadata>
		<RequestId>8d798a29-f083-11e1-bdfb-cb223EXAMPLE</RequestId>
	</ResponseMetadata>
</SuspendProcessesResponse> 
`
var TerminateInstanceInAutoScalingGroup = `
<TerminateInstanceInAutoScalingGroupResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
  <TerminateInstanceInAutoScalingGroupResult>
    <Activity>
      <StatusCode>InProgress</StatusCode>
      <ActivityId>cczc44a87-7d04-dsa15-31-d27c219864c5</ActivityId>
      <Progress>0</Progress>
      <StartTime>2014-01-26T14:08:30.560Z</StartTime>
      <Cause>At 2014-01-26T14:08:30Z instance i-br234123 was taken out of service in response to a user request.</Cause>
      <Details>{&quot;Availability Zone&quot;:&quot;us-east-1b&quot;}</Details>
      <Description>Terminating EC2 instance: i-br234123</Description>
    </Activity>
  </TerminateInstanceInAutoScalingGroupResult>
  <ResponseMetadata>
    <RequestId>8d798a29-f083-11e1-bdfb-cb223EXAMPLE</RequestId>
  </ResponseMetadata>
</TerminateInstanceInAutoScalingGroupResponse>
`

var UpdateAutoScalingGroup = `
<UpdateAutoScalingGroupResponse xmlns="http://autoscaling.amazonaws.com/doc/2011-01-01/">
	<ResponseMetadata>
		<RequestId>8d798a29-f083-11e1-bdfb-cb223EXAMPLE</RequestId>
	</ResponseMetadata>
</UpdateAutoScalingGroupResponse> 
`
