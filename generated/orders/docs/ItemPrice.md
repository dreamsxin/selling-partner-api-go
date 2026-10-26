# ItemPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitPrice** | Pointer to [**Money**](Money.md) |  | [optional] 
**PriceDesignation** | Pointer to **string** | Indicates that the selling price is a special price that is only available for Amazon Business orders. For more information about the Amazon Business Seller Program, refer to the [Amazon Business website](https://www.amazon.com/b2b/info/amazon-business).  **Possible value**: &#x60;BUSINESS_PRICE&#x60; | [optional] 

## Methods

### NewItemPrice

`func NewItemPrice() *ItemPrice`

NewItemPrice instantiates a new ItemPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemPriceWithDefaults

`func NewItemPriceWithDefaults() *ItemPrice`

NewItemPriceWithDefaults instantiates a new ItemPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitPrice

`func (o *ItemPrice) GetUnitPrice() Money`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *ItemPrice) GetUnitPriceOk() (*Money, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *ItemPrice) SetUnitPrice(v Money)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *ItemPrice) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetPriceDesignation

`func (o *ItemPrice) GetPriceDesignation() string`

GetPriceDesignation returns the PriceDesignation field if non-nil, zero value otherwise.

### GetPriceDesignationOk

`func (o *ItemPrice) GetPriceDesignationOk() (*string, bool)`

GetPriceDesignationOk returns a tuple with the PriceDesignation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDesignation

`func (o *ItemPrice) SetPriceDesignation(v string)`

SetPriceDesignation sets PriceDesignation field to given value.

### HasPriceDesignation

`func (o *ItemPrice) HasPriceDesignation() bool`

HasPriceDesignation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


