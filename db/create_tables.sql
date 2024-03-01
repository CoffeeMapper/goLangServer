CREATE TABLE coffee_point (
                             coffee_point_id int PRIMARY KEY,
                             logo varchar,
                             description varchar,
                             address varchar,
                             email varchar,
                             brand_id int,
                             work_time varchar,
                             order_id int,
                             avg_rating int,
                             feedback_id int
);

CREATE TABLE brand (
                       brand_id int PRIMARY KEY,
                       name varchar,
                       owner_id int,
                       phone varchar
);

CREATE TABLE users (
                       user_id int PRIMARY KEY,
                       login varchar,
                       email varchar,
                       password varchar,
                       sub_id int
);

CREATE TABLE menu (
                      menu_id int PRIMARY KEY,
                      product_id int
);

CREATE TABLE product (
                         product_id int PRIMARY KEY,
                         category_id int,
                         title varchar,
                         picture varchar,
                         description varchar,
                         price int
);

CREATE TABLE category (
                          category_id int PRIMARY KEY,
                          title varchar,
                          description varchar
);

CREATE TABLE feedback (
                          feedback_id int PRIMARY KEY,
                          user_id int,
                          feedback varchar,
                          rate int
);


CREATE TABLE order (
                       order_id int PRIMARY KEY,
                       product_id int,
                       started_at timestamp,
                       will_done_at timestamp,
                       status bool
);

CREATE TABLE combo (
                       combo_id int PRIMARY KEY,
                       title int,
                       price int,
                       product_id int
);

ALTER TABLE coffee_point ADD FOREIGN KEY (brand_id) REFERENCES brand (brand_id);

ALTER TABLE brand ADD FOREIGN KEY (owner_id) REFERENCES users (user_id);

ALTER TABLE brand ADD FOREIGN KEY (brand_id) REFERENCES users (sub_id);

ALTER TABLE coffee_point ADD FOREIGN KEY (coffee_point_id) REFERENCES menu (menu_id);

ALTER TABLE product ADD FOREIGN KEY (product_id) REFERENCES menu (product_id);

ALTER TABLE category ADD FOREIGN KEY (category_id) REFERENCES product (category_id);

ALTER TABLE feedback ADD FOREIGN KEY (feedback_id) REFERENCES coffee_point (feedback_id);

ALTER TABLE users ADD FOREIGN KEY (user_id) REFERENCES feedback (user_id);

ALTER TABLE product ADD FOREIGN KEY (product_id) REFERENCES order (product_id);

ALTER TABLE product ADD FOREIGN KEY (product_id) REFERENCES combo (product_id);

ALTER TABLE order ADD FOREIGN KEY (order_id) REFERENCES coffee_point (order_id);