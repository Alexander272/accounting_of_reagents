import type { DialogProps as MuiDialogProps } from '@mui/material/Dialog'
import MuiDialog from '@mui/material/Dialog'
import MuiDialogActions from '@mui/material/DialogActions'
import MuiDialogContent from '@mui/material/DialogContent'
import MuiDialogContentText from '@mui/material/DialogContentText'
import MuiDialogTitle from '@mui/material/DialogTitle'

interface IDialogProps extends MuiDialogProps {
	title: string
	body: string | JSX.Element
	actions?: JSX.Element
}

export const Dialog = (props: IDialogProps) => {
	const { actions, title, body, ...other } = props

	return (
		<MuiDialog {...other}>
			<MuiDialogTitle>{title}</MuiDialogTitle>
			<MuiDialogContent>
				{typeof body === 'string' ? <MuiDialogContentText>{body}</MuiDialogContentText> : body}
			</MuiDialogContent>
			<MuiDialogActions>{actions}</MuiDialogActions>
		</MuiDialog>
	)
}
