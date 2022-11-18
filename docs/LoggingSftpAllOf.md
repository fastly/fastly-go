# LoggingSftpAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | The password for the server. If both `password` and `secret_key` are passed, `secret_key` will be used in preference. | [optional] 
**Path** | Pointer to **NullableString** | The path to upload logs to. | [optional] [default to "null"]
**Port** | Pointer to **int32** | The port number. | [optional] [default to 22]
**PublicKey** | Pointer to **NullableString** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [optional] [default to "null"]
**SecretKey** | Pointer to **NullableString** | The SSH private key for the server. If both `password` and `secret_key` are passed, `secret_key` will be used in preference. | [optional] [default to "null"]
**SSHKnownHosts** | Pointer to **string** | A list of host keys for all hosts we can connect to over SFTP. | [optional] 
**User** | Pointer to **string** | The username for the server. | [optional] 

## Methods

### NewLoggingSftpAllOf

`func NewLoggingSftpAllOf() *LoggingSftpAllOf`

NewLoggingSftpAllOf instantiates a new LoggingSftpAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingSftpAllOfWithDefaults

`func NewLoggingSftpAllOfWithDefaults() *LoggingSftpAllOf`

NewLoggingSftpAllOfWithDefaults instantiates a new LoggingSftpAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *LoggingSftpAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoggingSftpAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoggingSftpAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LoggingSftpAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPath

`func (o *LoggingSftpAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LoggingSftpAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LoggingSftpAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LoggingSftpAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *LoggingSftpAllOf) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *LoggingSftpAllOf) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPort

`func (o *LoggingSftpAllOf) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoggingSftpAllOf) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoggingSftpAllOf) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *LoggingSftpAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPublicKey

`func (o *LoggingSftpAllOf) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *LoggingSftpAllOf) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *LoggingSftpAllOf) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *LoggingSftpAllOf) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *LoggingSftpAllOf) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *LoggingSftpAllOf) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetSecretKey

`func (o *LoggingSftpAllOf) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *LoggingSftpAllOf) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *LoggingSftpAllOf) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *LoggingSftpAllOf) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### SetSecretKeyNil

`func (o *LoggingSftpAllOf) SetSecretKeyNil(b bool)`

 SetSecretKeyNil sets the value for SecretKey to be an explicit nil

### UnsetSecretKey
`func (o *LoggingSftpAllOf) UnsetSecretKey()`

UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
### GetSSHKnownHosts

`func (o *LoggingSftpAllOf) GetSSHKnownHosts() string`

GetSSHKnownHosts returns the SSHKnownHosts field if non-nil, zero value otherwise.

### GetSSHKnownHostsOk

`func (o *LoggingSftpAllOf) GetSSHKnownHostsOk() (*string, bool)`

GetSSHKnownHostsOk returns a tuple with the SSHKnownHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSHKnownHosts

`func (o *LoggingSftpAllOf) SetSSHKnownHosts(v string)`

SetSSHKnownHosts sets SSHKnownHosts field to given value.

### HasSSHKnownHosts

`func (o *LoggingSftpAllOf) HasSSHKnownHosts() bool`

HasSSHKnownHosts returns a boolean if a field has been set.

### GetUser

`func (o *LoggingSftpAllOf) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoggingSftpAllOf) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoggingSftpAllOf) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LoggingSftpAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
