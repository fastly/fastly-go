# LoggingScalyrAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | The region that log data will be sent to. | [optional] [default to "US"]
**Token** | Pointer to **string** | The token to use for authentication. | [optional] 
**ProjectID** | Pointer to **string** | The name of the logfile within Scalyr. | [optional] [default to "logplex"]

## Methods

### NewLoggingScalyrAdditional

`func NewLoggingScalyrAdditional() *LoggingScalyrAdditional`

NewLoggingScalyrAdditional instantiates a new LoggingScalyrAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingScalyrAdditionalWithDefaults

`func NewLoggingScalyrAdditionalWithDefaults() *LoggingScalyrAdditional`

NewLoggingScalyrAdditionalWithDefaults instantiates a new LoggingScalyrAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *LoggingScalyrAdditional) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LoggingScalyrAdditional) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LoggingScalyrAdditional) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LoggingScalyrAdditional) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetToken

`func (o *LoggingScalyrAdditional) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LoggingScalyrAdditional) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LoggingScalyrAdditional) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LoggingScalyrAdditional) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetProjectID

`func (o *LoggingScalyrAdditional) GetProjectID() string`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *LoggingScalyrAdditional) GetProjectIDOk() (*string, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *LoggingScalyrAdditional) SetProjectID(v string)`

SetProjectID sets ProjectID field to given value.

### HasProjectID

`func (o *LoggingScalyrAdditional) HasProjectID() bool`

HasProjectID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
