# Director

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backends** | Pointer to [**[]Backend**](Backend.md) | List of backends associated to a director. | [optional] 
**Capacity** | Pointer to **int32** | Unused. | [optional] 
**Comment** | Pointer to **NullableString** | A freeform descriptive note. | [optional] 
**Name** | Pointer to **string** | Name for the Director. | [optional] 
**Quorum** | Pointer to **int32** | The percentage of capacity that needs to be up for a director to be considered up. `0` to `100`. | [optional] [default to 75]
**Shield** | Pointer to **NullableString** | Selected POP to serve as a shield for the backends. Defaults to `null` meaning no origin shielding if not set. Refer to the [POPs API endpoint](https://www.fastly.com/documentation/reference/api/utils/pops/) to get a list of available POPs used for shielding. | [optional] [default to "null"]
**Type** | Pointer to **int32** | What type of load balance group to use. | [optional] [default to 1]
**Retries** | Pointer to **int32** | How many backends to search if it fails. | [optional] [default to 5]

## Methods

### NewDirector

`func NewDirector() *Director`

NewDirector instantiates a new Director object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectorWithDefaults

`func NewDirectorWithDefaults() *Director`

NewDirectorWithDefaults instantiates a new Director object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackends

`func (o *Director) GetBackends() []Backend`

GetBackends returns the Backends field if non-nil, zero value otherwise.

### GetBackendsOk

`func (o *Director) GetBackendsOk() (*[]Backend, bool)`

GetBackendsOk returns a tuple with the Backends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackends

`func (o *Director) SetBackends(v []Backend)`

SetBackends sets Backends field to given value.

### HasBackends

`func (o *Director) HasBackends() bool`

HasBackends returns a boolean if a field has been set.

### GetCapacity

`func (o *Director) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *Director) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *Director) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *Director) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetComment

`func (o *Director) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Director) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Director) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Director) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *Director) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *Director) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetName

`func (o *Director) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Director) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Director) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Director) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuorum

`func (o *Director) GetQuorum() int32`

GetQuorum returns the Quorum field if non-nil, zero value otherwise.

### GetQuorumOk

`func (o *Director) GetQuorumOk() (*int32, bool)`

GetQuorumOk returns a tuple with the Quorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorum

`func (o *Director) SetQuorum(v int32)`

SetQuorum sets Quorum field to given value.

### HasQuorum

`func (o *Director) HasQuorum() bool`

HasQuorum returns a boolean if a field has been set.

### GetShield

`func (o *Director) GetShield() string`

GetShield returns the Shield field if non-nil, zero value otherwise.

### GetShieldOk

`func (o *Director) GetShieldOk() (*string, bool)`

GetShieldOk returns a tuple with the Shield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShield

`func (o *Director) SetShield(v string)`

SetShield sets Shield field to given value.

### HasShield

`func (o *Director) HasShield() bool`

HasShield returns a boolean if a field has been set.

### SetShieldNil

`func (o *Director) SetShieldNil(b bool)`

 SetShieldNil sets the value for Shield to be an explicit nil

### UnsetShield
`func (o *Director) UnsetShield()`

UnsetShield ensures that no value is present for Shield, not even an explicit nil
### GetType

`func (o *Director) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Director) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Director) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *Director) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRetries

`func (o *Director) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *Director) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *Director) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *Director) HasRetries() bool`

HasRetries returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


