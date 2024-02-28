# ServiceResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] [readonly] 
**PublishKey** | Pointer to **string** | Unused at this time. | [optional] 
**Paused** | Pointer to **bool** | Whether the service is paused. Services are paused due to a lack of traffic for an extended period of time. Services are resumed either when a draft version is activated or a locked version is cloned and reactivated. | [optional] 
**Versions** | Pointer to [**[]SchemasVersionResponse**](SchemasVersionResponse.md) | A list of [versions](https://www.fastly.com/documentation/reference/api/services/version/) associated with the service. | [optional] 

## Methods

### NewServiceResponseAllOf

`func NewServiceResponseAllOf() *ServiceResponseAllOf`

NewServiceResponseAllOf instantiates a new ServiceResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceResponseAllOfWithDefaults

`func NewServiceResponseAllOfWithDefaults() *ServiceResponseAllOf`

NewServiceResponseAllOfWithDefaults instantiates a new ServiceResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *ServiceResponseAllOf) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ServiceResponseAllOf) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ServiceResponseAllOf) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *ServiceResponseAllOf) HasID() bool`

HasID returns a boolean if a field has been set.

### GetPublishKey

`func (o *ServiceResponseAllOf) GetPublishKey() string`

GetPublishKey returns the PublishKey field if non-nil, zero value otherwise.

### GetPublishKeyOk

`func (o *ServiceResponseAllOf) GetPublishKeyOk() (*string, bool)`

GetPublishKeyOk returns a tuple with the PublishKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishKey

`func (o *ServiceResponseAllOf) SetPublishKey(v string)`

SetPublishKey sets PublishKey field to given value.

### HasPublishKey

`func (o *ServiceResponseAllOf) HasPublishKey() bool`

HasPublishKey returns a boolean if a field has been set.

### GetPaused

`func (o *ServiceResponseAllOf) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *ServiceResponseAllOf) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *ServiceResponseAllOf) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *ServiceResponseAllOf) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetVersions

`func (o *ServiceResponseAllOf) GetVersions() []SchemasVersionResponse`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ServiceResponseAllOf) GetVersionsOk() (*[]SchemasVersionResponse, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ServiceResponseAllOf) SetVersions(v []SchemasVersionResponse)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ServiceResponseAllOf) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
