// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetManagedScalingPolicyInput struct {
	_ struct{} `type:"structure"`

	// Specifies the ID of the cluster for which the managed scaling policy will
	// be fetched.
	//
	// ClusterId is a required field
	ClusterId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetManagedScalingPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetManagedScalingPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetManagedScalingPolicyInput"}

	if s.ClusterId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetManagedScalingPolicyOutput struct {
	_ struct{} `type:"structure"`

	// Specifies the managed scaling policy that is attached to an Amazon EMR cluster.
	ManagedScalingPolicy *ManagedScalingPolicy `type:"structure"`
}

// String returns the string representation
func (s GetManagedScalingPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetManagedScalingPolicy = "GetManagedScalingPolicy"

// GetManagedScalingPolicyRequest returns a request value for making API operation for
// Amazon Elastic MapReduce.
//
// Fetches the attached managed scaling policy for an Amazon EMR cluster.
//
//    // Example sending a request using GetManagedScalingPolicyRequest.
//    req := client.GetManagedScalingPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/GetManagedScalingPolicy
func (c *Client) GetManagedScalingPolicyRequest(input *GetManagedScalingPolicyInput) GetManagedScalingPolicyRequest {
	op := &aws.Operation{
		Name:       opGetManagedScalingPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetManagedScalingPolicyInput{}
	}

	req := c.newRequest(op, input, &GetManagedScalingPolicyOutput{})
	return GetManagedScalingPolicyRequest{Request: req, Input: input, Copy: c.GetManagedScalingPolicyRequest}
}

// GetManagedScalingPolicyRequest is the request type for the
// GetManagedScalingPolicy API operation.
type GetManagedScalingPolicyRequest struct {
	*aws.Request
	Input *GetManagedScalingPolicyInput
	Copy  func(*GetManagedScalingPolicyInput) GetManagedScalingPolicyRequest
}

// Send marshals and sends the GetManagedScalingPolicy API request.
func (r GetManagedScalingPolicyRequest) Send(ctx context.Context) (*GetManagedScalingPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetManagedScalingPolicyResponse{
		GetManagedScalingPolicyOutput: r.Request.Data.(*GetManagedScalingPolicyOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetManagedScalingPolicyResponse is the response type for the
// GetManagedScalingPolicy API operation.
type GetManagedScalingPolicyResponse struct {
	*GetManagedScalingPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetManagedScalingPolicy request.
func (r *GetManagedScalingPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}