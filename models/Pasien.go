package models

type Pasien struct {
	Id           int64  `gorm:"primaryKey" json:"id"`
	NamaPasien   string `gorm:"type:varchar(255)" json:"nama_pasien"`
	AlamatPasien string `gorm:"type:text" json:"alamat_pasien"`
}
