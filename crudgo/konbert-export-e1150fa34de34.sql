CREATE TABLE IF NOT EXISTS `usuarios` (
`Name` VARCHAR(60) NULL,
`CPF` VARCHAR(60) PRIMARY KEY NOT NULL,
`PhoneNumber` VARCHAR (60) NULL,
`Celphones` VARCHAR (60) NOT NULL,
`Faxphones` VARCHAR (60) NULL,
`email` VARCHAR(60) NULL,
`cep` VARCHAR(60) NULL
);

INSERT INTO usuarios VALUES
('Eric Torres','921.889.550-80',[{\"CelNumber\":\" (81) 98785-0100 \"\,\"FaxNumber\":\" (81) 3550-6874 \"\,\"PhoneNumber\":\" (81) 3049-0100 \"}],'erictorres@gmail.com','52051-320'),
('Cosmina Santos','921.889.550-80',[{\"CelNumber\":\" (81) 98871-3521 \"\,\"PhoneNumber\":\" (81) 4875-3541 \"}],'cosmina@hotmail.com',NULL),
('Bala Dantas Silva','921.889.550-80',[{\"FaxNumber\":\" (81) 4842-6587 \"}],'bd123@yahoo.com.br','51000-870');