# swagger_client.ApplycontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**apply_using_post**](ApplycontrollerApi.md#apply_using_post) | **POST** /apply | apply


# **apply_using_post**
> ResponseEntity apply_using_post(apply)

apply

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.ApplycontrollerApi()
apply = swagger_client.Apply() # Apply | apply

try: 
    # apply
    api_response = api_instance.apply_using_post(apply)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ApplycontrollerApi->apply_using_post: %s\n" % e)
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

