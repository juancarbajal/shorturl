package entities
import ("time"
)
type Shorten struct{
	ID string `gorm:"primaryKey"`
	Url string
	CreatedAt time.Time
	UpdatedAt time.Time
}
