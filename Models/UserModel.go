//created a model for User that consists following fields:
/*
name
email
phone
address
*/

package Models

type User struct {
	Id      uint   `json:"id,omitempty" validate:"omitempty,uuid"` //don’t have an ID, but if we did it would need to be a UUID
	Name    string `json:"name" validate:"required,alpha"`         // CUstomer name mandatory
	Email   string `json:"email" validate:"required,email"`        // email required.
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (b *User) TableName() string {

	return "user"
}
