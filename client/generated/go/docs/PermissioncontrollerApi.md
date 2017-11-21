# \PermissioncontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPermissionUsingPOST**](PermissioncontrollerApi.md#AddPermissionUsingPOST) | **Post** /permission/_permission | addPermission
[**EditPermissionUsingPUT**](PermissioncontrollerApi.md#EditPermissionUsingPUT) | **Put** /permission/_permission/{id} | editPermission
[**FindUserUsingGET**](PermissioncontrollerApi.md#FindUserUsingGET) | **Get** /permission/_permission/{id} | findUser
[**ListUsingGET1**](PermissioncontrollerApi.md#ListUsingGET1) | **Get** /permission/_permission | list


# **AddPermissionUsingPOST**
> Permission AddPermissionUsingPOST($permission)

addPermission


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permission** | [**Permission**](Permission.md)| permission | 

### Return type

[**Permission**](Permission.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPermissionUsingPUT**
> Permission EditPermissionUsingPUT($id, $permission)

editPermission


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| id | 
 **permission** | [**Permission**](Permission.md)| permission | 

### Return type

[**Permission**](Permission.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindUserUsingGET**
> Permission FindUserUsingGET($id)

findUser


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| id | 

### Return type

[**Permission**](Permission.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET1**
> []Permission ListUsingGET1($roleId)

list


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleId** | **string**| roleId | 

### Return type

[**[]Permission**](Permission.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

