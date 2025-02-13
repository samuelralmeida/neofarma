package user

import "fmt"

type HierarchyRole int

const (
	AdminRoleName                = "admin"
	AdminHierarchy HierarchyRole = 30

	ClientRoleName                = "client"
	ClientHierarchy HierarchyRole = 10

	SellerRoleName                = "seller"
	SellerHierarchy HierarchyRole = 15
)

type Permission struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Role struct {
	Name        string        `json:"name"`
	Hierarchy   HierarchyRole `json:"-"`
	Permissions []Permission  `json:"permissions"`
}

var roles = map[string]Role{
	AdminRoleName:  {Name: AdminRoleName, Hierarchy: AdminHierarchy},
	ClientRoleName: {Name: ClientRoleName, Hierarchy: ClientHierarchy},
	SellerRoleName: {Name: SellerRoleName, Hierarchy: SellerHierarchy},
}

func (r *Role) Description() string {
	return fmt.Sprintf("Role: %s, Hierarchy: %d", r.Name, r.Hierarchy)
}

func GetRoleByName(roleName string) (Role, bool) {
	r, exists := roles[roleName]
	return r, exists
}

func (r *Role) Validate() error {
	if len(r.Permissions) == 0 {
		return fmt.Errorf("role %s must have at least one permission", r.Name)
	}
	return nil
}

func (r *Role) HasPermission(permissionName string) bool {
	for _, permission := range r.Permissions {
		if permission.Name == permissionName {
			return true
		}
	}
	return false
}

func (r *Role) HasHierarchy(hierarchyTarget HierarchyRole) bool {
	return r.Hierarchy >= hierarchyTarget
}
