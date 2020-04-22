// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package detective

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateGraphInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateGraphInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateGraphInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	return nil
}

type CreateGraphOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the new behavior graph.
	GraphArn *string `type:"string"`
}

// String returns the string representation
func (s CreateGraphOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateGraphOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.GraphArn != nil {
		v := *s.GraphArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GraphArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateGraph = "CreateGraph"

// CreateGraphRequest returns a request value for making API operation for
// Amazon Detective.
//
// Creates a new behavior graph for the calling account, and sets that account
// as the master account. This operation is called by the account that is enabling
// Detective.
//
// Before you try to enable Detective, make sure that your account has been
// enrolled in Amazon GuardDuty for at least 48 hours. If you do not meet this
// requirement, you cannot enable Detective. If you do meet the GuardDuty prerequisite,
// then when you make the request to enable Detective, it checks whether your
// data volume is within the Detective quota. If it exceeds the quota, then
// you cannot enable Detective.
//
// The operation also enables Detective for the calling account in the currently
// selected Region. It returns the ARN of the new behavior graph.
//
// CreateGraph triggers a process to create the corresponding data tables for
// the new behavior graph.
//
// An account can only be the master account for one behavior graph within a
// Region. If the same account calls CreateGraph with the same master account,
// it always returns the same behavior graph ARN. It does not create a new behavior
// graph.
//
//    // Example sending a request using CreateGraphRequest.
//    req := client.CreateGraphRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/detective-2018-10-26/CreateGraph
func (c *Client) CreateGraphRequest(input *CreateGraphInput) CreateGraphRequest {
	op := &aws.Operation{
		Name:       opCreateGraph,
		HTTPMethod: "POST",
		HTTPPath:   "/graph",
	}

	if input == nil {
		input = &CreateGraphInput{}
	}

	req := c.newRequest(op, input, &CreateGraphOutput{})
	return CreateGraphRequest{Request: req, Input: input, Copy: c.CreateGraphRequest}
}

// CreateGraphRequest is the request type for the
// CreateGraph API operation.
type CreateGraphRequest struct {
	*aws.Request
	Input *CreateGraphInput
	Copy  func(*CreateGraphInput) CreateGraphRequest
}

// Send marshals and sends the CreateGraph API request.
func (r CreateGraphRequest) Send(ctx context.Context) (*CreateGraphResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateGraphResponse{
		CreateGraphOutput: r.Request.Data.(*CreateGraphOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateGraphResponse is the response type for the
// CreateGraph API operation.
type CreateGraphResponse struct {
	*CreateGraphOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateGraph request.
func (r *CreateGraphResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
