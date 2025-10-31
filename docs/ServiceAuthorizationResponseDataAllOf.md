# ServiceAuthorizationResponseDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Attributes** | Pointer to [**Timestamps**](Timestamps.md) |  | [optional] 

## Methods

### NewServiceAuthorizationResponseDataAllOf

`func NewServiceAuthorizationResponseDataAllOf() *ServiceAuthorizationResponseDataAllOf`

NewServiceAuthorizationResponseDataAllOf instantiates a new ServiceAuthorizationResponseDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAuthorizationResponseDataAllOfWithDefaults

`func NewServiceAuthorizationResponseDataAllOfWithDefaults() *ServiceAuthorizationResponseDataAllOf`

NewServiceAuthorizationResponseDataAllOfWithDefaults instantiates a new ServiceAuthorizationResponseDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceAuthorizationResponseDataAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceAuthorizationResponseDataAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceAuthorizationResponseDataAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceAuthorizationResponseDataAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAttributes

`func (o *ServiceAuthorizationResponseDataAllOf) GetAttributes() Timestamps`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ServiceAuthorizationResponseDataAllOf) GetAttributesOk() (*Timestamps, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ServiceAuthorizationResponseDataAllOf) SetAttributes(v Timestamps)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ServiceAuthorizationResponseDataAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


