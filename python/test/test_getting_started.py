import unittest
from pprint import pprint

import acme_client
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
        client = acme_client.ApiClient(configuration=acme_client.Configuration(
            host="http://petstore.swagger.io/v2",
            access_token='YOUR_ACCESS_TOKEN'
        ))

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
        )  # Pet | Pet object that needs to be added to the store

        try:
            # Add a new pet to the store
            api_response = client.pet.add_pet(pet)
            pprint(api_response)
        except acme_client.ApiException as e:
            print("Exception when calling PetApi->add_pet: %s\n" % e)


if __name__ == '__main__':
    unittest.main()
