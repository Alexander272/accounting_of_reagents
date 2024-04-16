import { createListenerMiddleware, TypedStartListening } from '@reduxjs/toolkit'

import { apiSlice } from '@/app/apiSlice'
import { AppDispatch, RootState } from '@/app/store'
import { resetUser } from '@/features/user/userSlice'
// import { resetDataTableState } from '@/features/dataTable/dataTableSlice'
// import { resetModalState } from '@/features/modal/modalSlice'
// import { resetEmployeeState } from '@/features/employees/employeeSlice'

export const resetStoreListener = createListenerMiddleware()

const startResetStoreListener = resetStoreListener.startListening as TypedStartListening<RootState, AppDispatch>

startResetStoreListener({
	actionCreator: resetUser,
	effect: async (_, listenerApi) => {
		await listenerApi.delay(100)
		//TODO
		// listenerApi.dispatch(resetDataTableState())
		// listenerApi.dispatch(resetModalState())
		// listenerApi.dispatch(resetEmployeeState())
		listenerApi.dispatch(apiSlice.util.resetApiState())
	},
})
