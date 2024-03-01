INSERT INTO users (user_id, login, email, password, has_premium)
VALUES
    (1, 'john_doe', 'johndoe@example.com', 'password123', true),
    (2, 'alice_smith', 'alicesmith@example.com', 'password456', false),
    (3, 'bob_ross', 'bobross@example.com', 'password789', true),
    (4, 'emma_watson', 'emmawatson@example.com', 'password987', true),
    (5, 'michael_jackson', 'michaeljackson@example.com', 'password654', true),
    (6, 'laura_coleman', 'lauracoleman@example.com', 'password321', false),
    (7, 'david_bowie', 'davidbowie@example.com', 'password987', false),
    (8, 'sarah_parker', 'sarahparker@example.com', 'password246', true),
    (9, 'robert_green', 'robertgreen@example.com', 'password135', false),
    (10, 'olivia_moore', 'oliviamoore@example.com', 'password579', true);

INSERT INTO brand VALUES (1, 'CoffePapa', 1, '555-1234');
INSERT INTO brand VALUES (2, 'CoffeMama', 2, '555-5678');
INSERT INTO brand VALUES (3,'CoffeOtchim', 3, '555-9101');
INSERT INTO brand VALUES (4,'CoffeTea', 4, '555-1122');

INSERT INTO coffe_point (coffe_point_id, name, logo, description, address, email, brand_id, work_time) VALUES
                                                                                                           (1, 'CoffePapa', 'image1.jpg', 'Cozy cafe with a wide range of coffees', '123 Main Street', 'info@coffeehouse.com', 1, 'Mon-Sun: 8am-10pm'),
                                                                                                           (2, 'CoffeMama', 'image2.jpg', 'Trendy spot known for its strong espresso drinks', '456 Oak Avenue', 'hello@espressobar.com', 2, 'Mon-Fri: 7am-7pm, Sat-Sun: 8am-6pm'),
                                                                                                           (3, 'CoffePapa', 'image3.jpg', 'Relaxing cafe specializing in lattes and pastries', '789 Elm Street', 'contact@lattelounge.com', 1, 'Mon-Sun: 9am-9pm'),
                                                                                                           (4, 'CoffeOtchim', 'image4.jpg', 'Magical coffee shop with a variety of mochas', '234 Maple Road', 'magic@mochamagic.com', 3, 'Mon-Sat: 7am-8pm, Sun: 8am-6pm'),
                                                                                                           (5, 'CoffeMama', 'image5.jpg', 'Cozy corner cafe serving delicious cappuccinos', '567 Pine Lane', 'hello@cappuccinocorner.com', 2, 'Mon-Fri: 6:30am-6pm, Sat: 8am-4pm'),
                                                                                                           (6, 'CoffePapa', 'image6.jpg', 'Buzzing coffee shop with a focus on espresso drinks', '789 Walnut Street', 'buzz@beanbuzz.com', 1, 'Mon-Sun: 7am-7pm'),
                                                                                                           (7, 'CoffeMama', 'image7.jpg', 'Popular spot with a wide selection of coffee blends', '890 Cedar Avenue', 'info@javajunction.com', 2, 'Mon-Fri: 6am-8pm, Sat-Sun: 7am-6pm'),
                                                                                                           (8, 'CoffeOtchim', 'image8.jpg', 'Perfect place for getting your daily caffeine fix', '432 Birch Road', 'contact@caffeinefix.com', 3, 'Mon-Sat: 6:30am-7:30pm, Sun: 8am-6pm'),
                                                                                                           (9, 'CoffeTea', 'image9.jpg', 'Cozy and welcoming cafe with a variety of coffee options', '654 Cherry Street', 'cozy@coffeecozy.com', 4, 'Mon-Sun: 8am-9pm'),
                                                                                                           (10, 'CoffeMama', 'image10.jpg', 'Showing love through delicious lattes and pastries', '321 Cedar Lane', 'love@lattelove.com', 2, 'Mon-Fri: 7am-7pm, Sat: 8am-6pm, Sun: 9am-5pm');