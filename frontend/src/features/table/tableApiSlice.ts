import type { ICreateDataItem, IDataItem, IParams } from './types/data'
import { apiSlice } from '@/app/apiSlice'
import { API } from '@/app/api'
import { buildSiUrlParams } from './utils/buildUrlParams'

const tableApiSlice = apiSlice.injectEndpoints({
	overrideExisting: false,
	endpoints: builder => ({
		getAll: builder.query<{ data: IDataItem[]; total: number }, IParams>({
			query: params => ({
				url: API.reagents,
				method: 'GET',
				params: buildSiUrlParams(params),
			}),
			providesTags: [{ type: 'DataItems', id: 'ALL' }],
		}),
		// не знаю правильно ли получать данные одним куском
		getById: builder.query<{ data: IDataItem }, string>({
			query: id => `${API.reagents}/${id}`,
			providesTags: (_res, _err, id) => [{ type: 'DataItems', id: id }],
		}),

		create: builder.mutation<null, ICreateDataItem>({
			query: data => ({
				url: API.reagents,
				method: 'POST',
				body: data,
			}),
			invalidatesTags: [{ type: 'DataItems', id: 'ALL' }],
		}),
	}),

	// TODO когда удается реагент можно делать ему метку, а потом проверять если метка старше 30 дней, то запись удаляется
})

export const { useGetAllQuery, useGetByIdQuery, useCreateMutation } = tableApiSlice
