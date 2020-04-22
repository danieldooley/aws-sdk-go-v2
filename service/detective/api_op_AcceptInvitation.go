// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package detective

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type AcceptInvitationInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the behavior graph that the member account is accepting the invitation
	// for.
	//
	// The member account status in the behavior graph must be INVITED.
	//
	// GraphArn is a required field
	GraphArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AcceptInvitationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AcceptInvitationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AcceptInvitationInput"}

	if s.GraphArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("GraphArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AcceptInvitationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GraphArn != nil {
		v := *s.GraphArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GraphArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type AcceptInvitationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AcceptInvitationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AcceptInvitationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opAcceptInvitation = "AcceptInvitation"

// AcceptInvitationRequest returns a request value for making API operation for
// Amazon Detective.
//
// Accepts an invitation for the member account to contribute data to a behavior
// graph. This operation can only be called by an invited member account.
//
// The request provides the ARN of behavior graph.
//
// The member account status in the graph must be INVITED.
//
//    // Example sending a request using AcceptInvitationRequest.
//    req := client.AcceptInvitationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/detective-2018-10-26/AcceptInvitation
func (c *Client) AcceptInvitationRequest(input *AcceptInvitationInput) AcceptInvitationRequest {
	op := &aws.Operation{
		Name:       opAcceptInvitation,
		HTTPMethod: "PUT",
		HTTPPath:   "/invitation",
	}

	if input == nil {
		input = &AcceptInvitationInput{}
	}

	req := c.newRequest(op, input, &AcceptInvitationOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return AcceptInvitationRequest{Request: req, Input: input, Copy: c.AcceptInvitationRequest}
}

// AcceptInvitationRequest is the request type for the
// AcceptInvitation API operation.
type AcceptInvitationRequest struct {
	*aws.Request
	Input *AcceptInvitationInput
	Copy  func(*AcceptInvitationInput) AcceptInvitationRequest
}

// Send marshals and sends the AcceptInvitation API request.
func (r AcceptInvitationRequest) Send(ctx context.Context) (*AcceptInvitationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcceptInvitationResponse{
		AcceptInvitationOutput: r.Request.Data.(*AcceptInvitationOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcceptInvitationResponse is the response type for the
// AcceptInvitation API operation.
type AcceptInvitationResponse struct {
	*AcceptInvitationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcceptInvitation request.
func (r *AcceptInvitationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
