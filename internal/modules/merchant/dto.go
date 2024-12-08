package merchant

type (
	CreateMerchantRequestDTO struct {
		Nama string `gorm:"column:nama" json:"nama" binding:"required"`
		Email string  `gorm:"column:email" json:"email" binding:"required"`
		Alamat string `gorm:"column:alamat" json:"alamat" binding:"required"`
	}

	CreateMerchantResponseDTO struct {
		ID int32 `gorm:"id"`
		Nama string `gorm:"column:nama"`
		Email string  `gorm:"column:email"`
		Alamat string `gorm:"column:alamat"`
	}

	DeleteMerchantResponseDTO struct {
		Message string `json:"message"`
	}

	CreateMerchantEditRequestDTO struct {
		Nama string `gorm:"column:nama" json:"nama"`
		Email string  `gorm:"column:email" json:"email"`
		Alamat string `gorm:"column:alamat" json:"alamat"`
	}
	
)
