# LoggingDigitaloceanAllOf

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

### NewLoggingDigitaloceanAllOf

`func NewLoggingDigitaloceanAllOf() *LoggingDigitaloceanAllOf`

NewLoggingDigitaloceanAllOf instantiates a new LoggingDigitaloceanAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingDigitaloceanAllOfWithDefaults

`func NewLoggingDigitaloceanAllOfWithDefaults() *LoggingDigitaloceanAllOf`

NewLoggingDigitaloceanAllOfWithDefaults instantiates a new LoggingDigitaloceanAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *LoggingDigitaloceanAllOf) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *LoggingDigitaloceanAllOf) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *LoggingDigitaloceanAllOf) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *LoggingDigitaloceanAllOf) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetAccessKey

`func (o *LoggingDigitaloceanAllOf) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *LoggingDigitaloceanAllOf) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *LoggingDigitaloceanAllOf) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *LoggingDigitaloceanAllOf) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *LoggingDigitaloceanAllOf) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *LoggingDigitaloceanAllOf) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *LoggingDigitaloceanAllOf) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *LoggingDigitaloceanAllOf) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetDomain

`func (o *LoggingDigitaloceanAllOf) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *LoggingDigitaloceanAllOf) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *LoggingDigitaloceanAllOf) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *LoggingDigitaloceanAllOf) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPath

`func (o *LoggingDigitaloceanAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LoggingDigitaloceanAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LoggingDigitaloceanAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LoggingDigitaloceanAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *LoggingDigitaloceanAllOf) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *LoggingDigitaloceanAllOf) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPublicKey

`func (o *LoggingDigitaloceanAllOf) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *LoggingDigitaloceanAllOf) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *LoggingDigitaloceanAllOf) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *LoggingDigitaloceanAllOf) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *LoggingDigitaloceanAllOf) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *LoggingDigitaloceanAllOf) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
