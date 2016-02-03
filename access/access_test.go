package access_test

import (
	"errors"
	"testing"

	"upspin.googlesource.com/upspin.git/access"
	"upspin.googlesource.com/upspin.git/upspin"
)

func TestSwitch(t *testing.T) {
	// These should succeed.
	if err := access.Switch.RegisterUser("dummy", &dummyUser{}); err != nil {
		t.Errorf("registerUser failed")
	}
	if err := access.Switch.RegisterStore("dummy", &dummyStore{}); err != nil {
		t.Errorf("registerStore failed")
	}
	if err := access.Switch.RegisterDirectory("dummy", &dummyDirectory{}); err != nil {
		t.Errorf("registerDirectory failed")
	}

	// These should fail.
	if err := access.Switch.RegisterUser("dummy", &dummyUser{}); err == nil {
		t.Errorf("registerUser should have failed")
	}
	if err := access.Switch.RegisterStore("dummy", &dummyStore{}); err == nil {
		t.Errorf("registerStore should have failed")
	}
	if err := access.Switch.RegisterDirectory("dummy", &dummyDirectory{}); err == nil {
		t.Errorf("registerDirectory should have failed")
	}

	// These should return different NetAddrs
	s1, _ := access.Switch.BindStore(nil, upspin.Endpoint{Transport: "dummy", NetAddr: "addr1"})
	s2, _ := access.Switch.BindStore(nil, upspin.Endpoint{Transport: "dummy", NetAddr: "addr2"})
	if s1.Endpoint().NetAddr != "addr1" || s2.Endpoint().NetAddr != "addr2" {
		t.Errorf("got %s %s, expected addr1 addr2", s1.Endpoint().NetAddr, s2.Endpoint().NetAddr)
	}

	// This should fail.
	if _, err := access.Switch.BindStore(nil, upspin.Endpoint{Transport: "undefined"}); err == nil {
		t.Errorf("expected BindStore of undefined to fail")
	}
}

// Some dummy interfaces.
type dummyUser struct {
	endpoint upspin.Endpoint
}
type dummyStore struct {
	endpoint upspin.Endpoint
}
type dummyDirectory struct {
	endpoint upspin.Endpoint
}
type dummyContext int

func (d *dummyContext) Name() string {
	return "george"
}

func (d *dummyUser) Lookup(userName upspin.UserName) ([]upspin.Endpoint, error) {
	return nil, errors.New("dummyUser.Lookup not implemented")
}
func (d *dummyUser) Dial(cc upspin.ClientContext, e upspin.Endpoint) (interface{}, error) {
	user := &dummyUser{endpoint: e}
	return user, nil
}
func (d *dummyUser) ServerUserName() string {
	return "userUser"
}

func (d *dummyStore) Get(location upspin.Location) ([]byte, []upspin.Location, error) {
	return nil, nil, errors.New("dummyStore.Get not implemented")
}
func (d *dummyStore) Put(ref upspin.Reference, data []byte) (upspin.Location, error) {
	return upspin.Location{}, errors.New("dummyStore.Put not implemented")
}
func (d *dummyStore) Dial(cc upspin.ClientContext, e upspin.Endpoint) (interface{}, error) {
	store := &dummyStore{endpoint: e}
	return store, nil
}
func (d *dummyStore) Endpoint() upspin.Endpoint {
	return d.endpoint
}
func (d *dummyStore) ServerUserName() string {
	return "userStore"
}

func (d *dummyDirectory) Lookup(name upspin.PathName) (*upspin.DirEntry, error) {
	return nil, errors.New("dummyDirectory.Lookup not implemented")
}
func (d *dummyDirectory) Put(name upspin.PathName, data, packdata []byte) (upspin.Location, error) {
	return upspin.Location{}, errors.New("dummyDirectory.Lookup not implemented")
}
func (d *dummyDirectory) MakeDirectory(dirName upspin.PathName) (upspin.Location, error) {
	return upspin.Location{}, errors.New("dummyDirectory.MakeDirectory not implemented")
}
func (d *dummyDirectory) Glob(pattern string) ([]*upspin.DirEntry, error) {
	return nil, errors.New("dummyDirectory.GLob not implemented")
}
func (d *dummyDirectory) Dial(cc upspin.ClientContext, e upspin.Endpoint) (interface{}, error) {
	dir := &dummyDirectory{endpoint: e}
	return dir, nil
}
func (d *dummyDirectory) ServerUserName() string {
	return "userDirectory"
}