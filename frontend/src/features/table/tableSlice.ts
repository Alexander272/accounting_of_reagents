import { type PayloadAction, createSlice } from '@reduxjs/toolkit'

import type { ISort, ISearch, IFilter } from './types/data'
import type { ICoordinates } from './types/table'
import { Size } from '@/constants/defaultValues'
import { RootState } from '@/app/store'
import { localKeys } from './constants/localKeys'

interface ITableSlice {
	page: number
	size: number
	sort: ISort
	filters: IFilter[]
	search: ISearch
	selected: { [id: string]: boolean }
	contextMenu?: ICoordinates
}

const initialState: ITableSlice = {
	page: +(localStorage.getItem(localKeys.page) || 1),
	size: +(localStorage.getItem(localKeys.size) || Size),
	sort: {},
	filters: [],
	search: {
		value: '',
		fields: ['name', 'uname'],
	},
	selected: {},
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

		setFilters: (state, action: PayloadAction<IFilter[]>) => {
			state.filters = action.payload
		},

		setSearch: (state, action: PayloadAction<string>) => {
			state.search.value = action.payload
		},
		setSearchFields: (state, action: PayloadAction<string[]>) => {
			state.search.fields = action.payload
		},

		setSelected: (state, action: PayloadAction<string | string[] | undefined>) => {
			if (action.payload) {
				if (typeof action.payload == 'string') {
					if (state.selected[action.payload]) delete state.selected[action.payload]
					else state.selected[action.payload] = true
				} else {
					state.selected = action.payload.reduce((a, v) => ({ ...a, [v]: true }), {})
				}
			} else state.selected = {}
		},

		setContextMenu: (state, action: PayloadAction<ICoordinates | undefined>) => {
			state.contextMenu = action.payload
		},
	},
})

export const getTablePage = (state: RootState) => state.table.page
export const getTableSize = (state: RootState) => state.table.size
export const getTableSort = (state: RootState) => state.table.sort
export const getFilters = (state: RootState) => state.table.filters
export const getSearch = (state: RootState) => state.table.search
export const getSelected = (state: RootState) => state.table.selected
export const getContextMenu = (state: RootState) => state.table.contextMenu

export const tablePath = tableSlice.name
export const tableReducer = tableSlice.reducer

export const { setPage, setSize, setSort, setSearch, setFilters, setSearchFields, setSelected, setContextMenu } =
	tableSlice.actions
