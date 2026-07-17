# SecurityHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**ObservedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSecurityHeader

`func NewSecurityHeader() *SecurityHeader`

NewSecurityHeader instantiates a new SecurityHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityHeaderWithDefaults

`func NewSecurityHeaderWithDefaults() *SecurityHeader`

NewSecurityHeaderWithDefaults instantiates a new SecurityHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecurityHeader) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityHeader) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityHeader) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityHeader) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *SecurityHeader) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SecurityHeader) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SecurityHeader) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SecurityHeader) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetObservedAt

`func (o *SecurityHeader) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *SecurityHeader) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *SecurityHeader) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.

### HasObservedAt

`func (o *SecurityHeader) HasObservedAt() bool`

HasObservedAt returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


