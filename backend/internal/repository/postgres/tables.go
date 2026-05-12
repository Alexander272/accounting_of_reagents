package postgres

var Tables = struct {
	Permissions   string
	RolePermissions string
	Roles         string
	RoleHierarchy string
	Users         string
	Realms        string
	UserRealms    string
	ReagentTypes  string
	AmountType    string
	Reagents      string
	Spending      string
	Extending     string
	Notes         string
}{
	Permissions:   "permissions",
	RolePermissions: "role_permissions",
	Roles:         "roles",
	RoleHierarchy: "role_hierarchy",
	Users:         "users",
	Realms:        "realms",
	UserRealms:    "user_realms",
	ReagentTypes:  "reagent_types",
	AmountType:    "amount_types",
	Reagents:      "reagents",
	Spending:      "spending",
	Extending:     "extending",
	Notes:         "notes",
}
