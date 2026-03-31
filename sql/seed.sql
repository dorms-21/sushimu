DELETE FROM sales;
DELETE FROM order_items;
DELETE FROM orders;
DELETE FROM products;
DELETE FROM role_permissions;
DELETE FROM users;

INSERT INTO users (username, password_hash, name, role, active)
VALUES
('admin', '$2a$10$QfVv7yumxg/rEUyHvpYTluwFwo7gsb2O2o6Ts4ETmKX3cK2ZVeqNe', 'Administrador', 'admin', true),
('mesero', '$2a$10$QfVv7yumxg/rEUyHvpYTluwFwo7gsb2O2o6Ts4ETmKX3cK2ZVeqNe', 'Mesero Principal', 'mesero', true),
('cliente', '$2a$10$QfVv7yumxg/rEUyHvpYTluwFwo7gsb2O2o6Ts4ETmKX3cK2ZVeqNe', 'Cliente Demo', 'cliente', true);

INSERT INTO role_permissions (role, module, can_view, can_create, can_edit, can_delete) VALUES
('admin', 'dashboard', true, false, false, false),
('admin', 'productos', true, true, true, true),
('admin', 'ordenes', true, true, true, true),
('admin', 'ventas', true, true, true, true),
('admin', 'admin', true, true, true, true),
('admin', 'pos', true, true, true, true),

('mesero', 'dashboard', true, false, false, false),
('mesero', 'productos', true, false, false, false),
('mesero', 'ordenes', true, true, true, false),
('mesero', 'ventas', true, true, false, false),
('mesero', 'admin', false, false, false, false),
('mesero', 'pos', true, true, true, false),

('cliente', 'dashboard', false, false, false, false),
('cliente', 'productos', true, false, false, false),
('cliente', 'ordenes', true, true, false, false),
('cliente', 'ventas', false, false, false, false),
('cliente', 'admin', false, false, false, false),
('cliente', 'pos', false, false, false, false);

INSERT INTO products (name, description, price, stock, image_url, active)
VALUES
('Hamburguesa Clásica', 'Carne, queso y vegetales', 89.00, 20, '', true),
('Papas Fritas', 'Porción mediana', 45.00, 35, '', true),
('Refresco', 'Bebida de 600 ml', 25.00, 40, '', true);

INSERT INTO orders (table_no, user_id, status, total)
VALUES
('1', 1, 'abierta', 114.00),
('2', 2, 'preparando', 89.00);

INSERT INTO sales (order_id, payment_method, amount, status)
VALUES
(1, 'efectivo', 114.00, 'pagado');