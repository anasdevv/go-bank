package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Storage interface{
	CreateAccount(*Account) error
	DeleteAccount(int) error
	GetAccountById(int) (*Account , error)
	UpdateAccount(*Account) error
	GetAccounts()([]*Account,error)
}

type PostgresStore	struct{
	db *gorm.DB
}


func (s *PostgresStore) Init() error {
	err := s.db.AutoMigrate(&Account{})
	if err != nil {
		fmt.Printf("AutoMigrate failed: +%v\n", err)
		return err
	}
	return nil
}



func (s * PostgresStore)CreateAccount(acc *Account) error{
	result := s.db.Create(acc)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s * PostgresStore) DeleteAccount(id int) error{
	 result := s.db.Delete(&Account{},id); if result.Error != nil {
		fmt.Printf("Error finding acount with id %d error: +%v\n",id,result.Error )
		return result.Error
	 }
	return nil
}

func (s *PostgresStore) GetAccountById(id int) (*Account, error) {
	var account Account
	result := s.db.First(&account, id)

	if result.Error != nil {
		fmt.Printf("Error in finding account with id %d: %v\n", id, result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		fmt.Printf("No account found with id %d\n", id)
		return nil, fmt.Errorf("no account found with id %d", id)
	}

	return &account, nil
}

func (s * PostgresStore) UpdateAccount(acc *Account) error {
	result := s.db.Save(acc); if result.Error != nil {
		fmt.Printf("Error in updating account %v\n" , result.Error)
		return result.Error
	}
	return nil
}


func(s *PostgresStore) GetAccounts() ([]*Account,error) {
	accounts := []*Account{}
	return accounts,nil
}

func NewPostgresStore()(*PostgresStore , error) {
	// connStr := "user=app dbname=etest sslmode=disable password=1q2w3e4r"
	dsn := "host=localhost user=app password=1q2w3e4r dbname=etest port=5432 sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "t_",
			SingularTable: true,
		},
	})

	if err != nil {
		return nil,err
	}
	
	return &PostgresStore{
		db: db,
	},nil
}
