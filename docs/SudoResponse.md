# SudoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiryTime** | Pointer to **time.Time** | A UTC time-stamp of when sudo access will expire. If blank, sudo access expires five minutes after the request. | [optional] 

## Methods

### NewSudoResponse

`func NewSudoResponse() *SudoResponse`

NewSudoResponse instantiates a new SudoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSudoResponseWithDefaults

`func NewSudoResponseWithDefaults() *SudoResponse`

NewSudoResponseWithDefaults instantiates a new SudoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiryTime

`func (o *SudoResponse) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *SudoResponse) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *SudoResponse) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *SudoResponse) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
