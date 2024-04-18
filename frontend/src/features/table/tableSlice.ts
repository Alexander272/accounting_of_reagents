import { type PayloadAction, createSlice } from '@reduxjs/toolkit'

import type { ISort } from './types/table'
import { Size } from '@/constants/defaultValues'
import { RootState } from '@/app/store'
import { localKeys } from './constants/localKeys'

interface ITableSlice {
	page: number
	size: number
	sort: ISort
}

const initialState: ITableSlice = {
	page: +(localStorage.getItem(localKeys.page) || 1),
	size: +(localStorage.getItem(localKeys.size) || Size),
	sort: {},
}

const tableSlice = createSlice({
	name: 'table',
	initialState,
	reducers: {
		setPage: (state, action: PayloadAction<number>) => {
			state.page = action.payload
			localStorage.setItem(localKeys.page, action.payload.toString())
		},
		setSize: (state, action: PayloadAction<number>) => {
			state.size = action.payload
			localStorage.setItem(localKeys.size, action.payload.toString())
		},

		setSort: (state, action: PayloadAction<string>) => {
			if (!state.sort[action.payload]) {
				state.sort = { ...(state.sort || {}), [action.payload]: 'ASC' }
				return
			}

			if (state.sort[action.payload] == 'ASC') state.sort[action.payload] = 'DESC'
			else {
				delete state.sort[action.payload]
			}
		},
	},
})

export const getTablePage = (state: RootState) => state.table.page
export const getTableSize = (state: RootState) => state.table.size
export const getTableSort = (state: RootState) => state.table.sort

export const tablePath = tableSlice.name
export const tableReducer = tableSlice.reducer

export const { setPage, setSize, setSort } = tableSlice.actions
