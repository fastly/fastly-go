# LoggingS3AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **NullableString** | The access key for your S3 account. Not required if `iam_role` is provided. | [optional] 
**ACL** | Pointer to **string** | The access control list (ACL) specific request header. See the AWS documentation for [Access Control List (ACL) Specific Request Headers](https://docs.aws.amazon.com/AmazonS3/latest/API/mpUploadInitiate.html#initiate-mpu-acl-specific-request-headers) for more information. | [optional] 
**BucketName** | Pointer to **string** | The bucket name for S3 account. | [optional] 
**Domain** | Pointer to **string** | The domain of the Amazon S3 endpoint. | [optional] 
**IamRole** | Pointer to **NullableString** | The Amazon Resource Name (ARN) for the IAM role granting Fastly access to S3. Not required if `access_key` and `secret_key` are provided. | [optional] 
**Path** | Pointer to **NullableString** | The path to upload logs to. | [optional] [default to "null"]
**PublicKey** | Pointer to **NullableString** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [optional] [default to "null"]
**Redundancy** | Pointer to **NullableString** | The S3 redundancy level. | [optional] [default to "null"]
**SecretKey** | Pointer to **NullableString** | The secret key for your S3 account. Not required if `iam_role` is provided. | [optional] 
**ServerSideEncryptionKmsKeyID** | Pointer to **NullableString** | Optional server-side KMS Key ID. Must be set if `server_side_encryption` is set to `aws:kms` or `AES256`. | [optional] [default to "null"]
**ServerSideEncryption** | Pointer to **NullableString** | Set this to `AES256` or `aws:kms` to enable S3 Server Side Encryption. | [optional] [default to "null"]

## Methods

### NewLoggingS3AllOf

`func NewLoggingS3AllOf() *LoggingS3AllOf`

NewLoggingS3AllOf instantiates a new LoggingS3AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingS3AllOfWithDefaults

`func NewLoggingS3AllOfWithDefaults() *LoggingS3AllOf`

NewLoggingS3AllOfWithDefaults instantiates a new LoggingS3AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *LoggingS3AllOf) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *LoggingS3AllOf) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *LoggingS3AllOf) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *LoggingS3AllOf) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### SetAccessKeyNil

`func (o *LoggingS3AllOf) SetAccessKeyNil(b bool)`

 SetAccessKeyNil sets the value for AccessKey to be an explicit nil

### UnsetAccessKey
`func (o *LoggingS3AllOf) UnsetAccessKey()`

UnsetAccessKey ensures that no value is present for AccessKey, not even an explicit nil
### GetACL

`func (o *LoggingS3AllOf) GetACL() string`

GetACL returns the ACL field if non-nil, zero value otherwise.

### GetACLOk

`func (o *LoggingS3AllOf) GetACLOk() (*string, bool)`

GetACLOk returns a tuple with the ACL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACL

`func (o *LoggingS3AllOf) SetACL(v string)`

SetACL sets ACL field to given value.

### HasACL

`func (o *LoggingS3AllOf) HasACL() bool`

HasACL returns a boolean if a field has been set.

### GetBucketName

`func (o *LoggingS3AllOf) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *LoggingS3AllOf) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *LoggingS3AllOf) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *LoggingS3AllOf) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetDomain

`func (o *LoggingS3AllOf) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *LoggingS3AllOf) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *LoggingS3AllOf) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *LoggingS3AllOf) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetIamRole

`func (o *LoggingS3AllOf) GetIamRole() string`

GetIamRole returns the IamRole field if non-nil, zero value otherwise.

### GetIamRoleOk

`func (o *LoggingS3AllOf) GetIamRoleOk() (*string, bool)`

GetIamRoleOk returns a tuple with the IamRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRole

`func (o *LoggingS3AllOf) SetIamRole(v string)`

SetIamRole sets IamRole field to given value.

### HasIamRole

`func (o *LoggingS3AllOf) HasIamRole() bool`

HasIamRole returns a boolean if a field has been set.

### SetIamRoleNil

`func (o *LoggingS3AllOf) SetIamRoleNil(b bool)`

 SetIamRoleNil sets the value for IamRole to be an explicit nil

### UnsetIamRole
`func (o *LoggingS3AllOf) UnsetIamRole()`

