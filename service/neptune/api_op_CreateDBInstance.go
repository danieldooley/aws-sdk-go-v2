// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/CreateDBInstanceMessage
type CreateDBInstanceInput struct {
	_ struct{} `type:"structure"`

	// The amount of storage (in gibibytes) to allocate for the DB instance.
	//
	// Type: Integer
	//
	// Not applicable. Neptune cluster volumes automatically grow as the amount
	// of data in your database increases, though you are only charged for the space
	// that you use in a Neptune cluster volume.
	AllocatedStorage *int64 `type:"integer"`

	// Indicates that minor engine upgrades are applied automatically to the DB
	// instance during the maintenance window.
	//
	// Default: true
	AutoMinorVersionUpgrade *bool `type:"boolean"`

	// The EC2 Availability Zone that the DB instance is created in.
	//
	// Default: A random, system-chosen Availability Zone in the endpoint's AWS
	// Region.
	//
	// Example: us-east-1d
	//
	// Constraint: The AvailabilityZone parameter can't be specified if the MultiAZ
	// parameter is set to true. The specified Availability Zone must be in the
	// same AWS Region as the current endpoint.
	AvailabilityZone *string `type:"string"`

	// The number of days for which automated backups are retained.
	//
	// Not applicable. The retention period for automated backups is managed by
	// the DB cluster. For more information, see CreateDBCluster.
	//
	// Default: 1
	//
	// Constraints:
	//
	//    * Must be a value from 0 to 35
	//
	//    * Cannot be set to 0 if the DB instance is a source to Read Replicas
	BackupRetentionPeriod *int64 `type:"integer"`

	// Indicates that the DB instance should be associated with the specified CharacterSet.
	//
	// Not applicable. The character set is managed by the DB cluster. For more
	// information, see CreateDBCluster.
	CharacterSetName *string `type:"string"`

	// True to copy all tags from the DB instance to snapshots of the DB instance,
	// and otherwise false. The default is false.
	CopyTagsToSnapshot *bool `type:"boolean"`

	// The identifier of the DB cluster that the instance will belong to.
	//
	// For information on creating a DB cluster, see CreateDBCluster.
	//
	// Type: String
	DBClusterIdentifier *string `type:"string"`

	// The compute and memory capacity of the DB instance, for example, db.m4.large.
	// Not all DB instance classes are available in all AWS Regions.
	//
	// DBInstanceClass is a required field
	DBInstanceClass *string `type:"string" required:"true"`

	// The DB instance identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 letters, numbers, or hyphens.
	//
	//    * First character must be a letter.
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example: mydbinstance
	//
	// DBInstanceIdentifier is a required field
	DBInstanceIdentifier *string `type:"string" required:"true"`

	// The database name.
	//
	// Type: String
	DBName *string `type:"string"`

	// The name of the DB parameter group to associate with this DB instance. If
	// this argument is omitted, the default DBParameterGroup for the specified
	// engine is used.
	//
	// Constraints:
	//
	//    * Must be 1 to 255 letters, numbers, or hyphens.
	//
	//    * First character must be a letter
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens
	DBParameterGroupName *string `type:"string"`

	// A list of DB security groups to associate with this DB instance.
	//
	// Default: The default DB security group for the database engine.
	DBSecurityGroups []string `locationNameList:"DBSecurityGroupName" type:"list"`

	// A DB subnet group to associate with this DB instance.
	//
	// If there is no DB subnet group, then it is a non-VPC DB instance.
	DBSubnetGroupName *string `type:"string"`

	// Specify the Active Directory Domain to create the instance in.
	Domain *string `type:"string"`

	// Specify the name of the IAM role to be used when making API calls to the
	// Directory Service.
	DomainIAMRoleName *string `type:"string"`

	// The list of log types that need to be enabled for exporting to CloudWatch
	// Logs.
	EnableCloudwatchLogsExports []string `type:"list"`

	// True to enable AWS Identity and Access Management (IAM) authentication for
	// Neptune.
	//
	// Default: false
	EnableIAMDatabaseAuthentication *bool `type:"boolean"`

	// True to enable Performance Insights for the DB instance, and otherwise false.
	EnablePerformanceInsights *bool `type:"boolean"`

	// The name of the database engine to be used for this instance.
	//
	// Valid Values: neptune
	//
	// Engine is a required field
	Engine *string `type:"string" required:"true"`

	// The version number of the database engine to use.
	EngineVersion *string `type:"string"`

	// The amount of Provisioned IOPS (input/output operations per second) to be
	// initially allocated for the DB instance.
	Iops *int64 `type:"integer"`

	// The AWS KMS key identifier for an encrypted DB instance.
	//
	// The KMS key identifier is the Amazon Resource Name (ARN) for the KMS encryption
	// key. If you are creating a DB instance with the same AWS account that owns
	// the KMS encryption key used to encrypt the new DB instance, then you can
	// use the KMS key alias instead of the ARN for the KM encryption key.
	//
	// Not applicable. The KMS key identifier is managed by the DB cluster. For
	// more information, see CreateDBCluster.
	//
	// If the StorageEncrypted parameter is true, and you do not specify a value
	// for the KmsKeyId parameter, then Amazon Neptune will use your default encryption
	// key. AWS KMS creates the default encryption key for your AWS account. Your
	// AWS account has a different default encryption key for each AWS Region.
	KmsKeyId *string `type:"string"`

	// License model information for this DB instance.
	//
	// Valid values: license-included | bring-your-own-license | general-public-license
	LicenseModel *string `type:"string"`

	// The password for the master user. The password can include any printable
	// ASCII character except "/", """, or "@".
	//
	// Not used.
	MasterUserPassword *string `type:"string"`

	// The name for the master user. Not used.
	MasterUsername *string `type:"string"`

	// The interval, in seconds, between points when Enhanced Monitoring metrics
	// are collected for the DB instance. To disable collecting Enhanced Monitoring
	// metrics, specify 0. The default is 0.
	//
	// If MonitoringRoleArn is specified, then you must also set MonitoringInterval
	// to a value other than 0.
	//
	// Valid Values: 0, 1, 5, 10, 15, 30, 60
	MonitoringInterval *int64 `type:"integer"`

	// The ARN for the IAM role that permits Neptune to send enhanced monitoring
	// metrics to Amazon CloudWatch Logs. For example, arn:aws:iam:123456789012:role/emaccess.
	//
	// If MonitoringInterval is set to a value other than 0, then you must supply
	// a MonitoringRoleArn value.
	MonitoringRoleArn *string `type:"string"`

	// Specifies if the DB instance is a Multi-AZ deployment. You can't set the
	// AvailabilityZone parameter if the MultiAZ parameter is set to true.
	MultiAZ *bool `type:"boolean"`

	// Indicates that the DB instance should be associated with the specified option
	// group.
	//
	// Permanent options, such as the TDE option for Oracle Advanced Security TDE,
	// can't be removed from an option group, and that option group can't be removed
	// from a DB instance once it is associated with a DB instance
	OptionGroupName *string `type:"string"`

	// The AWS KMS key identifier for encryption of Performance Insights data. The
	// KMS key ID is the Amazon Resource Name (ARN), KMS key identifier, or the
	// KMS key alias for the KMS encryption key.
	PerformanceInsightsKMSKeyId *string `type:"string"`

	// The port number on which the database accepts connections.
	//
	// Not applicable. The port is managed by the DB cluster. For more information,
	// see CreateDBCluster.
	//
	// Default: 8182
	//
	// Type: Integer
	Port *int64 `type:"integer"`

	// The daily time range during which automated backups are created.
	//
	// Not applicable. The daily time range for creating automated backups is managed
	// by the DB cluster. For more information, see CreateDBCluster.
	PreferredBackupWindow *string `type:"string"`

	// The time range each week during which system maintenance can occur, in Universal
	// Coordinated Time (UTC).
	//
	// Format: ddd:hh24:mi-ddd:hh24:mi
	//
	// The default is a 30-minute window selected at random from an 8-hour block
	// of time for each AWS Region, occurring on a random day of the week.
	//
	// Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `type:"string"`

	// A value that specifies the order in which an Read Replica is promoted to
	// the primary instance after a failure of the existing primary instance.
	//
	// Default: 1
	//
	// Valid Values: 0 - 15
	PromotionTier *int64 `type:"integer"`

	// This parameter is not supported.
	PubliclyAccessible *bool `deprecated:"true" type:"boolean"`

	// Specifies whether the DB instance is encrypted.
	//
	// Not applicable. The encryption for DB instances is managed by the DB cluster.
	// For more information, see CreateDBCluster.
	//
	// Default: false
	StorageEncrypted *bool `type:"boolean"`

	// Specifies the storage type to be associated with the DB instance.
	//
	// Not applicable. Storage is managed by the DB Cluster.
	StorageType *string `type:"string"`

	// A list of tags. For more information, see Tagging Amazon Neptune Resources
	// (http://docs.aws.amazon.com/neptune/latest/UserGuide/tagging.ARN.html).
	Tags []Tag `locationNameList:"Tag" type:"list"`

	// The ARN from the key store with which to associate the instance for TDE encryption.
	TdeCredentialArn *string `type:"string"`

	// The password for the given ARN from the key store in order to access the
	// device.
	TdeCredentialPassword *string `type:"string"`

	// The time zone of the DB instance.
	Timezone *string `type:"string"`

	// A list of EC2 VPC security groups to associate with this DB instance.
	//
	// Not applicable. The associated list of EC2 VPC security groups is managed
	// by the DB cluster. For more information, see CreateDBCluster.
	//
	// Default: The default EC2 VPC security group for the DB subnet group's VPC.
	VpcSecurityGroupIds []string `locationNameList:"VpcSecurityGroupId" type:"list"`
}

