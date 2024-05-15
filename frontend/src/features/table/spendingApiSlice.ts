import type { CreateSpending, ISpending } from './types/spending'
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
	}),
})

export const { useGetSpendingQuery, useCreateSpendingMutation } = spendingApiSlice
