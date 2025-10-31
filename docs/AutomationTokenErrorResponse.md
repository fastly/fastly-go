# AutomationTokenErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **[]interface{}** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewAutomationTokenErrorResponse

`func NewAutomationTokenErrorResponse() *AutomationTokenErrorResponse`

NewAutomationTokenErrorResponse instantiates a new AutomationTokenErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomationTokenErrorResponseWithDefaults

`func NewAutomationTokenErrorResponseWithDefaults() *AutomationTokenErrorResponse`

NewAutomationTokenErrorResponseWithDefaults instantiates a new AutomationTokenErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *AutomationTokenErrorResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *AutomationTokenErrorResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *AutomationTokenErrorResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *AutomationTokenErrorResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetErrors

`func (o *AutomationTokenErrorResponse) GetErrors() []interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AutomationTokenErrorResponse) GetErrorsOk() (*[]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AutomationTokenErrorResponse) SetErrors(v []interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AutomationTokenErrorResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetStatus

`func (o *AutomationTokenErrorResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutomationTokenErrorResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutomationTokenErrorResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AutomationTokenErrorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *AutomationTokenErrorResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AutomationTokenErrorResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AutomationTokenErrorResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AutomationTokenErrorResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


