DELETE FROM order_items;
DELETE FROM sales;
DELETE FROM orders;
DELETE FROM products;
DELETE FROM users;

INSERT INTO users (username, password_hash, name, role, active)
VALUES
('admin', '$2a$10$QfVv7yumxg/rEUyHvpYTluwFwo7gsb2O2o6Ts4ETmKX3cK2ZVeqNe', 'Administrador', 'admin', true),
('mesero', '$2a$10$QfVv7yumxg/rEUyHvpYTluwFwo7gsb2O2o6Ts4ETmKX3cK2ZVeqNe', 'Mesero Principal', 'mesero', true),
('cliente', '$2a$10$QfVv7yumxg/rEUyHvpYTluwFwo7gsb2O2o6Ts4ETmKX3cK2ZVeqNe', 'Cliente Demo', 'cliente', true);

INSERT INTO products (name, description, price, stock, image_url, active)
VALUES
('Hamburguesa Clásica', 'Carne, queso y vegetales', 89.00, 20, '', true),
('Papas Fritas', 'Porción mediana', 45.00, 35, '', true),
('Refresco', 'Bebida de 600 ml', 25.00, 40, '', true);