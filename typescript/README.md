<div align="center">

[![Visit Acme](./header.png)](https://konfigthis.com)

# [Acme](https://konfigthis.com)<a id="acme"></a>

This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.

[![npm](https://img.shields.io/badge/npm-v1.0.7-blue)](https://www.npmjs.com/package/acme-typescript-sdk/v/1.0.7)

</div>

## Table of Contents<a id="table-of-contents"></a>

<!-- toc -->

- [Installation](#installation)
- [Getting Started](#getting-started)
- [Reference](#reference)
  * [`acme.miscellaneous.paginate`](#acmemiscellaneouspaginate)
  * [`acme.pet.add`](#acmepetadd)
  * [`acme.pet.delete`](#acmepetdelete)
  * [`acme.pet.findByStatus`](#acmepetfindbystatus)
  * [`acme.pet.findByTags`](#acmepetfindbytags)
  * [`acme.pet.getById`](#acmepetgetbyid)
  * [`acme.pet.update`](#acmepetupdate)
  * [`acme.pet.updateWithForm`](#acmepetupdatewithform)
  * [`acme.pet.uploadImage`](#acmepetuploadimage)
  * [`acme.store.deleteOrder`](#acmestoredeleteorder)
  * [`acme.store.getInventory`](#acmestoregetinventory)
  * [`acme.store.getOrderById`](#acmestoregetorderbyid)
  * [`acme.store.placeOrder`](#acmestoreplaceorder)
  * [`acme.user.create`](#acmeusercreate)
  * [`acme.user.createWithArray`](#acmeusercreatewitharray)
  * [`acme.user.createWithList`](#acmeusercreatewithlist)
  * [`acme.user.delete`](#acmeuserdelete)
  * [`acme.user.getByName`](#acmeusergetbyname)
  * [`acme.user.login`](#acmeuserlogin)
  * [`acme.user.logout`](#acmeuserlogout)
  * [`acme.user.update`](#acmeuserupdate)

<!-- tocstop -->

## Installation<a id="installation"></a>

<table>
<tr>
<th width="292px"><code>npm</code></th>
<th width="293px"><code>pnpm</code></th>
<th width="292px"><code>yarn</code></th>
</tr>
<tr>
<td>

```bash
npm i acme-typescript-sdk
```

</td>
<td>

```bash
pnpm i acme-typescript-sdk
```

</td>
<td>

```bash
yarn add acme-typescript-sdk
```

</td>
</tr>
</table>

## Getting Started<a id="getting-started"></a>

```typescript
import { Acme } from "acme-typescript-sdk";

const acme = new Acme({
  // Defining the base path is optional and defaults to https://petstore.swagger.io/v2
  // basePath: "https://petstore.swagger.io/v2",
  apiKey: "API_KEY",
  accessToken: "ACCESS_TOKEN",
});

const paginateResponse = await acme.miscellaneous.paginate({});

console.log(paginateResponse);
```

## Reference<a id="reference"></a>


### `acme.miscellaneous.paginate`<a id="acmemiscellaneouspaginate"></a>

Iterate through a bunch of items

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const paginateResponse = await acme.miscellaneous.paginate({});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### first: `number`<a id="first-number"></a>

##### after: `string`<a id="after-string"></a>

#### 🔄 Return<a id="🔄-return"></a>

[PaginateResponse](./models/paginate-response.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/pagination` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `acme.pet.add`<a id="acmepetadd"></a>



#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const addResponse = await acme.pet.add({
  name: "doggie",
  photoUrls: ["photoUrls_example"],
  status: "available",
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### name: `string`<a id="name-string"></a>

##### photoUrls: `string`[]<a id="photourls-string"></a>

##### tags: [`Tag`](./models/tag.ts)[]<a id="tags-tagmodelstagts"></a>

##### id: `number`<a id="id-number"></a>

##### category: [`Category`](./models/category.ts)<a id="category-categorymodelscategoryts"></a>

##### status: `string`<a id="status-string"></a>

pet status in the store

#### 🔄 Return<a id="🔄-return"></a>

[Pet](./models/pet.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/pet` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `acme.pet.delete`<a id="acmepetdelete"></a>



#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const deleteResponse = await acme.pet.delete({
  petId: 1,
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### petId: `number`<a id="petid-number"></a>

Pet id to delete

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/pet/{petId}` `DELETE`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `acme.pet.findByStatus`<a id="acmepetfindbystatus"></a>

Multiple status values can be provided with comma separated strings

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const findByStatusResponse = await acme.pet.findByStatus({
  status: ["available"],
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### status: `string`[]<a id="status-string"></a>

Status values that need to be considered for filter

#### 🔄 Return<a id="🔄-return"></a>

[Pet](./models/pet.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/pet/findByStatus` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `acme.pet.findByTags`<a id="acmepetfindbytags"></a>

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const findByTagsResponse = await acme.pet.findByTags({
  tags: ["tags_example"],
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### tags: `string`[]<a id="tags-string"></a>

Tags to filter by

#### 🔄 Return<a id="🔄-return"></a>

[Pet](./models/pet.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/pet/findByTags` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `acme.pet.getById`<a id="acmepetgetbyid"></a>

Returns a single pet

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getByIdResponse = await acme.pet.getById({
  petId: 1,
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### petId: `number`<a id="petid-number"></a>

ID of pet to return

#### 🔄 Return<a id="🔄-return"></a>

[Pet](./models/pet.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/pet/{petId}` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `acme.pet.update`<a id="acmepetupdate"></a>



#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const updateResponse = await acme.pet.update({
  name: "doggie",
  photoUrls: ["photoUrls_example"],
  status: "available",
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### name: `string`<a id="name-string"></a>

##### photoUrls: `string`[]<a id="photourls-string"></a>

##### tags: [`Tag`](./models/tag.ts)[]<a id="tags-tagmodelstagts"></a>

##### id: `number`<a id="id-number"></a>

##### category: [`Category`](./models/category.ts)<a id="category-categorymodelscategoryts"></a>

##### status: `string`<a id="status-string"></a>

pet status in the store

#### 🔄 Return<a id="🔄-return"></a>

[Pet](./models/pet.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/pet` `PUT`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `acme.pet.updateWithForm`<a id="acmepetupdatewithform"></a>



#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const updateWithFormResponse = await acme.pet.updateWithForm({
  petId: 1,
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### petId: `number`<a id="petid-number"></a>

ID of pet that needs to be updated

##### name: `string`<a id="name-string"></a>

Updated name of the pet

##### status: `string`<a id="status-string"></a>

Updated status of the pet

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/pet/{petId}` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `acme.pet.uploadImage`<a id="acmepetuploadimage"></a>



#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const uploadImageResponse = await acme.pet.uploadImage({
  petId: 1,
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### petId: `number`<a id="petid-number"></a>

ID of pet to update

##### additionalMetadata: `string`<a id="additionalmetadata-string"></a>

Additional data to pass to server

##### file: `Uint8Array | File | buffer.File`<a id="file-uint8array--file--bufferfile"></a>

file to upload

#### 🔄 Return<a id="🔄-return"></a>

[ApiResponse](./models/api-response.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/pet/{petId}/uploadImage` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `acme.store.deleteOrder`<a id="acmestoredeleteorder"></a>

For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const deleteOrderResponse = await acme.store.deleteOrder({
  orderId: "orderId_example",
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### orderId: `string`<a id="orderid-string"></a>

ID of the order that needs to be deleted

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/store/order/{orderId}` `DELETE`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `acme.store.getInventory`<a id="acmestoregetinventory"></a>

Returns a map of status codes to quantities

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getInventoryResponse = await acme.store.getInventory();
```

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/store/inventory` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `acme.store.getOrderById`<a id="acmestoregetorderbyid"></a>

For valid response try integer IDs with value <= 5 or > 10. Other values will generated exceptions

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getOrderByIdResponse = await acme.store.getOrderById({
  orderId: 1,
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### orderId: `number`<a id="orderid-number"></a>

ID of pet that needs to be fetched

#### 🔄 Return<a id="🔄-return"></a>

[Order](./models/order.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/store/order/{orderId}` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `acme.store.placeOrder`<a id="acmestoreplaceorder"></a>



#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const placeOrderResponse = await acme.store.placeOrder({
  status: "placed",
  complete: false,
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### id: `number`<a id="id-number"></a>

##### petId: `number`<a id="petid-number"></a>

##### quantity: `number`<a id="quantity-number"></a>

##### shipDate: `string`<a id="shipdate-string"></a>

##### status: `string`<a id="status-string"></a>

Order Status

##### complete: `boolean`<a id="complete-boolean"></a>

#### 🔄 Return<a id="🔄-return"></a>

[Order](./models/order.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/store/order` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `acme.user.create`<a id="acmeusercreate"></a>

This can only be done by the logged in user.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const createResponse = await acme.user.create({});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### id: `number`<a id="id-number"></a>

##### username: `string`<a id="username-string"></a>

##### firstName: `string`<a id="firstname-string"></a>

##### lastName: `string`<a id="lastname-string"></a>

##### email: `string`<a id="email-string"></a>

##### password: `string`<a id="password-string"></a>

##### phone: `string`<a id="phone-string"></a>

##### userStatus: `number`<a id="userstatus-number"></a>

User Status

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/user` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `acme.user.createWithArray`<a id="acmeusercreatewitharray"></a>



#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const createWithArrayResponse = await acme.user.createWithArray([{}]);
```

#### ⚙️ Request Body<a id="⚙️-request-body"></a>

[`User`](./models/user.ts)[]

List of user object

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/user/createWithArray` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `acme.user.createWithList`<a id="acmeusercreatewithlist"></a>



#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const createWithListResponse = await acme.user.createWithList([{}]);
```

#### ⚙️ Request Body<a id="⚙️-request-body"></a>

[`User`](./models/user.ts)[]

List of user object

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/user/createWithList` `POST`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `acme.user.delete`<a id="acmeuserdelete"></a>

This can only be done by the logged in user.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const deleteResponse = await acme.user.delete({
  username: "username_example",
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### username: `string`<a id="username-string"></a>

The name that needs to be deleted

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/user/{username}` `DELETE`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `acme.user.getByName`<a id="acmeusergetbyname"></a>



#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const getByNameResponse = await acme.user.getByName({
  username: "username_example",
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### username: `string`<a id="username-string"></a>

The name that needs to be fetched. Use user1 for testing.

#### 🔄 Return<a id="🔄-return"></a>

[User](./models/user.ts)

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/user/{username}` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `acme.user.login`<a id="acmeuserlogin"></a>



#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const loginResponse = await acme.user.login({
  username:
    "CbUUGjjNSwg0_bs9ZayIMrKdgNvb6gvxmPb9GcsM61ate1RA89q3w1l4eH4XxEz.5awLMdeXylwK0lMGUSM4jsrh4dstlnQUN5vVdMLPA",
  password: "password_example",
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### username: `string`<a id="username-string"></a>

The user name for login

##### password: `string`<a id="password-string"></a>

The password for login in clear text

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/user/login` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `acme.user.logout`<a id="acmeuserlogout"></a>



#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const logoutResponse = await acme.user.logout();
```

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/user/logout` `GET`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


### `acme.user.update`<a id="acmeuserupdate"></a>

This can only be done by the logged in user.

#### 🛠️ Usage<a id="🛠️-usage"></a>

```typescript
const updateResponse = await acme.user.update({
  username: "username_example",
});
```

#### ⚙️ Parameters<a id="⚙️-parameters"></a>

##### username: `string`<a id="username-string"></a>

name that need to be deleted

##### id: `number`<a id="id-number"></a>

##### firstName: `string`<a id="firstname-string"></a>

##### lastName: `string`<a id="lastname-string"></a>

##### email: `string`<a id="email-string"></a>

##### password: `string`<a id="password-string"></a>

##### phone: `string`<a id="phone-string"></a>

##### userStatus: `number`<a id="userstatus-number"></a>

User Status

#### 🌐 Endpoint<a id="🌐-endpoint"></a>

`/user/{username}` `PUT`

[🔙 **Back to Table of Contents**](#table-of-contents)

---


## Author<a id="author"></a>
This TypeScript package is automatically generated by [Konfig](https://konfigthis.com)
