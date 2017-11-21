# \UsercontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserUsingPOST**](UsercontrollerApi.md#AddUserUsingPOST) | **Post** /user/_users | addUser
[**EditFieldUsingPUT1**](UsercontrollerApi.md#EditFieldUsingPUT1) | **Put** /user/_users/put/{id} | editField
[**FindUserUsingGET1**](UsercontrollerApi.md#FindUserUsingGET1) | **Get** /user/_users/{userId} | findUser
[**GetAuthenticatedUserUsingGET**](UsercontrollerApi.md#GetAuthenticatedUserUsingGET) | **Get** /user/me | getAuthenticatedUser
[**ListUsingGET3**](UsercontrollerApi.md#ListUsingGET3) | **Get** /user/_users | list


# **AddUserUsingPOST**
> User AddUserUsingPOST($user)

addUser


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | [**User**](User.md)| user | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditFieldUsingPUT1**
> User EditFieldUsingPUT1($id, $user)

editField


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| id | 
 **user** | [**User**](User.md)| user | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindUserUsingGET1**
> User FindUserUsingGET1($userId)

findUser


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| userId | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthenticatedUserUsingGET**
> JwtUser GetAuthenticatedUserUsingGET()

getAuthenticatedUser


### Parameters
This endpoint does not need any parameter.

### Return type

[**JwtUser**](JwtUser.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET3**
> []User ListUsingGET3()

list


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

