# \RolecontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRoleUsingPOST**](RolecontrollerApi.md#AddRoleUsingPOST) | **Post** /role/_roles | addRole
[**EditRoleUsingPUT**](RolecontrollerApi.md#EditRoleUsingPUT) | **Put** /role/_roles/put/{id} | editRole
[**FindRoleUsingGET1**](RolecontrollerApi.md#FindRoleUsingGET1) | **Get** /role/_roles/{roleId} | findRole
[**ListUsingGET2**](RolecontrollerApi.md#ListUsingGET2) | **Get** /role/_roles | list


# **AddRoleUsingPOST**
> Role AddRoleUsingPOST($role)

addRole


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | [**Role**](Role.md)| role | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditRoleUsingPUT**
> Role EditRoleUsingPUT($id, $role)

editRole


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| id | 
 **role** | [**Role**](Role.md)| role | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindRoleUsingGET1**
> Role FindRoleUsingGET1($roleId)

findRole


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleId** | **string**| roleId | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET2**
> []Role ListUsingGET2()

list


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

