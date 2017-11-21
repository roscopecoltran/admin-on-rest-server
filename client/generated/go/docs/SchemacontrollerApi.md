# \SchemacontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEntityUsingPOST**](SchemacontrollerApi.md#AddEntityUsingPOST) | **Post** /schemas/_entitys | addEntity
[**AddFieldUsingPOST**](SchemacontrollerApi.md#AddFieldUsingPOST) | **Post** /schemas/_fields | addField
[**EditEntityUsingPUT**](SchemacontrollerApi.md#EditEntityUsingPUT) | **Put** /schemas/_entitys/put/{id} | editEntity
[**EditFieldUsingPUT**](SchemacontrollerApi.md#EditFieldUsingPUT) | **Put** /schemas/_fields/put/{id} | editField
[**FindAllFieldsUsingGET**](SchemacontrollerApi.md#FindAllFieldsUsingGET) | **Get** /schemas/_fields | findAllFields
[**FindEntityFieldsUsingGET**](SchemacontrollerApi.md#FindEntityFieldsUsingGET) | **Get** /schemas/fields | findEntityFields
[**FindOneFieldUsingGET**](SchemacontrollerApi.md#FindOneFieldUsingGET) | **Get** /schemas/_fields/{fid} | findOneField
[**FindOneUsingGET1**](SchemacontrollerApi.md#FindOneUsingGET1) | **Get** /schemas/_entitys/{eid} | findOne
[**GetSchemasUsingGET**](SchemacontrollerApi.md#GetSchemasUsingGET) | **Get** /schemas/_entitys | getSchemas
[**ResetCurrentDsUsingPUT**](SchemacontrollerApi.md#ResetCurrentDsUsingPUT) | **Put** /schemas/resetCurrentDs/{dataSourceId} | resetCurrentDs
[**SyncSchemasUsingPUT**](SchemacontrollerApi.md#SyncSchemasUsingPUT) | **Put** /schemas/sync/{dataSourceId} | syncSchemas


# **AddEntityUsingPOST**
> Entity AddEntityUsingPOST($entity)

addEntity


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | [**Entity**](Entity.md)| entity | 

### Return type

[**Entity**](Entity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddFieldUsingPOST**
> Field AddFieldUsingPOST($field)

addField


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **field** | [**Field**](Field.md)| field | 

### Return type

[**Field**](Field.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditEntityUsingPUT**
> Entity EditEntityUsingPUT($id, $entity)

editEntity


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| id | 
 **entity** | [**Entity**](Entity.md)| entity | 

### Return type

[**Entity**](Entity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditFieldUsingPUT**
> Field EditFieldUsingPUT($id, $field)

editField


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| id | 
 **field** | [**Field**](Field.md)| field | 

### Return type

[**Field**](Field.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindAllFieldsUsingGET**
> []Field FindAllFieldsUsingGET($eid)

findAllFields


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eid** | **string**| eid | 

### Return type

[**[]Field**](Field.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindEntityFieldsUsingGET**
> []Field FindEntityFieldsUsingGET($entityName)

findEntityFields


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityName** | **string**| entityName | 

### Return type

[**[]Field**](Field.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOneFieldUsingGET**
> Field FindOneFieldUsingGET($fid)

findOneField


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fid** | **string**| fid | 

### Return type

[**Field**](Field.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOneUsingGET1**
> Entity FindOneUsingGET1($eid)

findOne


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eid** | **string**| eid | 

### Return type

[**Entity**](Entity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSchemasUsingGET**
> []Entity GetSchemasUsingGET()

getSchemas


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]Entity**](Entity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetCurrentDsUsingPUT**
> ResponseEntity ResetCurrentDsUsingPUT($dataSourceId)

resetCurrentDs


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSourceId** | **string**| dataSourceId | 

### Return type

[**ResponseEntity**](ResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncSchemasUsingPUT**
> ResponseEntity SyncSchemasUsingPUT($dataSourceId)

syncSchemas


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataSourceId** | **string**| dataSourceId | 

### Return type

[**ResponseEntity**](ResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

