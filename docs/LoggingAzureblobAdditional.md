# LoggingAzureblobAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **NullableString** | The path to upload logs to. | [optional] [default to "null"]
**AccountName** | Pointer to **string** | The unique Azure Blob Storage namespace in which your data objects are stored. Required. | [optional] 
**Container** | Pointer to **string** | The name of the Azure Blob Storage container in which to store logs. Required. | [optional] 
**SasToken** | Pointer to **string** | The Azure shared access signature providing write access to the blob service objects. Be sure to update your token before it expires or the logging functionality will not work. Required. | [optional] 
**PublicKey** | Pointer to **NullableString** | A PGP public key that Fastly will use to encrypt your log files before writing them to disk. | [optional] [default to "null"]
**FileMaxBytes** | Pointer to **int32** | The maximum number of bytes for each uploaded file. A value of 0 can be used to indicate there is no limit on the size of uploaded files, otherwise the minimum value is 1048576 bytes (1 MiB). Note that Microsoft Azure Storage has [block size limits](https://learn.microsoft.com/en-us/rest/api/storageservices/put-block?tabs&#x3D;microsoft-entra-id#remarks). | [optional] 

## Methods

### NewLoggingAzureblobAdditional

`func NewLoggingAzureblobAdditional() *LoggingAzureblobAdditional`

NewLoggingAzureblobAdditional instantiates a new LoggingAzureblobAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingAzureblobAdditionalWithDefaults

`func NewLoggingAzureblobAdditionalWithDefaults() *LoggingAzureblobAdditional`

NewLoggingAzureblobAdditionalWithDefaults instantiates a new LoggingAzureblobAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *LoggingAzureblobAdditional) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LoggingAzureblobAdditional) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LoggingAzureblobAdditional) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LoggingAzureblobAdditional) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *LoggingAzureblobAdditional) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *LoggingAzureblobAdditional) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetAccountName

`func (o *LoggingAzureblobAdditional) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *LoggingAzureblobAdditional) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *LoggingAzureblobAdditional) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *LoggingAzureblobAdditional) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetContainer

`func (o *LoggingAzureblobAdditional) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *LoggingAzureblobAdditional) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *LoggingAzureblobAdditional) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *LoggingAzureblobAdditional) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetSasToken

`func (o *LoggingAzureblobAdditional) GetSasToken() string`

GetSasToken returns the SasToken field if non-nil, zero value otherwise.

### GetSasTokenOk

`func (o *LoggingAzureblobAdditional) GetSasTokenOk() (*string, bool)`

GetSasTokenOk returns a tuple with the SasToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasToken

`func (o *LoggingAzureblobAdditional) SetSasToken(v string)`

SetSasToken sets SasToken field to given value.

### HasSasToken

`func (o *LoggingAzureblobAdditional) HasSasToken() bool`

HasSasToken returns a boolean if a field has been set.

### GetPublicKey

`func (o *LoggingAzureblobAdditional) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *LoggingAzureblobAdditional) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *LoggingAzureblobAdditional) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *LoggingAzureblobAdditional) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *LoggingAzureblobAdditional) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *LoggingAzureblobAdditional) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetFileMaxBytes

`func (o *LoggingAzureblobAdditional) GetFileMaxBytes() int32`

GetFileMaxBytes returns the FileMaxBytes field if non-nil, zero value otherwise.

### GetFileMaxBytesOk

`func (o *LoggingAzureblobAdditional) GetFileMaxBytesOk() (*int32, bool)`

GetFileMaxBytesOk returns a tuple with the FileMaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileMaxBytes

`func (o *LoggingAzureblobAdditional) SetFileMaxBytes(v int32)`

SetFileMaxBytes sets FileMaxBytes field to given value.

### HasFileMaxBytes

`func (o *LoggingAzureblobAdditional) HasFileMaxBytes() bool`

HasFileMaxBytes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


