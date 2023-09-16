DROP TABLE IF EXISTS `thought`;
DROP TABLE IF EXISTS `mind`;
DROP TABLE IF EXISTS `user`;

CREATE TABLE IF NOT EXISTS `user` (
  `id`          BIGINT(19) NOT NULL,
  `uuid`        VARCHAR(255) NOT NULL,
  `name`        VARCHAR(255) NOT NULL,
  `username`    VARCHAR(255) NOT NULL,
  `email`       VARCHAR(255) NOT NULL,
  `token`       VARCHAR(255) NOT NULL,
  `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS `mind` (
  `id`          BIGINT(19) NOT NULL,
  `uuid`        VARCHAR(255) NOT NULL,
  `user_id`     BIGINT(19) NOT NULL,  
  `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX usr_idx (user_id),
  FOREIGN KEY (user_id) REFERENCES user(id)
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS `thought` (
  `id`          BIGINT(19) NOT NULL,
  `uuid`        VARCHAR(255) NOT NULL,
  `body`        VARCHAR(255) NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  `kind`        VARCHAR(100) NOT NULL,
  `public`      TINYINT(1) DEFAULT 0 NOT NULL,
  `mind_id`     BIGINT(19) NOT NULL,
  `user_id`     BIGINT(19) NOT NULL,
  `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX m_idx (mind_id),
  FOREIGN KEY (mind_id) REFERENCES mind(id),
  INDEX u_idx (user_id),
  FOREIGN KEY (user_id) REFERENCES user(id)
) ENGINE=INNODB;

INSERT INTO `user` (id, uuid, name, username, email, token)
VALUES
(1, "gd35a0e0-1b0c-87e0-0336-5dr3bd85jg64", "Test User", "testuser", "", "testtoken");

INSERT INTO `mind` (id, uuid, user_id)
VALUES
(1, "e9b1a0e0-4b0a-11e9-8647-d663bd873d93", 1);

INSERT INTO `thought`
(id, uuid, body, description, kind, public, mind_id, user_id)
VALUES
(1, "vq35a2q3-1b0c-21e9-3188-f58xbd8gk383", "Initial thought", "Test tought", "text", 0, 1, 1),
(2, "oa53a1a2-6b2a-45m7-8120-4slxbd85cmf3", "Second thought", "Test tought", "text", 1, 1, 1),
(3, "aq21o6g3-2a1z-80w1-1832-2dyubd83dqg5", "Third thought", "Test tought", "text", 1, 1, 1),
(4, "mt95k3x7-1k6l-07d7-2790-0pzmre34mtq0", "Fourth thought", "Test tought", "text", 0, 1, 1);