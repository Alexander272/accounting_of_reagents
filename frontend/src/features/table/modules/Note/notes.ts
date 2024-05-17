export interface INote {
	id: string
	reagentId: string
	comment: string
	note: string
}

export type CreateNote = Omit<INote, 'id'>

export type NoteForm = Omit<INote, 'id' | 'reagentId'>
