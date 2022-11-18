# Pop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Coordinates** | Pointer to [**PopCoordinates**](PopCoordinates.md) |  | [optional] 
**Shield** | Pointer to **string** |  | [optional] 

## Methods

### NewPop

`func NewPop() *Pop`

NewPop instantiates a new Pop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPopWithDefaults

`func NewPopWithDefaults() *Pop`

NewPopWithDefaults instantiates a new Pop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Pop) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Pop) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Pop) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Pop) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *Pop) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pop) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pop) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Pop) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGroup

`func (o *Pop) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Pop) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Pop) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Pop) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCoordinates

`func (o *Pop) GetCoordinates() PopCoordinates`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *Pop) GetCoordinatesOk() (*PopCoordinates, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *Pop) SetCoordinates(v PopCoordinates)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *Pop) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### GetShield

`func (o *Pop) GetShield() string`

GetShield returns the Shield field if non-nil, zero value otherwise.

### GetShieldOk

`func (o *Pop) GetShieldOk() (*string, bool)`

GetShieldOk returns a tuple with the Shield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShield

`func (o *Pop) SetShield(v string)`

SetShield sets Shield field to given value.

### HasShield

`func (o *Pop) HasShield() bool`

HasShield returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
