# Getting Started

This is an interactive tutorial the Pet Store API.

## 1) Log user into the system

::::form

::input{name="KEY" label="API Key" placeholder="Your API Key"}
::input{name=USERNAME label="Username" placeholder="Username"}
::input{name=PASSWORD label="Password" placeholder="Password"}

```python
from pprint import pprint
from petstore import Petstore

pestore = Petstore()

response = pestore.user.login(
    username="USERNAME",
    password="PASSWORD"
)
pprint(response.body)
```

::button[Login]
::::