// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateRegexPatternSetInput struct {
	_ struct{} `type:"structure"`

	// A description of the set that helps with identification. You cannot change
	// the description of a set after you create it.
	Description *string `min:"1" type:"string"`

	// The name of the set. You cannot change the name after you create the set.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// Array of regular expression strings.
	//
	// RegularExpressionList is a required field
	RegularExpressionList []Regex `type:"list" required:"true"`

	// Specifies whether this is for an AWS CloudFront distribution or for a regional
	// application. A regional application can be an Application Load Balancer (ALB)
	// or an API Gateway stage.
	//
	// To work with CloudFront, you must also specify the Region US East (N. Virginia)
	// as follows:
	//
	//    * CLI - Specify the Region when you use the CloudFront scope: --scope=CLOUDFRONT
	//    --region=us-east-1.
	//
	//    * API and SDKs - For all calls, use the Region endpoint us-east-1.
	//
	// Scope is a required field
	Scope Scope `type:"string" required:"true" enum:"true"`

	// An array of key:value pairs to associate with the resource.
	Tags []Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s CreateRegexPatternSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRegexPatternSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRegexPatternSetInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.RegularExpressionList == nil {
		invalidParams.Add(aws.NewErrParamRequired("RegularExpressionList"))
	}
	if len(s.Scope) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Scope"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.RegularExpressionList != nil {
		for i, v := range s.RegularExpressionList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "RegularExpressionList", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateRegexPatternSetOutput struct {
	_ struct{} `type:"structure"`

	// High-level information about a RegexPatternSet, returned by operations like
	// create and list. This provides information like the ID, that you can use
	// to retrieve and manage a RegexPatternSet, and the ARN, that you provide to
	// the RegexPatternSetReferenceStatement to use the pattern set in a Rule.
	Summary *RegexPatternSetSummary `type:"structure"`
}

// String returns the string representation
func (s CreateRegexPatternSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateRegexPatternSet = "CreateRegexPatternSet"

// CreateRegexPatternSetRequest returns a request value for making API operation for
// AWS WAFV2.
//
//
// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from
// the prior release, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
//
// Creates a RegexPatternSet, which you reference in a RegexPatternSetReferenceStatement,
// to have AWS WAF inspect a web request component for the specified patterns.
//
//    // Example sending a request using CreateRegexPatternSetRequest.
//    req := client.CreateRegexPatternSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/wafv2-2019-07-29/CreateRegexPatternSet
func (c *Client) CreateRegexPatternSetRequest(input *CreateRegexPatternSetInput) CreateRegexPatternSetRequest {
	op := &aws.Operation{
		Name:       opCreateRegexPatternSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateRegexPatternSetInput{}
	}

	req := c.newRequest(op, input, &CreateRegexPatternSetOutput{})
	return CreateRegexPatternSetRequest{Request: req, Input: input, Copy: c.CreateRegexPatternSetRequest}
}

// CreateRegexPatternSetRequest is the request type for the
// CreateRegexPatternSet API operation.
type CreateRegexPatternSetRequest struct {
	*aws.Request
	Input *CreateRegexPatternSetInput
	Copy  func(*CreateRegexPatternSetInput) CreateRegexPatternSetRequest
}

// Send marshals and sends the CreateRegexPatternSet API request.
func (r CreateRegexPatternSetRequest) Send(ctx context.Context) (*CreateRegexPatternSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRegexPatternSetResponse{
		CreateRegexPatternSetOutput: r.Request.Data.(*CreateRegexPatternSetOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRegexPatternSetResponse is the response type for the
// CreateRegexPatternSet API operation.
type CreateRegexPatternSetResponse struct {
	*CreateRegexPatternSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRegexPatternSet request.
func (r *CreateRegexPatternSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
