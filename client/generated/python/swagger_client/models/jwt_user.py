# coding: utf-8

"""
    DataHive RESTful APIs

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

    OpenAPI spec version: 1.0
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


from pprint import pformat
from six import iteritems
import re


class JwtUser(object):
    """
    NOTE: This class is auto generated by the swagger code generator program.
    Do not edit the class manually.
    """


    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'authorities': 'list[GrantedAuthority]',
        'email': 'str',
        'enabled': 'bool',
        'username': 'str'
    }

    attribute_map = {
        'authorities': 'authorities',
        'email': 'email',
        'enabled': 'enabled',
        'username': 'username'
    }

    def __init__(self, authorities=None, email=None, enabled=None, username=None):
        """
        JwtUser - a model defined in Swagger
        """

        self._authorities = None
        self._email = None
        self._enabled = None
        self._username = None

        if authorities is not None:
          self.authorities = authorities
        if email is not None:
          self.email = email
        if enabled is not None:
          self.enabled = enabled
        if username is not None:
          self.username = username

    @property
    def authorities(self):
        """
        Gets the authorities of this JwtUser.

        :return: The authorities of this JwtUser.
        :rtype: list[GrantedAuthority]
        """
        return self._authorities

    @authorities.setter
    def authorities(self, authorities):
        """
        Sets the authorities of this JwtUser.

        :param authorities: The authorities of this JwtUser.
        :type: list[GrantedAuthority]
        """

        self._authorities = authorities

    @property
    def email(self):
        """
        Gets the email of this JwtUser.

        :return: The email of this JwtUser.
        :rtype: str
        """
        return self._email

    @email.setter
    def email(self, email):
        """
        Sets the email of this JwtUser.

        :param email: The email of this JwtUser.
        :type: str
        """

        self._email = email

    @property
    def enabled(self):
        """
        Gets the enabled of this JwtUser.

        :return: The enabled of this JwtUser.
        :rtype: bool
        """
        return self._enabled

    @enabled.setter
    def enabled(self, enabled):
        """
        Sets the enabled of this JwtUser.

        :param enabled: The enabled of this JwtUser.
        :type: bool
        """

        self._enabled = enabled

    @property
    def username(self):
        """
        Gets the username of this JwtUser.

        :return: The username of this JwtUser.
        :rtype: str
        """
        return self._username

    @username.setter
    def username(self, username):
        """
        Sets the username of this JwtUser.

        :param username: The username of this JwtUser.
        :type: str
        """

        self._username = username

    def to_dict(self):
        """
        Returns the model properties as a dict
        """
        result = {}

        for attr, _ in iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """
        Returns the string representation of the model
        """
        return pformat(self.to_dict())

    def __repr__(self):
        """
        For `print` and `pprint`
        """
        return self.to_str()

    def __eq__(self, other):
        """
        Returns true if both objects are equal
        """
        if not isinstance(other, JwtUser):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """
        Returns true if both objects are not equal
        """
        return not self == other
