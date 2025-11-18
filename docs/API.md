# API Reference

Base URL: `/api`

## Auth
- `POST /api/auth/register`
  - Body: `{ email, password, full_name }`
  - 201: `{ id, email }`
- `POST /api/auth/login`
  - Body: `{ email, password }`
  - 200: `{ token }`

## Products
- `GET /api/products`
  - Query: `page`, `limit`, `category`, `search`
  - 200: `{ data: Product[], page, limit, total }`
- `GET /api/products/:id`
  - 200: `Product`

Admin (requires auth + admin):
- `POST /api/admin/products`
  - Body: `{ name, description, price, stock, category, sku, discount?, image? }`
  - 201: `Product`
- `PUT /api/admin/products/:id`
  - Body: same as create
  - 200: `Product`
- `DELETE /api/admin/products/:id`
  - 200: `{ message }`

## Cart (requires auth)
- `GET /api/cart`
  - 200: `{ data: CartItem[] }`
- `POST /api/cart/add`
  - Body: `{ product_id, quantity }`
  - 201: `CartItem`
- `DELETE /api/cart/:product_id`
  - 200: `{ message }`

## Orders (requires auth)
- `POST /api/orders`
  - Body: `{ items: [{ product_id, quantity }...], shipping_address }`
  - 201: `Order`
- `GET /api/orders`
  - Query: `page`, `limit`
  - 200: `{ data: Order[], page, limit }`
- `GET /api/orders/:id`
  - 200: `Order`

## Models
- Product: `{ id, name, description, price, stock, category, image, sku, discount, created_at, updated_at, reviews? }`
- CartItem: `{ id, user_id, product_id, product, quantity, added_at }`
- Order: `{ id, user_id, order_number, total_amount, status, shipping_address, order_items[] }`

## Auth & Roles
- JWT required for protected routes via `Authorization: Bearer <token>` header.
- Admin routes require `is_admin = true` on the user.