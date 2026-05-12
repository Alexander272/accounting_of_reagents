export const API = {
	auth: {
		signIn: `auth/sign-in`,
		refresh: `auth/refresh`,
		signOut: `auth/sign-out`,
	},
	amountTypes: 'amount-types',
	reagentTypes: 'reagent-types',
	reagents: 'reagents',
	spending: 'spending',
	extending: 'extending',
	notes: 'notes',
	users: {
		base: 'users',
		sync: 'users/sync',
	},
	roles: {
		base: 'roles',
		permissions: (id: string) => `roles/${id}/permissions`,
		stats: 'roles/all/stats',
	},
	permissions: {
		base: 'permissions',
	},
	realms: {
		base: 'realms',
		byId: (id: string) => `realms/${id}`,
	},
}
