import { toast } from 'react-toastify'

import type { IBaseFetchError } from '@/app/types/error'
import type { IReportParams, ISpendingReportItem } from './types'
import { apiSlice } from '@/app/apiSlice'
import { API } from '@/app/api'

const reportApiSlice = apiSlice.injectEndpoints({
	overrideExisting: false,
	endpoints: builder => ({
		getSpendingReport: builder.query<{ data: ISpendingReportItem[] }, IReportParams>({
			query: params => ({
				url: API.spendingReport,
				params: { from: params.from, to: params.to, realm: params.realm },
			}),
			providesTags: [{ type: 'Spending', id: 'ALL' }],
			onQueryStarted: async (_arg, api) => {
				try {
					await api.queryFulfilled
				} catch (error) {
					const fetchError = (error as IBaseFetchError).error
					toast.error(fetchError.data.message, { autoClose: false })
				}
			},
		}),
	}),
})

export const { useLazyGetSpendingReportQuery } = reportApiSlice
