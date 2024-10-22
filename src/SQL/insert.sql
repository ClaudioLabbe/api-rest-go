-- Insertar categorías en la tabla Categories
INSERT INTO Categories ("name", description) VALUES
('Electronics', 'Dispositivos electrónicos y gadgets'),
('Furniture', 'Muebles para el hogar y la oficina'),
('Clothing', 'Ropa y accesorios de vestir'),
('Books', 'Libros y material de lectura'),
('Toys', 'Juguetes y artículos para niños'),
('Beauty', 'Skincare, haircare, and makeup products'),
('Groceries', 'Food items and daily essentials');


-- Insertar productos en la tabla Products
INSERT INTO Products (name, description, price, quantity_in_stock, category_id) VALUES 
('iPhone 13', 'Último modelo de iPhone de Apple', 999.99, 100, 1),  -- Electrónica
('MacBook Pro', 'Laptop de Apple con chip M1', 1999.99, 50, 1), -- Electrónica
('Silla de Oficina', 'Silla ergonómica para oficina', 150.50, 200, 2),   -- Muebles
('Mesa de Comedor', 'Mesa de madera para 6 personas', 299.99, 30, 2),  -- Muebles
('Manzanas Orgánicas', 'Manzanas orgánicas frescas', 1.99, 500, 7),     -- Alimentos
('Leche', '1 litro de leche fresca', 0.99, 300, 7),              -- Alimentos
('Jeans', 'Pantalones vaqueros de mezclilla azul', 49.99, 100, 7),                 -- Ropa
('Camiseta', 'Camiseta de algodón en varios colores', 19.99, 150, 7), -- Ropa
('Harry Potter', 'Novela de fantasía de J.K. Rowling', 12.99, 200, 4), -- Libros
('El Método Lean Startup', 'Libro de negocios de Eric Ries', 14.99, 100, 4), -- Libros
('Crema Facial', 'Crema hidratante para el rostro', 24.99, 80, 6),        -- Belleza
('Champú', 'Champú para el cuidado del cabello para todo tipo de cabello', 8.99, 150, 6); -- Belleza

-- Insertar proveedores en la tabla Suppliers
INSERT INTO Suppliers (name, email, number_phone) VALUES 
('Proveedor Tecnológico', 'tecnologia@example.com', '+34 900 123 456'),
('Muebles Modernos', 'muebles@example.com', '+34 900 654 321'),
('Alimentos Frescos', 'alimentosfrescos@example.com', '+34 900 987 654'),
('Ropa Estilo', 'ropaestilo@example.com', '+34 900 456 789'),
('Librería Central', 'libreriacentral@example.com', '+34 900 321 654'),
('Belleza Natural', 'bellezanatural@example.com', '+34 900 789 123');
