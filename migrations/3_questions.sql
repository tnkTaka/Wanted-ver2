-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `questions` (
  `question_id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT ,
  `student_id`  INT(4) NOT NULL ,
  `genre_id` INT(2) NOT NULL ,
  `question_title` VARCHAR(50) NOT NULL ,
  `question_body` VARCHAR (200) NOT NULL ,
  `delete_flg` INT(1) NOT NULL DEFAULT 0 ,
  `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `questions`;