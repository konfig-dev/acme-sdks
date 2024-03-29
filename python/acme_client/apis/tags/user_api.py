# coding: utf-8

"""
    OpenAPI Petstore

    This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.

    The version of the OpenAPI document: 1.0.0
    Generated by: https://konfigthis.com
"""

from acme_client.paths.user.post import Create
from acme_client.paths.user_create_with_array.post import CreateWithArray
from acme_client.paths.user_create_with_list.post import CreateWithList
from acme_client.paths.user_username.delete import Delete
from acme_client.paths.user_username.get import GetByName
from acme_client.paths.user_login.get import Login
from acme_client.paths.user_logout.get import Logout
from acme_client.paths.user_username.put import Update


class UserApi(
    Create,
    CreateWithArray,
    CreateWithList,
    Delete,
    GetByName,
    Login,
    Logout,
    Update,
):
    """NOTE:
    This class is auto generated by Konfig (https://konfigthis.com)
    """
    pass
