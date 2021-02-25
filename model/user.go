package model

type (
	User struct {
		UserID   string `json:"userId" gorm:"column:id"`
		UserName string `json:"userName" gorm:"column:userName"`
		BirthDate string `json:"birthDate" gorm:"column:birthDate"`
		Pekerjaan    string `json:"pekerjaan" gorm:"column:Pekerjaan"`
		NoKtp  string `json:"noKtp" gorm:"column:noKtp"`
		Pendidikan    string `json:"pendidikan" gorm:"column:Pendidikan"`
	}
	RequestUser struct {
		UserName string `json:"userName" gorm:"column:userName"`
		BirthDate string `json:"birthDate" gorm:"column:birthDate"`
		Pekerjaan    string `json:"pekerjaan" gorm:"column:Pekerjaan"`
		NoKtp  string `json:"noKtp" gorm:"column:noKtp"`
		Pendidikan    string `json:"pendidikan" gorm:"column:Pendidikan"`
	}

)