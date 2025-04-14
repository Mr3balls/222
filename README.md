Логин и получение JWT
POST /auth/login
Content-Type: application/json

{
  "user_id": "12345"
}

Создание заказа
POST /orders
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "product_id": "abc123",
  "quantity": 2
}

Получение заказа по ID
GET /orders/:id
Authorization: Bearer <jwt_token>

Создание продукта
POST /products
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "name": "Product 1",
  "category": "Category 1",
  "stock": 100,
  "price": 29.99
}

Получение продукта по ID
GET /products/:id
Authorization: Bearer <jwt_token>

Список продуктов
GET /products
Authorization: Bearer <jwt_token>

