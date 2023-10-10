CREATE SCHEMA wisdom;
CREATE TABLE wisdom.services (
  id UUID PRIMARY KEY,
  name VARCHAR UNIQUE,
  price NUMERIC(12,2)
);

CREATE TABLE wisdom.customers (
   id UUID PRIMARY KEY,
   first_name VARCHAR,
   last_name VARCHAR,
   email VARCHAR UNIQUE,
   phone VARCHAR,
   address VARCHAR
);

CREATE UNIQUE INDEX IF NOT EXISTS customers_email_index ON wisdom.customers (email);

CREATE TABLE wisdom.vendors (
     id UUID PRIMARY KEY,
     name VARCHAR NOT NULL,
     contact VARCHAR,
     phone VARCHAR,
     email VARCHAR UNIQUE,
     address VARCHAR
);

CREATE TABLE wisdom.products (
      id UUID PRIMARY KEY,
      name VARCHAR UNIQUE,
      price NUMERIC (12,2),
      vendor_id UUID NOT NULL,
      FOREIGN KEY (VENDOR_ID) references wisdom.vendors(ID)
);