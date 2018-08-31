package model

// 教師テーブルの構造体
type Teacher struct {
	TeacherId            int    `json:"teacher_id"`
	TeacherName          string `json:"teacher_name"`
	TeacherGradeHomeroom int    `json:"teacher_grade_homeroom"`
	TeacherHomeroom      string `json:"teacher_homeroom"`
	TeacherPassword      string `json:"teacher_password"`
	DeleteFlg            int    `json:"delete_flf"`
}

var teacher Teacher
var teachers []Teacher

func CreateTeacher(teacherName, teacherPassword, teacherHomeroom string, teacherGradeHomeroom int) bool {
	db := GormConnect()
	teacher.TeacherName = teacherName
	teacher.TeacherPassword = teacherPassword
	teacher.TeacherHomeroom = teacherHomeroom
	teacher.TeacherGradeHomeroom = teacherGradeHomeroom

	db.Exec("INSERT INTO `teachers`(`teacher_name`, `teacher_grade_homeroom`, `teacher_homeroom`, `teacher_password`) VALUES (?,?,?,?)", teacher.TeacherName, teacher.TeacherGradeHomeroom, teacher.TeacherHomeroom, teacher.TeacherPassword)

	defer db.Close()
	return true
}
