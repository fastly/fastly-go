# LoggingSftpAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | The password for the server. If both `password` and `secret_key` are passed, `secret_key` will be used in preference. | [optional] 
**Path** | Pointer to **NullableString** | The path to upload logs to. | [optional] [default to "null"]
**PublicKey** | Pointer to **NullableString** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [optional] [default to "null"]
**SecretKey** | Pointer to **NullableString** | The SSH private key for the server. If both `password` and `secret_key` are passed, `secret_key` will be used in preference. | [optional] [default to "null"]
**SshKnownHosts** | Pointer to **string** | A list of host keys for all hosts we can connect to over SFTP. | [optional] 
**User** | Pointer to **string** | The username for the server. | [optional] 

## Methods

### NewLoggingSftpAdditional

`func NewLoggingSftpAdditional() *LoggingSftpAdditional`

NewLoggingSftpAdditional instantiates a new LoggingSftpAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingSftpAdditionalWithDefaults

`func NewLoggingSftpAdditionalWithDefaults() *LoggingSftpAdditional`

NewLoggingSftpAdditionalWithDefaults instantiates a new LoggingSftpAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *LoggingSftpAdditional) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoggingSftpAdditional) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoggingSftpAdditional) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LoggingSftpAdditional) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPath

`func (o *LoggingSftpAdditional) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LoggingSftpAdditional) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LoggingSftpAdditional) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LoggingSftpAdditional) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *LoggingSftpAdditional) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *LoggingSftpAdditional) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPublicKey

`func (o *LoggingSftpAdditional) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *LoggingSftpAdditional) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *LoggingSftpAdditional) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *LoggingSftpAdditional) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *LoggingSftpAdditional) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *LoggingSftpAdditional) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetSecretKey

`func (o *LoggingSftpAdditional) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *LoggingSftpAdditional) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *LoggingSftpAdditional) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *LoggingSftpAdditional) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### SetSecretKeyNil

`func (o *LoggingSftpAdditional) SetSecretKeyNil(b bool)`

 SetSecretKeyNil sets the value for SecretKey to be an explicit nil

### UnsetSecretKey
`func (o *LoggingSftpAdditional) UnsetSecretKey()`

UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
### GetSshKnownHosts

`func (o *LoggingSftpAdditional) GetSshKnownHosts() string`

GetSshKnownHosts returns the SshKnownHosts field if non-nil, zero value otherwise.

### GetSshKnownHostsOk

`func (o *LoggingSftpAdditional) GetSshKnownHostsOk() (*string, bool)`

GetSshKnownHostsOk returns a tuple with the SshKnownHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKnownHosts

`func (o *LoggingSftpAdditional) SetSshKnownHosts(v string)`

SetSshKnownHosts sets SshKnownHosts field to given value.

### HasSshKnownHosts

`func (o *LoggingSftpAdditional) HasSshKnownHosts() bool`

HasSshKnownHosts returns a boolean if a field has been set.

### GetUser

`func (o *LoggingSftpAdditional) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoggingSftpAdditional) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoggingSftpAdditional) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LoggingSftpAdditional) HasUser() bool`

HasUser returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


