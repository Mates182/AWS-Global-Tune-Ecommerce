CREATE DATABASE Invoicing;
GO

USE Invoicing;
GO

CREATE TABLE InvoiceDetails (
    invoice_id INT NOT NULL,
    product_id NVARCHAR(50) NOT NULL,
    distributor NVARCHAR(100),
    quantity INT NOT NULL,
    discount DECIMAL(18, 2),
    unit_price DECIMAL(18, 2) NOT NULL,
    total_price DECIMAL(18, 2) NOT NULL,
    PRIMARY KEY (invoice_id, product_id)
);
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
