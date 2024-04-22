export type ReagentType = 'main' | 'auxiliary' | 'experimental' | 'standard' | 'titling'

export interface IBaseInfForm {
	type: ReagentType
	name: string
	uname: string
	document: string
	purity: string
	dateOfManufacture: number
	consignment: string
	manufacturer: string
	shelfLife: number
	place_closet: string
	place_shelf: number
}

// (Вид реактива, Наименование по сопроводительной документации, Наименование по ИЮПАК, Нормативный документ, Степень чистоты, Дата изготовления, Партия, Производитель, Срок годности (в месяцах), Место хранения (шкаф и полка)
