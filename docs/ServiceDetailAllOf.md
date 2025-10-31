# ServiceDetailAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveVersion** | Pointer to [**NullableServiceVersionDetailOrNull**](ServiceVersionDetailOrNull.md) |  | [optional] 
**Version** | Pointer to [**ServiceVersionDetail**](ServiceVersionDetail.md) |  | [optional] 

## Methods

### NewServiceDetailAllOf

`func NewServiceDetailAllOf() *ServiceDetailAllOf`

NewServiceDetailAllOf instantiates a new ServiceDetailAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDetailAllOfWithDefaults

`func NewServiceDetailAllOfWithDefaults() *ServiceDetailAllOf`

NewServiceDetailAllOfWithDefaults instantiates a new ServiceDetailAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveVersion

`func (o *ServiceDetailAllOf) GetActiveVersion() ServiceVersionDetailOrNull`

GetActiveVersion returns the ActiveVersion field if non-nil, zero value otherwise.

### GetActiveVersionOk

`func (o *ServiceDetailAllOf) GetActiveVersionOk() (*ServiceVersionDetailOrNull, bool)`

GetActiveVersionOk returns a tuple with the ActiveVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveVersion

`func (o *ServiceDetailAllOf) SetActiveVersion(v ServiceVersionDetailOrNull)`

SetActiveVersion sets ActiveVersion field to given value.

### HasActiveVersion

`func (o *ServiceDetailAllOf) HasActiveVersion() bool`

HasActiveVersion returns a boolean if a field has been set.

### SetActiveVersionNil

`func (o *ServiceDetailAllOf) SetActiveVersionNil(b bool)`

 SetActiveVersionNil sets the value for ActiveVersion to be an explicit nil

### UnsetActiveVersion
`func (o *ServiceDetailAllOf) UnsetActiveVersion()`

UnsetActiveVersion ensures that no value is present for ActiveVersion, not even an explicit nil
### GetVersion

`func (o *ServiceDetailAllOf) GetVersion() ServiceVersionDetail`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceDetailAllOf) GetVersionOk() (*ServiceVersionDetail, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceDetailAllOf) SetVersion(v ServiceVersionDetail)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServiceDetailAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


