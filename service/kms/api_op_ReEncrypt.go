// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ReEncryptInput struct {
	_ struct{} `type:"structure"`

	// Ciphertext of the data to reencrypt.
	//
	// CiphertextBlob is automatically base64 encoded/decoded by the SDK.
	//
	// CiphertextBlob is a required field
	CiphertextBlob []byte `min:"1" type:"blob" required:"true"`

	// Specifies the encryption algorithm that AWS KMS will use to reecrypt the
	// data after it has decrypted it. The default value, SYMMETRIC_DEFAULT, represents
	// the encryption algorithm used for symmetric CMKs.
	//
	// This parameter is required only when the destination CMK is an asymmetric
	// CMK.
	DestinationEncryptionAlgorithm EncryptionAlgorithmSpec `type:"string" enum:"true"`

	// Specifies that encryption context to use when the reencrypting the data.
	//
	// A destination encryption context is valid only when the destination CMK is
	// a symmetric CMK. The standard ciphertext format for asymmetric CMKs does
	// not include fields for metadata.
	//
	// An encryption context is a collection of non-secret key-value pairs that
	// represents additional authenticated data. When you use an encryption context
	// to encrypt data, you must specify the same (an exact case-sensitive match)
	// encryption context to decrypt the data. An encryption context is optional
	// when encrypting with a symmetric CMK, but it is highly recommended.
	//
	// For more information, see Encryption Context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the AWS Key Management Service Developer Guide.
	DestinationEncryptionContext map[string]string `type:"map"`

	// A unique identifier for the CMK that is used to reencrypt the data. Specify
	// a symmetric or asymmetric CMK with a KeyUsage value of ENCRYPT_DECRYPT. To
	// find the KeyUsage value of a CMK, use the DescribeKey operation.
	//
	// To specify a CMK, use its key ID, Amazon Resource Name (ARN), alias name,
	// or alias ARN. When using an alias name, prefix it with "alias/". To specify
	// a CMK in a different AWS account, you must use the key ARN or alias ARN.
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Alias name: alias/ExampleAlias
	//
	//    * Alias ARN: arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// To get the key ID and key ARN for a CMK, use ListKeys or DescribeKey. To
	// get the alias name and alias ARN, use ListAliases.
	//
	// DestinationKeyId is a required field
	DestinationKeyId *string `min:"1" type:"string" required:"true"`

	// A list of grant tokens.
	//
	// For more information, see Grant Tokens (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantTokens []string `type:"list"`

	// Specifies the encryption algorithm that AWS KMS will use to decrypt the ciphertext
	// before it is reencrypted. The default value, SYMMETRIC_DEFAULT, represents
	// the algorithm used for symmetric CMKs.
	//
	// Specify the same algorithm that was used to encrypt the ciphertext. If you
	// specify a different algorithm, the decrypt attempt fails.
	//
	// This parameter is required only when the ciphertext was encrypted under an
	// asymmetric CMK.
	SourceEncryptionAlgorithm EncryptionAlgorithmSpec `type:"string" enum:"true"`

	// Specifies the encryption context to use to decrypt the ciphertext. Enter
	// the same encryption context that was used to encrypt the ciphertext.
	//
	// An encryption context is a collection of non-secret key-value pairs that
	// represents additional authenticated data. When you use an encryption context
	// to encrypt data, you must specify the same (an exact case-sensitive match)
	// encryption context to decrypt the data. An encryption context is optional
	// when encrypting with a symmetric CMK, but it is highly recommended.
	//
	// For more information, see Encryption Context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the AWS Key Management Service Developer Guide.
	SourceEncryptionContext map[string]string `type:"map"`

	// A unique identifier for the CMK that is used to decrypt the ciphertext before
	// it reencrypts it using the destination CMK.
	//
	// This parameter is required only when the ciphertext was encrypted under an
	// asymmetric CMK. Otherwise, AWS KMS uses the metadata that it adds to the
	// ciphertext blob to determine which CMK was used to encrypt the ciphertext.
	// However, you can use this parameter to ensure that a particular CMK (of any
	// kind) is used to decrypt the ciphertext before it is reencrypted.
	//
	// If you specify a KeyId value, the decrypt part of the ReEncrypt operation
	// succeeds only if the specified CMK was used to encrypt the ciphertext.
	//
	// To specify a CMK, use its key ID, Amazon Resource Name (ARN), alias name,
	// or alias ARN. When using an alias name, prefix it with "alias/".
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Alias name: alias/ExampleAlias
	//
	//    * Alias ARN: arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// To get the key ID and key ARN for a CMK, use ListKeys or DescribeKey. To
	// get the alias name and alias ARN, use ListAliases.
	SourceKeyId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ReEncryptInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReEncryptInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ReEncryptInput"}

	if s.CiphertextBlob == nil {
		invalidParams.Add(aws.NewErrParamRequired("CiphertextBlob"))
	}
	if s.CiphertextBlob != nil && len(s.CiphertextBlob) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CiphertextBlob", 1))
	}

	if s.DestinationKeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationKeyId"))
	}
	if s.DestinationKeyId != nil && len(*s.DestinationKeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DestinationKeyId", 1))
	}
	if s.SourceKeyId != nil && len(*s.SourceKeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SourceKeyId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ReEncryptOutput struct {
	_ struct{} `type:"structure"`

	// The reencrypted data. When you use the HTTP API or the AWS CLI, the value
	// is Base64-encoded. Otherwise, it is not Base64-encoded.
	//
	// CiphertextBlob is automatically base64 encoded/decoded by the SDK.
	CiphertextBlob []byte `min:"1" type:"blob"`

	// The encryption algorithm that was used to reencrypt the data.
	DestinationEncryptionAlgorithm EncryptionAlgorithmSpec `type:"string" enum:"true"`

	// Unique identifier of the CMK used to reencrypt the data.
	KeyId *string `min:"1" type:"string"`

	// The encryption algorithm that was used to decrypt the ciphertext before it
	// was reencrypted.
	SourceEncryptionAlgorithm EncryptionAlgorithmSpec `type:"string" enum:"true"`

	// Unique identifier of the CMK used to originally encrypt the data.
	SourceKeyId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ReEncryptOutput) String() string {
	return awsutil.Prettify(s)
}

const opReEncrypt = "ReEncrypt"

// ReEncryptRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Decrypts ciphertext and then reencrypts it entirely within AWS KMS. You can
// use this operation to change the customer master key (CMK) under which data
// is encrypted, such as when you manually rotate (https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html#rotate-keys-manually)
// a CMK or change the CMK that protects a ciphertext. You can also use it to
// reencrypt ciphertext under the same CMK, such as to change the encryption
// context of a ciphertext.
//
// The ReEncrypt operation can decrypt ciphertext that was encrypted by using
// an AWS KMS CMK in an AWS KMS operation, such as Encrypt or GenerateDataKey.
// It can also decrypt ciphertext that was encrypted by using the public key
// of an asymmetric CMK outside of AWS KMS. However, it cannot decrypt ciphertext
// produced by other libraries, such as the AWS Encryption SDK (https://docs.aws.amazon.com/encryption-sdk/latest/developer-guide/)
// or Amazon S3 client-side encryption (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html).
// These libraries return a ciphertext format that is incompatible with AWS
// KMS.
//
// When you use the ReEncrypt operation, you need to provide information for
// the decrypt operation and the subsequent encrypt operation.
//
//    * If your ciphertext was encrypted under an asymmetric CMK, you must identify
//    the source CMK, that is, the CMK that encrypted the ciphertext. You must
//    also supply the encryption algorithm that was used. This information is
//    required to decrypt the data.
//
//    * It is optional, but you can specify a source CMK even when the ciphertext
//    was encrypted under a symmetric CMK. This ensures that the ciphertext
//    is decrypted only by using a particular CMK. If the CMK that you specify
//    cannot decrypt the ciphertext, the ReEncrypt operation fails.
//
//    * To reencrypt the data, you must specify the destination CMK, that is,
//    the CMK that re-encrypts the data after it is decrypted. You can select
//    a symmetric or asymmetric CMK. If the destination CMK is an asymmetric
//    CMK, you must also provide the encryption algorithm. The algorithm that
//    you choose must be compatible with the CMK. When you use an asymmetric
//    CMK to encrypt or reencrypt data, be sure to record the CMK and encryption
//    algorithm that you choose. You will be required to provide the same CMK
//    and encryption algorithm when you decrypt the data. If the CMK and algorithm
//    do not match the values used to encrypt the data, the decrypt operation
//    fails. You are not required to supply the CMK ID and encryption algorithm
//    when you decrypt with symmetric CMKs because AWS KMS stores this information
//    in the ciphertext blob. AWS KMS cannot store metadata in ciphertext generated
//    with asymmetric keys. The standard format for asymmetric key ciphertext
//    does not include configurable fields.
//
// Unlike other AWS KMS API operations, ReEncrypt callers must have two permissions:
//
//    * kms:EncryptFrom permission on the source CMK
//
//    * kms:EncryptTo permission on the destination CMK
//
// To permit reencryption from
//
// or to a CMK, include the "kms:ReEncrypt*" permission in your key policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html).
// This permission is automatically included in the key policy when you use
// the console to create a CMK. But you must include it manually when you create
// a CMK programmatically or when you use the PutKeyPolicy operation set a key
// policy.
//
// The CMK that you use for this operation must be in a compatible key state.
// For details, see How Key State Affects Use of a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the AWS Key Management Service Developer Guide.
//
//    // Example sending a request using ReEncryptRequest.
//    req := client.ReEncryptRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/ReEncrypt
func (c *Client) ReEncryptRequest(input *ReEncryptInput) ReEncryptRequest {
	op := &aws.Operation{
		Name:       opReEncrypt,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ReEncryptInput{}
	}

	req := c.newRequest(op, input, &ReEncryptOutput{})

	return ReEncryptRequest{Request: req, Input: input, Copy: c.ReEncryptRequest}
}

// ReEncryptRequest is the request type for the
// ReEncrypt API operation.
type ReEncryptRequest struct {
	*aws.Request
	Input *ReEncryptInput
	Copy  func(*ReEncryptInput) ReEncryptRequest
}

// Send marshals and sends the ReEncrypt API request.
func (r ReEncryptRequest) Send(ctx context.Context) (*ReEncryptResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReEncryptResponse{
		ReEncryptOutput: r.Request.Data.(*ReEncryptOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReEncryptResponse is the response type for the
// ReEncrypt API operation.
type ReEncryptResponse struct {
	*ReEncryptOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReEncrypt request.
func (r *ReEncryptResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
