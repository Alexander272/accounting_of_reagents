import type { IAmountType, IReagentType } from './types'
import { API } from '@/app/api'
import { apiSlice } from '@/app/apiSlice'

export const typesApiSlice = apiSlice.injectEndpoints({
	overrideExisting: false,
	endpoints: builder => ({
		getReagentTypes: builder.query<{ data: IReagentType[] }, null>({
			query: () => API.reagentTypes,
			providesTags: [{ type: 'ReagentTypes', id: 'ALL' }],
		}),

		getAmountTypes: builder.query<{ data: IAmountType[] }, null>({
			query: () => API.amountTypes,
			providesTags: [{ type: 'AmountTypes', id: 'ALL' }],
		}),
	}),
})

export const { useGetReagentTypesQuery, useGetAmountTypesQuery } = typesApiSlice
