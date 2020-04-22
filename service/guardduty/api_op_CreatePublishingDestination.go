// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreatePublishingDestinationInput struct {
	_ struct{} `type:"structure"`

	// The idempotency token for the request.
	ClientToken *string `locationName:"clientToken" type:"string" idempotencyToken:"true"`

	// The properties of the publishing destination, including the ARNs for the
	// destination and the KMS key used for encryption.
	//
	// DestinationProperties is a required field
	DestinationProperties *DestinationProperties `locationName:"destinationProperties" type:"structure" required:"true"`

	// The type of resource for the publishing destination. Currently only Amazon
	// S3 buckets are supported.
	//
	// DestinationType is a required field
	DestinationType DestinationType `locationName:"destinationType" min:"1" type:"string" required:"true" enum:"true"`

	// The ID of the GuardDuty detector associated with the publishing destination.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreatePublishingDestinationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePublishingDestinationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePublishingDestinationInput"}

	if s.DestinationProperties == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationProperties"))
	}
	if len(s.DestinationType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("DestinationType"))
	}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreatePublishingDestinationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	var ClientToken string
	if s.ClientToken != nil {
		ClientToken = *s.ClientToken
	} else {
		ClientToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DestinationProperties != nil {
		v := s.DestinationProperties

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "destinationProperties", v, metadata)
	}
	if len(s.DestinationType) > 0 {
		v := s.DestinationType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "destinationType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreatePublishingDestinationOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the publishing destination that is created.
	//
	// DestinationId is a required field
	DestinationId *string `locationName:"destinationId" type:"string" required:"true"`
}

// String returns the string representation
func (s CreatePublishingDestinationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreatePublishingDestinationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DestinationId != nil {
		v := *s.DestinationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "destinationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreatePublishingDestination = "CreatePublishingDestination"

// CreatePublishingDestinationRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Creates a publishing destination to export findings to. The resource to export
// findings to must exist before you use this operation.
//
//    // Example sending a request using CreatePublishingDestinationRequest.
//    req := client.CreatePublishingDestinationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/CreatePublishingDestination
func (c *Client) CreatePublishingDestinationRequest(input *CreatePublishingDestinationInput) CreatePublishingDestinationRequest {
	op := &aws.Operation{
		Name:       opCreatePublishingDestination,
		HTTPMethod: "POST",
		HTTPPath:   "/detector/{detectorId}/publishingDestination",
	}

	if input == nil {
		input = &CreatePublishingDestinationInput{}
	}

	req := c.newRequest(op, input, &CreatePublishingDestinationOutput{})
	return CreatePublishingDestinationRequest{Request: req, Input: input, Copy: c.CreatePublishingDestinationRequest}
}

// CreatePublishingDestinationRequest is the request type for the
// CreatePublishingDestination API operation.
type CreatePublishingDestinationRequest struct {
	*aws.Request
	Input *CreatePublishingDestinationInput
	Copy  func(*CreatePublishingDestinationInput) CreatePublishingDestinationRequest
}

// Send marshals and sends the CreatePublishingDestination API request.
func (r CreatePublishingDestinationRequest) Send(ctx context.Context) (*CreatePublishingDestinationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePublishingDestinationResponse{
		CreatePublishingDestinationOutput: r.Request.Data.(*CreatePublishingDestinationOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePublishingDestinationResponse is the response type for the
// CreatePublishingDestination API operation.
type CreatePublishingDestinationResponse struct {
	*CreatePublishingDestinationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePublishingDestination request.
func (r *CreatePublishingDestinationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
