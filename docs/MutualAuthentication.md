# MutualAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**MutualAuthenticationData**](MutualAuthenticationData.md) |  | [optional] 

## Methods

### NewMutualAuthentication

`func NewMutualAuthentication() *MutualAuthentication`

NewMutualAuthentication instantiates a new MutualAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualAuthenticationWithDefaults

`func NewMutualAuthenticationWithDefaults() *MutualAuthentication`

NewMutualAuthenticationWithDefaults instantiates a new MutualAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MutualAuthentication) GetData() MutualAuthenticationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MutualAuthentication) GetDataOk() (*MutualAuthenticationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MutualAuthentication) SetData(v MutualAuthenticationData)`

SetData sets Data field to given value.

### HasData

`func (o *MutualAuthentication) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


