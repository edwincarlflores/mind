DROP TABLE IF EXISTS `thought`;
DROP TABLE IF EXISTS `mind`;
DROP TABLE IF EXISTS `user`;

CREATE TABLE IF NOT EXISTS `user` (
  `id`              VARCHAR(60) NOT NULL,
  `uuid`            VARCHAR(60) NOT NULL,
  `username`        VARCHAR(32) NOT NULL,
  `auth_id`         VARCHAR(60),
  `auth_name`       VARCHAR(70),
  `email`           VARCHAR(70),
  `name`            VARCHAR(70) NOT NULL,
  `provider`        VARCHAR(20),
  `created_at`      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at`      TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=INNODB;


CREATE TABLE IF NOT EXISTS `thought` (
  `id`          VARCHAR(60) NOT NULL,
  `uuid`        VARCHAR(60) NOT NULL,
  `body`        VARCHAR(255) NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  `kind`        VARCHAR(60) NOT NULL,
  `public`      TINYINT(1) DEFAULT 0 NOT NULL,
  `user_id`     VARCHAR(60) NOT NULL,
  `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX u_idx (user_id),
  FOREIGN KEY (user_id) REFERENCES user(id)
) ENGINE=INNODB;

INSERT INTO `user` (id, uuid, name, username)
VALUES
("02a09c59-ac90-4716-b4d6-1477c6d8a780", "gd35a0e0-1b0c-87e0-0336-5dr3bd85jg64", "Test User", "testuser");

INSERT INTO `thought`
(id, uuid, body, description, kind, public, user_id)
VALUES
("166a6a3d-eca7-4224-8955-325e3ebfe5b4", "vq35a2q3-1b0c-21e9-3188-f58xbd8gk383", "Initial thought", "Test tought", "text", 0, "02a09c59-ac90-4716-b4d6-1477c6d8a780"),
("9cc5c9c6-cb77-4bf2-a2e9-bda47aaf4b7d", "oa53a1a2-6b2a-45m7-8120-4slxbd85cmf3", "Second thought", "Test tought", "text", 1, "02a09c59-ac90-4716-b4d6-1477c6d8a780"),
("8e6cba0f-f34b-4dc7-b387-03c132d6a068", "aq21o6g3-2a1z-80w1-1832-2dyubd83dqg5", "Third thought", "Test tought", "text", 1, "02a09c59-ac90-4716-b4d6-1477c6d8a780"),
("2f336fe0-09d1-404a-ad5b-86c4f099ae71", "mt95k3x7-1k6l-07d7-2790-0pzmre34mtq0", "Fourth thought", "Test tought", "text", 0, "02a09c59-ac90-4716-b4d6-1477c6d8a780");
