import { useCheckPermissions } from '../hooks/checkPerms'

type CanProps = {
	anyOf?: string[]
	allOf?: string[]
	children: React.ReactNode
}

export const Can = ({ anyOf, allOf, children }: CanProps) => {
	const allowed = useCheckPermissions({ anyOf, allOf })

	if (!allowed) {
		return null
	}

	return <>{children}</>
}
