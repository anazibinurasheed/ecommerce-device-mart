package helper

import (
	"fmt"
	"sync"
	"time"
)

type data struct {
	Phone      string
	IsVerified bool
}

// A phone implements the user auth related fields
type phone struct {
	Details map[string]data //store phone for otp verification
	Mutex   *sync.Mutex     // unMutex is mutex for unverifiedData

}

// NewPhone returns a initialized new phone
func NewPhone() phone {

	return phone{
		Details: make(map[string]data),
		Mutex:   &sync.Mutex{},
	}
}

// Clean deletes the details after a specific duration if user is not verified
func (p *phone) Clean(uuid string) {

	time.Sleep(1 * time.Minute)
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	if !p.Details[uuid].IsVerified {
		delete(p.Details, uuid)
	}

}

// Delete deletes specified data
func (p *phone) Delete(uuid string) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	delete(p.Details, uuid)
}

// Verified mark the Details as verified
func (p *phone) Verified(uuid, phone string) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	p.Details[uuid] = data{
		Phone:      phone,
		IsVerified: true,
	}
}

// NotVerified sets the Details as not verified
func (p *phone) NotVerified(uuid, phone string) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	p.Details[uuid] = data{
		Phone:      phone,
		IsVerified: false,
	}
}

// Set sets the phone to Details
func (p *phone) Set(uuid, phone string) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	p.Details[uuid] = data{
		Phone: phone,
	}
}

// Get gets the phone
func (p *phone) Get(uuid string) (phone string, ok bool, isVerified bool) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()
	v, ok := p.Details[uuid]
	phone = v.Phone

	return phone, ok, v.IsVerified
}

// Print prints the specific Detail
func (p *phone) Print(uuid string) {
	p.Mutex.Lock()
	fmt.Println(p.Details)
	p.Mutex.Unlock()
}
