import { RegisterOptions } from 'react-hook-form'

export type FieldBase<Keys, Type, ExtraProps> = {
	key: Keys
	label: string
	type: Type
	rules?: RegisterOptions
} & ExtraProps

export type StringType = 'String'
export type NumberType = 'Number'
export type DateType = 'Date'
export type FileType = 'File'
export type BoolType = 'Bool'
export type ListType = 'List'
export type LinkType = 'Link'
export type RadioType = 'Radio'

export type Options = {
	value: string
	label: string
}

export type StringField<Keys> = FieldBase<Keys, StringType, { multiline?: boolean; minRows?: number }>
export type NumberField<Keys> = FieldBase<Keys, NumberType, unknown>
export type DateField<Keys> = FieldBase<Keys, DateType, unknown>
export type FileField<Keys> = FieldBase<Keys, FileType, unknown>
export type BoolField<Keys> = FieldBase<Keys, BoolType, unknown>
export type ListField<Keys> = FieldBase<Keys, ListType, { options: Options[] }>
export type LinkField<Keys> = FieldBase<Keys, LinkType, unknown>
export type RadioField<Keys> = FieldBase<Keys, RadioType, { options: Options[] }>

export type Field<Keys> =
	| StringField<Keys>
	| NumberField<Keys>
	| DateField<Keys>
	| FileField<Keys>
	| BoolField<Keys>
	| ListField<Keys>
	| LinkField<Keys>
	| RadioField<Keys>
