import type { CreateSpending, ISpending } from './spending'
import { apiSlice } from '@/app/apiSlice'
import { API } from '@/app/api'

const spendingApiSlice = apiSlice.injectEndpoints({
	overrideExisting: false,
	endpoints: builder => ({
		getSpending: builder.query<{ data: ISpending[] }, string>({
			query: reagentId => `${API.spending}/${reagentId}`,
			providesTags: (_res, _err, id) => [{ type: 'Spending', id: id }],
		}),

		createSpending: builder.mutation<null, CreateSpending>({
			query: data => ({
				url: API.spending,
				method: 'POST',
				body: data,
			}),
			invalidatesTags: (_res, _err, data) => [
				{ type: 'Spending', id: data.reagentId },
				{ type: 'DataItems', id: 'ALL' },
			],
		}),

		updateSpending: builder.mutation<null, ISpending>({
			query: data => ({
				url: `${API.spending}/${data.id}`,
				method: 'PUT',
				body: data,
			}),
			invalidatesTags: (_res, _err, data) => [
				{ type: 'Spending', id: data.reagentId },
				{ type: 'DataItems', id: 'ALL' },
			],
		}),

		deleteSpending: builder.mutation<null, ISpending>({
			query: data => ({
				url: `${API.spending}/${data.id}`,
				params: new URLSearchParams({ reagentId: data.reagentId }),
				method: 'DELETE',
			}),
			invalidatesTags: (_res, _err, data) => [
				{ type: 'Spending', id: data.reagentId },
				{ type: 'DataItems', id: 'ALL' },
			],
		}),
	}),
})

export const { useGetSpendingQuery, useCreateSpendingMutation, useUpdateSpendingMutation, useDeleteSpendingMutation } =
	spendingApiSlice
