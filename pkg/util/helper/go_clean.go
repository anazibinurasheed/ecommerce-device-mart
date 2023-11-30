package helper

import (
	"fmt"
	"sync"
	"time"
)

// A passwordManager implements passwordManager related functionalities.
type passwordManager struct {
	details map[int]string
	mutex   *sync.Mutex
}

func NewPasswordManager() passwordManager {

	return passwordManager{
		details: make(map[int]string),
		mutex:   &sync.Mutex{},
	}
}

func (p *passwordManager) Set(userID int, uuid string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.details[userID] = uuid
}

func (p *passwordManager) Remove(userID int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	delete(p.details, userID)
}

func (p *passwordManager) Check(uuid string, userID int) (ok bool) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if userID == 0 {
		return false
	}
	val, ok := p.details[userID]
	fmt.Println(val)

	return val == uuid
}

// NewPhone returns a initialized new phone
func NewPhone() phone {

	return phone{
		details: make(map[string]data),
		mutex:   &sync.Mutex{},
	}
}

type data struct {
	Phone      string
	IsVerified bool
}

// A phone implements the user auth related functionalities.
type phone struct {
	details map[string]data //store phone for otp verification
	mutex   *sync.Mutex     // unMutex is mutex for unverifiedData

}

// Clean deletes the details after a specific duration if user is not verified
func (p *phone) Clean(uuid string) {

	time.Sleep(3 * time.Minute)
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if !p.details[uuid].IsVerified {
		delete(p.details, uuid)
	}

}

// Delete deletes specified data
func (p *phone) Delete(uuid string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	delete(p.details, uuid)
}

// Verified mark the Details as verified
func (p *phone) Verified(uuid, phone string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.details[uuid] = data{
		Phone:      phone,
		IsVerified: true,
	}

}

// NotVerified sets the Details as not verified
func (p *phone) NotVerified(uuid, phone string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.details[uuid] = data{
		Phone:      phone,
		IsVerified: false,
	}
}

// Set sets the phone to Details
func (p *phone) Set(uuid, phone string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.details[uuid] = data{
		Phone: phone,
	}
}

// Get gets the phone
func (p *phone) Get(uuid string) (phone string, ok bool, isVerified bool) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	v, ok := p.details[uuid]
	phone = v.Phone

	return phone, ok, v.IsVerified
}

// Print prints the specific Detail
func (p *phone) Print(uuid string) {
	p.mutex.Lock()
	fmt.Println(p.details)
	p.mutex.Unlock()
}
