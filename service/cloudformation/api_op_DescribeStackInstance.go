// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeStackInstanceInput struct {
	_ struct{} `type:"structure"`

	// The ID of an AWS account that's associated with this stack instance.
	//
	// StackInstanceAccount is a required field
	StackInstanceAccount *string `type:"string" required:"true"`

	// The name of a Region that's associated with this stack instance.
	//
	// StackInstanceRegion is a required field
	StackInstanceRegion *string `type:"string" required:"true"`

	// The name or the unique stack ID of the stack set that you want to get stack
	// instance information for.
	//
	// StackSetName is a required field
	StackSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeStackInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStackInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStackInstanceInput"}

	if s.StackInstanceAccount == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackInstanceAccount"))
	}

	if s.StackInstanceRegion == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackInstanceRegion"))
	}

	if s.StackSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeStackInstanceOutput struct {
	_ struct{} `type:"structure"`

	// The stack instance that matches the specified request parameters.
	StackInstance *StackInstance `type:"structure"`
}

// String returns the string representation
func (s DescribeStackInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeStackInstance = "DescribeStackInstance"

// DescribeStackInstanceRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns the stack instance that's associated with the specified stack set,
// AWS account, and Region.
//
// For a list of stack instances that are associated with a specific stack set,
// use ListStackInstances.
//
//    // Example sending a request using DescribeStackInstanceRequest.
//    req := client.DescribeStackInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeStackInstance
func (c *Client) DescribeStackInstanceRequest(input *DescribeStackInstanceInput) DescribeStackInstanceRequest {
	op := &aws.Operation{
		Name:       opDescribeStackInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeStackInstanceInput{}
	}

	req := c.newRequest(op, input, &DescribeStackInstanceOutput{})
	return DescribeStackInstanceRequest{Request: req, Input: input, Copy: c.DescribeStackInstanceRequest}
}

// DescribeStackInstanceRequest is the request type for the
// DescribeStackInstance API operation.
type DescribeStackInstanceRequest struct {
	*aws.Request
	Input *DescribeStackInstanceInput
	Copy  func(*DescribeStackInstanceInput) DescribeStackInstanceRequest
}

// Send marshals and sends the DescribeStackInstance API request.
func (r DescribeStackInstanceRequest) Send(ctx context.Context) (*DescribeStackInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStackInstanceResponse{
		DescribeStackInstanceOutput: r.Request.Data.(*DescribeStackInstanceOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeStackInstanceResponse is the response type for the
// DescribeStackInstance API operation.
type DescribeStackInstanceResponse struct {
	*DescribeStackInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStackInstance request.
func (r *DescribeStackInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
