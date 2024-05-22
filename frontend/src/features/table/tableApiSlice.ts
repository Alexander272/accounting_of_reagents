import type { ICreateDataItem, IDataItem, IParams, IUpdateDataItem } from './types/data'
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
		getById: builder.query<{ data: IUpdateDataItem }, string>({
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
		update: builder.mutation<null, IUpdateDataItem>({
			query: data => ({
				url: `${API.reagents}/${data.id}`,
				method: 'PUT',
				body: data,
			}),
			invalidatesTags: (_res, _err, data) => [
				{ type: 'DataItems', id: data.id },
				{ type: 'DataItems', id: 'ALL' },
			],
		}),
		delete: builder.mutation<null, string>({
			query: id => ({
				url: `${API.reagents}/${id}`,
				method: 'DELETE',
			}),
			invalidatesTags: (_res, _err, id) => [
				{ type: 'DataItems', id: id },
				{ type: 'DataItems', id: 'ALL' },
			],
		}),

		prepareOrder: builder.mutation<null, { list: string[] }>({
			query: data => ({
				url: `${API.reagents}/order`,
				method: 'POST',
				body: data,
			}),
		}),
	}),

	// TODO когда удается реагент можно делать ему метку, а потом проверять если метка старше 30 дней, то запись удаляется
})

export const {
	useGetAllQuery,
	useGetByIdQuery,
	useCreateMutation,
	useUpdateMutation,
	useDeleteMutation,
	usePrepareOrderMutation,
} = tableApiSlice