UnsetIamRole ensures that no value is present for IamRole, not even an explicit nil
### GetPath

`func (o *LoggingS3AllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LoggingS3AllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LoggingS3AllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LoggingS3AllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *LoggingS3AllOf) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *LoggingS3AllOf) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPublicKey

`func (o *LoggingS3AllOf) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *LoggingS3AllOf) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *LoggingS3AllOf) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *LoggingS3AllOf) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *LoggingS3AllOf) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *LoggingS3AllOf) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetRedundancy

`func (o *LoggingS3AllOf) GetRedundancy() string`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *LoggingS3AllOf) GetRedundancyOk() (*string, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *LoggingS3AllOf) SetRedundancy(v string)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *LoggingS3AllOf) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### SetRedundancyNil

`func (o *LoggingS3AllOf) SetRedundancyNil(b bool)`

 SetRedundancyNil sets the value for Redundancy to be an explicit nil

### UnsetRedundancy
`func (o *LoggingS3AllOf) UnsetRedundancy()`

UnsetRedundancy ensures that no value is present for Redundancy, not even an explicit nil
### GetSecretKey

`func (o *LoggingS3AllOf) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *LoggingS3AllOf) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *LoggingS3AllOf) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *LoggingS3AllOf) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### SetSecretKeyNil

`func (o *LoggingS3AllOf) SetSecretKeyNil(b bool)`

 SetSecretKeyNil sets the value for SecretKey to be an explicit nil

### UnsetSecretKey
`func (o *LoggingS3AllOf) UnsetSecretKey()`

UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
### GetServerSideEncryptionKmsKeyID

`func (o *LoggingS3AllOf) GetServerSideEncryptionKmsKeyID() string`

GetServerSideEncryptionKmsKeyID returns the ServerSideEncryptionKmsKeyID field if non-nil, zero value otherwise.

### GetServerSideEncryptionKmsKeyIDOk

`func (o *LoggingS3AllOf) GetServerSideEncryptionKmsKeyIDOk() (*string, bool)`

GetServerSideEncryptionKmsKeyIDOk returns a tuple with the ServerSideEncryptionKmsKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSideEncryptionKmsKeyID

`func (o *LoggingS3AllOf) SetServerSideEncryptionKmsKeyID(v string)`

SetServerSideEncryptionKmsKeyID sets ServerSideEncryptionKmsKeyID field to given value.

### HasServerSideEncryptionKmsKeyID

`func (o *LoggingS3AllOf) HasServerSideEncryptionKmsKeyID() bool`

HasServerSideEncryptionKmsKeyID returns a boolean if a field has been set.

### SetServerSideEncryptionKmsKeyIDNil

`func (o *LoggingS3AllOf) SetServerSideEncryptionKmsKeyIDNil(b bool)`

 SetServerSideEncryptionKmsKeyIDNil sets the value for ServerSideEncryptionKmsKeyID to be an explicit nil

### UnsetServerSideEncryptionKmsKeyID
`func (o *LoggingS3AllOf) UnsetServerSideEncryptionKmsKeyID()`

UnsetServerSideEncryptionKmsKeyID ensures that no value is present for ServerSideEncryptionKmsKeyID, not even an explicit nil
### GetServerSideEncryption

`func (o *LoggingS3AllOf) GetServerSideEncryption() string`

GetServerSideEncryption returns the ServerSideEncryption field if non-nil, zero value otherwise.

### GetServerSideEncryptionOk

`func (o *LoggingS3AllOf) GetServerSideEncryptionOk() (*string, bool)`

GetServerSideEncryptionOk returns a tuple with the ServerSideEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSideEncryption

`func (o *LoggingS3AllOf) SetServerSideEncryption(v string)`

SetServerSideEncryption sets ServerSideEncryption field to given value.

### HasServerSideEncryption

`func (o *LoggingS3AllOf) HasServerSideEncryption() bool`

HasServerSideEncryption returns a boolean if a field has been set.

### SetServerSideEncryptionNil

`func (o *LoggingS3AllOf) SetServerSideEncryptionNil(b bool)`

 SetServerSideEncryptionNil sets the value for ServerSideEncryption to be an explicit nil

### UnsetServerSideEncryption
`func (o *LoggingS3AllOf) UnsetServerSideEncryption()`

UnsetServerSideEncryption ensures that no value is present for ServerSideEncryption, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
