# GenericTokenError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** |  | [optional] 

## Methods

### NewGenericTokenError

`func NewGenericTokenError() *GenericTokenError`

NewGenericTokenError instantiates a new GenericTokenError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericTokenErrorWithDefaults

`func NewGenericTokenErrorWithDefaults() *GenericTokenError`

NewGenericTokenErrorWithDefaults instantiates a new GenericTokenError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *GenericTokenError) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GenericTokenError) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GenericTokenError) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GenericTokenError) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
