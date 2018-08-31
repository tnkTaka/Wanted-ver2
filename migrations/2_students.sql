-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `students` (
  `student_id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT ,
  `student_name` VARCHAR(20) NOT NULL ,
  `student_grade` INT(1) NOT NULL ,
  `student_class` VARCHAR(10) NOT NULL ,
  `student_class_number` INT(4) NOT NULL ,
  `student_login_id` VARCHAR(6) NOT NULL ,
  `student_login_password` CHAR(8) NOT NULL ,
  `student_nick_name` VARCHAR(10) NOT NULL DEFAULT 'ななし' ,
  `student_profile_image` VARCHAR(30) NOT NULL DEFAULT 'default.png' ,
  `delete_flg` INT(1) NOT NULL DEFAULT 0 ,
  `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO `students`(`student_name`, `student_grade`, `student_class`, `student_class_number`, `student_login_id`, `student_login_password`) VALUES ("伊勢田",3,"A",2,"3A2",1);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `students`;