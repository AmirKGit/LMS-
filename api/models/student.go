package models

type Student struct {
	Id         string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
	ExternalId string `json:"external_id"`
	Phone      string `json:"phone"`
	Mail       string `json:"mail"`
	Pasword    string `json:"pasword"`
	Active     bool   `json:"active"`
	Phones     []StudentNumbers`json:"student_phones"`
}

type StudentTimeTable struct {
	Id        string `json:"id"`
	Subject   string `json:"subject"`
	Teacher   string `json:"teacher"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type StudentSubjects struct {
	Id   string `json:"id"`
	Name string `josn:"name"`
	Type string `json:"type"`
}

type StudentTeacher struct {
	Name string `json:"name"`
}

type GetStudent struct {
	Id         string             `json:"id"`
	FirstName  string             `json:"first_name"`
	LastName   string             `json:"last_name"`
	Age        int                `json:"age"`
	ExternalId string             `json:"external_id"`
	Phone      string             `json:"phone"`
	Mail       string             `json:"mail,omitempty"`
	Teacher    []StudentTeacher   `json:"teacher"`
	TimeTable  []StudentTimeTable `json:"time_table"`
	Subjects   []StudentSubjects  `json:"subjects"`
	Phones     []StudentNumbers   `json:"student_phones"`
}

type StudentNumbers struct {
	Phone     string `json:"phone"`
	StudentId string `json:"student_id"`
}

type GetAllStudentsRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type GetAllStudentsResponse struct {
	Students []GetStudent `json:"students"`
	Count    int          `json:"count"`
}

type IsActiveResponse struct {
	Active bool `json:"active"`
}

type CheckLessonStudent struct {
	TeacherName string `json:"teacher_name"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	SubjectName string `json:"subject_name"`
}

type GetAllStudentsAttandenceReportRequest struct {
	StudentId string `json:"student_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	TeacherId string `json:"teacher_id"`
	Page      uint64 `json:"page"`
	Limit     uint64 `json:"limit"`
}

type GetAllStudentsAttandenceReportResponse struct {
	Students []StudentAttandenceReport `json:"students"`
	Count    int64                     `json:"count"`
}

type StudentAttandenceReport struct {
	StudentId        string  `json:"student_id"`
	StudentName      string  `json:"student_name"`
	StudentCreatedAt string  `json:"student_created_at"`
	TeacherName      string  `json:"teacher_name"`
	StudyTime        float64 `json:"study_time"`
	AvgStudyTime     float64 `json:"avg_study_time"`
}

