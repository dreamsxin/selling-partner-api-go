# Buyer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerName** | Pointer to **string** | The full name of the customer who placed the order. | [optional] 
**BuyerEmail** | Pointer to **string** | The anonymized email address of the buyer. **Note:** Only available for merchant-fulfilled (FBM) orders. | [optional] 
**BuyerCompanyName** | Pointer to **string** | The name of the company or organization for a business order. | [optional] 
**BuyerPurchaseOrderNumber** | Pointer to **string** | The purchase order (PO) number entered by the buyer at checkout. Only returned for orders where the buyer entered a PO number at checkout. | [optional] 

## Methods

### NewBuyer

`func NewBuyer() *Buyer`

NewBuyer instantiates a new Buyer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyerWithDefaults

`func NewBuyerWithDefaults() *Buyer`

NewBuyerWithDefaults instantiates a new Buyer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerName

`func (o *Buyer) GetBuyerName() string`

GetBuyerName returns the BuyerName field if non-nil, zero value otherwise.

### GetBuyerNameOk

`func (o *Buyer) GetBuyerNameOk() (*string, bool)`

GetBuyerNameOk returns a tuple with the BuyerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerName

`func (o *Buyer) SetBuyerName(v string)`

SetBuyerName sets BuyerName field to given value.

### HasBuyerName

`func (o *Buyer) HasBuyerName() bool`

HasBuyerName returns a boolean if a field has been set.

### GetBuyerEmail

`func (o *Buyer) GetBuyerEmail() string`

GetBuyerEmail returns the BuyerEmail field if non-nil, zero value otherwise.

### GetBuyerEmailOk

`func (o *Buyer) GetBuyerEmailOk() (*string, bool)`

GetBuyerEmailOk returns a tuple with the BuyerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerEmail

`func (o *Buyer) SetBuyerEmail(v string)`

SetBuyerEmail sets BuyerEmail field to given value.

### HasBuyerEmail

`func (o *Buyer) HasBuyerEmail() bool`

HasBuyerEmail returns a boolean if a field has been set.

### GetBuyerCompanyName

`func (o *Buyer) GetBuyerCompanyName() string`

GetBuyerCompanyName returns the BuyerCompanyName field if non-nil, zero value otherwise.

### GetBuyerCompanyNameOk

`func (o *Buyer) GetBuyerCompanyNameOk() (*string, bool)`

GetBuyerCompanyNameOk returns a tuple with the BuyerCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerCompanyName

`func (o *Buyer) SetBuyerCompanyName(v string)`

SetBuyerCompanyName sets BuyerCompanyName field to given value.

### HasBuyerCompanyName

`func (o *Buyer) HasBuyerCompanyName() bool`

HasBuyerCompanyName returns a boolean if a field has been set.

### GetBuyerPurchaseOrderNumber

`func (o *Buyer) GetBuyerPurchaseOrderNumber() string`

GetBuyerPurchaseOrderNumber returns the BuyerPurchaseOrderNumber field if non-nil, zero value otherwise.

### GetBuyerPurchaseOrderNumberOk

`func (o *Buyer) GetBuyerPurchaseOrderNumberOk() (*string, bool)`

GetBuyerPurchaseOrderNumberOk returns a tuple with the BuyerPurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerPurchaseOrderNumber

`func (o *Buyer) SetBuyerPurchaseOrderNumber(v string)`

SetBuyerPurchaseOrderNumber sets BuyerPurchaseOrderNumber field to given value.

### HasBuyerPurchaseOrderNumber

`func (o *Buyer) HasBuyerPurchaseOrderNumber() bool`

HasBuyerPurchaseOrderNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


