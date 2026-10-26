# ItemSearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfResults** | **int32** | The total number of selling partner listings items found for the search criteria (only results up to the page count limit is returned per request regardless of the number found).  Note: The maximum number of items (SKUs) that can be returned and paged through is 1000. | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Items** | [**[]Item**](Item.md) | A list of listings items. | 

## Methods

### NewItemSearchResults

`func NewItemSearchResults(numberOfResults int32, items []Item, ) *ItemSearchResults`

NewItemSearchResults instantiates a new ItemSearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemSearchResultsWithDefaults

`func NewItemSearchResultsWithDefaults() *ItemSearchResults`

NewItemSearchResultsWithDefaults instantiates a new ItemSearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfResults

`func (o *ItemSearchResults) GetNumberOfResults() int32`

GetNumberOfResults returns the NumberOfResults field if non-nil, zero value otherwise.

### GetNumberOfResultsOk

`func (o *ItemSearchResults) GetNumberOfResultsOk() (*int32, bool)`

GetNumberOfResultsOk returns a tuple with the NumberOfResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfResults

`func (o *ItemSearchResults) SetNumberOfResults(v int32)`

SetNumberOfResults sets NumberOfResults field to given value.


### GetPagination

`func (o *ItemSearchResults) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ItemSearchResults) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ItemSearchResults) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ItemSearchResults) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetItems

`func (o *ItemSearchResults) GetItems() []Item`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ItemSearchResults) GetItemsOk() (*[]Item, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ItemSearchResults) SetItems(v []Item)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


