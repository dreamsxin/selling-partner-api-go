# BusinessContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreName** | Pointer to **string** | The store name associated with the transaction. | [optional] 

## Methods

### NewBusinessContext

`func NewBusinessContext() *BusinessContext`

NewBusinessContext instantiates a new BusinessContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessContextWithDefaults

`func NewBusinessContextWithDefaults() *BusinessContext`

NewBusinessContextWithDefaults instantiates a new BusinessContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreName

`func (o *BusinessContext) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *BusinessContext) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *BusinessContext) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.

### HasStoreName

`func (o *BusinessContext) HasStoreName() bool`

HasStoreName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


