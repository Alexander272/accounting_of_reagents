import { Size } from '@/constants/defaultValues'
import type { IParams } from '../types/data'

export const buildSiUrlParams = (req: IParams): URLSearchParams => {
	const params: string[][] = []

	if (req.page && req.page != 1) params.push(['page', req.page.toString()])
	if (req.size && req.size != Size) params.push(['size', req.size.toString()])

	if (req.search?.value) {
		params.push([`search[${req.search.fields.join(',')}]`, req.search.value])
	}

	if (Object.keys(req.sort).length) {
		const sort: string[] = []
		Object.keys(req.sort).forEach(k => {
			sort.push(`${req.sort[k] == 'DESC' ? '-' : ''}${k}`)
		})
		params.push(['sort_by', sort.join(',')])
	}

	// if (req.filter) {
	// 	req.filter.forEach(f => {
	// 		params.push([`filters[${f.field}]`, f.fieldType])
	// 		f.values.forEach(v => {
	// 			params.push([`${f.field}[${v.compareType}]`, v.value])
	// 		})
	// 	})
	// }

	return new URLSearchParams(params)
}
