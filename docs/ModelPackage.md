# ModelPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**PackageMetadata**](PackageMetadata.md) |  | [optional] 

## Methods

### NewModelPackage

`func NewModelPackage() *ModelPackage`

NewModelPackage instantiates a new ModelPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPackageWithDefaults

`func NewModelPackageWithDefaults() *ModelPackage`

NewModelPackageWithDefaults instantiates a new ModelPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ModelPackage) GetMetadata() PackageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelPackage) GetMetadataOk() (*PackageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelPackage) SetMetadata(v PackageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ModelPackage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
