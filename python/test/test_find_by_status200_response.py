"""
    OpenAPI Petstore

    This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://konfigthis.com
"""


import sys
import unittest

import acme_client
from acme_client.model.pet import Pet
globals()['Pet'] = Pet
from acme_client.model.find_by_status200_response import FindByStatus200Response


class TestFindByStatus200Response(unittest.TestCase):
    """FindByStatus200Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testFindByStatus200Response(self):
        """Test FindByStatus200Response"""
        # FIXME: construct object with mandatory attributes with example values
        # model = FindByStatus200Response()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()