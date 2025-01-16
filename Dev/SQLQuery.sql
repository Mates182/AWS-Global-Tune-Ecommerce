CREATE DATABASE Invoicing;
GO

USE Invoicing;
GO

-- Create Invoice table
CREATE TABLE Invoice (
    id INT IDENTITY(1,1) PRIMARY KEY,
    billing_date DATETIME NOT NULL,
    payment_method NVARCHAR(50) NOT NULL,
    amount DECIMAL(18, 2) NOT NULL,
    shipping_fee DECIMAL(18, 2),
    client_id INT NOT NULL,
    tax DECIMAL(18, 2)
);
GO

-- Create InvoiceDetails table with a composite primary key
CREATE TABLE InvoiceDetails (
    invoice_id INT NOT NULL,
    product_id NVARCHAR(50) NOT NULL, -- ID from MongoDB
    distributor NVARCHAR(100),
    quantity INT NOT NULL,
    discount DECIMAL(18, 2),
    unit_price DECIMAL(18, 2) NOT NULL,
    total_price DECIMAL(18, 2) NOT NULL,
    PRIMARY KEY (invoice_id, product_id),
    FOREIGN KEY (invoice_id) REFERENCES Invoice(id)
);
GO

CREATE PROCEDURE CreateInvoice
    @billing_date DATETIME,
    @payment_method NVARCHAR(50),
    @amount DECIMAL(18, 2),
    @shipping_fee DECIMAL(18, 2),
    @client_id INT,
    @tax DECIMAL(18, 2)
AS
BEGIN
    INSERT INTO Invoice (billing_date, payment_method, amount, shipping_fee, client_id, tax)
    VALUES (@billing_date, @payment_method, @amount, @shipping_fee, @client_id, @tax);
END;
GO

CREATE PROCEDURE GetInvoice
    @id INT
AS
BEGIN
    SELECT * FROM Invoice WHERE id = @id;
END;
GO

CREATE PROCEDURE UpdateInvoice
    @id INT,
    @billing_date DATETIME,
    @payment_method NVARCHAR(50),
    @amount DECIMAL(18, 2),
    @shipping_fee DECIMAL(18, 2),
    @client_id INT,
    @tax DECIMAL(18, 2)
AS
BEGIN
    UPDATE Invoice
    SET billing_date = @billing_date,
        payment_method = @payment_method,
        amount = @amount,
        shipping_fee = @shipping_fee,
        client_id = @client_id,
        tax = @tax
    WHERE id = @id;
END;
GO

CREATE PROCEDURE DeleteInvoice
    @id INT
AS
BEGIN
    DELETE FROM Invoice WHERE id = @id;
END;
GO

CREATE PROCEDURE CreateInvoiceDetails
    @invoice_id INT,
    @product_id NVARCHAR(50),
    @distributor NVARCHAR(100),
    @quantity INT,
    @discount DECIMAL(18, 2),
    @unit_price DECIMAL(18, 2),
    @total_price DECIMAL(18, 2)
AS
BEGIN
    INSERT INTO InvoiceDetails (invoice_id, product_id, distributor, quantity, discount, unit_price, total_price)
    VALUES (@invoice_id, @product_id, @distributor, @quantity, @discount, @unit_price, @total_price);
END;
GO
CREATE PROCEDURE GetInvoiceDetails
    @invoice_id INT,
    @product_id NVARCHAR(50)
AS
BEGIN
    SELECT * FROM InvoiceDetails WHERE invoice_id = @invoice_id AND product_id = @product_id;
END;
GO
CREATE PROCEDURE UpdateInvoiceDetails
    @invoice_id INT,
    @product_id NVARCHAR(50),
    @distributor NVARCHAR(100),
    @quantity INT,
    @discount DECIMAL(18, 2),
    @unit_price DECIMAL(18, 2),
    @total_price DECIMAL(18, 2)
AS
BEGIN
    UPDATE InvoiceDetails
    SET distributor = @distributor,
        quantity = @quantity,
        discount = @discount,
        unit_price = @unit_price,
        total_price = @total_price
    WHERE invoice_id = @invoice_id AND product_id = @product_id;
END;
GO
CREATE PROCEDURE DeleteInvoiceDetails
    @invoice_id INT,
    @product_id NVARCHAR(50)
AS
BEGIN
    DELETE FROM InvoiceDetails WHERE invoice_id = @invoice_id AND product_id = @product_id;
END;
GO