// String returns the string representation
func (s CreateDBInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDBInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDBInstanceInput"}

	if s.DBInstanceClass == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBInstanceClass"))
	}

	if s.DBInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBInstanceIdentifier"))
	}

	if s.Engine == nil {
		invalidParams.Add(aws.NewErrParamRequired("Engine"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/CreateDBInstanceResult
type CreateDBInstanceOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon Neptune DB instance.
	//
	// This data type is used as a response element in the DescribeDBInstances action.
	DBInstance *DBInstance `type:"structure"`
}

// String returns the string representation
func (s CreateDBInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDBInstance = "CreateDBInstance"

// CreateDBInstanceRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Creates a new DB instance.
//
//    // Example sending a request using CreateDBInstanceRequest.
//    req := client.CreateDBInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/CreateDBInstance
func (c *Client) CreateDBInstanceRequest(input *CreateDBInstanceInput) CreateDBInstanceRequest {
	op := &aws.Operation{
		Name:       opCreateDBInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDBInstanceInput{}
	}

	req := c.newRequest(op, input, &CreateDBInstanceOutput{})
	return CreateDBInstanceRequest{Request: req, Input: input, Copy: c.CreateDBInstanceRequest}
}

// CreateDBInstanceRequest is the request type for the
// CreateDBInstance API operation.
type CreateDBInstanceRequest struct {
	*aws.Request
	Input *CreateDBInstanceInput
	Copy  func(*CreateDBInstanceInput) CreateDBInstanceRequest
}

// Send marshals and sends the CreateDBInstance API request.
func (r CreateDBInstanceRequest) Send(ctx context.Context) (*CreateDBInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDBInstanceResponse{
		CreateDBInstanceOutput: r.Request.Data.(*CreateDBInstanceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDBInstanceResponse is the response type for the
// CreateDBInstance API operation.
type CreateDBInstanceResponse struct {
	*CreateDBInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDBInstance request.
func (r *CreateDBInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}