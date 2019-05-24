// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to list users.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ListUsersRequest
type ListUsersInput struct {
	_ struct{} `type:"structure"`

	// An array of strings, where each string is the name of a user attribute to
	// be returned for each user in the search results. If the array is null, all
	// attributes are returned.
	AttributesToGet []string `type:"list"`

	// A filter string of the form "AttributeNameFilter-Type "AttributeValue"".
	// Quotation marks within the filter string must be escaped using the backslash
	// (\) character. For example, "family_name = \"Reddy\"".
	//
	//    * AttributeName: The name of the attribute to search for. You can only
	//    search for one attribute at a time.
	//
	//    * Filter-Type: For an exact match, use =, for example, "given_name = \"Jon\"".
	//    For a prefix ("starts with") match, use ^=, for example, "given_name ^=
	//    \"Jon\"".
	//
	//    * AttributeValue: The attribute value that must be matched for each user.
	//
	// If the filter string is empty, ListUsers returns all users in the user pool.
	//
	// You can only search for the following standard attributes:
	//
	//    * username (case-sensitive)
	//
	//    * email
	//
	//    * phone_number
	//
	//    * name
	//
	//    * given_name
	//
	//    * family_name
	//
	//    * preferred_username
	//
	//    * cognito:user_status (called Status in the Console) (case-insensitive)
	//
	//    * status (called Enabled in the Console) (case-sensitive)
	//
	//    * sub
	//
	// Custom attributes are not searchable.
	//
	// For more information, see Searching for Users Using the ListUsers API (http://docs.aws.amazon.com/cognito/latest/developerguide/how-to-manage-user-accounts.html#cognito-user-pools-searching-for-users-using-listusers-api)
	// and Examples of Using the ListUsers API (http://docs.aws.amazon.com/cognito/latest/developerguide/how-to-manage-user-accounts.html#cognito-user-pools-searching-for-users-listusers-api-examples)
	// in the Amazon Cognito Developer Guide.
	Filter *string `type:"string"`

	// Maximum number of users to be returned.
	Limit *int64 `type:"integer"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	PaginationToken *string `min:"1" type:"string"`

	// The user pool ID for the user pool on which the search should be performed.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListUsersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListUsersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListUsersInput"}
	if s.PaginationToken != nil && len(*s.PaginationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PaginationToken", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The response from the request to list users.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ListUsersResponse
type ListUsersOutput struct {
	_ struct{} `type:"structure"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	PaginationToken *string `min:"1" type:"string"`

	// The users returned in the request to list users.
	Users []UserType `type:"list"`
}

// String returns the string representation
func (s ListUsersOutput) String() string {
	return awsutil.Prettify(s)
}

const opListUsers = "ListUsers"

// ListUsersRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Lists the users in the Amazon Cognito user pool.
//
//    // Example sending a request using ListUsersRequest.
//    req := client.ListUsersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/ListUsers
func (c *Client) ListUsersRequest(input *ListUsersInput) ListUsersRequest {
	op := &aws.Operation{
		Name:       opListUsers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListUsersInput{}
	}

	req := c.newRequest(op, input, &ListUsersOutput{})
	return ListUsersRequest{Request: req, Input: input, Copy: c.ListUsersRequest}
}

// ListUsersRequest is the request type for the
// ListUsers API operation.
type ListUsersRequest struct {
	*aws.Request
	Input *ListUsersInput
	Copy  func(*ListUsersInput) ListUsersRequest
}

// Send marshals and sends the ListUsers API request.
func (r ListUsersRequest) Send(ctx context.Context) (*ListUsersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListUsersResponse{
		ListUsersOutput: r.Request.Data.(*ListUsersOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListUsersResponse is the response type for the
// ListUsers API operation.
type ListUsersResponse struct {
	*ListUsersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListUsers request.
func (r *ListUsersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}