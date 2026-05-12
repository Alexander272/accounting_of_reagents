package access

const (
	ResourcePrivateReagent ResourceSlug = "private_reagent"
	ResourcePublicReagent  ResourceSlug = "public_reagent"
	ResourceAmountTypes    ResourceSlug = "amount_types"
	ResourceReagentTypes   ResourceSlug = "reagent_types"

	ResourceExtending ResourceSlug = "extending"
	ResourceSpending  ResourceSlug = "spending"

	ResourceRealm ResourceSlug = "realm"
	ResourcePerm  ResourceSlug = "permission"
	ResourceRoles ResourceSlug = "roles"
	ResourceUsers ResourceSlug = "users"
)

var OrderOfResources = map[ResourceSlug]int{
	ResourcePublicReagent:  1,
	ResourcePrivateReagent: 2,
	ResourceAmountTypes:    3,
	ResourceReagentTypes:   4,
	ResourceExtending:      5,
	ResourceSpending:       6,
	ResourceRealm:          9,
	ResourceUsers:          10,
	ResourceRoles:          11,
	ResourcePerm:           12,
}

var Reg = NewRegistry(
	Resource{
		Slug:           ResourcePublicReagent,
		Name:           "Общедоступные реагенты",
		Group:          "Операции",
		Description:    "Управление общими реагентами",
		AllowedActions: actions(All),
	},
	Resource{
		Slug:           ResourcePrivateReagent,
		Name:           "Реагенты",
		Group:          "Операции",
		Description:    "Управление реагентами",
		AllowedActions: actions(All),
	},

	Resource{
		Slug:           ResourceAmountTypes,
		Name:           "Типы единиц измерений",
		Group:          "Справочники",
		Description:    "Управление типами единиц измерений",
		AllowedActions: actions(All),
	},
	Resource{
		Slug:           ResourceReagentTypes,
		Name:           "Типы реагентов",
		Group:          "Справочники",
		Description:    "Управление типами реагентов",
		AllowedActions: actions(All),
	},

	Resource{
		Slug:           ResourceExtending,
		Name:           "Продление срока годности",
		Group:          "Операции",
		Description:    "",
		AllowedActions: actions(All),
	},
	Resource{
		Slug:           ResourceSpending,
		Name:           "Расходование реактивов",
		Group:          "Операции",
		Description:    "",
		AllowedActions: actions(All),
	},

	Resource{
		Slug:           ResourceRealm,
		Name:           "Области",
		Group:          "Администрирование",
		Description:    "Управление областями",
		AllowedActions: actions(All),
	},
	Resource{
		Slug:           ResourceUsers,
		Name:           "Пользователи",
		Group:          "Администрирование",
		Description:    "Управление пользователями",
		AllowedActions: actions(Read, Write),
	},

	Resource{
		Slug:           ResourceRoles,
		Name:           "Роли",
		Group:          "Администрирование",
		Description:    "Управление ролями пользователей",
		AllowedActions: actions(All),
	},

	Resource{
		Slug:           ResourcePerm,
		Name:           "Права",
		Group:          "Администрирование",
		Description:    "Управление правами пользователей",
		AllowedActions: actions(All),
	},
)
