# coding: utf-8

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from flyteadmin.models.sql_dialect import SqlDialect  # noqa: F401,E501


class CoreSql(object):
    """NOTE: This class is auto generated by the swagger code generator program.

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
        'statement': 'str',
        'dialect': 'SqlDialect'
    }

    attribute_map = {
        'statement': 'statement',
        'dialect': 'dialect'
    }

    def __init__(self, statement=None, dialect=None):  # noqa: E501
        """CoreSql - a model defined in Swagger"""  # noqa: E501

        self._statement = None
        self._dialect = None
        self.discriminator = None

        if statement is not None:
            self.statement = statement
        if dialect is not None:
            self.dialect = dialect

    @property
    def statement(self):
        """Gets the statement of this CoreSql.  # noqa: E501


        :return: The statement of this CoreSql.  # noqa: E501
        :rtype: str
        """
        return self._statement

    @statement.setter
    def statement(self, statement):
        """Sets the statement of this CoreSql.


        :param statement: The statement of this CoreSql.  # noqa: E501
        :type: str
        """

        self._statement = statement

    @property
    def dialect(self):
        """Gets the dialect of this CoreSql.  # noqa: E501


        :return: The dialect of this CoreSql.  # noqa: E501
        :rtype: SqlDialect
        """
        return self._dialect

    @dialect.setter
    def dialect(self, dialect):
        """Sets the dialect of this CoreSql.


        :param dialect: The dialect of this CoreSql.  # noqa: E501
        :type: SqlDialect
        """

        self._dialect = dialect

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
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
        if issubclass(CoreSql, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, CoreSql):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
