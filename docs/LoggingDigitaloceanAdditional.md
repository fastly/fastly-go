# LoggingDigitaloceanAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | Pointer to **string** | The name of the DigitalOcean Space. | [optional] 
**AccessKey** | Pointer to **string** | Your DigitalOcean Spaces account access key. | [optional] 
**SecretKey** | Pointer to **string** | Your DigitalOcean Spaces account secret key. | [optional] 
**Domain** | Pointer to **string** | The domain of the DigitalOcean Spaces endpoint. | [optional] [default to "nyc3.digitaloceanspaces.com"]
**Path** | Pointer to **NullableString** | The path to upload logs to. | [optional] [default to "null"]
**PublicKey** | Pointer to **NullableString** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [optional] [default to "null"]

## Methods

### NewLoggingDigitaloceanAdditional

`func NewLoggingDigitaloceanAdditional() *LoggingDigitaloceanAdditional`

NewLoggingDigitaloceanAdditional instantiates a new LoggingDigitaloceanAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingDigitaloceanAdditionalWithDefaults

`func NewLoggingDigitaloceanAdditionalWithDefaults() *LoggingDigitaloceanAdditional`

NewLoggingDigitaloceanAdditionalWithDefaults instantiates a new LoggingDigitaloceanAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *LoggingDigitaloceanAdditional) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *LoggingDigitaloceanAdditional) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *LoggingDigitaloceanAdditional) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *LoggingDigitaloceanAdditional) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetAccessKey

`func (o *LoggingDigitaloceanAdditional) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *LoggingDigitaloceanAdditional) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *LoggingDigitaloceanAdditional) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *LoggingDigitaloceanAdditional) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *LoggingDigitaloceanAdditional) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *LoggingDigitaloceanAdditional) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *LoggingDigitaloceanAdditional) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *LoggingDigitaloceanAdditional) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetDomain

`func (o *LoggingDigitaloceanAdditional) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *LoggingDigitaloceanAdditional) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *LoggingDigitaloceanAdditional) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *LoggingDigitaloceanAdditional) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPath

`func (o *LoggingDigitaloceanAdditional) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LoggingDigitaloceanAdditional) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LoggingDigitaloceanAdditional) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LoggingDigitaloceanAdditional) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *LoggingDigitaloceanAdditional) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *LoggingDigitaloceanAdditional) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPublicKey

`func (o *LoggingDigitaloceanAdditional) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *LoggingDigitaloceanAdditional) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *LoggingDigitaloceanAdditional) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *LoggingDigitaloceanAdditional) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *LoggingDigitaloceanAdditional) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *LoggingDigitaloceanAdditional) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
