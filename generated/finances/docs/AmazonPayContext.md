# AmazonPayContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreName** | Pointer to **string** | The name of the store that is related to the transaction. | [optional] 
**OrderType** | Pointer to **string** | The order type of the transaction. | [optional] 
**Channel** | Pointer to **string** | The channel of the transaction. For example: &#x60;MFN&#x60;. | [optional] 

## Methods

### NewAmazonPayContext

`func NewAmazonPayContext() *AmazonPayContext`

NewAmazonPayContext instantiates a new AmazonPayContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonPayContextWithDefaults

`func NewAmazonPayContextWithDefaults() *AmazonPayContext`

NewAmazonPayContextWithDefaults instantiates a new AmazonPayContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreName

`func (o *AmazonPayContext) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *AmazonPayContext) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *AmazonPayContext) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.

### HasStoreName

`func (o *AmazonPayContext) HasStoreName() bool`

HasStoreName returns a boolean if a field has been set.

### GetOrderType

`func (o *AmazonPayContext) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *AmazonPayContext) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *AmazonPayContext) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *AmazonPayContext) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetChannel

`func (o *AmazonPayContext) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *AmazonPayContext) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *AmazonPayContext) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *AmazonPayContext) HasChannel() bool`

HasChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


