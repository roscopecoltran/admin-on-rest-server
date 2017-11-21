# \AuthenticationrestcontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthenticationTokenUsingPOST**](AuthenticationrestcontrollerApi.md#CreateAuthenticationTokenUsingPOST) | **Post** /auth | createAuthenticationToken
[**RefreshAndGetAuthenticationTokenUsingGET**](AuthenticationrestcontrollerApi.md#RefreshAndGetAuthenticationTokenUsingGET) | **Get** /refresh | refreshAndGetAuthenticationToken


# **CreateAuthenticationTokenUsingPOST**
> interface{} CreateAuthenticationTokenUsingPOST($authenticationRequest)

createAuthenticationToken


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticationRequest** | [**JwtAuthenticationRequest**](JwtAuthenticationRequest.md)| authenticationRequest | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshAndGetAuthenticationTokenUsingGET**
> interface{} RefreshAndGetAuthenticationTokenUsingGET()

refreshAndGetAuthenticationToken


### Parameters
This endpoint does not need any parameter.

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

