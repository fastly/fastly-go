# LoggingCloudfilesAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | Your Cloud Files account access key. | [optional] 
**BucketName** | Pointer to **string** | The name of your Cloud Files container. | [optional] 
**Path** | Pointer to **NullableString** | The path to upload logs to. | [optional] [default to "null"]
**Region** | Pointer to **NullableString** | The region to stream logs to. | [optional] 
**PublicKey** | Pointer to **NullableString** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [optional] [default to "null"]
**User** | Pointer to **string** | The username for your Cloud Files account. | [optional] 

## Methods

### NewLoggingCloudfilesAdditional

`func NewLoggingCloudfilesAdditional() *LoggingCloudfilesAdditional`

NewLoggingCloudfilesAdditional instantiates a new LoggingCloudfilesAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingCloudfilesAdditionalWithDefaults

`func NewLoggingCloudfilesAdditionalWithDefaults() *LoggingCloudfilesAdditional`

NewLoggingCloudfilesAdditionalWithDefaults instantiates a new LoggingCloudfilesAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *LoggingCloudfilesAdditional) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *LoggingCloudfilesAdditional) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *LoggingCloudfilesAdditional) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *LoggingCloudfilesAdditional) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetBucketName

`func (o *LoggingCloudfilesAdditional) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *LoggingCloudfilesAdditional) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *LoggingCloudfilesAdditional) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *LoggingCloudfilesAdditional) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetPath

`func (o *LoggingCloudfilesAdditional) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LoggingCloudfilesAdditional) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LoggingCloudfilesAdditional) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LoggingCloudfilesAdditional) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *LoggingCloudfilesAdditional) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *LoggingCloudfilesAdditional) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetRegion

`func (o *LoggingCloudfilesAdditional) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LoggingCloudfilesAdditional) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LoggingCloudfilesAdditional) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LoggingCloudfilesAdditional) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *LoggingCloudfilesAdditional) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *LoggingCloudfilesAdditional) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetPublicKey

`func (o *LoggingCloudfilesAdditional) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *LoggingCloudfilesAdditional) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *LoggingCloudfilesAdditional) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *LoggingCloudfilesAdditional) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *LoggingCloudfilesAdditional) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *LoggingCloudfilesAdditional) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetUser

`func (o *LoggingCloudfilesAdditional) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoggingCloudfilesAdditional) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoggingCloudfilesAdditional) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LoggingCloudfilesAdditional) HasUser() bool`

HasUser returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


