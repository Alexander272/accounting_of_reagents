export const formatAmount = (value: number, unit: string) =>
	`${value.toLocaleString('ru-RU', { maximumFractionDigits: 5 })}${unit ? ` ${unit}` : ''}`

export const formatNumber = (value: number) => value.toLocaleString('ru-RU', { maximumFractionDigits: 5 })

export const csvCell = (value: string) => `"${value.replace(/"/g, '""')}"`
