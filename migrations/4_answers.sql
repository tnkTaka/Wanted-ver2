-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `answers` (
  `answer_id`   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT ,
  `question_id` INT(4) NOT NULL ,
  `student_id`  INT(4) NOT NULL ,
  `answer_body` VARCHAR (200) NOT NULL ,
  `answer_like` INT (4) NOT NULL DEFAULT 0 ,
  `delete_flg` INT(1) NOT NULL DEFAULT 0 ,
  `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `answers`;