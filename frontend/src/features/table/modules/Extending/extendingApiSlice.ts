import type { CreateExtending, IExtending } from './extending'
import { apiSlice } from '@/app/apiSlice'
import { API } from '@/app/api'

const extendingApiSlice = apiSlice.injectEndpoints({
	overrideExisting: false,
	endpoints: builder => ({
		getExtending: builder.query<{ data: IExtending[] }, string>({
			query: reagentId => `${API.extending}/${reagentId}`,
			providesTags: (_res, _err, id) => [{ type: 'Extending', id: id }],
		}),

		createExtending: builder.mutation<null, CreateExtending>({
			query: data => ({
				url: API.extending,
				method: 'POST',
				body: data,
			}),
			invalidatesTags: (_res, _err, data) => [
				{ type: 'Extending', id: data.reagentId },
				{ type: 'DataItems', id: 'ALL' },
			],
		}),

		updateExtending: builder.mutation<null, IExtending>({
			query: data => ({
				url: `${API.extending}/${data.id}`,
				method: 'PUT',
				body: data,
			}),
			invalidatesTags: (_res, _err, data) => [
				{ type: 'Extending', id: data.reagentId },
				{ type: 'DataItems', id: 'ALL' },
			],
		}),

		deleteExtending: builder.mutation<null, IExtending>({
			query: data => ({
				url: `${API.extending}/${data.id}`,
				method: 'DELETE',
			}),
			invalidatesTags: (_res, _err, data) => [
				{ type: 'Extending', id: data.reagentId },
				{ type: 'DataItems', id: 'ALL' },
			],
		}),
	}),
})

export const {
	useGetExtendingQuery,
	useCreateExtendingMutation,
	useUpdateExtendingMutation,
	useDeleteExtendingMutation,
} = extendingApiSlice
