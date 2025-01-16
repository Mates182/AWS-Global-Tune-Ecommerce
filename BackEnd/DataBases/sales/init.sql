USE master;

IF NOT EXISTS (SELECT * FROM sys.databases WHERE name = 'Invoicing')
BEGIN
    CREATE DATABASE Invoicing;
END;

USE Invoicing;

IF NOT EXISTS (SELECT * FROM sys.tables WHERE name = 'Invoice')
BEGIN
    CREATE TABLE Invoice (
        id INT IDENTITY(1,1) PRIMARY KEY,
        billing_date DATETIME NOT NULL,
        payment_method NVARCHAR(50) NOT NULL,
        amount DECIMAL(18, 2) NOT NULL,
        shipping_fee DECIMAL(18, 2),
        client_id INT NOT NULL,
        tax DECIMAL(18, 2)
    );
END;

IF NOT EXISTS (SELECT * FROM sys.tables WHERE name = 'InvoiceDetails')
BEGIN
    CREATE TABLE InvoiceDetails (
        id INT IDENTITY(1,1) PRIMARY KEY,
        invoice_id INT NOT NULL,
        product_id NVARCHAR(50) NOT NULL, -- ID from MongoDB
        distributor NVARCHAR(100),
        quantity INT NOT NULL,
        discount DECIMAL(18, 2),
        unit_price DECIMAL(18, 2) NOT NULL,
        total_price DECIMAL(18, 2) NOT NULL,
        FOREIGN KEY (invoice_id) REFERENCES Invoice(id)
    );
END;
