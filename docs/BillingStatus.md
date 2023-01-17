# BillingStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | What the current status of this invoice can be. | [optional] 
**SentAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### NewBillingStatus

`func NewBillingStatus() *BillingStatus`

NewBillingStatus instantiates a new BillingStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingStatusWithDefaults

`func NewBillingStatusWithDefaults() *BillingStatus`

NewBillingStatusWithDefaults instantiates a new BillingStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BillingStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillingStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSentAt

`func (o *BillingStatus) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *BillingStatus) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *BillingStatus) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *BillingStatus) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
