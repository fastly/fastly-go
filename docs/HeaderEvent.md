# HeaderEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] 
**HeaderName** | Pointer to **string** |  | [optional] 
**OldValue** | Pointer to **string** |  | [optional] 
**NewValue** | Pointer to **string** |  | [optional] 
**ChangedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewHeaderEvent

`func NewHeaderEvent() *HeaderEvent`

NewHeaderEvent instantiates a new HeaderEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeaderEventWithDefaults

`func NewHeaderEventWithDefaults() *HeaderEvent`

NewHeaderEventWithDefaults instantiates a new HeaderEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *HeaderEvent) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *HeaderEvent) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *HeaderEvent) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *HeaderEvent) HasID() bool`

HasID returns a boolean if a field has been set.

### GetHeaderName

`func (o *HeaderEvent) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *HeaderEvent) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *HeaderEvent) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.

### HasHeaderName

`func (o *HeaderEvent) HasHeaderName() bool`

HasHeaderName returns a boolean if a field has been set.

### GetOldValue

`func (o *HeaderEvent) GetOldValue() string`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *HeaderEvent) GetOldValueOk() (*string, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *HeaderEvent) SetOldValue(v string)`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *HeaderEvent) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.

### GetNewValue

`func (o *HeaderEvent) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *HeaderEvent) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *HeaderEvent) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *HeaderEvent) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### GetChangedAt

`func (o *HeaderEvent) GetChangedAt() time.Time`

GetChangedAt returns the ChangedAt field if non-nil, zero value otherwise.

### GetChangedAtOk

`func (o *HeaderEvent) GetChangedAtOk() (*time.Time, bool)`

GetChangedAtOk returns a tuple with the ChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAt

`func (o *HeaderEvent) SetChangedAt(v time.Time)`

SetChangedAt sets ChangedAt field to given value.

### HasChangedAt

`func (o *HeaderEvent) HasChangedAt() bool`

HasChangedAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
