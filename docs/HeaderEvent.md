# HeaderEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
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

### GetId

`func (o *HeaderEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HeaderEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HeaderEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HeaderEvent) HasId() bool`

HasId returns a boolean if a field has been set.

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


