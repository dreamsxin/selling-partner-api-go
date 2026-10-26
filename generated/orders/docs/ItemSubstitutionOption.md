# ItemSubstitutionOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asin** | Pointer to **string** | Amazon Standard Identification Number of the substitute product. | [optional] 
**QuantityOrdered** | Pointer to **int32** | Number of units of the substitute item to be selected if substitution occurs. | [optional] 
**SellerSku** | Pointer to **string** | The item&#39;s seller stock keeping unit (SKU). | [optional] 
**Title** | Pointer to **string** | Product name or title of the substitute item as displayed to customers. | [optional] 
**Measurement** | Pointer to [**Measurement**](Measurement.md) |  | [optional] 

## Methods

### NewItemSubstitutionOption

`func NewItemSubstitutionOption() *ItemSubstitutionOption`

NewItemSubstitutionOption instantiates a new ItemSubstitutionOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemSubstitutionOptionWithDefaults

`func NewItemSubstitutionOptionWithDefaults() *ItemSubstitutionOption`

NewItemSubstitutionOptionWithDefaults instantiates a new ItemSubstitutionOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsin

`func (o *ItemSubstitutionOption) GetAsin() string`

GetAsin returns the Asin field if non-nil, zero value otherwise.

### GetAsinOk

`func (o *ItemSubstitutionOption) GetAsinOk() (*string, bool)`

GetAsinOk returns a tuple with the Asin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsin

`func (o *ItemSubstitutionOption) SetAsin(v string)`

SetAsin sets Asin field to given value.

### HasAsin

`func (o *ItemSubstitutionOption) HasAsin() bool`

HasAsin returns a boolean if a field has been set.

### GetQuantityOrdered

`func (o *ItemSubstitutionOption) GetQuantityOrdered() int32`

GetQuantityOrdered returns the QuantityOrdered field if non-nil, zero value otherwise.

### GetQuantityOrderedOk

`func (o *ItemSubstitutionOption) GetQuantityOrderedOk() (*int32, bool)`

GetQuantityOrderedOk returns a tuple with the QuantityOrdered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOrdered

`func (o *ItemSubstitutionOption) SetQuantityOrdered(v int32)`

SetQuantityOrdered sets QuantityOrdered field to given value.

### HasQuantityOrdered

`func (o *ItemSubstitutionOption) HasQuantityOrdered() bool`

HasQuantityOrdered returns a boolean if a field has been set.

### GetSellerSku

`func (o *ItemSubstitutionOption) GetSellerSku() string`

GetSellerSku returns the SellerSku field if non-nil, zero value otherwise.

### GetSellerSkuOk

`func (o *ItemSubstitutionOption) GetSellerSkuOk() (*string, bool)`

GetSellerSkuOk returns a tuple with the SellerSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerSku

`func (o *ItemSubstitutionOption) SetSellerSku(v string)`

SetSellerSku sets SellerSku field to given value.

### HasSellerSku

`func (o *ItemSubstitutionOption) HasSellerSku() bool`

HasSellerSku returns a boolean if a field has been set.

### GetTitle

`func (o *ItemSubstitutionOption) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemSubstitutionOption) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemSubstitutionOption) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ItemSubstitutionOption) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMeasurement

`func (o *ItemSubstitutionOption) GetMeasurement() Measurement`

GetMeasurement returns the Measurement field if non-nil, zero value otherwise.

### GetMeasurementOk

`func (o *ItemSubstitutionOption) GetMeasurementOk() (*Measurement, bool)`

GetMeasurementOk returns a tuple with the Measurement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurement

`func (o *ItemSubstitutionOption) SetMeasurement(v Measurement)`

SetMeasurement sets Measurement field to given value.

### HasMeasurement

`func (o *ItemSubstitutionOption) HasMeasurement() bool`

HasMeasurement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


