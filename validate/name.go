package validate

import (
	"encoding/json"
	"errors"
	"os"
	"regexp"
	"sync"

	"github.com/x1rh/gopkg/fsx"
)

var (
	reservedUsernameMapping = make(map[string]bool)
	once                    sync.Once
	usernameReg             = regexp.MustCompile(`^[a-z0-9._-]{4,30}$`)
)

var (
	ErrFileNotExists = errors.New("file not exists")
)

type NameValidator struct{}

func NewNameValidator(path string) (*NameValidator, error) {
	if fsx.CheckFileExist(path) {
		return nil, ErrFileNotExists
	}

	once.Do(func() {
		initReservedUsername(path)
	})

	return &NameValidator{}, nil
}

func initReservedUsername(path string) {
	reservedUsernamesJsonFile, err := os.ReadFile(path)
	if err == nil {
		var usernames []string
		_ = json.Unmarshal(reservedUsernamesJsonFile, &usernames)
		for _, username := range usernames {
			reservedUsernameMapping[username] = true
		}
	}
}

// IsReserved checks whether the username is reserved
func (nv *NameValidator) IsReserved(username string) bool {
	return reservedUsernameMapping[username]
}

func (nv *NameValidator) IsInvalid(username string) bool {
	return !usernameReg.MatchString(username)
}
