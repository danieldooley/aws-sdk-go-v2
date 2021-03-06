// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetConnectionsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog in which the connections reside. If none is provided,
	// the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// A filter that controls which connections are returned.
	Filter *GetConnectionsFilter `type:"structure"`

	// Allows you to retrieve the connection metadata without returning the password.
	// For instance, the AWS Glue console uses this flag to retrieve the connection,
	// and does not display the password. Set this parameter when the caller might
	// not have permission to use the AWS KMS key to decrypt the password, but it
	// does have permission to access the rest of the connection properties.
	HidePassword *bool `type:"boolean"`

	// The maximum number of connections to return in one response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A continuation token, if this is a continuation call.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetConnectionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetConnectionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetConnectionsInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetConnectionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of requested connection definitions.
	ConnectionList []Connection `type:"list"`

	// A continuation token, if the list of connections returned does not include
	// the last of the filtered connections.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetConnectionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetConnections = "GetConnections"

// GetConnectionsRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves a list of connection definitions from the Data Catalog.
//
//    // Example sending a request using GetConnectionsRequest.
//    req := client.GetConnectionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetConnections
func (c *Client) GetConnectionsRequest(input *GetConnectionsInput) GetConnectionsRequest {
	op := &aws.Operation{
		Name:       opGetConnections,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetConnectionsInput{}
	}

	req := c.newRequest(op, input, &GetConnectionsOutput{})

	return GetConnectionsRequest{Request: req, Input: input, Copy: c.GetConnectionsRequest}
}

// GetConnectionsRequest is the request type for the
// GetConnections API operation.
type GetConnectionsRequest struct {
	*aws.Request
	Input *GetConnectionsInput
	Copy  func(*GetConnectionsInput) GetConnectionsRequest
}

// Send marshals and sends the GetConnections API request.
func (r GetConnectionsRequest) Send(ctx context.Context) (*GetConnectionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetConnectionsResponse{
		GetConnectionsOutput: r.Request.Data.(*GetConnectionsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetConnectionsRequestPaginator returns a paginator for GetConnections.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetConnectionsRequest(input)
//   p := glue.NewGetConnectionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetConnectionsPaginator(req GetConnectionsRequest) GetConnectionsPaginator {
	return GetConnectionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetConnectionsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// GetConnectionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetConnectionsPaginator struct {
	aws.Pager
}

func (p *GetConnectionsPaginator) CurrentPage() *GetConnectionsOutput {
	return p.Pager.CurrentPage().(*GetConnectionsOutput)
}

// GetConnectionsResponse is the response type for the
// GetConnections API operation.
type GetConnectionsResponse struct {
	*GetConnectionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetConnections request.
func (r *GetConnectionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
