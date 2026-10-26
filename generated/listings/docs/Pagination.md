# Pagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** | A token that can be used to fetch the next page. | [optional] 
**PreviousToken** | Pointer to **string** | A token that can be used to fetch the previous page. | [optional] 

## Methods

### NewPagination

`func NewPagination() *Pagination`

NewPagination instantiates a new Pagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationWithDefaults

`func NewPaginationWithDefaults() *Pagination`

NewPaginationWithDefaults instantiates a new Pagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *Pagination) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *Pagination) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *Pagination) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *Pagination) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetPreviousToken

`func (o *Pagination) GetPreviousToken() string`

GetPreviousToken returns the PreviousToken field if non-nil, zero value otherwise.

### GetPreviousTokenOk

`func (o *Pagination) GetPreviousTokenOk() (*string, bool)`

GetPreviousTokenOk returns a tuple with the PreviousToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousToken

`func (o *Pagination) SetPreviousToken(v string)`

SetPreviousToken sets PreviousToken field to given value.

### HasPreviousToken

`func (o *Pagination) HasPreviousToken() bool`

HasPreviousToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


