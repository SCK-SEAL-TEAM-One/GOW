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
    quotation_number VARCHAR(15),
    contact_name VARCHAR(512),
    contact_email VARCHAR(512),
    contact_phoneNumber VARCHAR(13),
	total_price VARCHAR(13),
    discount_price VARCHAR(13),
    price_after_discount VARCHAR(13),
    vat VARCHAR(13),
    net_total_price VARCHAR(13),
    total_price_thai VARCHAR(512),
	project_name VARCHAR(512),
    quotation_date DATETIME,
    vat_rate VARCHAR(512),
    vat_included VARCHAR(512),
    customer_taxid VARCHAR(13) ,
    company_taxid VARCHAR(13) ,
    FOREIGN KEY (company_taxid) REFERENCES company(company_taxid),
    FOREIGN KEY (customer_taxid) REFERENCES customer(customer_taxid),
    created_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_time DATETIME ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS orders (
    order_id INT(3) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    amount INT(3) UNSIGNED,
    price_per_unit DECIMAL(13,4),
    price DECIMAL(13,4),
    description VARCHAR(512),
    quotation_id INT(3) UNSIGNED,
    FOREIGN KEY (quotation_id) REFERENCES quotation(quotation_id),
    created_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_time DATETIME ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
