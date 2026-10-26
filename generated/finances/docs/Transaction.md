# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SellingPartnerMetadata** | Pointer to [**SellingPartnerMetadata**](SellingPartnerMetadata.md) |  | [optional] 
**RelatedIdentifiers** | Pointer to [**[]RelatedIdentifier**](RelatedIdentifier.md) | Related business identifiers of the transaction. | [optional] 
**TransactionType** | Pointer to **string** | The type of transaction.  **Possible value:** &#x60;Shipment&#x60; | [optional] 
**TransactionId** | Pointer to **string** | The unique identifier of the transaction. | [optional] 
**TransactionStatus** | Pointer to **string** | The status of the transaction.  **Possible values:**  * &#x60;DEFERRED&#x60;: the transaction is currently deferred. * &#x60;RELEASED&#x60;: the transaction is currently released. * &#x60;DEFERRED_RELEASED&#x60;: the transaction was deferred in the past, but is now released. The status of a deferred transaction is updated to &#x60;DEFERRED_RELEASED&#x60; when the transaction is released. | [optional] 
**Description** | Pointer to **string** | Describes the reasons for the transaction.  **Example:** &#x60;Order Payment&#x60;, &#x60;Refund Order&#x60; | [optional] 
**PostedDate** | Pointer to **time.Time** | A date in [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) date-time format. | [optional] 
**TotalAmount** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**MarketplaceDetails** | Pointer to [**MarketplaceDetails**](MarketplaceDetails.md) |  | [optional] 
**Items** | Pointer to [**[]Item**](Item.md) | A list of items in the transaction. | [optional] 
**Contexts** | Pointer to [**[]Context**](Context.md) | A list of additional information about the item. | [optional] 
**Breakdowns** | Pointer to [**[]Breakdown**](Breakdown.md) | A list of breakdowns that detail how the total amount is calculated for the transaction. | [optional] 

## Methods

### NewTransaction

`func NewTransaction() *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSellingPartnerMetadata

`func (o *Transaction) GetSellingPartnerMetadata() SellingPartnerMetadata`

GetSellingPartnerMetadata returns the SellingPartnerMetadata field if non-nil, zero value otherwise.

### GetSellingPartnerMetadataOk

`func (o *Transaction) GetSellingPartnerMetadataOk() (*SellingPartnerMetadata, bool)`

GetSellingPartnerMetadataOk returns a tuple with the SellingPartnerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingPartnerMetadata

`func (o *Transaction) SetSellingPartnerMetadata(v SellingPartnerMetadata)`

SetSellingPartnerMetadata sets SellingPartnerMetadata field to given value.

### HasSellingPartnerMetadata

`func (o *Transaction) HasSellingPartnerMetadata() bool`

HasSellingPartnerMetadata returns a boolean if a field has been set.

### GetRelatedIdentifiers

`func (o *Transaction) GetRelatedIdentifiers() []RelatedIdentifier`

GetRelatedIdentifiers returns the RelatedIdentifiers field if non-nil, zero value otherwise.

### GetRelatedIdentifiersOk

`func (o *Transaction) GetRelatedIdentifiersOk() (*[]RelatedIdentifier, bool)`

GetRelatedIdentifiersOk returns a tuple with the RelatedIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedIdentifiers

`func (o *Transaction) SetRelatedIdentifiers(v []RelatedIdentifier)`

SetRelatedIdentifiers sets RelatedIdentifiers field to given value.

### HasRelatedIdentifiers

`func (o *Transaction) HasRelatedIdentifiers() bool`

HasRelatedIdentifiers returns a boolean if a field has been set.

### GetTransactionType

`func (o *Transaction) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *Transaction) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *Transaction) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *Transaction) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetTransactionId

`func (o *Transaction) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Transaction) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Transaction) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *Transaction) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionStatus

`func (o *Transaction) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *Transaction) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *Transaction) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.

### HasTransactionStatus

`func (o *Transaction) HasTransactionStatus() bool`

HasTransactionStatus returns a boolean if a field has been set.

### GetDescription

`func (o *Transaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Transaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Transaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Transaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPostedDate

`func (o *Transaction) GetPostedDate() time.Time`

GetPostedDate returns the PostedDate field if non-nil, zero value otherwise.

### GetPostedDateOk

`func (o *Transaction) GetPostedDateOk() (*time.Time, bool)`

GetPostedDateOk returns a tuple with the PostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedDate

`func (o *Transaction) SetPostedDate(v time.Time)`

SetPostedDate sets PostedDate field to given value.

### HasPostedDate

`func (o *Transaction) HasPostedDate() bool`

HasPostedDate returns a boolean if a field has been set.

### GetTotalAmount

`func (o *Transaction) GetTotalAmount() Currency`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *Transaction) GetTotalAmountOk() (*Currency, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *Transaction) SetTotalAmount(v Currency)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *Transaction) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetMarketplaceDetails

`func (o *Transaction) GetMarketplaceDetails() MarketplaceDetails`

GetMarketplaceDetails returns the MarketplaceDetails field if non-nil, zero value otherwise.

### GetMarketplaceDetailsOk

`func (o *Transaction) GetMarketplaceDetailsOk() (*MarketplaceDetails, bool)`

GetMarketplaceDetailsOk returns a tuple with the MarketplaceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceDetails

`func (o *Transaction) SetMarketplaceDetails(v MarketplaceDetails)`

SetMarketplaceDetails sets MarketplaceDetails field to given value.

### HasMarketplaceDetails

`func (o *Transaction) HasMarketplaceDetails() bool`

HasMarketplaceDetails returns a boolean if a field has been set.

### GetItems

`func (o *Transaction) GetItems() []Item`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Transaction) GetItemsOk() (*[]Item, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Transaction) SetItems(v []Item)`

SetItems sets Items field to given value.

### HasItems

`func (o *Transaction) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetContexts

`func (o *Transaction) GetContexts() []Context`

GetContexts returns the Contexts field if non-nil, zero value otherwise.

### GetContextsOk

`func (o *Transaction) GetContextsOk() (*[]Context, bool)`

GetContextsOk returns a tuple with the Contexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContexts

`func (o *Transaction) SetContexts(v []Context)`

SetContexts sets Contexts field to given value.

### HasContexts

`func (o *Transaction) HasContexts() bool`

HasContexts returns a boolean if a field has been set.

### GetBreakdowns

`func (o *Transaction) GetBreakdowns() []Breakdown`

GetBreakdowns returns the Breakdowns field if non-nil, zero value otherwise.

### GetBreakdownsOk

`func (o *Transaction) GetBreakdownsOk() (*[]Breakdown, bool)`

GetBreakdownsOk returns a tuple with the Breakdowns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdowns

`func (o *Transaction) SetBreakdowns(v []Breakdown)`

SetBreakdowns sets Breakdowns field to given value.

### HasBreakdowns

`func (o *Transaction) HasBreakdowns() bool`

HasBreakdowns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


