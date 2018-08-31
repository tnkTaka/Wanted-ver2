-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `teachers` (
  `teacher_id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  `teacher_name` VARCHAR(20) NOT NULL ,
  `teacher_grade_homeroom` INT(1) NOT NULL ,
  `teacher_homeroom` VARCHAR(10) NOT NULL ,
  `teacher_password` CHAR(8) NOT NULL ,
  `delete_flg` INT(1) NOT NULL DEFAULT 0 ,
  `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO `teachers`(`teacher_name`, `teacher_grade_homeroom`, `teacher_homeroom`, `teacher_password`) VALUES ("山田忠明",3,"A","abcd1234");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `teachers`;
