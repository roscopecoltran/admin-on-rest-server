/**
 * DataHive RESTful APIs
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator 2.2.3.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */

/*
 * UsercontrollerApi.h
 *
 * 
 */

#ifndef UsercontrollerApi_H_
#define UsercontrollerApi_H_


#include "ApiClient.h"

#include "JwtUser.h"
#include "User.h"
#include <cpprest/details/basic_types.h>

namespace io {
namespace swagger {
namespace client {
namespace api {

using namespace io::swagger::client::model;

class  UsercontrollerApi
{
public:
    UsercontrollerApi( std::shared_ptr<ApiClient> apiClient );
    virtual ~UsercontrollerApi();
    /// <summary>
    /// addUser
    /// </summary>
    /// <remarks>
    /// 
    /// </remarks>
    /// <param name="user">user</param>
    pplx::task<std::shared_ptr<User>> addUserUsingPOST(std::shared_ptr<User> user);
    /// <summary>
    /// editField
    /// </summary>
    /// <remarks>
    /// 
    /// </remarks>
    /// <param name="id">id</param>/// <param name="user">user</param>
    pplx::task<std::shared_ptr<User>> editFieldUsingPUT1(utility::string_t id, std::shared_ptr<User> user);
    /// <summary>
    /// findUser
    /// </summary>
    /// <remarks>
    /// 
    /// </remarks>
    /// <param name="userId">userId</param>
    pplx::task<std::shared_ptr<User>> findUserUsingGET1(utility::string_t userId);
    /// <summary>
    /// getAuthenticatedUser
    /// </summary>
    /// <remarks>
    /// 
    /// </remarks>
    
    pplx::task<std::shared_ptr<JwtUser>> getAuthenticatedUserUsingGET();
    /// <summary>
    /// list
    /// </summary>
    /// <remarks>
    /// 
    /// </remarks>
    
    pplx::task<std::vector<std::shared_ptr<User>>> listUsingGET3();

protected:
    std::shared_ptr<ApiClient> m_ApiClient;
};

}
}
}
}

#endif /* UsercontrollerApi_H_ */

