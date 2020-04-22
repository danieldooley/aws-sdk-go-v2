// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteDetectorInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the detector that you want to delete.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDetectorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDetectorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDetectorInput"}

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
func (s DeleteDetectorInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteDetectorOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDetectorOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDetectorOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteDetector = "DeleteDetector"

// DeleteDetectorRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Deletes an Amazon GuardDuty detector that is specified by the detector ID.
//
//    // Example sending a request using DeleteDetectorRequest.
//    req := client.DeleteDetectorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/DeleteDetector
func (c *Client) DeleteDetectorRequest(input *DeleteDetectorInput) DeleteDetectorRequest {
	op := &aws.Operation{
		Name:       opDeleteDetector,
		HTTPMethod: "DELETE",
		HTTPPath:   "/detector/{detectorId}",
	}

	if input == nil {
		input = &DeleteDetectorInput{}
	}

	req := c.newRequest(op, input, &DeleteDetectorOutput{})
	return DeleteDetectorRequest{Request: req, Input: input, Copy: c.DeleteDetectorRequest}
}

// DeleteDetectorRequest is the request type for the
// DeleteDetector API operation.
type DeleteDetectorRequest struct {
	*aws.Request
	Input *DeleteDetectorInput
	Copy  func(*DeleteDetectorInput) DeleteDetectorRequest
}

// Send marshals and sends the DeleteDetector API request.
func (r DeleteDetectorRequest) Send(ctx context.Context) (*DeleteDetectorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDetectorResponse{
		DeleteDetectorOutput: r.Request.Data.(*DeleteDetectorOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDetectorResponse is the response type for the
// DeleteDetector API operation.
type DeleteDetectorResponse struct {
	*DeleteDetectorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDetector request.
func (r *DeleteDetectorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
