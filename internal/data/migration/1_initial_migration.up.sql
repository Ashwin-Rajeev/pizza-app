BEGIN;

CREATE TYPE PIZZA_SIZE AS ENUM ('medium');

CREATE TYPE PIZZA_TYPE AS ENUM ('veggie', 'meat');

CREATE TYPE  PIZZA_COOKING_STAGE AS ENUM ('start', 'dough-prep', 'oven-bake', 'topping-art', 'done');

-- pizza master table
CREATE TABLE IF NOT EXISTS pizza
  (
     id            SERIAL PRIMARY KEY,
     name          VARCHAR NOT NULL,
     type          PIZZA_TYPE NOT NULL
  );

-- pizza table seed data
INSERT INTO
  pizza(name, type)
VALUES
  ('Veg Magic', 'veggie'),
  ('Meat Magic', 'meat');

-- pizza_order table holds the order details
CREATE TABLE IF NOT EXISTS pizza_order
  (
     id            SERIAL PRIMARY KEY,
     pizza_id      INT NOT NULL REFERENCES pizza(id),
     pizza_size    PIZZA_SIZE NOT NULL,
     cooking_stage PIZZA_COOKING_STAGE NOT NULL,
     user_id       INT NOT NULL,
     start_time    TIMESTAMP NOT NULL,
     end_time      TIMESTAMP,
     is_active     BOOLEAN
  );

END;
