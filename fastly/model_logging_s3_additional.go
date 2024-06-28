// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://www.fastly.com/documentation/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// LoggingS3Additional struct for LoggingS3Additional
type LoggingS3Additional struct {
	// The access key for your S3 account. Not required if `iam_role` is provided.
	AccessKey NullableString `json:"access_key,omitempty"`
	// The access control list (ACL) specific request header. See the AWS documentation for [Access Control List (ACL) Specific Request Headers](https://docs.aws.amazon.com/AmazonS3/latest/API/mpUploadInitiate.html#initiate-mpu-acl-specific-request-headers) for more information.
	ACL *string `json:"acl,omitempty"`
	// The bucket name for S3 account.
	BucketName *string `json:"bucket_name,omitempty"`
	// The domain of the Amazon S3 endpoint.
	Domain *string `json:"domain,omitempty"`
	// The Amazon Resource Name (ARN) for the IAM role granting Fastly access to S3. Not required if `access_key` and `secret_key` are provided.
	IamRole NullableString `json:"iam_role,omitempty"`
	// The path to upload logs to.
	Path NullableString `json:"path,omitempty"`
	// A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
	PublicKey NullableString `json:"public_key,omitempty"`
	// The S3 redundancy level.
	Redundancy NullableString `json:"redundancy,omitempty"`
	// The secret key for your S3 account. Not required if `iam_role` is provided.
	SecretKey NullableString `json:"secret_key,omitempty"`
	// Optional server-side KMS Key ID. Must be set if `server_side_encryption` is set to `aws:kms` or `AES256`.
	ServerSideEncryptionKmsKeyID NullableString `json:"server_side_encryption_kms_key_id,omitempty"`
	// Set this to `AES256` or `aws:kms` to enable S3 Server Side Encryption.
	ServerSideEncryption NullableString `json:"server_side_encryption,omitempty"`
	// The maximum number of bytes for each uploaded file. A value of 0 can be used to indicate there is no limit on the size of uploaded files, otherwise the minimum value is 1048576 bytes (1 MiB.)
	FileMaxBytes *int32 `json:"file_max_bytes,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingS3Additional LoggingS3Additional

// NewLoggingS3Additional instantiates a new LoggingS3Additional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingS3Additional() *LoggingS3Additional {
	this := LoggingS3Additional{}
	var path string = "null"
	this.Path = *NewNullableString(&path)
	var publicKey string = "null"
	this.PublicKey = *NewNullableString(&publicKey)
	var redundancy string = "null"
	this.Redundancy = *NewNullableString(&redundancy)
	var serverSideEncryptionKmsKeyID string = "null"
	this.ServerSideEncryptionKmsKeyID = *NewNullableString(&serverSideEncryptionKmsKeyID)
	var serverSideEncryption string = "null"
	this.ServerSideEncryption = *NewNullableString(&serverSideEncryption)
	return &this
}

// NewLoggingS3AdditionalWithDefaults instantiates a new LoggingS3Additional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingS3AdditionalWithDefaults() *LoggingS3Additional {
	this := LoggingS3Additional{}
	var path string = "null"
	this.Path = *NewNullableString(&path)
	var publicKey string = "null"
	this.PublicKey = *NewNullableString(&publicKey)
	var redundancy string = "null"
	this.Redundancy = *NewNullableString(&redundancy)
	var serverSideEncryptionKmsKeyID string = "null"
	this.ServerSideEncryptionKmsKeyID = *NewNullableString(&serverSideEncryptionKmsKeyID)
	var serverSideEncryption string = "null"
	this.ServerSideEncryption = *NewNullableString(&serverSideEncryption)
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingS3Additional) GetAccessKey() string {
	if o == nil || o.AccessKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccessKey.Get()
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingS3Additional) GetAccessKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccessKey.Get(), o.AccessKey.IsSet()
}

// HasAccessKey returns a boolean if a field has been set.
func (o *LoggingS3Additional) HasAccessKey() bool {
	if o != nil && o.AccessKey.IsSet() {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given NullableString and assigns it to the AccessKey field.
func (o *LoggingS3Additional) SetAccessKey(v string) {
	o.AccessKey.Set(&v)
}
// SetAccessKeyNil sets the value for AccessKey to be an explicit nil
func (o *LoggingS3Additional) SetAccessKeyNil() {
	o.AccessKey.Set(nil)
}

// UnsetAccessKey ensures that no value is present for AccessKey, not even an explicit nil
func (o *LoggingS3Additional) UnsetAccessKey() {
	o.AccessKey.Unset()
}

// GetACL returns the ACL field value if set, zero value otherwise.
func (o *LoggingS3Additional) GetACL() string {
	if o == nil || o.ACL == nil {
		var ret string
		return ret
	}
	return *o.ACL
}

// GetACLOk returns a tuple with the ACL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingS3Additional) GetACLOk() (*string, bool) {
	if o == nil || o.ACL == nil {
		return nil, false
	}
	return o.ACL, true
}

// HasACL returns a boolean if a field has been set.
func (o *LoggingS3Additional) HasACL() bool {
	if o != nil && o.ACL != nil {
		return true
	}

	return false
}

// SetACL gets a reference to the given string and assigns it to the ACL field.
func (o *LoggingS3Additional) SetACL(v string) {
	o.ACL = &v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *LoggingS3Additional) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingS3Additional) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *LoggingS3Additional) HasBucketName() bool {
	if o != nil && o.BucketName != nil {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *LoggingS3Additional) SetBucketName(v string) {
	o.BucketName = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *LoggingS3Additional) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingS3Additional) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *LoggingS3Additional) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *LoggingS3Additional) SetDomain(v string) {
	o.Domain = &v
}

// GetIamRole returns the IamRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingS3Additional) GetIamRole() string {
	if o == nil || o.IamRole.Get() == nil {
		var ret string
		return ret
	}
	return *o.IamRole.Get()
}

// GetIamRoleOk returns a tuple with the IamRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingS3Additional) GetIamRoleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IamRole.Get(), o.IamRole.IsSet()
}

// HasIamRole returns a boolean if a field has been set.
func (o *LoggingS3Additional) HasIamRole() bool {
	if o != nil && o.IamRole.IsSet() {
		return true
	}

	return false
}

// SetIamRole gets a reference to the given NullableString and assigns it to the IamRole field.
func (o *LoggingS3Additional) SetIamRole(v string) {
	o.IamRole.Set(&v)
}
// SetIamRoleNil sets the value for IamRole to be an explicit nil
func (o *LoggingS3Additional) SetIamRoleNil() {
	o.IamRole.Set(nil)
}

// UnsetIamRole ensures that no value is present for IamRole, not even an explicit nil
func (o *LoggingS3Additional) UnsetIamRole() {
	o.IamRole.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingS3Additional) GetPath() string {
	if o == nil || o.Path.Get() == nil {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingS3Additional) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *LoggingS3Additional) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *LoggingS3Additional) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *LoggingS3Additional) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *LoggingS3Additional) UnsetPath() {
	o.Path.Unset()
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingS3Additional) GetPublicKey() string {
	if o == nil || o.PublicKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingS3Additional) GetPublicKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// HasPublicKey returns a boolean if a field has been set.
func (o *LoggingS3Additional) HasPublicKey() bool {
	if o != nil && o.PublicKey.IsSet() {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given NullableString and assigns it to the PublicKey field.
func (o *LoggingS3Additional) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}
// SetPublicKeyNil sets the value for PublicKey to be an explicit nil
func (o *LoggingS3Additional) SetPublicKeyNil() {
	o.PublicKey.Set(nil)
}

// UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
func (o *LoggingS3Additional) UnsetPublicKey() {
	o.PublicKey.Unset()
}

// GetRedundancy returns the Redundancy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingS3Additional) GetRedundancy() string {
	if o == nil || o.Redundancy.Get() == nil {
		var ret string
		return ret
	}
	return *o.Redundancy.Get()
}

// GetRedundancyOk returns a tuple with the Redundancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingS3Additional) GetRedundancyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Redundancy.Get(), o.Redundancy.IsSet()
}

// HasRedundancy returns a boolean if a field has been set.
func (o *LoggingS3Additional) HasRedundancy() bool {
	if o != nil && o.Redundancy.IsSet() {
		return true
	}

	return false
}

// SetRedundancy gets a reference to the given NullableString and assigns it to the Redundancy field.
func (o *LoggingS3Additional) SetRedundancy(v string) {
	o.Redundancy.Set(&v)
}
// SetRedundancyNil sets the value for Redundancy to be an explicit nil
func (o *LoggingS3Additional) SetRedundancyNil() {
	o.Redundancy.Set(nil)
}

// UnsetRedundancy ensures that no value is present for Redundancy, not even an explicit nil
func (o *LoggingS3Additional) UnsetRedundancy() {
	o.Redundancy.Unset()
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingS3Additional) GetSecretKey() string {
	if o == nil || o.SecretKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.SecretKey.Get()
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingS3Additional) GetSecretKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SecretKey.Get(), o.SecretKey.IsSet()
}

// HasSecretKey returns a boolean if a field has been set.
func (o *LoggingS3Additional) HasSecretKey() bool {
	if o != nil && o.SecretKey.IsSet() {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given NullableString and assigns it to the SecretKey field.
func (o *LoggingS3Additional) SetSecretKey(v string) {
	o.SecretKey.Set(&v)
}
// SetSecretKeyNil sets the value for SecretKey to be an explicit nil
func (o *LoggingS3Additional) SetSecretKeyNil() {
	o.SecretKey.Set(nil)
}

// UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
func (o *LoggingS3Additional) UnsetSecretKey() {
	o.SecretKey.Unset()
}

// GetServerSideEncryptionKmsKeyID returns the ServerSideEncryptionKmsKeyID field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingS3Additional) GetServerSideEncryptionKmsKeyID() string {
	if o == nil || o.ServerSideEncryptionKmsKeyID.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServerSideEncryptionKmsKeyID.Get()
}

// GetServerSideEncryptionKmsKeyIDOk returns a tuple with the ServerSideEncryptionKmsKeyID field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingS3Additional) GetServerSideEncryptionKmsKeyIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServerSideEncryptionKmsKeyID.Get(), o.ServerSideEncryptionKmsKeyID.IsSet()
}

// HasServerSideEncryptionKmsKeyID returns a boolean if a field has been set.
func (o *LoggingS3Additional) HasServerSideEncryptionKmsKeyID() bool {
	if o != nil && o.ServerSideEncryptionKmsKeyID.IsSet() {
		return true
	}

	return false
}

// SetServerSideEncryptionKmsKeyID gets a reference to the given NullableString and assigns it to the ServerSideEncryptionKmsKeyID field.
func (o *LoggingS3Additional) SetServerSideEncryptionKmsKeyID(v string) {
	o.ServerSideEncryptionKmsKeyID.Set(&v)
}
// SetServerSideEncryptionKmsKeyIDNil sets the value for ServerSideEncryptionKmsKeyID to be an explicit nil
func (o *LoggingS3Additional) SetServerSideEncryptionKmsKeyIDNil() {
	o.ServerSideEncryptionKmsKeyID.Set(nil)
}

// UnsetServerSideEncryptionKmsKeyID ensures that no value is present for ServerSideEncryptionKmsKeyID, not even an explicit nil
func (o *LoggingS3Additional) UnsetServerSideEncryptionKmsKeyID() {
	o.ServerSideEncryptionKmsKeyID.Unset()
}

// GetServerSideEncryption returns the ServerSideEncryption field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingS3Additional) GetServerSideEncryption() string {
	if o == nil || o.ServerSideEncryption.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServerSideEncryption.Get()
}

// GetServerSideEncryptionOk returns a tuple with the ServerSideEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingS3Additional) GetServerSideEncryptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServerSideEncryption.Get(), o.ServerSideEncryption.IsSet()
}

// HasServerSideEncryption returns a boolean if a field has been set.
func (o *LoggingS3Additional) HasServerSideEncryption() bool {
	if o != nil && o.ServerSideEncryption.IsSet() {
		return true
	}

	return false
}

// SetServerSideEncryption gets a reference to the given NullableString and assigns it to the ServerSideEncryption field.
func (o *LoggingS3Additional) SetServerSideEncryption(v string) {
	o.ServerSideEncryption.Set(&v)
}
// SetServerSideEncryptionNil sets the value for ServerSideEncryption to be an explicit nil
func (o *LoggingS3Additional) SetServerSideEncryptionNil() {
	o.ServerSideEncryption.Set(nil)
}

// UnsetServerSideEncryption ensures that no value is present for ServerSideEncryption, not even an explicit nil
func (o *LoggingS3Additional) UnsetServerSideEncryption() {
	o.ServerSideEncryption.Unset()
}

// GetFileMaxBytes returns the FileMaxBytes field value if set, zero value otherwise.
func (o *LoggingS3Additional) GetFileMaxBytes() int32 {
	if o == nil || o.FileMaxBytes == nil {
		var ret int32
		return ret
	}
	return *o.FileMaxBytes
}

// GetFileMaxBytesOk returns a tuple with the FileMaxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingS3Additional) GetFileMaxBytesOk() (*int32, bool) {
	if o == nil || o.FileMaxBytes == nil {
		return nil, false
	}
	return o.FileMaxBytes, true
}

// HasFileMaxBytes returns a boolean if a field has been set.
func (o *LoggingS3Additional) HasFileMaxBytes() bool {
	if o != nil && o.FileMaxBytes != nil {
		return true
	}

	return false
}

// SetFileMaxBytes gets a reference to the given int32 and assigns it to the FileMaxBytes field.
func (o *LoggingS3Additional) SetFileMaxBytes(v int32) {
	o.FileMaxBytes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingS3Additional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.AccessKey.IsSet() {
		toSerialize["access_key"] = o.AccessKey.Get()
	}
	if o.ACL != nil {
		toSerialize["acl"] = o.ACL
	}
	if o.BucketName != nil {
		toSerialize["bucket_name"] = o.BucketName
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.IamRole.IsSet() {
		toSerialize["iam_role"] = o.IamRole.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.PublicKey.IsSet() {
		toSerialize["public_key"] = o.PublicKey.Get()
	}
	if o.Redundancy.IsSet() {
		toSerialize["redundancy"] = o.Redundancy.Get()
	}
	if o.SecretKey.IsSet() {
		toSerialize["secret_key"] = o.SecretKey.Get()
	}
	if o.ServerSideEncryptionKmsKeyID.IsSet() {
		toSerialize["server_side_encryption_kms_key_id"] = o.ServerSideEncryptionKmsKeyID.Get()
	}
	if o.ServerSideEncryption.IsSet() {
		toSerialize["server_side_encryption"] = o.ServerSideEncryption.Get()
	}
	if o.FileMaxBytes != nil {
		toSerialize["file_max_bytes"] = o.FileMaxBytes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *LoggingS3Additional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingS3Additional := _LoggingS3Additional{}

	if err = json.Unmarshal(bytes, &varLoggingS3Additional); err == nil {
		*o = LoggingS3Additional(varLoggingS3Additional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "access_key")
		delete(additionalProperties, "acl")
		delete(additionalProperties, "bucket_name")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "iam_role")
		delete(additionalProperties, "path")
		delete(additionalProperties, "public_key")
		delete(additionalProperties, "redundancy")
		delete(additionalProperties, "secret_key")
		delete(additionalProperties, "server_side_encryption_kms_key_id")
		delete(additionalProperties, "server_side_encryption")
		delete(additionalProperties, "file_max_bytes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingS3Additional is a helper abstraction for handling nullable loggings3additional types. 
type NullableLoggingS3Additional struct {
	value *LoggingS3Additional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingS3Additional) Get() *LoggingS3Additional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingS3Additional) Set(val *LoggingS3Additional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingS3Additional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingS3Additional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingS3Additional returns a pointer to a new instance of NullableLoggingS3Additional.
func NewNullableLoggingS3Additional(val *LoggingS3Additional) *NullableLoggingS3Additional {
	return &NullableLoggingS3Additional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingS3Additional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingS3Additional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
