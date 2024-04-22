import { useCallback, useRef } from 'react'

export const useDebounce = <T>(callback: (...args: unknown[]) => T, delay: number) => {
	const timer = useRef<NodeJS.Timeout>()

	const debouncedCallback = useCallback(
		(...args: unknown[]) => {
			if (timer.current) {
				clearTimeout(timer.current)
			}
			timer.current = setTimeout(() => callback(...args), delay)
		},
		[callback, delay]
	)

	return debouncedCallback
}
