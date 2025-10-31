# ServiceAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ServiceAuthorizationData**](ServiceAuthorizationData.md) |  | [optional] 

## Methods

### NewServiceAuthorization

`func NewServiceAuthorization() *ServiceAuthorization`

NewServiceAuthorization instantiates a new ServiceAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAuthorizationWithDefaults

`func NewServiceAuthorizationWithDefaults() *ServiceAuthorization`

NewServiceAuthorizationWithDefaults instantiates a new ServiceAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServiceAuthorization) GetData() ServiceAuthorizationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceAuthorization) GetDataOk() (*ServiceAuthorizationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceAuthorization) SetData(v ServiceAuthorizationData)`

SetData sets Data field to given value.

### HasData

`func (o *ServiceAuthorization) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


