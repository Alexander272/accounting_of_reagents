import { apiSlice } from '@/app/apiSlice'
import type { IDataItem, IParams } from './types/data'
import { buildSiUrlParams } from './utils/buildUrlParams'

const tableApiSlice = apiSlice.injectEndpoints({
	overrideExisting: false,
	endpoints: builder => ({
		getAll: builder.query<{ data: IDataItem[]; total: number }, IParams>({
			query: params => ({
				//TODO дописать api
				url: '',
				method: 'GET',
				params: buildSiUrlParams(params),
			}),
			providesTags: [{ type: 'DataItems', id: 'ALL' }],
		}),
	}),
	// TODO когда удается реагент можно делать ему метку, а потом проверять если метка старше 30 дней, то запись удаляется
})

export const { useGetAllQuery } = tableApiSlice
