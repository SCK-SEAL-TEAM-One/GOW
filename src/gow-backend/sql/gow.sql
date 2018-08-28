CREATE DATABASE IF NOT EXISTS GOW;
CREATE TABLE IF NOT EXISTS customer (
customer_id INT(3) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
customer_name_th VARCHAR(512),
customer_name_en VARCHAR(512),
customer_branch_th VARCHAR(256),
customer_branch_en VARCHAR(256),
customer_address_th VARCHAR(256),
customer_address_en VARCHAR(256),
customer_taxid VARCHAR(13) UNIQUE,
created_time DATETIME DEFAULT CURRENT_TIMESTAMP,
updated_time DATETIME ON UPDATE CURRENT_TIMESTAMP
)

CREATE TABLE IF NOT EXISTS company (
company_id INT(3) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
company_name_th VARCHAR(512),
company_name_en VARCHAR(512),
company_branch_th VARCHAR(256),
company_branch_en VARCHAR(256),
company_address_th VARCHAR(256),
company_address_en VARCHAR(256),
company_taxid VARCHAR(13) UNIQUE,
company_phonenumber VARCHAR(13),
created_time DATETIME DEFAULT CURRENT_TIMESTAMP,
updated_time DATETIME ON UPDATE CURRENT_TIMESTAMP
)