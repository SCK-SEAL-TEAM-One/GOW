CREATE DATABASE IF NOT EXISTS GOW;
USE GOW;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS quotation (
    quotation_id INT(3) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    quotation_number VARCHAR(13),
    contact_name VARCHAR(512),
    contact_email VARCHAR(512),
    contact_phoneNumber VARCHAR(13),
	total_price VARCHAR(13),
    discount_price VARCHAR(13),
    price_after_discount VARCHAR(13),
    vat VARCHAR(13),
    net_total_price VARCHAR(13),
    total_price_thai VARCHAR(13),
	project_name VARCHAR(13),
    quotation_date VARCHAR(512),
    vat_rate VARCHAR(512),
    vat_included VARCHAR(512),
    customer_id INT(3) UNSIGNED ,
    company_id INT(3) UNSIGNED ,
    FOREIGN KEY (customer_id) REFERENCES customer(customer_id),
    FOREIGN KEY (company_id) REFERENCES company(company_id),
    created_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_time DATETIME ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
