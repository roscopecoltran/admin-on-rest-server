# coding: utf-8

"""
    DataHive RESTful APIs

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

    OpenAPI spec version: 1.0
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


from __future__ import absolute_import

import os
import sys
import unittest

import swagger_client
from swagger_client.rest import ApiException
from swagger_client.apis.rolecontroller_api import RolecontrollerApi


class TestRolecontrollerApi(unittest.TestCase):
    """ RolecontrollerApi unit test stubs """

    def setUp(self):
        self.api = swagger_client.apis.rolecontroller_api.RolecontrollerApi()

    def tearDown(self):
        pass

    def test_add_role_using_post(self):
        """
        Test case for add_role_using_post

        addRole
        """
        pass

    def test_edit_role_using_put(self):
        """
        Test case for edit_role_using_put

        editRole
        """
        pass

    def test_find_role_using_get1(self):
        """
        Test case for find_role_using_get1

        findRole
        """
        pass

    def test_list_using_get2(self):
        """
        Test case for list_using_get2

        list
        """
        pass


if __name__ == '__main__':
    unittest.main()