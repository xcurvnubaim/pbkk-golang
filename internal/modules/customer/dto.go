package customer

type (
	CreateCustomerRequestDTO struct {
		Nama string `gorm:"column:nama" json:"nama" binding:"required"`
		Umur int32  `gorm:"column:umur" json:"umur" binding:"required"`
		Asal string `gorm:"column:asal" json:"asal" binding:"required"`
		NoHp string `gorm:"column:no_hp" json:"no_hp" binding:"required"`
	}

	CreateCustomerResponseDTO struct {
		ID int32 `gorm:"id"`
		Nama string `gorm:"column:nama"`
		Umur int32  `gorm:"column:umur"`
		Asal string `gorm:"column:asal"`
		NoHp string `gorm:"column:no_hp"`
	}
)
