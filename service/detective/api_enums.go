// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package detective

type MemberDisabledReason string

// Enum values for MemberDisabledReason
const (
	MemberDisabledReasonVolumeTooHigh MemberDisabledReason = "VOLUME_TOO_HIGH"
	MemberDisabledReasonVolumeUnknown MemberDisabledReason = "VOLUME_UNKNOWN"
)

func (enum MemberDisabledReason) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MemberDisabledReason) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MemberStatus string

// Enum values for MemberStatus
const (
	MemberStatusInvited                MemberStatus = "INVITED"
	MemberStatusVerificationInProgress MemberStatus = "VERIFICATION_IN_PROGRESS"
	MemberStatusVerificationFailed     MemberStatus = "VERIFICATION_FAILED"
	MemberStatusEnabled                MemberStatus = "ENABLED"
	MemberStatusAcceptedButDisabled    MemberStatus = "ACCEPTED_BUT_DISABLED"
)

func (enum MemberStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MemberStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
