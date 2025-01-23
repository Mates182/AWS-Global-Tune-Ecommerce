CREATE DATABASE UserSystem;
GO

USE UserSystem;
GO

CREATE TABLE Users (
    id INT IDENTITY(1,1) NOT NULL,
    name NVARCHAR(100) NOT NULL,
    email NVARCHAR(100) NOT NULL,
    password NVARCHAR(100) NOT NULL,
    role NVARCHAR(50) NOT NULL,
    PRIMARY KEY (id)
);
GO

CREATE PROCEDURE CreateUser
    @name NVARCHAR(100),
    @email NVARCHAR(100),
    @password NVARCHAR(100),
    @role NVARCHAR(50)
AS
BEGIN
    INSERT INTO Users (name, email, password, role)
    VALUES (@name, @email, @password, @role);
END;
GO

CREATE PROCEDURE GetUser
    @id INT
AS
BEGIN
    SELECT * FROM Users WHERE id = @id;
END;
GO

CREATE PROCEDURE UpdateUser
    @id INT,
    @name NVARCHAR(100),
    @email NVARCHAR(100),
    @password NVARCHAR(100),
    @role NVARCHAR(50)
AS
BEGIN
    UPDATE Users
    SET name = @name,
        email = @email,
        password = @password,
        role = @role
    WHERE id = @id;
END;
GO

CREATE PROCEDURE DeleteUser
    @id INT
AS
BEGIN
    DELETE FROM Users WHERE id = @id;
END;
GO
