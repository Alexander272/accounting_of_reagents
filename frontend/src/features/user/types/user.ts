import { IRealm } from '@/features/realms/types/realm'
import { IRole } from './role'

export interface IUser {
	id: string
	name: string
	role: string
	permissions: Record<string, string[]>
	token: string
	realms: IUserRealm[]
}

export interface IUserShort {
	id: string
	firstName: string
	lastName: string
	email: string
}

export interface IUserData {
	id: string
	username: string
	firstName: string
	lastName: string
	email: string
	isActive: boolean
	createdAt: string

	realms: IUserRealm[]
}

export interface IUserRealm {
	id: string
	userId: string
	realmId: string
	roleId: string
	realm?: IRealm
	role?: IRole
	isActive: boolean
	createdAt: string
}

export interface IUserDataDTO {
	id: string
	username: string
	firstName: string
	lastName: string
	email: string
	isActive: boolean

	realms: IUserRealm[]
}
