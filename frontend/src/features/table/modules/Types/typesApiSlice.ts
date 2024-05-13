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
		editAmountType: builder.mutation<null, { data: IAmountType[]; deleted: string[] }>({
			query: data => ({
				url: `${API.amountTypes}/edit`,
				method: 'POST',
				body: data,
			}),
			invalidatesTags: [{ type: 'AmountTypes', id: 'ALL' }],
		}),
		// createAmountType: builder.mutation<null, CreateAmountType>({
		// 	query: data => ({
		// 		url: API.amountTypes,
		// 		method: 'POST',
		// 		body: data,
		// 	}),
		// 	invalidatesTags: [{ type: 'AmountTypes', id: 'ALL' }],
		// }),
		// updateAmountType: builder.mutation<null, IAmountType>({
		// 	query: data => ({
		// 		url: `${API.amountTypes}/${data.id}`,
		// 		method: 'PUT',
		// 		body: data,
		// 	}),
		// 	invalidatesTags: [{ type: 'AmountTypes', id: 'ALL' }],
		// }),
		// updateSeveralAmountType: builder.mutation<null, IAmountType[]>({
		// 	query: data => ({
		// 		url: API.amountTypes,
		// 		method: 'PUT',
		// 		body: data,
		// 	}),
		// 	invalidatesTags: [{ type: 'AmountTypes', id: 'ALL' }],
		// }),
		deleteAmountType: builder.mutation<null, string>({
			query: id => ({
				url: `${API.amountTypes}/${id}`,
				method: 'DELETE',
			}),
			invalidatesTags: [{ type: 'AmountTypes', id: 'ALL' }],
		}),
	}),
})

export const {
	useGetReagentTypesQuery,

	useGetAmountTypesQuery,
	useEditAmountTypeMutation,
	// useCreateAmountTypeMutation,
	// useUpdateAmountTypeMutation,
	// useUpdateSeveralAmountTypeMutation,
	useDeleteAmountTypeMutation,
} = typesApiSlice
