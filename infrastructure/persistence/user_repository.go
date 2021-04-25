package persistence

import (
	"context"
	"fmt"

	"danglingmind.com/ddd/v1/domain/entity"
	"danglingmind.com/ddd/v1/domain/repository"
	"danglingmind.com/ddd/v1/utils"
)

type UserRepo struct {
	db Database
}

// UserRepo implements user repository
var _ repository.UserRepository = &UserRepo{}

func NewUserRepository(db Database) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) GetById(ctx context.Context, id uint64) (*entity.User, error) {
	row, err := u.db.QueryRow(ctx, "select * from USERS where id=?", id)
	if err != nil {
		return nil, err
	}

	user := &entity.User{}
	err = row.Serialize(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) GetAll(ctx context.Context) ([]entity.User, error) {
	rows, err := u.db.Query(ctx, "select * from USERS")
	if err != nil {
		return nil, err
	}

	users := make([]entity.User, 0)
	for _, row := range rows {
		user := entity.User{}
		err = row.Serialize2(&user)
		// err = mapstructure.Decode(row, &user)
		fmt.Println(row)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepo) Update(ctx context.Context, id uint64, values map[string]interface{}) error {
	return nil
}

// TODO: use string builder here
func (u *UserRepo) Save(ctx context.Context, user entity.User) error {
	// senitize user struct
	user.PrepareToSave()
	fieldsMap := utils.GetJsonTagsWithValues2(user)
	args := make([]interface{}, 0)
	fieldNameString := ""
	fieldNamePlaceholder := ""
	i := 0
	for fieldJsonName, fieldValue := range fieldsMap {
		i++
		args = append(args, fieldValue)
		fieldNameString = fieldNameString + fieldJsonName
		fieldNamePlaceholder = fieldNamePlaceholder + "?"
		if i < len(fieldsMap) {
			fieldNameString = fieldNameString + ","
			fieldNamePlaceholder = fieldNamePlaceholder + ","
		}
	}
	insertStmt := "insert into USERS (" + fieldNameString + ") values (" + fieldNamePlaceholder + ")"

	return u.db.Save(ctx, insertStmt, args...)
}
