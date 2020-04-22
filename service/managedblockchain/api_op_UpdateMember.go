// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package managedblockchain

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateMemberInput struct {
	_ struct{} `type:"structure"`

	// Configuration properties for publishing to Amazon CloudWatch Logs.
	LogPublishingConfiguration *MemberLogPublishingConfiguration `type:"structure"`

	// The unique ID of the member.
	//
	// MemberId is a required field
	MemberId *string `location:"uri" locationName:"memberId" min:"1" type:"string" required:"true"`

	// The unique ID of the Managed Blockchain network to which the member belongs.
	//
	// NetworkId is a required field
	NetworkId *string `location:"uri" locationName:"networkId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateMemberInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateMemberInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateMemberInput"}

	if s.MemberId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MemberId"))
	}
	if s.MemberId != nil && len(*s.MemberId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MemberId", 1))
	}

	if s.NetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkId"))
	}
	if s.NetworkId != nil && len(*s.NetworkId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NetworkId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateMemberInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.LogPublishingConfiguration != nil {
		v := s.LogPublishingConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "LogPublishingConfiguration", v, metadata)
	}
	if s.MemberId != nil {
		v := *s.MemberId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "memberId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NetworkId != nil {
		v := *s.NetworkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "networkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateMemberOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateMemberOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateMemberOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateMember = "UpdateMember"

// UpdateMemberRequest returns a request value for making API operation for
// Amazon Managed Blockchain.
//
// Updates a member configuration with new parameters.
//
//    // Example sending a request using UpdateMemberRequest.
//    req := client.UpdateMemberRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/UpdateMember
func (c *Client) UpdateMemberRequest(input *UpdateMemberInput) UpdateMemberRequest {
	op := &aws.Operation{
		Name:       opUpdateMember,
		HTTPMethod: "PATCH",
		HTTPPath:   "/networks/{networkId}/members/{memberId}",
	}

	if input == nil {
		input = &UpdateMemberInput{}
	}

	req := c.newRequest(op, input, &UpdateMemberOutput{})
	return UpdateMemberRequest{Request: req, Input: input, Copy: c.UpdateMemberRequest}
}

// UpdateMemberRequest is the request type for the
// UpdateMember API operation.
type UpdateMemberRequest struct {
	*aws.Request
	Input *UpdateMemberInput
	Copy  func(*UpdateMemberInput) UpdateMemberRequest
}

// Send marshals and sends the UpdateMember API request.
func (r UpdateMemberRequest) Send(ctx context.Context) (*UpdateMemberResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateMemberResponse{
		UpdateMemberOutput: r.Request.Data.(*UpdateMemberOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateMemberResponse is the response type for the
// UpdateMember API operation.
type UpdateMemberResponse struct {
	*UpdateMemberOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateMember request.
func (r *UpdateMemberResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
