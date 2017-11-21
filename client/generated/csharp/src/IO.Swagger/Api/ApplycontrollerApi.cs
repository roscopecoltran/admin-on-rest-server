/* 
 * DataHive RESTful APIs
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using RestSharp;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace IO.Swagger.Api
{
    /// <summary>
    /// Represents a collection of functions to interact with the API endpoints
    /// </summary>
    public interface IApplycontrollerApi : IApiAccessor
    {
        #region Synchronous Operations
        /// <summary>
        /// apply
        /// </summary>
        /// <remarks>
        /// 
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="apply">apply</param>
        /// <returns>ResponseEntity</returns>
        ResponseEntity ApplyUsingPOST (Apply apply);

        /// <summary>
        /// apply
        /// </summary>
        /// <remarks>
        /// 
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="apply">apply</param>
        /// <returns>ApiResponse of ResponseEntity</returns>
        ApiResponse<ResponseEntity> ApplyUsingPOSTWithHttpInfo (Apply apply);
        #endregion Synchronous Operations
        #region Asynchronous Operations
        /// <summary>
        /// apply
        /// </summary>
        /// <remarks>
        /// 
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="apply">apply</param>
        /// <returns>Task of ResponseEntity</returns>
        System.Threading.Tasks.Task<ResponseEntity> ApplyUsingPOSTAsync (Apply apply);

        /// <summary>
        /// apply
        /// </summary>
        /// <remarks>
        /// 
        /// </remarks>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="apply">apply</param>
        /// <returns>Task of ApiResponse (ResponseEntity)</returns>
        System.Threading.Tasks.Task<ApiResponse<ResponseEntity>> ApplyUsingPOSTAsyncWithHttpInfo (Apply apply);
        #endregion Asynchronous Operations
    }

    /// <summary>
    /// Represents a collection of functions to interact with the API endpoints
    /// </summary>
    public partial class ApplycontrollerApi : IApplycontrollerApi
    {
        private IO.Swagger.Client.ExceptionFactory _exceptionFactory = (name, response) => null;

        /// <summary>
        /// Initializes a new instance of the <see cref="ApplycontrollerApi"/> class.
        /// </summary>
        /// <returns></returns>
        public ApplycontrollerApi(String basePath)
        {
            this.Configuration = new Configuration(new ApiClient(basePath));

            ExceptionFactory = IO.Swagger.Client.Configuration.DefaultExceptionFactory;

            // ensure API client has configuration ready
            if (Configuration.ApiClient.Configuration == null)
            {
                this.Configuration.ApiClient.Configuration = this.Configuration;
            }
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="ApplycontrollerApi"/> class
        /// using Configuration object
        /// </summary>
        /// <param name="configuration">An instance of Configuration</param>
        /// <returns></returns>
        public ApplycontrollerApi(Configuration configuration = null)
        {
            if (configuration == null) // use the default one in Configuration
                this.Configuration = Configuration.Default;
            else
                this.Configuration = configuration;

            ExceptionFactory = IO.Swagger.Client.Configuration.DefaultExceptionFactory;

            // ensure API client has configuration ready
            if (Configuration.ApiClient.Configuration == null)
            {
                this.Configuration.ApiClient.Configuration = this.Configuration;
            }
        }

        /// <summary>
        /// Gets the base path of the API client.
        /// </summary>
        /// <value>The base path</value>
        public String GetBasePath()
        {
            return this.Configuration.ApiClient.RestClient.BaseUrl.ToString();
        }

        /// <summary>
        /// Sets the base path of the API client.
        /// </summary>
        /// <value>The base path</value>
        [Obsolete("SetBasePath is deprecated, please do 'Configuration.ApiClient = new ApiClient(\"http://new-path\")' instead.")]
        public void SetBasePath(String basePath)
        {
            // do nothing
        }

        /// <summary>
        /// Gets or sets the configuration object
        /// </summary>
        /// <value>An instance of the Configuration</value>
        public Configuration Configuration {get; set;}

        /// <summary>
        /// Provides a factory method hook for the creation of exceptions.
        /// </summary>
        public IO.Swagger.Client.ExceptionFactory ExceptionFactory
        {
            get
            {
                if (_exceptionFactory != null && _exceptionFactory.GetInvocationList().Length > 1)
                {
                    throw new InvalidOperationException("Multicast delegate for ExceptionFactory is unsupported.");
                }
                return _exceptionFactory;
            }
            set { _exceptionFactory = value; }
        }

        /// <summary>
        /// Gets the default header.
        /// </summary>
        /// <returns>Dictionary of HTTP header</returns>
        [Obsolete("DefaultHeader is deprecated, please use Configuration.DefaultHeader instead.")]
        public Dictionary<String, String> DefaultHeader()
        {
            return this.Configuration.DefaultHeader;
        }

        /// <summary>
        /// Add default header.
        /// </summary>
        /// <param name="key">Header field name.</param>
        /// <param name="value">Header field value.</param>
        /// <returns></returns>
        [Obsolete("AddDefaultHeader is deprecated, please use Configuration.AddDefaultHeader instead.")]
        public void AddDefaultHeader(string key, string value)
        {
            this.Configuration.AddDefaultHeader(key, value);
        }

        /// <summary>
        /// apply 
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="apply">apply</param>
        /// <returns>ResponseEntity</returns>
        public ResponseEntity ApplyUsingPOST (Apply apply)
        {
             ApiResponse<ResponseEntity> localVarResponse = ApplyUsingPOSTWithHttpInfo(apply);
             return localVarResponse.Data;
        }

        /// <summary>
        /// apply 
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="apply">apply</param>
        /// <returns>ApiResponse of ResponseEntity</returns>
        public ApiResponse< ResponseEntity > ApplyUsingPOSTWithHttpInfo (Apply apply)
        {
            // verify the required parameter 'apply' is set
            if (apply == null)
                throw new ApiException(400, "Missing required parameter 'apply' when calling ApplycontrollerApi->ApplyUsingPOST");

            var localVarPath = "/apply";
            var localVarPathParams = new Dictionary<String, String>();
            var localVarQueryParams = new Dictionary<String, String>();
            var localVarHeaderParams = new Dictionary<String, String>(Configuration.DefaultHeader);
            var localVarFormParams = new Dictionary<String, String>();
            var localVarFileParams = new Dictionary<String, FileParameter>();
            Object localVarPostBody = null;

            // to determine the Content-Type header
            String[] localVarHttpContentTypes = new String[] {
                "application/json"
            };
            String localVarHttpContentType = Configuration.ApiClient.SelectHeaderContentType(localVarHttpContentTypes);

            // to determine the Accept header
            String[] localVarHttpHeaderAccepts = new String[] {
                "*/*"
            };
            String localVarHttpHeaderAccept = Configuration.ApiClient.SelectHeaderAccept(localVarHttpHeaderAccepts);
            if (localVarHttpHeaderAccept != null)
                localVarHeaderParams.Add("Accept", localVarHttpHeaderAccept);

            if (apply != null && apply.GetType() != typeof(byte[]))
            {
                localVarPostBody = Configuration.ApiClient.Serialize(apply); // http body (model) parameter
            }
            else
            {
                localVarPostBody = apply; // byte array
            }


            // make the HTTP request
            IRestResponse localVarResponse = (IRestResponse) Configuration.ApiClient.CallApi(localVarPath,
                Method.POST, localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarFormParams, localVarFileParams,
                localVarPathParams, localVarHttpContentType);

            int localVarStatusCode = (int) localVarResponse.StatusCode;

            if (ExceptionFactory != null)
            {
                Exception exception = ExceptionFactory("ApplyUsingPOST", localVarResponse);
                if (exception != null) throw exception;
            }

            return new ApiResponse<ResponseEntity>(localVarStatusCode,
                localVarResponse.Headers.ToDictionary(x => x.Name, x => x.Value.ToString()),
                (ResponseEntity) Configuration.ApiClient.Deserialize(localVarResponse, typeof(ResponseEntity)));
        }

        /// <summary>
        /// apply 
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="apply">apply</param>
        /// <returns>Task of ResponseEntity</returns>
        public async System.Threading.Tasks.Task<ResponseEntity> ApplyUsingPOSTAsync (Apply apply)
        {
             ApiResponse<ResponseEntity> localVarResponse = await ApplyUsingPOSTAsyncWithHttpInfo(apply);
             return localVarResponse.Data;

        }

        /// <summary>
        /// apply 
        /// </summary>
        /// <exception cref="IO.Swagger.Client.ApiException">Thrown when fails to make API call</exception>
        /// <param name="apply">apply</param>
        /// <returns>Task of ApiResponse (ResponseEntity)</returns>
        public async System.Threading.Tasks.Task<ApiResponse<ResponseEntity>> ApplyUsingPOSTAsyncWithHttpInfo (Apply apply)
        {
            // verify the required parameter 'apply' is set
            if (apply == null)
                throw new ApiException(400, "Missing required parameter 'apply' when calling ApplycontrollerApi->ApplyUsingPOST");

            var localVarPath = "/apply";
            var localVarPathParams = new Dictionary<String, String>();
            var localVarQueryParams = new Dictionary<String, String>();
            var localVarHeaderParams = new Dictionary<String, String>(Configuration.DefaultHeader);
            var localVarFormParams = new Dictionary<String, String>();
            var localVarFileParams = new Dictionary<String, FileParameter>();
            Object localVarPostBody = null;

            // to determine the Content-Type header
            String[] localVarHttpContentTypes = new String[] {
                "application/json"
            };
            String localVarHttpContentType = Configuration.ApiClient.SelectHeaderContentType(localVarHttpContentTypes);

            // to determine the Accept header
            String[] localVarHttpHeaderAccepts = new String[] {
                "*/*"
            };
            String localVarHttpHeaderAccept = Configuration.ApiClient.SelectHeaderAccept(localVarHttpHeaderAccepts);
            if (localVarHttpHeaderAccept != null)
                localVarHeaderParams.Add("Accept", localVarHttpHeaderAccept);

            if (apply != null && apply.GetType() != typeof(byte[]))
            {
                localVarPostBody = Configuration.ApiClient.Serialize(apply); // http body (model) parameter
            }
            else
            {
                localVarPostBody = apply; // byte array
            }


            // make the HTTP request
            IRestResponse localVarResponse = (IRestResponse) await Configuration.ApiClient.CallApiAsync(localVarPath,
                Method.POST, localVarQueryParams, localVarPostBody, localVarHeaderParams, localVarFormParams, localVarFileParams,
                localVarPathParams, localVarHttpContentType);

            int localVarStatusCode = (int) localVarResponse.StatusCode;

            if (ExceptionFactory != null)
            {
                Exception exception = ExceptionFactory("ApplyUsingPOST", localVarResponse);
                if (exception != null) throw exception;
            }

            return new ApiResponse<ResponseEntity>(localVarStatusCode,
                localVarResponse.Headers.ToDictionary(x => x.Name, x => x.Value.ToString()),
                (ResponseEntity) Configuration.ApiClient.Deserialize(localVarResponse, typeof(ResponseEntity)));
        }

    }
}