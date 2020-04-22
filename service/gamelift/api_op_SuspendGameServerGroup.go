// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SuspendGameServerGroupInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the game server group to stop activity on. Use either
	// the GameServerGroup name or ARN value.
	//
	// GameServerGroupName is a required field
	GameServerGroupName *string `min:"1" type:"string" required:"true"`

	// The action to suspend for this game server group.
	//
	// SuspendActions is a required field
	SuspendActions []GameServerGroupAction `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s SuspendGameServerGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SuspendGameServerGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SuspendGameServerGroupInput"}

	if s.GameServerGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GameServerGroupName"))
	}
	if s.GameServerGroupName != nil && len(*s.GameServerGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameServerGroupName", 1))
	}

	if s.SuspendActions == nil {
		invalidParams.Add(aws.NewErrParamRequired("SuspendActions"))
	}
	if s.SuspendActions != nil && len(s.SuspendActions) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SuspendActions", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SuspendGameServerGroupOutput struct {
	_ struct{} `type:"structure"`

	// An object that describes the game server group resource, with the SuspendedActions
	// property updated to reflect the suspended activity.
	GameServerGroup *GameServerGroup `type:"structure"`
}

// String returns the string representation
func (s SuspendGameServerGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opSuspendGameServerGroup = "SuspendGameServerGroup"

// SuspendGameServerGroupRequest returns a request value for making API operation for
// Amazon GameLift.
//
// This action is part of Amazon GameLift FleetIQ with game server groups, which
// is in preview release and is subject to change.
//
// Temporarily stops activity on a game server group without terminating instances
// or the game server group. Activity can be restarted by calling ResumeGameServerGroup.
// Activities that can suspended are:
//
//    * Instance type replacement. This activity evaluates the current Spot
//    viability of all instance types that are defined for the game server group.
//    It updates the Auto Scaling group to remove nonviable Spot instance types
//    (which have a higher chance of game server interruptions) and rebalances
//    capacity across the remaining viable Spot instance types. When this activity
//    is suspended, the Auto Scaling group continues with its current balance,
//    regardless of viability. Instance protection, utilization metrics, and
//    capacity autoscaling activities continue to be active.
//
// To suspend activity, specify a game server group ARN and the type of activity
// to be suspended.
//
// Learn more
//
// GameLift FleetIQ Guide (https://docs.aws.amazon.com/gamelift/latest/developerguide/gsg-intro.html)
//
// Related operations
//
//    * CreateGameServerGroup
//
//    * ListGameServerGroups
//
//    * DescribeGameServerGroup
//
//    * UpdateGameServerGroup
//
//    * DeleteGameServerGroup
//
//    * ResumeGameServerGroup
//
//    * SuspendGameServerGroup
//
//    // Example sending a request using SuspendGameServerGroupRequest.
//    req := client.SuspendGameServerGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/SuspendGameServerGroup
func (c *Client) SuspendGameServerGroupRequest(input *SuspendGameServerGroupInput) SuspendGameServerGroupRequest {
	op := &aws.Operation{
		Name:       opSuspendGameServerGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SuspendGameServerGroupInput{}
	}

	req := c.newRequest(op, input, &SuspendGameServerGroupOutput{})
	return SuspendGameServerGroupRequest{Request: req, Input: input, Copy: c.SuspendGameServerGroupRequest}
}

// SuspendGameServerGroupRequest is the request type for the
// SuspendGameServerGroup API operation.
type SuspendGameServerGroupRequest struct {
	*aws.Request
	Input *SuspendGameServerGroupInput
	Copy  func(*SuspendGameServerGroupInput) SuspendGameServerGroupRequest
}

// Send marshals and sends the SuspendGameServerGroup API request.
func (r SuspendGameServerGroupRequest) Send(ctx context.Context) (*SuspendGameServerGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SuspendGameServerGroupResponse{
		SuspendGameServerGroupOutput: r.Request.Data.(*SuspendGameServerGroupOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SuspendGameServerGroupResponse is the response type for the
// SuspendGameServerGroup API operation.
type SuspendGameServerGroupResponse struct {
	*SuspendGameServerGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SuspendGameServerGroup request.
func (r *SuspendGameServerGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
