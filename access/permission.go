package access

import "errors"

type Permission struct {
	UserID    string   `json:"userId"`
	Resources []string `json:"resources"`
}

func NewPermission(userId string, loader PermissionLoader) (Permission, error) {
	defer loader.Close()

	if loader == nil {
		return Permission{}, errors.New("PermissionLoader is nil")
	}

	var p Permission
	var err error

	p.UserID = userId
	p.Resources, err = loader.PermissioLoad(userId)
	if err != nil {
		return Permission{}, err
	}

	return p, nil
}

func (p Permission) CheckPermission(resource string) bool {
	for _, r := range p.Resources {
		if r == resource {
			return true
		}
	}
	return false
}
