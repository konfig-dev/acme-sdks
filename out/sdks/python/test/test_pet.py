"""
    OpenAPI Petstore

    This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import acme_client
from acme_client.model.category import Category
from acme_client.model.tag import Tag
globals()['Category'] = Category
globals()['Tag'] = Tag
from acme_client.model.pet import Pet


class TestPet(unittest.TestCase):
    """Pet unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testPet(self):
        """Test Pet"""
        # FIXME: construct object with mandatory attributes with example values
        # model = Pet()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
