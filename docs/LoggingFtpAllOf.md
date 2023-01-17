# LoggingFtpAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | An hostname or IPv4 address. | [optional] 
**Hostname** | Pointer to **string** | Hostname used. | [optional] 
**Ipv4** | Pointer to **string** | IPv4 address of the host. | [optional] 
**Password** | Pointer to **string** | The password for the server. For anonymous use an email address. | [optional] 
**Path** | Pointer to **string** | The path to upload log files to. If the path ends in `/` then it is treated as a directory. | [optional] 
**Port** | Pointer to **int32** | The port number. | [optional] [default to 21]
**PublicKey** | Pointer to **NullableString** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [optional] [default to "null"]
**User** | Pointer to **string** | The username for the server. Can be anonymous. | [optional] 

## Methods

### NewLoggingFtpAllOf

`func NewLoggingFtpAllOf() *LoggingFtpAllOf`

NewLoggingFtpAllOf instantiates a new LoggingFtpAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingFtpAllOfWithDefaults

`func NewLoggingFtpAllOfWithDefaults() *LoggingFtpAllOf`

NewLoggingFtpAllOfWithDefaults instantiates a new LoggingFtpAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *LoggingFtpAllOf) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LoggingFtpAllOf) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LoggingFtpAllOf) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *LoggingFtpAllOf) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetHostname

`func (o *LoggingFtpAllOf) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *LoggingFtpAllOf) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *LoggingFtpAllOf) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *LoggingFtpAllOf) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIpv4

`func (o *LoggingFtpAllOf) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *LoggingFtpAllOf) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *LoggingFtpAllOf) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *LoggingFtpAllOf) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetPassword

`func (o *LoggingFtpAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoggingFtpAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoggingFtpAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LoggingFtpAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPath

`func (o *LoggingFtpAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LoggingFtpAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LoggingFtpAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LoggingFtpAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPort

`func (o *LoggingFtpAllOf) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoggingFtpAllOf) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoggingFtpAllOf) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *LoggingFtpAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPublicKey

`func (o *LoggingFtpAllOf) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *LoggingFtpAllOf) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *LoggingFtpAllOf) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *LoggingFtpAllOf) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *LoggingFtpAllOf) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *LoggingFtpAllOf) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetUser

`func (o *LoggingFtpAllOf) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoggingFtpAllOf) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoggingFtpAllOf) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LoggingFtpAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
