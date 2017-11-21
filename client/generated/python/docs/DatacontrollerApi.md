# swagger_client.DatacontrollerApi

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**data_mutation_using_delete**](DatacontrollerApi.md#data_mutation_using_delete) | **DELETE** /api/{entity}/{id} | dataMutation
[**data_mutation_using_post**](DatacontrollerApi.md#data_mutation_using_post) | **POST** /api/{entity} | dataMutation
[**data_mutation_using_put**](DatacontrollerApi.md#data_mutation_using_put) | **PUT** /api/{entity}/{id} | dataMutation
[**data_query_using_get**](DatacontrollerApi.md#data_query_using_get) | **GET** /api/{entity} | dataQuery
[**find_one_using_get**](DatacontrollerApi.md#find_one_using_get) | **GET** /api/{entity}/{id} | findOne


# **data_mutation_using_delete**
> dict(str, object) data_mutation_using_delete(entity, id)

dataMutation

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.DatacontrollerApi()
entity = 'entity_example' # str | entity
id = 'id_example' # str | id

try: 
    # dataMutation
    api_response = api_instance.data_mutation_using_delete(entity, id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DatacontrollerApi->data_mutation_using_delete: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **str**| entity | 
 **id** | **str**| id | 

### Return type

[**dict(str, object)**](dict.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **data_mutation_using_post**
> dict(str, object) data_mutation_using_post(entity, all_request_params)

dataMutation

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.DatacontrollerApi()
entity = 'entity_example' # str | entity
all_request_params = NULL # object | allRequestParams

try: 
    # dataMutation
    api_response = api_instance.data_mutation_using_post(entity, all_request_params)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DatacontrollerApi->data_mutation_using_post: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **str**| entity | 
 **all_request_params** | **object**| allRequestParams | 

### Return type

[**dict(str, object)**](dict.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **data_mutation_using_put**
> dict(str, object) data_mutation_using_put(entity, id, all_request_params)

dataMutation

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.DatacontrollerApi()
entity = 'entity_example' # str | entity
id = 'id_example' # str | id
all_request_params = NULL # object | allRequestParams

try: 
    # dataMutation
    api_response = api_instance.data_mutation_using_put(entity, id, all_request_params)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DatacontrollerApi->data_mutation_using_put: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **str**| entity | 
 **id** | **str**| id | 
 **all_request_params** | **object**| allRequestParams | 

### Return type

[**dict(str, object)**](dict.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **data_query_using_get**
> dict(str, object) data_query_using_get(entity)

dataQuery

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.DatacontrollerApi()
entity = 'entity_example' # str | entity

try: 
    # dataQuery
    api_response = api_instance.data_query_using_get(entity)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DatacontrollerApi->data_query_using_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **str**| entity | 

### Return type

[**dict(str, object)**](dict.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **find_one_using_get**
> dict(str, object) find_one_using_get(entity, id)

findOne

### Example 
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.DatacontrollerApi()
entity = 'entity_example' # str | entity
id = 'id_example' # str | id

try: 
    # findOne
    api_response = api_instance.find_one_using_get(entity, id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DatacontrollerApi->find_one_using_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **str**| entity | 
 **id** | **str**| id | 

### Return type

[**dict(str, object)**](dict.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

