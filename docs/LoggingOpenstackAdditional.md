# LoggingOpenstackAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | Your OpenStack account access key. | [optional] 
**BucketName** | Pointer to **string** | The name of your OpenStack container. | [optional] 
**Path** | Pointer to **NullableString** | The path to upload logs to. | [optional] [default to "null"]
**PublicKey** | Pointer to **NullableString** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [optional] [default to "null"]
**URL** | Pointer to **string** | Your OpenStack auth url. | [optional] 
**User** | Pointer to **string** | The username for your OpenStack account. | [optional] 

## Methods

### NewLoggingOpenstackAdditional

`func NewLoggingOpenstackAdditional() *LoggingOpenstackAdditional`

NewLoggingOpenstackAdditional instantiates a new LoggingOpenstackAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingOpenstackAdditionalWithDefaults

`func NewLoggingOpenstackAdditionalWithDefaults() *LoggingOpenstackAdditional`

NewLoggingOpenstackAdditionalWithDefaults instantiates a new LoggingOpenstackAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *LoggingOpenstackAdditional) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *LoggingOpenstackAdditional) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *LoggingOpenstackAdditional) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *LoggingOpenstackAdditional) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetBucketName

`func (o *LoggingOpenstackAdditional) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *LoggingOpenstackAdditional) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *LoggingOpenstackAdditional) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *LoggingOpenstackAdditional) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetPath

`func (o *LoggingOpenstackAdditional) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LoggingOpenstackAdditional) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LoggingOpenstackAdditional) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LoggingOpenstackAdditional) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *LoggingOpenstackAdditional) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *LoggingOpenstackAdditional) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPublicKey

`func (o *LoggingOpenstackAdditional) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *LoggingOpenstackAdditional) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *LoggingOpenstackAdditional) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *LoggingOpenstackAdditional) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *LoggingOpenstackAdditional) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *LoggingOpenstackAdditional) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetURL

`func (o *LoggingOpenstackAdditional) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *LoggingOpenstackAdditional) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *LoggingOpenstackAdditional) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *LoggingOpenstackAdditional) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetUser

`func (o *LoggingOpenstackAdditional) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoggingOpenstackAdditional) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoggingOpenstackAdditional) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LoggingOpenstackAdditional) HasUser() bool`

HasUser returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
