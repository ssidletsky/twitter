INSERT INTO users (username, password, email, first_name, last_name, age) VALUES 
("rachel", MD5("rachel_pass"), "rachel@gmail.com", "Rachel", "Green", 53),
("monica", MD5("monica_pass"), "monica@gmail.com", "Monica", "Geller", 57),
("phoebe", MD5("buffay_pass"), "phoebe@gmail.com", "Phoebe", "Buffay", 58),
("joey", MD5("joey_pass"), "joey@gmail.com", "Joey", "Tribbiani", 54),
("chandler", MD5("chandler_pass"), "chandler@gmail.com", "Chandler", "Bing", 52),
("ross", MD5("Ross"), "ross@gmail.com", "Ross", "Geller", 55);

INSERT INTO followers (followed_user_id, follower_user_id) VALUES
(1, 2), (1, 3), (1, 4), (2, 1), (2, 3), (2, 4), (3, 2), (3, 4),
(3, 5), (4, 3), (4, 5), (4, 6), (6, 1), (6, 2), (6, 5);

INSERT INTO tweets (author_user_id, `text`, publication_date) VALUES
(1, "Oh, I’m Sorry. Did My Back Hurt Your Knife?", NOW()),
(1, "Everyone I Know Is Either Getting Married Or Getting Pregnant...", NOW()),
(2, "Having a heart attack is nature’s way of telling you to slow down.", NOW()),
(2, "They’re as different as night and... later that night.", NOW()),
(3, "If You Want To Receive Emails About My Upcoming Shows, Then Please Give Me Money So I Can Buy A Computer.", NOW()),
(3, "Oh My God, A Woman Flirting With A Single Man? We Must Alert The Church Elders!", NOW()),
(4, "Joey doesn’t share food!", NOW()),
(4, "Why do you have to break up with her? Be a man. Just stop calling.", NOW()),
(5, "I’m full, and yet I know if I stop eating this, I’ll regret it.", NOW()),
(5, "I’m not great at the advice. Can I interest you in a sarcastic comment?", NOW()),
(6, "I Grew Up In A House With Monica, Okay. If You Didn't Eat Fast, You Didn't Eat.", NOW()),
(6, "Unagi Is A Total State Of Awareness.", NOW());
