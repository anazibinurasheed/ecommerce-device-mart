package helper

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/anazibinurasheed/project-device-mart/pkg/util/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserIDFromContext(c *gin.Context) (int, error) {
	userIDStr := c.GetString("userId")
	userID, err := strconv.Atoi(userIDStr)
	return userID, err
}

func SetToCookie(Data int, cookieName string, c *gin.Context) {

	maxAge := int(time.Now().Add(time.Minute * 6).Unix())
	c.SetCookie(cookieName, fmt.Sprint(Data), maxAge, "", "", false, true)
	c.SetSameSite(http.SameSiteLaxMode)
}

func DeleteCookie(cookieName string, c *gin.Context) {

	c.SetCookie(cookieName, "", -1, "", "", false, true)
}

func GenerateUniqueID() string {
	id := uuid.New().String()
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%s-%d", id, timestamp)
}

func MakeSKU(name string) string {
	name = strings.ReplaceAll(name, " ", "-")
	return name
}

func SetupDB(db *gorm.DB) error {
	err := setupStates(db)
	if err != nil {
		return err
	}

	fmt.Println("db is now ok our accept requests")
	return nil
}

func setupStates(db *gorm.DB) error {
	var states = make([]response.States, 0)
	query := `SELECT * FROM states;`
	err := db.Raw(query).Scan(&states).Error
	if err != nil {
		return fmt.Errorf("failed while getting state data from the db : %s", err)
	}

	// 28 states needed
	if len(states) == 27 {
		return nil
	}

	if len(states) < 27 {
		err = checkupStates(db, states)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkupStates(db *gorm.DB, dbStates []response.States) error {
	statesNeeded := map[string]bool{
		"Andhra Pradesh":    false,
		"Arunachal Pradesh": false,
		"Assam":             false,
		"Bihar":             false,
		"Chhattisgarh":      false,
		"Goa":               false,
		"Gujarat":           false,
		"Haryana":           false,
		"Himachal Pradesh":  false,
		"Jharkhand":         false,
		"Karnataka":         false,
		"Kerala":            false,
		"Madhya Pradesh":    false,
		"Maharashtra":       false,
		"Manipur":           false,
		"Meghalaya":         false,
		"Mizoram":           false,
		"Nagaland":          false,
		"Odisha":            false,
		"Punjab":            false,
		"Rajasthan":         false,
		"Sikkim":            false,
		"Tamil Nadu":        false,
		"Telangana":         false,
		"Tripura":           false,
		"Uttar Pradesh":     false,
		"Uttarakhand":       false,
		"West Bengal":       false,
	}

	for _, val := range dbStates {
		statesNeeded[val.Name] = true
	}

	var wg *sync.WaitGroup
	var count = 0
	errChan := make(chan error, len(statesNeeded))

	for stateName, ok := range statesNeeded {

		if !ok {
			count++
			wg.Add(count)
			go func(stateName string) {
				defer wg.Done()
				err := insertState(db, stateName)
				if err != nil {
					errChan <- err
				}

			}(stateName)
		}

	}

	wg.Wait()
	return nil

}

func insertState(db *gorm.DB, stateName string) (err error) {

	var insertedState response.States
	query := `INSERT INTO states (name)VALUES($1) RETURNING name ;`
	err = db.Raw(query, stateName).Scan(&insertedState).Error
	if err != nil {
		return fmt.Errorf("failed while executing insertState( ) first query execution : %s", err)
	}

	if stateName != insertedState.Name {
		return fmt.Errorf("failed while validating inserted state")
	}
	return nil
}

// func SetupStates(db *gorm.DB)
// func SetupStates(db *gorm.DB)
