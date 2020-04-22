// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateOrganizationConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether to automatically enable member accounts in the organization.
	//
	// AutoEnable is a required field
	AutoEnable *bool `locationName:"autoEnable" type:"boolean" required:"true"`

	// The ID of the detector to update the delegated administrator for.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateOrganizationConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateOrganizationConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateOrganizationConfigurationInput"}

	if s.AutoEnable == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutoEnable"))
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
func (s UpdateOrganizationConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AutoEnable != nil {
		v := *s.AutoEnable

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "autoEnable", protocol.BoolValue(v), metadata)
	}
	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateOrganizationConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateOrganizationConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateOrganizationConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateOrganizationConfiguration = "UpdateOrganizationConfiguration"

// UpdateOrganizationConfigurationRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Updates the delegated administrator account with the values provided.
//
//    // Example sending a request using UpdateOrganizationConfigurationRequest.
//    req := client.UpdateOrganizationConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/UpdateOrganizationConfiguration
func (c *Client) UpdateOrganizationConfigurationRequest(input *UpdateOrganizationConfigurationInput) UpdateOrganizationConfigurationRequest {
	op := &aws.Operation{
		Name:       opUpdateOrganizationConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/detector/{detectorId}/admin",
	}

	if input == nil {
		input = &UpdateOrganizationConfigurationInput{}
	}

	req := c.newRequest(op, input, &UpdateOrganizationConfigurationOutput{})
	return UpdateOrganizationConfigurationRequest{Request: req, Input: input, Copy: c.UpdateOrganizationConfigurationRequest}
}

// UpdateOrganizationConfigurationRequest is the request type for the
// UpdateOrganizationConfiguration API operation.
type UpdateOrganizationConfigurationRequest struct {
	*aws.Request
	Input *UpdateOrganizationConfigurationInput
	Copy  func(*UpdateOrganizationConfigurationInput) UpdateOrganizationConfigurationRequest
}

// Send marshals and sends the UpdateOrganizationConfiguration API request.
func (r UpdateOrganizationConfigurationRequest) Send(ctx context.Context) (*UpdateOrganizationConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateOrganizationConfigurationResponse{
		UpdateOrganizationConfigurationOutput: r.Request.Data.(*UpdateOrganizationConfigurationOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateOrganizationConfigurationResponse is the response type for the
// UpdateOrganizationConfiguration API operation.
type UpdateOrganizationConfigurationResponse struct {
	*UpdateOrganizationConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateOrganizationConfiguration request.
func (r *UpdateOrganizationConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
