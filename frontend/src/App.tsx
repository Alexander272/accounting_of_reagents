import { Provider } from 'react-redux'
import { ThemeProvider, CssBaseline } from '@mui/material'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import { LocalizationProvider } from '@mui/x-date-pickers'
import { ToastContainer } from 'react-toastify'
import dayjs from 'dayjs'
// import customParseFormat from 'dayjs/plugin/customParseFormat'
import 'dayjs/locale/ru'
import 'react-toastify/dist/ReactToastify.min.css'

import { store } from '@/app/store'
import { AppRouter } from '@/pages/router/AppRouter'
import { theme } from '@/theme/theme'

// dayjs.extend(customParseFormat)
dayjs.locale('ru') // глобальная локализация дат

function App() {
	return (
		<Provider store={store}>
			<ThemeProvider theme={theme}>
				<CssBaseline />
				<LocalizationProvider dateAdapter={AdapterDayjs} adapterLocale='ru'>
					<AppRouter />
				</LocalizationProvider>
			</ThemeProvider>
			<ToastContainer />
		</Provider>
	)
}

export default App
