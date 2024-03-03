package entities

import (
	"errors"
	"regexp"
)

type Manager struct {
	ID       int
	Phone    string
	Password string
}

func (m *Manager) validate() error {
	// 필수값 체크
	if m.Phone == "" || m.Password == "" {
		return errors.New("필수값을 입력해주세요.")
	}
	// 휴대폰번호 정규식 체크
	if !isValidPhoneNumber(m.Phone) {
		return errors.New("올바른 휴대폰번호 형식이 아닙니다. 01012345678 형태로 입력해주세요.")
	}
	return nil
}

func NewManager(id int, phone string, password string) *Manager {
	return &Manager{
		ID: id,
		Phone: phone,
		Password: password,
	}
}

func isValidPhoneNumber(phoneNumber string) bool {
	//^: 문자열의 시작
	//01: 휴대폰 번호가 010으로 시작
	//[0-9]{8,9}: 8 또는 9자리의 숫자 그룹
	//$: 문자열의 끝
	regexPattern := `^01[0-9]{8,9}$`
	regex := regexp.MustCompile(regexPattern)
	return regex.MatchString(phoneNumber)
}