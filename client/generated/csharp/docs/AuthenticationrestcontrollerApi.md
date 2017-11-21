# IO.Swagger.Api.AuthenticationrestcontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthenticationTokenUsingPOST**](AuthenticationrestcontrollerApi.md#createauthenticationtokenusingpost) | **POST** /auth | createAuthenticationToken
[**RefreshAndGetAuthenticationTokenUsingGET**](AuthenticationrestcontrollerApi.md#refreshandgetauthenticationtokenusingget) | **GET** /refresh | refreshAndGetAuthenticationToken


<a name="createauthenticationtokenusingpost"></a>
# **CreateAuthenticationTokenUsingPOST**
> Object CreateAuthenticationTokenUsingPOST (JwtAuthenticationRequest authenticationRequest)

createAuthenticationToken

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class CreateAuthenticationTokenUsingPOSTExample
    {
        public void main()
        {
            var apiInstance = new AuthenticationrestcontrollerApi();
            var authenticationRequest = new JwtAuthenticationRequest(); // JwtAuthenticationRequest | authenticationRequest

            try
            {
                // createAuthenticationToken
                Object result = apiInstance.CreateAuthenticationTokenUsingPOST(authenticationRequest);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling AuthenticationrestcontrollerApi.CreateAuthenticationTokenUsingPOST: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticationRequest** | [**JwtAuthenticationRequest**](JwtAuthenticationRequest.md)| authenticationRequest | 

### Return type

**Object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="refreshandgetauthenticationtokenusingget"></a>
# **RefreshAndGetAuthenticationTokenUsingGET**
> Object RefreshAndGetAuthenticationTokenUsingGET ()

refreshAndGetAuthenticationToken

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class RefreshAndGetAuthenticationTokenUsingGETExample
    {
        public void main()
        {
            var apiInstance = new AuthenticationrestcontrollerApi();

            try
            {
                // refreshAndGetAuthenticationToken
                Object result = apiInstance.RefreshAndGetAuthenticationTokenUsingGET();
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling AuthenticationrestcontrollerApi.RefreshAndGetAuthenticationTokenUsingGET: " + e.Message );
            }
        }
    }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

**Object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

