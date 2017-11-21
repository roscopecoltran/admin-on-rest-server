# IO.Swagger.Api.ApplycontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyUsingPOST**](ApplycontrollerApi.md#applyusingpost) | **POST** /apply | apply


<a name="applyusingpost"></a>
# **ApplyUsingPOST**
> ResponseEntity ApplyUsingPOST (Apply apply)

apply

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ApplyUsingPOSTExample
    {
        public void main()
        {
            var apiInstance = new ApplycontrollerApi();
            var apply = new Apply(); // Apply | apply

            try
            {
                // apply
                ResponseEntity result = apiInstance.ApplyUsingPOST(apply);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling ApplycontrollerApi.ApplyUsingPOST: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apply** | [**Apply**](Apply.md)| apply | 

### Return type

[**ResponseEntity**](ResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

