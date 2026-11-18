# HTTP3AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureRevision** | Pointer to **int32** | Revision number of the HTTP/3 feature implementation. Defaults to the most recent revision. | [optional] 

## Methods

### NewHTTP3AllOf

`func NewHTTP3AllOf() *HTTP3AllOf`

NewHTTP3AllOf instantiates a new HTTP3AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTP3AllOfWithDefaults

`func NewHTTP3AllOfWithDefaults() *HTTP3AllOf`

NewHTTP3AllOfWithDefaults instantiates a new HTTP3AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureRevision

`func (o *HTTP3AllOf) GetFeatureRevision() int32`

GetFeatureRevision returns the FeatureRevision field if non-nil, zero value otherwise.

### GetFeatureRevisionOk

`func (o *HTTP3AllOf) GetFeatureRevisionOk() (*int32, bool)`

GetFeatureRevisionOk returns a tuple with the FeatureRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureRevision

`func (o *HTTP3AllOf) SetFeatureRevision(v int32)`

SetFeatureRevision sets FeatureRevision field to given value.

### HasFeatureRevision

`func (o *HTTP3AllOf) HasFeatureRevision() bool`

HasFeatureRevision returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
