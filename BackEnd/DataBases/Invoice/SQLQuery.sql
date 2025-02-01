CREATE DATABASE Invoicing;
GO

USE Invoicing;
GO

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