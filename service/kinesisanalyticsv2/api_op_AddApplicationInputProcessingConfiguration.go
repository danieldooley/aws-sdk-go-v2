// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AddApplicationInputProcessingConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The name of the application to which you want to add the input processing
	// configuration.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// The version of the application to which you want to add the input processing
	// configuration. You can use the DescribeApplication operation to get the current
	// application version. If the version specified is not the current version,
	// the ConcurrentModificationException is returned.
	//
	// CurrentApplicationVersionId is a required field
	CurrentApplicationVersionId *int64 `min:"1" type:"long" required:"true"`

	// The ID of the input configuration to add the input processing configuration
	// to. You can get a list of the input IDs for an application using the DescribeApplication
	// operation.
	//
	// InputId is a required field
	InputId *string `min:"1" type:"string" required:"true"`

	// The InputProcessingConfiguration to add to the application.
	//
	// InputProcessingConfiguration is a required field
	InputProcessingConfiguration *InputProcessingConfiguration `type:"structure" required:"true"`
}

// String returns the string representation
func (s AddApplicationInputProcessingConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddApplicationInputProcessingConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddApplicationInputProcessingConfigurationInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if s.CurrentApplicationVersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CurrentApplicationVersionId"))
	}
	if s.CurrentApplicationVersionId != nil && *s.CurrentApplicationVersionId < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("CurrentApplicationVersionId", 1))
	}

	if s.InputId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputId"))
	}
	if s.InputId != nil && len(*s.InputId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InputId", 1))
	}

	if s.InputProcessingConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputProcessingConfiguration"))
	}
	if s.InputProcessingConfiguration != nil {
		if err := s.InputProcessingConfiguration.Validate(); err != nil {
			invalidParams.AddNested("InputProcessingConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AddApplicationInputProcessingConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the application.
	ApplicationARN *string `min:"1" type:"string"`

	// Provides the current application version.
	ApplicationVersionId *int64 `min:"1" type:"long"`

	// The input ID that is associated with the application input. This is the ID
	// that Amazon Kinesis Data Analytics assigns to each input configuration that
	// you add to your application.
	InputId *string `min:"1" type:"string"`

	// The description of the preprocessor that executes on records in this input
	// before the application's code is run.
	InputProcessingConfigurationDescription *InputProcessingConfigurationDescription `type:"structure"`
}

// String returns the string representation
func (s AddApplicationInputProcessingConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddApplicationInputProcessingConfiguration = "AddApplicationInputProcessingConfiguration"

// AddApplicationInputProcessingConfigurationRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
// Adds an InputProcessingConfiguration to an SQL-based Kinesis Data Analytics
// application. An input processor pre-processes records on the input stream
// before the application's SQL code executes. Currently, the only input processor
// available is AWS Lambda (https://aws.amazon.com/documentation/lambda/).
//
//    // Example sending a request using AddApplicationInputProcessingConfigurationRequest.
//    req := client.AddApplicationInputProcessingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/AddApplicationInputProcessingConfiguration
func (c *Client) AddApplicationInputProcessingConfigurationRequest(input *AddApplicationInputProcessingConfigurationInput) AddApplicationInputProcessingConfigurationRequest {
	op := &aws.Operation{
		Name:       opAddApplicationInputProcessingConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddApplicationInputProcessingConfigurationInput{}
	}

	req := c.newRequest(op, input, &AddApplicationInputProcessingConfigurationOutput{})

	return AddApplicationInputProcessingConfigurationRequest{Request: req, Input: input, Copy: c.AddApplicationInputProcessingConfigurationRequest}
}

// AddApplicationInputProcessingConfigurationRequest is the request type for the
// AddApplicationInputProcessingConfiguration API operation.
type AddApplicationInputProcessingConfigurationRequest struct {
	*aws.Request
	Input *AddApplicationInputProcessingConfigurationInput
	Copy  func(*AddApplicationInputProcessingConfigurationInput) AddApplicationInputProcessingConfigurationRequest
}

// Send marshals and sends the AddApplicationInputProcessingConfiguration API request.
func (r AddApplicationInputProcessingConfigurationRequest) Send(ctx context.Context) (*AddApplicationInputProcessingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddApplicationInputProcessingConfigurationResponse{
		AddApplicationInputProcessingConfigurationOutput: r.Request.Data.(*AddApplicationInputProcessingConfigurationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddApplicationInputProcessingConfigurationResponse is the response type for the
// AddApplicationInputProcessingConfiguration API operation.
type AddApplicationInputProcessingConfigurationResponse struct {
	*AddApplicationInputProcessingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddApplicationInputProcessingConfiguration request.
func (r *AddApplicationInputProcessingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
