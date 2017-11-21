# swagger_client.AuthenticationrestcontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_authentication_token_using_post**](AuthenticationrestcontrollerApi.md#create_authentication_token_using_post) | **POST** /auth | createAuthenticationToken
[**refresh_and_get_authentication_token_using_get**](AuthenticationrestcontrollerApi.md#refresh_and_get_authentication_token_using_get) | **GET** /refresh | refreshAndGetAuthenticationToken


# **create_authentication_token_using_post**
> object create_authentication_token_using_post(authentication_request)

createAuthenticationToken

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.AuthenticationrestcontrollerApi()
authentication_request = swagger_client.JwtAuthenticationRequest() # JwtAuthenticationRequest | authenticationRequest

try: 
    # createAuthenticationToken
    api_response = api_instance.create_authentication_token_using_post(authentication_request)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling AuthenticationrestcontrollerApi->create_authentication_token_using_post: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authentication_request** | [**JwtAuthenticationRequest**](JwtAuthenticationRequest.md)| authenticationRequest | 

### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **refresh_and_get_authentication_token_using_get**
> object refresh_and_get_authentication_token_using_get()

refreshAndGetAuthenticationToken

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.AuthenticationrestcontrollerApi()

try: 
    # refreshAndGetAuthenticationToken
    api_response = api_instance.refresh_and_get_authentication_token_using_get()
    pprint(api_response)
except ApiException as e:
    print("Exception when calling AuthenticationrestcontrollerApi->refresh_and_get_authentication_token_using_get: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

