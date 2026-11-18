# LoggingFtpAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | An hostname or IPv4 address. | [optional] 
**Hostname** | Pointer to **string** | Hostname used. | [optional] 
**Ipv4** | Pointer to **string** | IPv4 address of the host. | [optional] 
**Password** | Pointer to **string** | The password for the server. For anonymous use an email address. | [optional] 
**Path** | Pointer to **string** | The path to upload log files to. If the path ends in `/` then it is treated as a directory. | [optional] 
**PublicKey** | Pointer to **NullableString** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [optional] [default to "null"]
**User** | Pointer to **string** | The username for the server. Can be anonymous. | [optional] 

## Methods

### NewLoggingFtpAdditional

`func NewLoggingFtpAdditional() *LoggingFtpAdditional`

NewLoggingFtpAdditional instantiates a new LoggingFtpAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingFtpAdditionalWithDefaults

`func NewLoggingFtpAdditionalWithDefaults() *LoggingFtpAdditional`

NewLoggingFtpAdditionalWithDefaults instantiates a new LoggingFtpAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *LoggingFtpAdditional) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LoggingFtpAdditional) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LoggingFtpAdditional) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *LoggingFtpAdditional) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetHostname

`func (o *LoggingFtpAdditional) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *LoggingFtpAdditional) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *LoggingFtpAdditional) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *LoggingFtpAdditional) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIpv4

`func (o *LoggingFtpAdditional) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *LoggingFtpAdditional) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *LoggingFtpAdditional) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *LoggingFtpAdditional) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetPassword

`func (o *LoggingFtpAdditional) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoggingFtpAdditional) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoggingFtpAdditional) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LoggingFtpAdditional) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPath

`func (o *LoggingFtpAdditional) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LoggingFtpAdditional) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LoggingFtpAdditional) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LoggingFtpAdditional) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPublicKey

`func (o *LoggingFtpAdditional) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *LoggingFtpAdditional) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *LoggingFtpAdditional) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *LoggingFtpAdditional) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *LoggingFtpAdditional) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *LoggingFtpAdditional) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetUser

`func (o *LoggingFtpAdditional) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoggingFtpAdditional) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoggingFtpAdditional) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LoggingFtpAdditional) HasUser() bool`

HasUser returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
