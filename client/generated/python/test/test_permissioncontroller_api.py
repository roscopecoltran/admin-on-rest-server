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
from swagger_client.apis.permissioncontroller_api import PermissioncontrollerApi


class TestPermissioncontrollerApi(unittest.TestCase):
    """ PermissioncontrollerApi unit test stubs """

    def setUp(self):
        self.api = swagger_client.apis.permissioncontroller_api.PermissioncontrollerApi()

    def tearDown(self):
        pass

    def test_add_permission_using_post(self):
        """
        Test case for add_permission_using_post

        addPermission
        """
        pass

    def test_edit_permission_using_put(self):
        """
        Test case for edit_permission_using_put

        editPermission
        """
        pass

    def test_find_user_using_get(self):
        """
        Test case for find_user_using_get

        findUser
        """
        pass

    def test_list_using_get1(self):
        """
        Test case for list_using_get1

        list
        """
        pass


if __name__ == '__main__':
    unittest.main()
