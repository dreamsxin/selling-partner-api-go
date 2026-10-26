# ProductType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Amazon product type. | 
**DisplayName** | **string** | Human-readable and localized description of the Amazon product type. | 
**MarketplaceIds** | **[]string** | The Amazon marketplace identifiers for which the product type definition is available. | 

## Methods

### NewProductType

`func NewProductType(name string, displayName string, marketplaceIds []string, ) *ProductType`

NewProductType instantiates a new ProductType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductTypeWithDefaults

`func NewProductTypeWithDefaults() *ProductType`

NewProductTypeWithDefaults instantiates a new ProductType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProductType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductType) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *ProductType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ProductType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ProductType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetMarketplaceIds

`func (o *ProductType) GetMarketplaceIds() []string`

GetMarketplaceIds returns the MarketplaceIds field if non-nil, zero value otherwise.

### GetMarketplaceIdsOk

`func (o *ProductType) GetMarketplaceIdsOk() (*[]string, bool)`

GetMarketplaceIdsOk returns a tuple with the MarketplaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceIds

`func (o *ProductType) SetMarketplaceIds(v []string)`

SetMarketplaceIds sets MarketplaceIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


