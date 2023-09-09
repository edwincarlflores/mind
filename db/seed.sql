DROP TABLE IF EXISTS `mind`;
DROP TABLE IF EXISTS `thought`;

CREATE TABLE IF NOT EXISTS `mind` (
  `id`          BIGINT(19) NOT NULL,
  `owner_token` VARCHAR(255) NOT NULL,
  `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS `thought` (
  `id`          BIGINT(19) NOT NULL,
  `body`        VARCHAR(255) NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  `kind`        VARCHAR(100) NOT NULL,
  `public`      TINYINT(1) DEFAULT 0 NOT NULL,
  `mind_id`     BIGINT(19) NOT NULL,
  `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX m_idx (mind_id),
  FOREIGN KEY (mind_id) REFERENCES mind(id)
) ENGINE=INNODB;

INSERT INTO `mind` (id, owner_token)
VALUES
(1, "ed_token_000");

INSERT INTO `thought`
(id, body, description, kind, public, mind_id)
VALUES
(1, "Initial thought", "Test tought", "text", 0, 1),
(2, "Second thought", "Test tought", "text", 1, 1),
(3, "Third thought", "Test tought", "text", 1, 1),
(4, "Fourth thought", "Test tought", "text", 0, 1);