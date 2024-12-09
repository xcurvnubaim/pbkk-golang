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

	CustomerDTO struct {
		ID int32 `gorm:"id"`
		Nama string `gorm:"column:nama"`
		Umur int32  `gorm:"column:umur"`
		Asal string `gorm:"column:asal"`
		NoHp string `gorm:"column:no_hp"`
	}

	FindAllCustomerResponseDTO struct {
		Customers []CustomerDTO `json:"customers"`
	}

	UpdateCustomerRequestDTO struct {
		ID   int32  `json:"id" binding:"required"`
		Nama string `json:"nama" binding:"required"`
		Umur int32  `json:"umur" binding:"required"`
		Asal string `json:"asal" binding:"required"`
		NoHp string `json:"no_hp" binding:"required"`
	}

	UpdateCustomerResponseDTO struct {
		CustomerDTO
	}

	DeleteCustomerRequestDTO struct {
		ID int32 `json:"id" binding:"required"`
	}

	DeleteCustomerResponseDTO struct {
		ID int32 `json:"id"`
	}
)
