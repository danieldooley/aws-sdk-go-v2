// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type PutVoiceConnectorProxyInput struct {
	_ struct{} `type:"structure"`

	// The default number of minutes allowed for proxy sessions.
	//
	// DefaultSessionExpiryMinutes is a required field
	DefaultSessionExpiryMinutes *int64 `type:"integer" required:"true"`

	// When true, stops proxy sessions from being created on the specified Amazon
	// Chime Voice Connector.
	Disabled *bool `type:"boolean"`

	// The phone number to route calls to after a proxy session expires.
	FallBackPhoneNumber *string `type:"string" sensitive:"true"`

	// The countries for proxy phone numbers to be selected from.
	//
	// PhoneNumberPoolCountries is a required field
	PhoneNumberPoolCountries []string `min:"1" type:"list" required:"true"`

	// The Amazon Chime voice connector ID.
	//
	// VoiceConnectorId is a required field
	VoiceConnectorId *string `location:"uri" locationName:"voiceConnectorId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutVoiceConnectorProxyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutVoiceConnectorProxyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutVoiceConnectorProxyInput"}

	if s.DefaultSessionExpiryMinutes == nil {
		invalidParams.Add(aws.NewErrParamRequired("DefaultSessionExpiryMinutes"))
	}

	if s.PhoneNumberPoolCountries == nil {
		invalidParams.Add(aws.NewErrParamRequired("PhoneNumberPoolCountries"))
	}
	if s.PhoneNumberPoolCountries != nil && len(s.PhoneNumberPoolCountries) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PhoneNumberPoolCountries", 1))
	}

	if s.VoiceConnectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorId"))
	}
	if s.VoiceConnectorId != nil && len(*s.VoiceConnectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VoiceConnectorId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutVoiceConnectorProxyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DefaultSessionExpiryMinutes != nil {
		v := *s.DefaultSessionExpiryMinutes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DefaultSessionExpiryMinutes", protocol.Int64Value(v), metadata)
	}
	if s.Disabled != nil {
		v := *s.Disabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Disabled", protocol.BoolValue(v), metadata)
	}
	if s.FallBackPhoneNumber != nil {
		v := *s.FallBackPhoneNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FallBackPhoneNumber", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PhoneNumberPoolCountries != nil {
		v := s.PhoneNumberPoolCountries

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "PhoneNumberPoolCountries", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.VoiceConnectorId != nil {
		v := *s.VoiceConnectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "voiceConnectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type PutVoiceConnectorProxyOutput struct {
	_ struct{} `type:"structure"`

	// The proxy configuration details.
	Proxy *Proxy `type:"structure"`
}

// String returns the string representation
func (s PutVoiceConnectorProxyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutVoiceConnectorProxyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Proxy != nil {
		v := s.Proxy

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Proxy", v, metadata)
	}
	return nil
}

const opPutVoiceConnectorProxy = "PutVoiceConnectorProxy"

// PutVoiceConnectorProxyRequest returns a request value for making API operation for
// Amazon Chime.
//
// Puts the specified proxy configuration to the specified Amazon Chime Voice
// Connector.
//
//    // Example sending a request using PutVoiceConnectorProxyRequest.
//    req := client.PutVoiceConnectorProxyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/PutVoiceConnectorProxy
func (c *Client) PutVoiceConnectorProxyRequest(input *PutVoiceConnectorProxyInput) PutVoiceConnectorProxyRequest {
	op := &aws.Operation{
		Name:       opPutVoiceConnectorProxy,
		HTTPMethod: "PUT",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}/programmable-numbers/proxy",
	}

	if input == nil {
		input = &PutVoiceConnectorProxyInput{}
	}

	req := c.newRequest(op, input, &PutVoiceConnectorProxyOutput{})
	return PutVoiceConnectorProxyRequest{Request: req, Input: input, Copy: c.PutVoiceConnectorProxyRequest}
}

// PutVoiceConnectorProxyRequest is the request type for the
// PutVoiceConnectorProxy API operation.
type PutVoiceConnectorProxyRequest struct {
	*aws.Request
	Input *PutVoiceConnectorProxyInput
	Copy  func(*PutVoiceConnectorProxyInput) PutVoiceConnectorProxyRequest
}

// Send marshals and sends the PutVoiceConnectorProxy API request.
func (r PutVoiceConnectorProxyRequest) Send(ctx context.Context) (*PutVoiceConnectorProxyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutVoiceConnectorProxyResponse{
		PutVoiceConnectorProxyOutput: r.Request.Data.(*PutVoiceConnectorProxyOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutVoiceConnectorProxyResponse is the response type for the
// PutVoiceConnectorProxy API operation.
type PutVoiceConnectorProxyResponse struct {
	*PutVoiceConnectorProxyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutVoiceConnectorProxy request.
func (r *PutVoiceConnectorProxyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
