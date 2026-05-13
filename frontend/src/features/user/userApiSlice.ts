import { toast } from 'react-toastify'

import type { IBaseFetchError } from '@/app/types/error'
import type { IUserData, IUserDataDTO } from './types/user'
import { API } from '@/app/api'
import { apiSlice } from '@/app/apiSlice'

const userApiSlice = apiSlice.injectEndpoints({
	overrideExisting: false,
	endpoints: builder => ({
		getAllUsers: builder.query<{ data: IUserData[] }, void>({
			query: () => ({
				url: API.users.base,
				method: 'GET',
			}),
			onQueryStarted: async (_arg, api) => {
				try {
					await api.queryFulfilled
				} catch (error) {
					const fetchError = (error as IBaseFetchError).error
					toast.error(fetchError.data.message, { autoClose: false })
				}
			},
			providesTags: [{ type: 'Users', id: 'ALL' }],
		}),

		syncUsers: builder.mutation<null, void>({
			query: () => ({
				url: API.users.sync,
				method: 'POST',
			}),
			invalidatesTags: ['Users'],
		}),

		updateUser: builder.mutation<null, IUserDataDTO>({
			query: data => ({
				url: `${API.users.base}/${data.id}`,
				method: 'PUT',
				body: data,
			}),
			invalidatesTags: ['Users'],
		}),
	}),
})

export const { useGetAllUsersQuery, useSyncUsersMutation, useUpdateUserMutation } = userApiSlice
