import unittest
from pprint import pprint

import acme_client
from acme_client.api import pet_api
from acme_client.model.api_response import ApiResponse
from acme_client.model.pet import Pet
from acme_client.model.category import Category
from acme_client.model.tag import Tag


class Test(unittest.TestCase):
    """ApiResponse unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def test(self):
        configuration = acme_client.Configuration(
            host = "http://petstore.swagger.io/v2"
        )
        configuration.access_token = 'YOUR_ACCESS_TOKEN'


        # Enter a context with an instance of the API client
        with acme_client.ApiClient(configuration) as api_client:
            api_instance = pet_api.PetApi(api_client)
            pet = Pet(
                id=1,
                category=Category(
                    id=1,
                    name="CbUUGjjNSwg0_bs9ZayIMrKdgNvb6gvxmPb9GcsM61ate1RA89q3w1l4eH4XxEz.5awLMdeXylwK0lMGUSM4jsrh4dstlnQUN5vVdMLPA",
                ),
                name="doggie",
                photo_urls=[
                    "photo_urls_example",
                ],
                tags=[
                    Tag(
                        id=1,
                        name="name_example",
                    ),
                ],
                status="available",
            ) # Pet | Pet object that needs to be added to the store

            try:
                # Add a new pet to the store
                api_response = api_instance.add_pet(pet)
                pprint(api_response)
            except acme_client.ApiException as e:
                print("Exception when calling PetApi->add_pet: %s\n" % e)



if __name__ == '__main__':
    unittest.main()
