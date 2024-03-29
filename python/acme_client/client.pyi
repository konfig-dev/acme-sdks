# coding: utf-8
"""
    OpenAPI Petstore

    This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.

    The version of the OpenAPI document: 1.0.0
    Generated by: https://konfigthis.com
"""

import typing
import inspect
from datetime import date, datetime
from acme_client.client_custom import ClientCustom
from acme_client.configuration import Configuration
from acme_client.api_client import ApiClient
from acme_client.type_util import copy_signature
from acme_client.apis.tags.miscellaneous_api import MiscellaneousApi
from acme_client.apis.tags.pet_api import PetApi
from acme_client.apis.tags.store_api import StoreApi
from acme_client.apis.tags.user_api import UserApi



class Acme(ClientCustom):

    def __init__(self, configuration: typing.Union[Configuration, None] = None, **kwargs):
        super().__init__(configuration, **kwargs)
        if (len(kwargs) > 0):
            configuration = Configuration(**kwargs)
        if (configuration is None):
            raise Exception("configuration is required")
        api_client = ApiClient(configuration)
        self.miscellaneous: MiscellaneousApi = MiscellaneousApi(api_client)
        self.pet: PetApi = PetApi(api_client)
        self.store: StoreApi = StoreApi(api_client)
        self.user: UserApi = UserApi(api_client)
