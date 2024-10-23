CREATE TABLE Categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE Products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    quantity_in_stock INT NOT NULL,
    category_id INT REFERENCES Categories(id) ON DELETE SET NULL,  -- FK hacia Categories
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE Suppliers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    number_phone VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE PurchaseOrders (
    id SERIAL PRIMARY KEY,
    supplier_id INT REFERENCES Suppliers(id) ON DELETE SET NULL,  -- FK hacia Suppliers
    order_date TIMESTAMPTZ DEFAULT NOW(),
    status VARCHAR(50),
    total DECIMAL(10, 2),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE PurchaseOrderDetails (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES PurchaseOrders(id) ON DELETE CASCADE,  -- FK hacia PurchaseOrders
    product_id INT REFERENCES Products(id) ON DELETE CASCADE,    -- FK hacia Products
    quantity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE Sales (
    id SERIAL PRIMARY KEY,
    sale_date TIMESTAMPTZ DEFAULT NOW(),
    name VARCHAR(255) NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE SalesDetails (
    id SERIAL PRIMARY KEY,
    sale_id INT REFERENCES Sales(id) ON DELETE CASCADE,   -- FK hacia Sales
    product_id INT REFERENCES Products(id) ON DELETE CASCADE,  -- FK hacia Products
    quantity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);
