# ApexRedirectAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int32** | HTTP status code used to redirect the client. | [optional] 
**Domains** | Pointer to **[]string** | Array of apex domains that should redirect to their WWW subdomain. | [optional] 
**FeatureRevision** | Pointer to **int32** | Revision number of the apex redirect feature implementation. Defaults to the most recent revision. | [optional] 

## Methods

### NewApexRedirectAllOf

`func NewApexRedirectAllOf() *ApexRedirectAllOf`

NewApexRedirectAllOf instantiates a new ApexRedirectAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApexRedirectAllOfWithDefaults

`func NewApexRedirectAllOfWithDefaults() *ApexRedirectAllOf`

NewApexRedirectAllOfWithDefaults instantiates a new ApexRedirectAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *ApexRedirectAllOf) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ApexRedirectAllOf) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ApexRedirectAllOf) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ApexRedirectAllOf) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetDomains

`func (o *ApexRedirectAllOf) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *ApexRedirectAllOf) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *ApexRedirectAllOf) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *ApexRedirectAllOf) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetFeatureRevision

`func (o *ApexRedirectAllOf) GetFeatureRevision() int32`

GetFeatureRevision returns the FeatureRevision field if non-nil, zero value otherwise.

### GetFeatureRevisionOk

`func (o *ApexRedirectAllOf) GetFeatureRevisionOk() (*int32, bool)`

GetFeatureRevisionOk returns a tuple with the FeatureRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureRevision

`func (o *ApexRedirectAllOf) SetFeatureRevision(v int32)`

SetFeatureRevision sets FeatureRevision field to given value.

### HasFeatureRevision

`func (o *ApexRedirectAllOf) HasFeatureRevision() bool`

HasFeatureRevision returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
