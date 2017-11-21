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
 * GrantedAuthority.h
 *
 * 
 */

#ifndef GrantedAuthority_H_
#define GrantedAuthority_H_


#include "ModelBase.h"

#include <cpprest/details/basic_types.h>

namespace io {
namespace swagger {
namespace client {
namespace model {

/// <summary>
/// 
/// </summary>
class  GrantedAuthority
    : public ModelBase
{
public:
    GrantedAuthority();
    virtual ~GrantedAuthority();

    /////////////////////////////////////////////
    /// ModelBase overrides

    void validate() override;

    web::json::value toJson() const override;
    void fromJson(web::json::value& json) override;

    void toMultipart(std::shared_ptr<MultipartFormData> multipart, const utility::string_t& namePrefix) const override;
    void fromMultiPart(std::shared_ptr<MultipartFormData> multipart, const utility::string_t& namePrefix) override;

    /////////////////////////////////////////////
    /// GrantedAuthority members

    /// <summary>
    /// 
    /// </summary>
    utility::string_t getAuthority() const;
    bool authorityIsSet() const;
    void unsetAuthority();
    void setAuthority(utility::string_t value);

protected:
    utility::string_t m_Authority;
    bool m_AuthorityIsSet;
};

}
}
}
}

#endif /* GrantedAuthority_H_ */