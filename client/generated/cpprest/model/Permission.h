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
 * Permission.h
 *
 * 
 */

#ifndef Permission_H_
#define Permission_H_


#include "ModelBase.h"

#include <cpprest/details/basic_types.h>

namespace io {
namespace swagger {
namespace client {
namespace model {

/// <summary>
/// 
/// </summary>
class  Permission
    : public ModelBase
{
public:
    Permission();
    virtual ~Permission();

    /////////////////////////////////////////////
    /// ModelBase overrides

    void validate() override;

    web::json::value toJson() const override;
    void fromJson(web::json::value& json) override;

    void toMultipart(std::shared_ptr<MultipartFormData> multipart, const utility::string_t& namePrefix) const override;
    void fromMultiPart(std::shared_ptr<MultipartFormData> multipart, const utility::string_t& namePrefix) override;

    /////////////////////////////////////////////
    /// Permission members

    /// <summary>
    /// 
    /// </summary>
    bool getC() const;
    bool CIsSet() const;
    void unsetc();
    void setC(bool value);
    /// <summary>
    /// 
    /// </summary>
    bool getD() const;
    bool DIsSet() const;
    void unsetd();
    void setD(bool value);
    /// <summary>
    /// 
    /// </summary>
    utility::string_t getEid() const;
    bool eidIsSet() const;
    void unsetEid();
    void setEid(utility::string_t value);
    /// <summary>
    /// 
    /// </summary>
    utility::string_t getId() const;
    bool idIsSet() const;
    void unsetId();
    void setId(utility::string_t value);
    /// <summary>
    /// 
    /// </summary>
    bool getR() const;
    bool RIsSet() const;
    void unsetr();
    void setR(bool value);
    /// <summary>
    /// 
    /// </summary>
    utility::string_t getRoleId() const;
    bool roleIdIsSet() const;
    void unsetRoleId();
    void setRoleId(utility::string_t value);
    /// <summary>
    /// 
    /// </summary>
    bool getU() const;
    bool UIsSet() const;
    void unsetu();
    void setU(bool value);

protected:
    bool m_c;
    bool m_cIsSet;
    bool m_d;
    bool m_dIsSet;
    utility::string_t m_Eid;
    bool m_EidIsSet;
    utility::string_t m_Id;
    bool m_IdIsSet;
    bool m_r;
    bool m_rIsSet;
    utility::string_t m_RoleId;
    bool m_RoleIdIsSet;
    bool m_u;
    bool m_uIsSet;
};

}
}
}
}

#endif /* Permission_H_ */
