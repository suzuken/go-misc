# MyServiceproto.YourServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**echo**](YourServiceApi.md#echo) | **POST** /v1/example/echo | 


<a name="echo"></a>
# **echo**
> ExampleStringMessage echo(body)



### Example
```javascript
var MyServiceproto = require('my_serviceproto');

var apiInstance = new MyServiceproto.YourServiceApi();

var body = new MyServiceproto.ExampleStringMessage(); // ExampleStringMessage | 


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.echo(body, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ExampleStringMessage**](ExampleStringMessage.md)|  | 

### Return type

[**ExampleStringMessage**](ExampleStringMessage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

