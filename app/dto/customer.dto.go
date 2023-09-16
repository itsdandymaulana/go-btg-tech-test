package dto

type GetCustomerDTO struct {
	CstName     string      `json:"nama"`
	CstDob      string      `json:"tanggal_lahir"`
	CstPhoneNum string      `json:"telepon"`
	Nationality string      `json:"kewarganegaraan"`
	CstEmail    string      `json:"email"`
	Family      []FamilyDTO `json:"keluarga"`
}

type FamilyDTO struct {
	FlRelation string `json:"hubungan"`
	FlName     string `json:"nama"`
	FlDob      string `json:"tanggal_lahir"`
}
