# BackendResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locked** | Pointer to **bool** | Indicates whether the version of the service this backend is attached to accepts edits. | [optional] [readonly] 

## Methods

### NewBackendResponseAllOf

`func NewBackendResponseAllOf() *BackendResponseAllOf`

NewBackendResponseAllOf instantiates a new BackendResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackendResponseAllOfWithDefaults

`func NewBackendResponseAllOfWithDefaults() *BackendResponseAllOf`

NewBackendResponseAllOfWithDefaults instantiates a new BackendResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocked

`func (o *BackendResponseAllOf) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *BackendResponseAllOf) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *BackendResponseAllOf) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *BackendResponseAllOf) HasLocked() bool`

HasLocked returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
