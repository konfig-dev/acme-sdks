
# flake8: noqa

# Import all APIs into this package.
# If you have many APIs here with many many models used in each API this may
# raise a `RecursionError`.
# In order to avoid this, import only the API that you directly need like:
#
#   from acme_client.api.pet_api import PetApi
#
# or import this package, but before doing it, use:
#
#   import sys
#   sys.setrecursionlimit(n)

# Import APIs into API package:
from acme_client.api.pet_api import PetApi
from acme_client.api.store_api import StoreApi
from acme_client.api.user_api import UserApi

