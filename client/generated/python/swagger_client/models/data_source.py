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


class DataSource(object):
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
        'cluster_name': 'str',
        'current': 'bool',
        'db_name': 'str',
        'id': 'str',
        'index_name': 'str',
        'jdbc_url': 'str',
        'my_sql_db_name': 'str',
        'password': 'str',
        'type': 'str',
        'username': 'str'
    }

    attribute_map = {
        'cluster_name': 'clusterName',
        'current': 'current',
        'db_name': 'dbName',
        'id': 'id',
        'index_name': 'indexName',
        'jdbc_url': 'jdbcUrl',
        'my_sql_db_name': 'mySqlDbName',
        'password': 'password',
        'type': 'type',
        'username': 'username'
    }

    def __init__(self, cluster_name=None, current=None, db_name=None, id=None, index_name=None, jdbc_url=None, my_sql_db_name=None, password=None, type=None, username=None):
        """
        DataSource - a model defined in Swagger
        """

        self._cluster_name = None
        self._current = None
        self._db_name = None
        self._id = None
        self._index_name = None
        self._jdbc_url = None
        self._my_sql_db_name = None
        self._password = None
        self._type = None
        self._username = None

        if cluster_name is not None:
          self.cluster_name = cluster_name
        if current is not None:
          self.current = current
        if db_name is not None:
          self.db_name = db_name
        if id is not None:
          self.id = id
        if index_name is not None:
          self.index_name = index_name
        if jdbc_url is not None:
          self.jdbc_url = jdbc_url
        if my_sql_db_name is not None:
          self.my_sql_db_name = my_sql_db_name
        if password is not None:
          self.password = password
        if type is not None:
          self.type = type
        if username is not None:
          self.username = username

    @property
    def cluster_name(self):
        """
        Gets the cluster_name of this DataSource.

        :return: The cluster_name of this DataSource.
        :rtype: str
        """
        return self._cluster_name

    @cluster_name.setter
    def cluster_name(self, cluster_name):
        """
        Sets the cluster_name of this DataSource.

        :param cluster_name: The cluster_name of this DataSource.
        :type: str
        """

        self._cluster_name = cluster_name

    @property
    def current(self):
        """
        Gets the current of this DataSource.

        :return: The current of this DataSource.
        :rtype: bool
        """
        return self._current

    @current.setter
    def current(self, current):
        """
        Sets the current of this DataSource.

        :param current: The current of this DataSource.
        :type: bool
        """

        self._current = current

    @property
    def db_name(self):
        """
        Gets the db_name of this DataSource.

        :return: The db_name of this DataSource.
        :rtype: str
        """
        return self._db_name

    @db_name.setter
    def db_name(self, db_name):
        """
        Sets the db_name of this DataSource.

        :param db_name: The db_name of this DataSource.
        :type: str
        """

        self._db_name = db_name

    @property
    def id(self):
        """
        Gets the id of this DataSource.

        :return: The id of this DataSource.
        :rtype: str
        """
        return self._id

    @id.setter
    def id(self, id):
        """
        Sets the id of this DataSource.

        :param id: The id of this DataSource.
        :type: str
        """

        self._id = id

    @property
    def index_name(self):
        """
        Gets the index_name of this DataSource.

        :return: The index_name of this DataSource.
        :rtype: str
        """
        return self._index_name

    @index_name.setter
    def index_name(self, index_name):
        """
        Sets the index_name of this DataSource.

        :param index_name: The index_name of this DataSource.
        :type: str
        """

        self._index_name = index_name

    @property
    def jdbc_url(self):
        """
        Gets the jdbc_url of this DataSource.

        :return: The jdbc_url of this DataSource.
        :rtype: str
        """
        return self._jdbc_url

    @jdbc_url.setter
    def jdbc_url(self, jdbc_url):
        """
        Sets the jdbc_url of this DataSource.

        :param jdbc_url: The jdbc_url of this DataSource.
        :type: str
        """

        self._jdbc_url = jdbc_url

    @property
    def my_sql_db_name(self):
        """
        Gets the my_sql_db_name of this DataSource.

        :return: The my_sql_db_name of this DataSource.
        :rtype: str
        """
        return self._my_sql_db_name

    @my_sql_db_name.setter
    def my_sql_db_name(self, my_sql_db_name):
        """
        Sets the my_sql_db_name of this DataSource.

        :param my_sql_db_name: The my_sql_db_name of this DataSource.
        :type: str
        """

        self._my_sql_db_name = my_sql_db_name

    @property
    def password(self):
        """
        Gets the password of this DataSource.

        :return: The password of this DataSource.
        :rtype: str
        """
        return self._password

    @password.setter
    def password(self, password):
        """
        Sets the password of this DataSource.

        :param password: The password of this DataSource.
        :type: str
        """

        self._password = password

    @property
    def type(self):
        """
        Gets the type of this DataSource.

        :return: The type of this DataSource.
        :rtype: str
        """
        return self._type

    @type.setter
    def type(self, type):
        """
        Sets the type of this DataSource.

        :param type: The type of this DataSource.
        :type: str
        """
        allowed_values = ["mongo", "mysql", "cds", "es"]
        if type not in allowed_values:
            raise ValueError(
                "Invalid value for `type` ({0}), must be one of {1}"
                .format(type, allowed_values)
            )

        self._type = type

    @property
    def username(self):
        """
        Gets the username of this DataSource.

        :return: The username of this DataSource.
        :rtype: str
        """
        return self._username

    @username.setter
    def username(self, username):
        """
        Sets the username of this DataSource.

        :param username: The username of this DataSource.
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
        if not isinstance(other, DataSource):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """
        Returns true if both objects are not equal
        """
        return not self == other