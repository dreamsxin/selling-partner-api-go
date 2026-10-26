# GetSubscriptionsPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscriptions** | Pointer to [**[]Subscription**](Subscription.md) | A list of subscriptions. | [optional] 
**NextToken** | Pointer to **string** | A token that you can use to retrieve the next page of results. When this field is not empty, pass its value in the &#x60;nextToken&#x60; query parameter of the next request. | [optional] 

## Methods

### NewGetSubscriptionsPayload

`func NewGetSubscriptionsPayload() *GetSubscriptionsPayload`

NewGetSubscriptionsPayload instantiates a new GetSubscriptionsPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSubscriptionsPayloadWithDefaults

`func NewGetSubscriptionsPayloadWithDefaults() *GetSubscriptionsPayload`

NewGetSubscriptionsPayloadWithDefaults instantiates a new GetSubscriptionsPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptions

`func (o *GetSubscriptionsPayload) GetSubscriptions() []Subscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *GetSubscriptionsPayload) GetSubscriptionsOk() (*[]Subscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *GetSubscriptionsPayload) SetSubscriptions(v []Subscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *GetSubscriptionsPayload) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetNextToken

`func (o *GetSubscriptionsPayload) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *GetSubscriptionsPayload) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *GetSubscriptionsPayload) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *GetSubscriptionsPayload) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


