# Product Service

This is a RESTful service that handles product-related operations for an e-commerce platform. It provides APIs for filtering products, retrieving product details, and performing sorting based on various criteria like price, rating, and creation timestamp.

## Features

- Filter products based on categories, price ranges, stock availability, and more.
- Sort products by price, rating, or timestamp.
- Support for both ascending and descending order for sorting.
- Handle product details retrieval.

## Technologies

- **Golang**: Backend service implemented in Go.
- **GORM**: ORM used for interacting with the database.
- **gRPC**: Used for communication between services.
- **Protocol Buffers**: Define service methods and data models.
- **SQLite**: Database for storing product information.

## Endpoints

### Product.FilterProducts

#### Parameters

##### ProductCategory
Defines the category of products for filtering.

| String Value           | Numeric Value | Description                                |
|------------------------|---------------|--------------------------------------------|
| `CATEGORY_UNSPECIFIED` | 0             | No specific category filter applied.       |
| `CATEGORY_CLOTHING`    | 1             | Filters products in the clothing category. |
| `CATEGORY_BAGS`        | 2             | Filters products in the bags category.     |

##### ProductSort
Defines the sorting criteria for the product list.

| String Value       | Numeric Value | Description                                         |
|--------------------|---------------|-----------------------------------------------------|
| `SORT_UNSPECIFIED` | 0             | No specific sorting applied (default order).        |
| `SORT_TIMESTAMP`   | 1             | Sorts products by creation timestamp.               |
| `SORT_RATING`      | 2             | Sorts products by customer rating (highest/lowest). |
| `SORT_PRICE`       | 3             | Sorts products by price (highest/lowest).           |

##### ProductStock
Defines the stock availability filter.

| String Value        | Numeric Value | Description                                  |
|---------------------|---------------|----------------------------------------------|
| `STOCK_UNSPECIFIED` | 0             | No specific stock filter applied.            |
| `STOCK_IN`          | 1             | Filters only products that are in stock.     |
| `STOCK_OUT`         | 2             | Filters only products that are out of stock. |

---

#### Request

##### Install CLI Dependencies

Install `grpcurl` to run grpc requests on terminal.

```bash

sudo curl -sSL https://github.com/fullstorydev/grpcurl/releases/latest/download/grpcurl-linux-amd64 -o /usr/local/bin/grpcurl && sudo chmod +x /usr/local/bin/grpcurl

```

**If you get error: `curl: (23) Failure writing output to destination`, add `sudo`.

##### Description

The request message for fetching a filtered list of product summaries.

| Field      | Type              | Description                                                                 |
|------------|-------------------|-----------------------------------------------------------------------------|
| `Sort`     | `ProductSort`     | Defines the sorting order (e.g., by price, rating, timestamp).              |
| `Category` | `ProductCategory` | Filters products by category (e.g., clothing, bags).                        |
| `Stock`    | `ProductStock`    | Filters products based on stock availability.                               |
| `MinPrice` | `uint32`          | Minimum price filter (inclusive).                                           |
| `MaxPrice` | `uint32`          | Maximum price filter (inclusive).                                           |
| `Desc`     | `bool`            | Determines the sorting order: `true` for descending, `false` for ascending. |


##### Commands

Lists GRPC services:

```bash

grpcurl -plaintext localhost:50051 list

```

You can also list the handlers available in the GRPC service:

```bash

grpcurl -plaintext localhost:50051 list Product

```

Description:
- Sort by: `SORT_PRICE`
- Category: `CATEGORY_BAGS`
- Stock: `STOCK_UNSPECIFIED`
- MinPrice: `0`
- MaxPrice: `0`
- Desc: `false`
- Address: `localhost:50051`
- Handlers: Product.FilterProducts

```bash

grpcurl -plaintext -d '{"Sort":3,"Category":2,"Stock":0,"MinPrice":0,"MaxPrice":0,"Desc":false}' localhost:50051 Product.FilterProducts

```

##### Example

**Request Body**:

```json
{
  "Sort": 3,
  "Category": 2,
  "Stock": 0,
  "MinPrice": 0,
  "MaxPrice": 100,
  "Desc": false
}**
```

**Response Body**:

```json
{
  "products": [
    {
      "ID": 8,
      "Name": "Messenger Bag",
      "ThumbnailURL": "https://yourblobstorageurl.blob.core.windows.net/images/messenger.jpg",
      "Price": "59.99"
    },
    {
      "ID": 7,
      "Name": "Travel Duffel Bag",
      "ThumbnailURL": "https://yourblobstorageurl.blob.core.windows.net/images/duffel.jpg",
      "Price": "69.99"
    },
  ]
}
```