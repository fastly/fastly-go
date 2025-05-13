# DdosProtectionEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** | Unique ID of the event. | [optional] 
**Name** | Pointer to **string** | A human-readable name for the event. | [optional] 
**CustomerID** | Pointer to **string** | Alphanumeric string identifying the customer. | [optional] 
**ServiceID** | Pointer to **string** | Alphanumeric string identifying the service. | [optional] 
**StartedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 
**EndedAt** | Pointer to **NullableTime** | Date and time in ISO 8601 format. | [optional] [readonly] 

## Methods

### NewDdosProtectionEventAllOf

`func NewDdosProtectionEventAllOf() *DdosProtectionEventAllOf`

NewDdosProtectionEventAllOf instantiates a new DdosProtectionEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdosProtectionEventAllOfWithDefaults

`func NewDdosProtectionEventAllOfWithDefaults() *DdosProtectionEventAllOf`

NewDdosProtectionEventAllOfWithDefaults instantiates a new DdosProtectionEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *DdosProtectionEventAllOf) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *DdosProtectionEventAllOf) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *DdosProtectionEventAllOf) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *DdosProtectionEventAllOf) HasID() bool`

HasID returns a boolean if a field has been set.

### GetName

`func (o *DdosProtectionEventAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DdosProtectionEventAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DdosProtectionEventAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DdosProtectionEventAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCustomerID

`func (o *DdosProtectionEventAllOf) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *DdosProtectionEventAllOf) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *DdosProtectionEventAllOf) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *DdosProtectionEventAllOf) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetServiceID

`func (o *DdosProtectionEventAllOf) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *DdosProtectionEventAllOf) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *DdosProtectionEventAllOf) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *DdosProtectionEventAllOf) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetStartedAt

`func (o *DdosProtectionEventAllOf) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DdosProtectionEventAllOf) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DdosProtectionEventAllOf) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DdosProtectionEventAllOf) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *DdosProtectionEventAllOf) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *DdosProtectionEventAllOf) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEndedAt

`func (o *DdosProtectionEventAllOf) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *DdosProtectionEventAllOf) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *DdosProtectionEventAllOf) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *DdosProtectionEventAllOf) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *DdosProtectionEventAllOf) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *DdosProtectionEventAllOf) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
