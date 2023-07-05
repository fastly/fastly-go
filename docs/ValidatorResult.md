# ValidatorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ValidatorResultData**](ValidatorResultData.md) |  | [optional] 

## Methods

### NewValidatorResult

`func NewValidatorResult() *ValidatorResult`

NewValidatorResult instantiates a new ValidatorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatorResultWithDefaults

`func NewValidatorResultWithDefaults() *ValidatorResult`

NewValidatorResultWithDefaults instantiates a new ValidatorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ValidatorResult) GetData() ValidatorResultData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ValidatorResult) GetDataOk() (*ValidatorResultData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ValidatorResult) SetData(v ValidatorResultData)`

SetData sets Data field to given value.

### HasData

`func (o *ValidatorResult) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
