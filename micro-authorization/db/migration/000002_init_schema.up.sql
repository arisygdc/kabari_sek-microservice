INSERT INTO role (name) VALUES ("basic");
INSERT INTO role (name) VALUES ("advance");
INSERT INTO permission (name) VALUES ("chat-text");
INSERT INTO permission (name) VALUES ("chat-special-emote");

INSERT INTO role_permission (role_id, permission_id) 
VALUES (
    (SELECT id FROM role WHERE name = "basic"), 
    (SELECT id FROM permission WHERE name = "chat-text")
);

INSERT INTO role_permission (role_id, permission_id) 
VALUES (
    (SELECT id FROM role WHERE name = "advance"), 
    (SELECT id FROM permission WHERE name = "chat-special-emote")
);
