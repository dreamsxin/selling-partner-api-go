# ProductTypeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductTypes** | [**[]ProductType**](ProductType.md) | A list of product types. | 
**ProductTypeVersion** | **string** | Amazon product type version identifier. | 

## Methods

### NewProductTypeList

`func NewProductTypeList(productTypes []ProductType, productTypeVersion string, ) *ProductTypeList`

NewProductTypeList instantiates a new ProductTypeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductTypeListWithDefaults

`func NewProductTypeListWithDefaults() *ProductTypeList`

NewProductTypeListWithDefaults instantiates a new ProductTypeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductTypes

`func (o *ProductTypeList) GetProductTypes() []ProductType`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *ProductTypeList) GetProductTypesOk() (*[]ProductType, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *ProductTypeList) SetProductTypes(v []ProductType)`

SetProductTypes sets ProductTypes field to given value.


### GetProductTypeVersion

`func (o *ProductTypeList) GetProductTypeVersion() string`

GetProductTypeVersion returns the ProductTypeVersion field if non-nil, zero value otherwise.

### GetProductTypeVersionOk

`func (o *ProductTypeList) GetProductTypeVersionOk() (*string, bool)`

GetProductTypeVersionOk returns a tuple with the ProductTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeVersion

`func (o *ProductTypeList) SetProductTypeVersion(v string)`

SetProductTypeVersion sets ProductTypeVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


