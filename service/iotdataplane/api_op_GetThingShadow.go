// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotdataplane

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the GetThingShadow operation.
type GetThingShadowInput struct {
	_ struct{} `type:"structure"`

	// The name of the thing.
	//
	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetThingShadowInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetThingShadowInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetThingShadowInput"}

	if s.ThingName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetThingShadowInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The output from the GetThingShadow operation.
type GetThingShadowOutput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// The state information, in JSON format.
	Payload []byte `locationName:"payload" type:"blob"`
}

// String returns the string representation
func (s GetThingShadowOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetThingShadowOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Payload != nil {
		v := s.Payload

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "payload", protocol.BytesStream(v), metadata)
	}
	return nil
}

const opGetThingShadow = "GetThingShadow"

// GetThingShadowRequest returns a request value for making API operation for
// AWS IoT Data Plane.
//
// Gets the thing shadow for the specified thing.
//
// For more information, see GetThingShadow (http://docs.aws.amazon.com/iot/latest/developerguide/API_GetThingShadow.html)
// in the AWS IoT Developer Guide.
//
//    // Example sending a request using GetThingShadowRequest.
//    req := client.GetThingShadowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetThingShadowRequest(input *GetThingShadowInput) GetThingShadowRequest {
	op := &aws.Operation{
		Name:       opGetThingShadow,
		HTTPMethod: "GET",
		HTTPPath:   "/things/{thingName}/shadow",
	}

	if input == nil {
		input = &GetThingShadowInput{}
	}

	req := c.newRequest(op, input, &GetThingShadowOutput{})

	return GetThingShadowRequest{Request: req, Input: input, Copy: c.GetThingShadowRequest}
}

// GetThingShadowRequest is the request type for the
// GetThingShadow API operation.
type GetThingShadowRequest struct {
	*aws.Request
	Input *GetThingShadowInput
	Copy  func(*GetThingShadowInput) GetThingShadowRequest
}

// Send marshals and sends the GetThingShadow API request.
func (r GetThingShadowRequest) Send(ctx context.Context) (*GetThingShadowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetThingShadowResponse{
		GetThingShadowOutput: r.Request.Data.(*GetThingShadowOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetThingShadowResponse is the response type for the
// GetThingShadow API operation.
type GetThingShadowResponse struct {
	*GetThingShadowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetThingShadow request.
func (r *GetThingShadowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
