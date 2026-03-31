INSERT INTO users (username, password_hash, name, role, active)
VALUES
('admin', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'Administrador', 'admin', true),
('mesero', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'Mesero Principal', 'mesero', true),
('cliente', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'Cliente Demo', 'cliente', true);

INSERT INTO products (name, description, price, stock, image_url, active)
VALUES
('Hamburguesa Clásica', 'Carne, queso y vegetales', 89.00, 20, '', true),
('Papas Fritas', 'Porción mediana', 45.00, 35, '', true),
('Refresco', 'Bebida de 600 ml', 25.00, 40, '', true);