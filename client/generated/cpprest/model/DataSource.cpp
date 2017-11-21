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



#include "DataSource.h"

namespace io {
namespace swagger {
namespace client {
namespace model {

DataSource::DataSource()
{
    m_ClusterName = U("");
    m_ClusterNameIsSet = false;
    m_Current = false;
    m_CurrentIsSet = false;
    m_DbName = U("");
    m_DbNameIsSet = false;
    m_Id = U("");
    m_IdIsSet = false;
    m_IndexName = U("");
    m_IndexNameIsSet = false;
    m_JdbcUrl = U("");
    m_JdbcUrlIsSet = false;
    m_MySqlDbName = U("");
    m_MySqlDbNameIsSet = false;
    m_Password = U("");
    m_PasswordIsSet = false;
    m_Type = U("");
    m_TypeIsSet = false;
    m_Username = U("");
    m_UsernameIsSet = false;
}

DataSource::~DataSource()
{
}

void DataSource::validate()
{
    // TODO: implement validation
}

web::json::value DataSource::toJson() const
{
    web::json::value val = web::json::value::object();

    if(m_ClusterNameIsSet)
    {
        val[U("clusterName")] = ModelBase::toJson(m_ClusterName);
    }
    if(m_CurrentIsSet)
    {
        val[U("current")] = ModelBase::toJson(m_Current);
    }
    if(m_DbNameIsSet)
    {
        val[U("dbName")] = ModelBase::toJson(m_DbName);
    }
    if(m_IdIsSet)
    {
        val[U("id")] = ModelBase::toJson(m_Id);
    }
    if(m_IndexNameIsSet)
    {
        val[U("indexName")] = ModelBase::toJson(m_IndexName);
    }
    if(m_JdbcUrlIsSet)
    {
        val[U("jdbcUrl")] = ModelBase::toJson(m_JdbcUrl);
    }
    if(m_MySqlDbNameIsSet)
    {
        val[U("mySqlDbName")] = ModelBase::toJson(m_MySqlDbName);
    }
    if(m_PasswordIsSet)
    {
        val[U("password")] = ModelBase::toJson(m_Password);
    }
    if(m_TypeIsSet)
    {
        val[U("type")] = ModelBase::toJson(m_Type);
    }
    if(m_UsernameIsSet)
    {
        val[U("username")] = ModelBase::toJson(m_Username);
    }

    return val;
}

void DataSource::fromJson(web::json::value& val)
{
    if(val.has_field(U("clusterName")))
    {
        setClusterName(ModelBase::stringFromJson(val[U("clusterName")]));
    }
    if(val.has_field(U("current")))
    {
        setCurrent(ModelBase::boolFromJson(val[U("current")]));
    }
    if(val.has_field(U("dbName")))
    {
        setDbName(ModelBase::stringFromJson(val[U("dbName")]));
    }
    if(val.has_field(U("id")))
    {
        setId(ModelBase::stringFromJson(val[U("id")]));
    }
    if(val.has_field(U("indexName")))
    {
        setIndexName(ModelBase::stringFromJson(val[U("indexName")]));
    }
    if(val.has_field(U("jdbcUrl")))
    {
        setJdbcUrl(ModelBase::stringFromJson(val[U("jdbcUrl")]));
    }
    if(val.has_field(U("mySqlDbName")))
    {
        setMySqlDbName(ModelBase::stringFromJson(val[U("mySqlDbName")]));
    }
    if(val.has_field(U("password")))
    {
        setPassword(ModelBase::stringFromJson(val[U("password")]));
    }
    if(val.has_field(U("type")))
    {
        setType(ModelBase::stringFromJson(val[U("type")]));
    }
    if(val.has_field(U("username")))
    {
        setUsername(ModelBase::stringFromJson(val[U("username")]));
    }
}

void DataSource::toMultipart(std::shared_ptr<MultipartFormData> multipart, const utility::string_t& prefix) const
{
    utility::string_t namePrefix = prefix;
    if(namePrefix.size() > 0 && namePrefix[namePrefix.size() - 1] != U('.'))
    {
        namePrefix += U(".");
    }

    if(m_ClusterNameIsSet)
    {
        multipart->add(ModelBase::toHttpContent(namePrefix + U("clusterName"), m_ClusterName));
        
    }
    if(m_CurrentIsSet)
    {
        multipart->add(ModelBase::toHttpContent(namePrefix + U("current"), m_Current));
    }
    if(m_DbNameIsSet)
    {
        multipart->add(ModelBase::toHttpContent(namePrefix + U("dbName"), m_DbName));
        
    }
    if(m_IdIsSet)
    {
        multipart->add(ModelBase::toHttpContent(namePrefix + U("id"), m_Id));
        
    }
    if(m_IndexNameIsSet)
    {
        multipart->add(ModelBase::toHttpContent(namePrefix + U("indexName"), m_IndexName));
        
    }
    if(m_JdbcUrlIsSet)
    {
        multipart->add(ModelBase::toHttpContent(namePrefix + U("jdbcUrl"), m_JdbcUrl));
        
    }
    if(m_MySqlDbNameIsSet)
    {
        multipart->add(ModelBase::toHttpContent(namePrefix + U("mySqlDbName"), m_MySqlDbName));
        
    }
    if(m_PasswordIsSet)
    {
        multipart->add(ModelBase::toHttpContent(namePrefix + U("password"), m_Password));
        
    }
    if(m_TypeIsSet)
    {
        multipart->add(ModelBase::toHttpContent(namePrefix + U("type"), m_Type));
        
    }
    if(m_UsernameIsSet)
    {
        multipart->add(ModelBase::toHttpContent(namePrefix + U("username"), m_Username));
        
    }
}

void DataSource::fromMultiPart(std::shared_ptr<MultipartFormData> multipart, const utility::string_t& prefix)
{
    utility::string_t namePrefix = prefix;
    if(namePrefix.size() > 0 && namePrefix[namePrefix.size() - 1] != U('.'))
    {
        namePrefix += U(".");
    }

    if(multipart->hasContent(U("clusterName")))
    {
        setClusterName(ModelBase::stringFromHttpContent(multipart->getContent(U("clusterName"))));
    }
    if(multipart->hasContent(U("current")))
    {
        setCurrent(ModelBase::boolFromHttpContent(multipart->getContent(U("current"))));
    }
    if(multipart->hasContent(U("dbName")))
    {
        setDbName(ModelBase::stringFromHttpContent(multipart->getContent(U("dbName"))));
    }
    if(multipart->hasContent(U("id")))
    {
        setId(ModelBase::stringFromHttpContent(multipart->getContent(U("id"))));
    }
    if(multipart->hasContent(U("indexName")))
    {
        setIndexName(ModelBase::stringFromHttpContent(multipart->getContent(U("indexName"))));
    }
    if(multipart->hasContent(U("jdbcUrl")))
    {
        setJdbcUrl(ModelBase::stringFromHttpContent(multipart->getContent(U("jdbcUrl"))));
    }
    if(multipart->hasContent(U("mySqlDbName")))
    {
        setMySqlDbName(ModelBase::stringFromHttpContent(multipart->getContent(U("mySqlDbName"))));
    }
    if(multipart->hasContent(U("password")))
    {
        setPassword(ModelBase::stringFromHttpContent(multipart->getContent(U("password"))));
    }
    if(multipart->hasContent(U("type")))
    {
        setType(ModelBase::stringFromHttpContent(multipart->getContent(U("type"))));
    }
    if(multipart->hasContent(U("username")))
    {
        setUsername(ModelBase::stringFromHttpContent(multipart->getContent(U("username"))));
    }
}

utility::string_t DataSource::getClusterName() const
{
    return m_ClusterName;
}


void DataSource::setClusterName(utility::string_t value)
{
    m_ClusterName = value;
    m_ClusterNameIsSet = true;
}
bool DataSource::clusterNameIsSet() const
{
    return m_ClusterNameIsSet;
}

void DataSource::unsetClusterName()
{
    m_ClusterNameIsSet = false;
}

bool DataSource::getCurrent() const
{
    return m_Current;
}


void DataSource::setCurrent(bool value)
{
    m_Current = value;
    m_CurrentIsSet = true;
}
bool DataSource::currentIsSet() const
{
    return m_CurrentIsSet;
}

void DataSource::unsetCurrent()
{
    m_CurrentIsSet = false;
}

utility::string_t DataSource::getDbName() const
{
    return m_DbName;
}


void DataSource::setDbName(utility::string_t value)
{
    m_DbName = value;
    m_DbNameIsSet = true;
}
bool DataSource::dbNameIsSet() const
{
    return m_DbNameIsSet;
}

void DataSource::unsetDbName()
{
    m_DbNameIsSet = false;
}

utility::string_t DataSource::getId() const
{
    return m_Id;
}


void DataSource::setId(utility::string_t value)
{
    m_Id = value;
    m_IdIsSet = true;
}
bool DataSource::idIsSet() const
{
    return m_IdIsSet;
}

void DataSource::unsetId()
{
    m_IdIsSet = false;
}

utility::string_t DataSource::getIndexName() const
{
    return m_IndexName;
}


void DataSource::setIndexName(utility::string_t value)
{
    m_IndexName = value;
    m_IndexNameIsSet = true;
}
bool DataSource::indexNameIsSet() const
{
    return m_IndexNameIsSet;
}

void DataSource::unsetIndexName()
{
    m_IndexNameIsSet = false;
}

utility::string_t DataSource::getJdbcUrl() const
{
    return m_JdbcUrl;
}


void DataSource::setJdbcUrl(utility::string_t value)
{
    m_JdbcUrl = value;
    m_JdbcUrlIsSet = true;
}
bool DataSource::jdbcUrlIsSet() const
{
    return m_JdbcUrlIsSet;
}

void DataSource::unsetJdbcUrl()
{
    m_JdbcUrlIsSet = false;
}

utility::string_t DataSource::getMySqlDbName() const
{
    return m_MySqlDbName;
}


void DataSource::setMySqlDbName(utility::string_t value)
{
    m_MySqlDbName = value;
    m_MySqlDbNameIsSet = true;
}
bool DataSource::mySqlDbNameIsSet() const
{
    return m_MySqlDbNameIsSet;
}

void DataSource::unsetMySqlDbName()
{
    m_MySqlDbNameIsSet = false;
}

utility::string_t DataSource::getPassword() const
{
    return m_Password;
}


void DataSource::setPassword(utility::string_t value)
{
    m_Password = value;
    m_PasswordIsSet = true;
}
bool DataSource::passwordIsSet() const
{
    return m_PasswordIsSet;
}

void DataSource::unsetPassword()
{
    m_PasswordIsSet = false;
}

utility::string_t DataSource::getType() const
{
    return m_Type;
}


void DataSource::setType(utility::string_t value)
{
    m_Type = value;
    m_TypeIsSet = true;
}
bool DataSource::typeIsSet() const
{
    return m_TypeIsSet;
}

void DataSource::unsetType()
{
    m_TypeIsSet = false;
}

utility::string_t DataSource::getUsername() const
{
    return m_Username;
}


void DataSource::setUsername(utility::string_t value)
{
    m_Username = value;
    m_UsernameIsSet = true;
}
bool DataSource::usernameIsSet() const
{
    return m_UsernameIsSet;
}

void DataSource::unsetUsername()
{
    m_UsernameIsSet = false;
}

}
}
}
}

