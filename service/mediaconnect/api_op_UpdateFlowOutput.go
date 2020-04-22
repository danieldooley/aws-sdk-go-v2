// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The updates that you want to make to an existing output of an existing flow.
type UpdateFlowOutputInput struct {
	_ struct{} `type:"structure"`

	// The range of IP addresses that should be allowed to initiate output requests
	// to this flow. These IP addresses should be in the form of a Classless Inter-Domain
	// Routing (CIDR) block; for example, 10.0.0.0/16.
	CidrAllowList []string `locationName:"cidrAllowList" type:"list"`

	// A description of the output. This description appears only on the AWS Elemental
	// MediaConnect console and will not be seen by the end user.
	Description *string `locationName:"description" type:"string"`

	// The IP address where you want to send the output.
	Destination *string `locationName:"destination" type:"string"`

	// The type of key used for the encryption. If no keyType is provided, the service
	// will use the default setting (static-key).
	Encryption *UpdateEncryption `locationName:"encryption" type:"structure"`

	// FlowArn is a required field
	FlowArn *string `location:"uri" locationName:"flowArn" type:"string" required:"true"`

	// The maximum latency in milliseconds for Zixi-based streams.
	MaxLatency *int64 `locationName:"maxLatency" type:"integer"`

	// OutputArn is a required field
	OutputArn *string `location:"uri" locationName:"outputArn" type:"string" required:"true"`

	// The port to use when content is distributed to this output.
	Port *int64 `locationName:"port" type:"integer"`

	// The protocol to use for the output.
	Protocol Protocol `locationName:"protocol" type:"string" enum:"true"`

	// The remote ID for the Zixi-pull stream.
	RemoteId *string `locationName:"remoteId" type:"string"`

	// The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.
	SmoothingLatency *int64 `locationName:"smoothingLatency" type:"integer"`

	// The stream ID that you want to use for this transport. This parameter applies
	// only to Zixi-based streams.
	StreamId *string `locationName:"streamId" type:"string"`

	// The name of the VPC interface attachment to use for this output.
	VpcInterfaceAttachment *VpcInterfaceAttachment `locationName:"vpcInterfaceAttachment" type:"structure"`
}

// String returns the string representation
func (s UpdateFlowOutputInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateFlowOutputInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateFlowOutputInput"}

	if s.FlowArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FlowArn"))
	}

	if s.OutputArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("OutputArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFlowOutputInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CidrAllowList != nil {
		v := s.CidrAllowList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "cidrAllowList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Destination != nil {
		v := *s.Destination

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "destination", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Encryption != nil {
		v := s.Encryption

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "encryption", v, metadata)
	}
	if s.MaxLatency != nil {
		v := *s.MaxLatency

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxLatency", protocol.Int64Value(v), metadata)
	}
	if s.Port != nil {
		v := *s.Port

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "port", protocol.Int64Value(v), metadata)
	}
	if len(s.Protocol) > 0 {
		v := s.Protocol

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "protocol", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.RemoteId != nil {
		v := *s.RemoteId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "remoteId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SmoothingLatency != nil {
		v := *s.SmoothingLatency

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "smoothingLatency", protocol.Int64Value(v), metadata)
	}
	if s.StreamId != nil {
		v := *s.StreamId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "streamId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VpcInterfaceAttachment != nil {
		v := s.VpcInterfaceAttachment

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "vpcInterfaceAttachment", v, metadata)
	}
	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OutputArn != nil {
		v := *s.OutputArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "outputArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The result of a successful UpdateFlowOutput request including the flow ARN
// and the updated output.
type UpdateFlowOutputOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the flow that is associated with the updated output.
	FlowArn *string `locationName:"flowArn" type:"string"`

	// The settings for an output.
	Output *Output `locationName:"output" type:"structure"`
}

// String returns the string representation
func (s UpdateFlowOutputOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFlowOutputOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Output != nil {
		v := s.Output

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "output", v, metadata)
	}
	return nil
}

const opUpdateFlowOutput = "UpdateFlowOutput"

// UpdateFlowOutputRequest returns a request value for making API operation for
// AWS MediaConnect.
//
// Updates an existing flow output.
//
//    // Example sending a request using UpdateFlowOutputRequest.
//    req := client.UpdateFlowOutputRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/UpdateFlowOutput
func (c *Client) UpdateFlowOutputRequest(input *UpdateFlowOutputInput) UpdateFlowOutputRequest {
	op := &aws.Operation{
		Name:       opUpdateFlowOutput,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/flows/{flowArn}/outputs/{outputArn}",
	}

	if input == nil {
		input = &UpdateFlowOutputInput{}
	}

	req := c.newRequest(op, input, &UpdateFlowOutputOutput{})
	return UpdateFlowOutputRequest{Request: req, Input: input, Copy: c.UpdateFlowOutputRequest}
}

// UpdateFlowOutputRequest is the request type for the
// UpdateFlowOutput API operation.
type UpdateFlowOutputRequest struct {
	*aws.Request
	Input *UpdateFlowOutputInput
	Copy  func(*UpdateFlowOutputInput) UpdateFlowOutputRequest
}

// Send marshals and sends the UpdateFlowOutput API request.
func (r UpdateFlowOutputRequest) Send(ctx context.Context) (*UpdateFlowOutputResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateFlowOutputResponse{
		UpdateFlowOutputOutput: r.Request.Data.(*UpdateFlowOutputOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateFlowOutputResponse is the response type for the
// UpdateFlowOutput API operation.
type UpdateFlowOutputResponse struct {
	*UpdateFlowOutputOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateFlowOutput request.
func (r *UpdateFlowOutputResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
