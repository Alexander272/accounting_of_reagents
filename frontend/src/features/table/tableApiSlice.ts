import { toast } from 'react-toastify'

import type { IBaseFetchError } from '@/app/types/error'
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
			onQueryStarted: async (_arg, api) => {
				try {
					await api.queryFulfilled
				} catch (error) {
					console.log(error)
					const fetchError = (error as IBaseFetchError).error
					toast.error(fetchError.data.message, { autoClose: false })
				}
			},
			providesTags: [{ type: 'DataItems', id: 'ALL' }],
		}),
		getById: builder.query<{ data: IUpdateDataItem }, string>({
			query: id => `${API.reagents}/${id}`,
			providesTags: (_res, _err, id) => [{ type: 'DataItems', id: id }],
		}),
		getUniqueField: builder.query<{ data: string[] }, string>({
			query: field => ({
				url: `${API.reagents}/unique/${field}`,
			}),
			providesTags: [{ type: 'DataItems', id: 'Unique' }],
			onQueryStarted: async (_arg, api) => {
				try {
					await api.queryFulfilled
				} catch (error) {
					const fetchError = (error as IBaseFetchError).error
					toast.error(fetchError.data.message, { autoClose: false })
				}
			},
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
})

export const {
	useGetAllQuery,
	useGetByIdQuery,
	useLazyGetByIdQuery,
	useGetUniqueFieldQuery,
	useLazyGetUniqueFieldQuery,
	useCreateMutation,
	useUpdateMutation,
	useDeleteMutation,
	usePrepareOrderMutation,
} = tableApiSlice
