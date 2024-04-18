import type { FC } from 'react'

import type { IColumn, IGroup } from './types'
import { TableHeadContainer } from './style'

type Props = {
	columns: IColumn[]
	groups?: IGroup[]
}

export const TableHeader: FC<Props> = ({ columns, groups }) => {
	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	const header: any[] = []

	// columns.forEach(c => {
	// 	if (!c.visible) continue

	// 	header.push(`${c.name} ${c.group}`)
	// })

	for (let i = 0; i < columns.length; i++) {
		const c = columns[i]
		if (c.visible === false) continue

		if (c.group) {
			const group = groups?.find(g => g.name == c.group)
			if (!group) continue

			if (header[header.length - 1].name == group.name) {
				header[header.length - 1].children.push(c)
			} else {
				group.children = [c]
				header.push(group)
			}
		} else {
			header.push(c)
		}
	}

	console.log('header', header)

	return <TableHeadContainer>TableHeader</TableHeadContainer>
}
