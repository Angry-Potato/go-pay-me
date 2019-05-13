package payments

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	//postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ValidationError describes the incorrectness of a payment operation
type ValidationError struct {
	err string
}

func (e *ValidationError) Error() string {
	return e.err
}

// All payments
func All(DB *gorm.DB) ([]Payment, error) {
	allPayments := []Payment{}
	err := DB.Find(&allPayments).Error
	return allPayments, err
}

// Create a new payment
func Create(DB *gorm.DB, payment *Payment) (*Payment, error) {
	validationErrors := Validate(payment)
	if len(validationErrors) == 0 {
		if err := DB.Create(&payment).Error; err != nil {
			return nil, err
		}
		return payment, nil
	}

	return nil, consolidateValidationErrors(validationErrors, fmt.Sprintf("Validation errors whilst creating payment with id %s", payment.ID))
}

// DeleteAll payments
func DeleteAll(DB *gorm.DB) error {
	allPayments := []Payment{}
	err := DB.Delete(&allPayments).Error
	return err
}

// SetAll payments
func SetAll(DB *gorm.DB, payments []Payment) ([]Payment, error) {
	var consolidatedValidation []error
	for _, payment := range payments {
		validationErrors := Validate(&payment)
		if len(validationErrors) != 0 {
			consolidatedValidation = append(consolidatedValidation, consolidateValidationErrors(validationErrors, fmt.Sprintf("Validation errors whilst creating payment with id %s", payment.ID)))
		}
	}

	if len(consolidatedValidation) == 0 {
		DB.Lock()
		defer DB.Unlock()
		err := DeleteAll(DB)
		if err != nil {
			return nil, err
		}
		count, err := batchInsert(DB, payments)
		if err != nil {
			return nil, err
		} else if count != int64(len(payments)) {
			return nil, fmt.Errorf("Could only insert %d out of %d", count, len(payments))
		}
		newPayments, err := All(DB)
		if err != nil {
			return nil, err
		}
		return newPayments, nil
	}

	return nil, consolidateValidationErrors(consolidatedValidation, "Errors")
}

// Get a payment by ID
func Get(DB *gorm.DB, ID string) (*Payment, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, &ValidationError{err.Error()}
	}

	payment := Payment{}
	if err = DB.Where(&Payment{ID: ID}).First(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

// Delete a payment by ID
func Delete(DB *gorm.DB, ID string) error {
	err := ValidateID(ID)
	if err != nil {
		return &ValidationError{err.Error()}
	}

	DB = DB.Delete(&Payment{ID: ID})
	if err = DB.Error; err != nil {
		return err
	}
	if DB.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// Update a payment by ID
func Update(DB *gorm.DB, ID string, payment *Payment) (*Payment, error) {
	err := ValidateID(ID)
	if err != nil {
		return nil, &ValidationError{err.Error()}
	}
	errs := Validate(payment)
	if len(errs) != 0 {
		return nil, consolidateValidationErrors(errs, fmt.Sprintf("Validation errors whilst Updating payment %s", payment.ID))
	}

	exisingPayment, err := Get(DB, ID)
	if err != nil {
		return nil, err
	}
	if *exisingPayment == *payment {
		return nil, nil
	}

	DB = DB.Save(payment)
	if err = DB.Error; err != nil {
		return nil, err
	}
	if DB.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return payment, nil
}

func consolidateValidationErrors(errs []error, prefix string) error {
	var errstrings []string
	for _, err := range errs {
		errstrings = append(errstrings, err.Error())
	}
	return &ValidationError{fmt.Sprintf("%s: %s", prefix, strings.Join(errstrings, ", "))}
}

// shamelessly stolen from https://github.com/jinzhu/gorm/issues/255#issuecomment-481159929
func batchInsert(DB *gorm.DB, objArr []Payment) (int64, error) {
	// If there is no data, nothing to do.
	if len(objArr) == 0 {
		return 0, errors.New("insert a slice length of 0")
	}

	mainObj := objArr[0]
	mainScope := DB.NewScope(mainObj)
	mainFields := mainScope.Fields()
	quoted := make([]string, 0, len(mainFields))
	for i := range mainFields {
		// If primary key has blank value (0 for int, "" for string, nil for interface ...), skip it.
		// If field is ignore field, skip it.
		if (mainFields[i].IsPrimaryKey && mainFields[i].IsBlank) || (mainFields[i].IsIgnored) {
			continue
		}
		quoted = append(quoted, mainScope.Quote(mainFields[i].DBName))
	}

	placeholdersArr := make([]string, 0, len(objArr))

	for _, obj := range objArr {
		scope := DB.NewScope(obj)
		fields := scope.Fields()
		placeholders := make([]string, 0, len(fields))
		for i := range fields {
			if (fields[i].IsPrimaryKey && fields[i].IsBlank) || (fields[i].IsIgnored) {
				continue
			}
			var vars interface{}
			if (fields[i].Name == "CreatedAt" || fields[i].Name == "UpdatedAt") && fields[i].IsBlank {
				vars = gorm.NowFunc()
			} else {
				vars = fields[i].Field.Interface()
			}
			placeholders = append(placeholders, mainScope.AddToVars(vars))
		}
		placeholdersStr := "(" + strings.Join(placeholders, ", ") + ")"
		placeholdersArr = append(placeholdersArr, placeholdersStr)
		// add real variables for the replacement of placeholders' '?' letter later.
		mainScope.SQLVars = append(mainScope.SQLVars, scope.SQLVars...)
	}
	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s",
		mainScope.QuotedTableName(),
		strings.Join(quoted, ", "),
		strings.Join(placeholdersArr, ", "),
	)
	log.Println(sql)
	mainScope.Raw(sql)
	//Execute and Log
	if err := mainScope.Exec().DB().Error; err != nil {
		return 0, err
	}
	return mainScope.DB().RowsAffected, nil
}
