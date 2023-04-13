package entities

type Logs struct {
	ID        int64   `gorm:"column:logs_id;primaryKey;autoIncrement;not null"`
	ValueA    int64   `gorm:"column:value_a;not null"`
	ValueB    int64   `gorm:"column:value_b;not null"`
	Operation string  `gorm:"column:operation;not null"`
	Result    float64 `gorm:"column:result;not null"`
}
