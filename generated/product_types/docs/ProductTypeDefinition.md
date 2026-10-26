# ProductTypeDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaSchema** | Pointer to [**SchemaLink**](SchemaLink.md) |  | [optional] 
**Schema** | [**SchemaLink**](SchemaLink.md) |  | 
**Requirements** | **string** | Name of the requirements set represented in this product type definition. | 
**RequirementsEnforced** | **string** | Identifies if the required attributes for a requirements set are enforced by the product type definition schema. Non-enforced requirements enable structural validation of individual attributes without all of the required attributes being present (such as for partial updates). | 
**PropertyGroups** | [**map[string]PropertyGroup**](PropertyGroup.md) | Mapping of property group names to property groups. Property groups represent logical groupings of schema properties that can be used for display or informational purposes. | 
**Locale** | **string** | Locale of the display elements contained in the product type definition. | 
**MarketplaceIds** | **[]string** | Amazon marketplace identifiers for which the product type definition is applicable. | 
**ProductType** | **string** | The name of the Amazon product type that this product type definition applies to. | 
**DisplayName** | **string** | Human-readable and localized description of the Amazon product type. | 
**ProductTypeVersion** | [**ProductTypeVersion**](ProductTypeVersion.md) |  | 

## Methods

### NewProductTypeDefinition

`func NewProductTypeDefinition(schema SchemaLink, requirements string, requirementsEnforced string, propertyGroups map[string]PropertyGroup, locale string, marketplaceIds []string, productType string, displayName string, productTypeVersion ProductTypeVersion, ) *ProductTypeDefinition`

NewProductTypeDefinition instantiates a new ProductTypeDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductTypeDefinitionWithDefaults

`func NewProductTypeDefinitionWithDefaults() *ProductTypeDefinition`

NewProductTypeDefinitionWithDefaults instantiates a new ProductTypeDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaSchema

`func (o *ProductTypeDefinition) GetMetaSchema() SchemaLink`

GetMetaSchema returns the MetaSchema field if non-nil, zero value otherwise.

### GetMetaSchemaOk

`func (o *ProductTypeDefinition) GetMetaSchemaOk() (*SchemaLink, bool)`

GetMetaSchemaOk returns a tuple with the MetaSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaSchema

`func (o *ProductTypeDefinition) SetMetaSchema(v SchemaLink)`

SetMetaSchema sets MetaSchema field to given value.

### HasMetaSchema

`func (o *ProductTypeDefinition) HasMetaSchema() bool`

HasMetaSchema returns a boolean if a field has been set.

### GetSchema

`func (o *ProductTypeDefinition) GetSchema() SchemaLink`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ProductTypeDefinition) GetSchemaOk() (*SchemaLink, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ProductTypeDefinition) SetSchema(v SchemaLink)`

SetSchema sets Schema field to given value.


### GetRequirements

`func (o *ProductTypeDefinition) GetRequirements() string`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *ProductTypeDefinition) GetRequirementsOk() (*string, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *ProductTypeDefinition) SetRequirements(v string)`

SetRequirements sets Requirements field to given value.


### GetRequirementsEnforced

`func (o *ProductTypeDefinition) GetRequirementsEnforced() string`

GetRequirementsEnforced returns the RequirementsEnforced field if non-nil, zero value otherwise.

### GetRequirementsEnforcedOk

`func (o *ProductTypeDefinition) GetRequirementsEnforcedOk() (*string, bool)`

GetRequirementsEnforcedOk returns a tuple with the RequirementsEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementsEnforced

`func (o *ProductTypeDefinition) SetRequirementsEnforced(v string)`

SetRequirementsEnforced sets RequirementsEnforced field to given value.


### GetPropertyGroups

`func (o *ProductTypeDefinition) GetPropertyGroups() map[string]PropertyGroup`

GetPropertyGroups returns the PropertyGroups field if non-nil, zero value otherwise.

### GetPropertyGroupsOk

`func (o *ProductTypeDefinition) GetPropertyGroupsOk() (*map[string]PropertyGroup, bool)`

GetPropertyGroupsOk returns a tuple with the PropertyGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyGroups

`func (o *ProductTypeDefinition) SetPropertyGroups(v map[string]PropertyGroup)`

SetPropertyGroups sets PropertyGroups field to given value.


### GetLocale

`func (o *ProductTypeDefinition) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *ProductTypeDefinition) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *ProductTypeDefinition) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetMarketplaceIds

`func (o *ProductTypeDefinition) GetMarketplaceIds() []string`

GetMarketplaceIds returns the MarketplaceIds field if non-nil, zero value otherwise.

### GetMarketplaceIdsOk

`func (o *ProductTypeDefinition) GetMarketplaceIdsOk() (*[]string, bool)`

GetMarketplaceIdsOk returns a tuple with the MarketplaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceIds

`func (o *ProductTypeDefinition) SetMarketplaceIds(v []string)`

SetMarketplaceIds sets MarketplaceIds field to given value.


### GetProductType

`func (o *ProductTypeDefinition) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ProductTypeDefinition) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ProductTypeDefinition) SetProductType(v string)`

SetProductType sets ProductType field to given value.


### GetDisplayName

`func (o *ProductTypeDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ProductTypeDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ProductTypeDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetProductTypeVersion

`func (o *ProductTypeDefinition) GetProductTypeVersion() ProductTypeVersion`

GetProductTypeVersion returns the ProductTypeVersion field if non-nil, zero value otherwise.

### GetProductTypeVersionOk

`func (o *ProductTypeDefinition) GetProductTypeVersionOk() (*ProductTypeVersion, bool)`

GetProductTypeVersionOk returns a tuple with the ProductTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeVersion

`func (o *ProductTypeDefinition) SetProductTypeVersion(v ProductTypeVersion)`

SetProductTypeVersion sets ProductTypeVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


