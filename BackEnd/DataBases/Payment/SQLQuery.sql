CREATE DATABASE PaymentSystem;
GO

USE PaymentSystem;
GO

CREATE TABLE Payment (
    id INT IDENTITY(1,1) NOT NULL,
    method NVARCHAR(100) NOT NULL,
    amount DECIMAL(18, 2) NOT NULL,
    date DATETIME NOT NULL,
    user_id INT NOT NULL,
    PRIMARY KEY (id)
);
GO

CREATE PROCEDURE CreatePayment
    @method NVARCHAR(100),
    @amount DECIMAL(18, 2),
    @date DATETIME,
    @user_id INT
AS
BEGIN
    INSERT INTO Payment (method, amount, date, user_id)
    VALUES (@method, @amount, @date, @user_id);
END;
GO

CREATE PROCEDURE GetPayment
    @id INT
AS
BEGIN
    SELECT * FROM Payment WHERE id = @id;
END;
GO

CREATE PROCEDURE UpdatePayment
    @id INT,
    @method NVARCHAR(100),
    @amount DECIMAL(18, 2),
    @date DATETIME,
    @user_id INT
AS
BEGIN
    UPDATE Payment
    SET method = @method,
        amount = @amount,
        date = @date,
        user_id = @user_id
    WHERE id = @id;
END;
GO

CREATE PROCEDURE DeletePayment
    @id INT
AS
BEGIN
    DELETE FROM Payment WHERE id = @id;
END;
GO
