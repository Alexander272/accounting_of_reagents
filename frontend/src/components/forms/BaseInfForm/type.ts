export interface IBaseInfForm {
	type: string
	name: string
	uname: string
	document: string
	purity: string
	dateOfManufacture: number
	consignment: string
	manufacturer: string
	shelfLife: string
	place_closet: string
	place_shelf: string
}

// (Вид реактива, Наименование по сопроводительной документации, Наименование по ИЮПАК, Нормативный документ, Степень чистоты, Дата изготовления, Партия, Производитель, Срок годности (в месяцах), Место хранения (шкаф и полка)
