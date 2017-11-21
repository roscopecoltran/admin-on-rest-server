# \DatacontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataMutationUsingDELETE**](DatacontrollerApi.md#DataMutationUsingDELETE) | **Delete** /api/{entity}/{id} | dataMutation
[**DataMutationUsingPOST**](DatacontrollerApi.md#DataMutationUsingPOST) | **Post** /api/{entity} | dataMutation
[**DataMutationUsingPUT**](DatacontrollerApi.md#DataMutationUsingPUT) | **Put** /api/{entity}/{id} | dataMutation
[**DataQueryUsingGET**](DatacontrollerApi.md#DataQueryUsingGET) | **Get** /api/{entity} | dataQuery
[**FindOneUsingGET**](DatacontrollerApi.md#FindOneUsingGET) | **Get** /api/{entity}/{id} | findOne


# **DataMutationUsingDELETE**
> map[string]interface{} DataMutationUsingDELETE($entity, $id)

dataMutation


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **string**| entity | 
 **id** | **string**| id | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DataMutationUsingPOST**
> map[string]interface{} DataMutationUsingPOST($entity, $allRequestParams)

dataMutation


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **string**| entity | 
 **allRequestParams** | [**interface{}**](interface{}.md)| allRequestParams | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DataMutationUsingPUT**
> map[string]interface{} DataMutationUsingPUT($entity, $id, $allRequestParams)

dataMutation


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **string**| entity | 
 **id** | **string**| id | 
 **allRequestParams** | [**interface{}**](interface{}.md)| allRequestParams | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DataQueryUsingGET**
> map[string]interface{} DataQueryUsingGET($entity)

dataQuery


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **string**| entity | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOneUsingGET**
> map[string]interface{} FindOneUsingGET($entity, $id)

findOne


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **string**| entity | 
 **id** | **string**| id | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

