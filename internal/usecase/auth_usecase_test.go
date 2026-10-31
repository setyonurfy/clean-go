package usecase

import (
	"testing"

	"clean-go/internal/entity"
	"clean-go/internal/repository"
)

// repository palsu yang menyimpan data di memori
type mockRepo struct {
	users map[string]*entity.User
}

func newMockRepo() *mockRepo {
	return &mockRepo{users: map[string]*entity.User{}}
}

func (m *mockRepo) All() ([]*entity.User, error) {
	var list []*entity.User
	for _, u := range m.users {
		list = append(list, u)
	}
	return list, nil
}

func (m *mockRepo) GetByEmail(email string) (*entity.User, error) {
	u, ok := m.users[email]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return u, nil
}

func (m *mockRepo) GetById(id string) (*entity.User, error) {
	for _, u := range m.users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, repository.ErrNotFound
}

func (m *mockRepo) Create(u *entity.User) error {
	u.ID = "1"
	m.users[u.Email] = u
	return nil
}

func TestRegisterAndLogin(t *testing.T) {
	uc := NewAuthUsecase(newMockRepo(), "secret-untuk-test")

	if err := uc.Register("budi", "budi@example.com", "secret123"); err != nil {
		t.Fatalf("register seharusnya berhasil, dapat error: %v", err)
	}

	token, err := uc.Login("budi@example.com", "secret123")
	if err != nil || token == "" {
		t.Fatalf("login seharusnya berhasil, dapat token=%q err=%v", token, err)
	}
}

func TestRegisterDuplicateEmail(t *testing.T) {
	uc := NewAuthUsecase(newMockRepo(), "secret-untuk-test")

	_ = uc.Register("budi", "budi@example.com", "secret123")
	err := uc.Register("budi2", "budi@example.com", "lain123")
	if err == nil || err.Error() != "email already exists" {
		t.Fatalf("seharusnya error 'email already exists', dapat: %v", err)
	}
}

func TestLoginWrongPassword(t *testing.T) {
	uc := NewAuthUsecase(newMockRepo(), "secret-untuk-test")

	_ = uc.Register("budi", "budi@example.com", "secret123")
	_, err := uc.Login("budi@example.com", "salah")
	if err == nil || err.Error() != "invalid email or password" {
		t.Fatalf("seharusnya error 'invalid email or password', dapat: %v", err)
	}
}

func TestLoginUnknownEmail(t *testing.T) {
	uc := NewAuthUsecase(newMockRepo(), "secret-untuk-test")

	_, err := uc.Login("tidakada@example.com", "secret123")
	if err == nil || err.Error() != "invalid email or password" {
		t.Fatalf("seharusnya error 'invalid email or password', dapat: %v", err)
	}
}
