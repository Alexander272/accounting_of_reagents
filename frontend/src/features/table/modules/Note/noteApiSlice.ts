import type { CreateNote, INote } from './notes'
import { apiSlice } from '@/app/apiSlice'
import { API } from '@/app/api'

const noteApiSlice = apiSlice.injectEndpoints({
	overrideExisting: false,
	endpoints: builder => ({
		getNotes: builder.query<{ data: INote[] }, string>({
			query: reagentId => `${API.notes}/${reagentId}`,
			providesTags: (_res, _err, id) => [{ type: 'Notes', id: id }],
		}),

		createNotes: builder.mutation<null, CreateNote>({
			query: data => ({
				url: API.notes,
				method: 'POST',
				body: data,
			}),
			invalidatesTags: (_res, _err, data) => [
				{ type: 'Notes', id: data.reagentId },
				{ type: 'DataItems', id: 'ALL' },
			],
		}),

		updateNotes: builder.mutation<null, INote>({
			query: data => ({
				url: `${API.notes}/${data.id}`,
				method: 'PUT',
				body: data,
			}),
			invalidatesTags: (_res, _err, data) => [
				{ type: 'Notes', id: data.reagentId },
				{ type: 'DataItems', id: 'ALL' },
			],
		}),

		deleteNotes: builder.mutation<null, INote>({
			query: data => ({
				url: `${API.notes}/${data.id}`,
				method: 'DELETE',
			}),
			invalidatesTags: (_res, _err, data) => [
				{ type: 'Notes', id: data.reagentId },
				{ type: 'DataItems', id: 'ALL' },
			],
		}),
	}),
})

export const { useGetNotesQuery, useCreateNotesMutation, useUpdateNotesMutation, useDeleteNotesMutation } = noteApiSlice
