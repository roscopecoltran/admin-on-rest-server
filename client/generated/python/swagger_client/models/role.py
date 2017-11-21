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


class Role(object):
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
        'id': 'str',
        'name': 'str',
        'users': 'list[User]'
    }

    attribute_map = {
        'id': 'id',
        'name': 'name',
        'users': 'users'
    }

    def __init__(self, id=None, name=None, users=None):
        """
        Role - a model defined in Swagger
        """

        self._id = None
        self._name = None
        self._users = None

        if id is not None:
          self.id = id
        if name is not None:
          self.name = name
        if users is not None:
          self.users = users

    @property
    def id(self):
        """
        Gets the id of this Role.

        :return: The id of this Role.
        :rtype: str
        """
        return self._id

    @id.setter
    def id(self, id):
        """
        Sets the id of this Role.

        :param id: The id of this Role.
        :type: str
        """

        self._id = id

    @property
    def name(self):
        """
        Gets the name of this Role.

        :return: The name of this Role.
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """
        Sets the name of this Role.

        :param name: The name of this Role.
        :type: str
        """

        self._name = name

    @property
    def users(self):
        """
        Gets the users of this Role.

        :return: The users of this Role.
        :rtype: list[User]
        """
        return self._users

    @users.setter
    def users(self, users):
        """
        Sets the users of this Role.

        :param users: The users of this Role.
        :type: list[User]
        """

        self._users = users

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
        if not isinstance(other, Role):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """
        Returns true if both objects are not equal
        """
        return not self == other