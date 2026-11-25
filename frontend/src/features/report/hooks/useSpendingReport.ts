import { useCallback, useEffect, useState } from 'react'
import dayjs from 'dayjs'

import type { IReportPeriod, ISpendingReportItem } from '../types'
import { useLazyGetSpendingReportQuery } from '../reportApiSlice'
import { downloadReportCsv } from '../utils/downloadCsv'
import { getRealm } from '@/features/realms/realmSlice'
import { useAppSelector } from '@/hooks/redux'

const getDefaultPeriod = (): IReportPeriod => ({
	from: dayjs().startOf('month'),
	to: dayjs().endOf('month'),
})
const EMPTY_ITEMS: ISpendingReportItem[] = []

export const useSpendingReport = () => {
	const realmId = useAppSelector(getRealm)?.id

	const [period, setPeriod] = useState<IReportPeriod>(getDefaultPeriod)
	const [getReport, { data, isFetching, isUninitialized }] = useLazyGetSpendingReportQuery()

	const items: ISpendingReportItem[] = data?.data ?? EMPTY_ITEMS

	const generate = useCallback(
		(next: IReportPeriod) => {
			if (!realmId) return

			setPeriod(next)
			void getReport({
				from: next.from.startOf('day').unix(),
				to: next.to.endOf('day').unix(),
				realm: realmId,
			})
		},
		[realmId, getReport],
	)

	// Отчёт строится сразу при открытии страницы (или смене realm) за текущий месяц
	useEffect(() => {
		if (!realmId) return
		generate(getDefaultPeriod())
	}, [realmId, generate])

	const download = useCallback(() => {
		if (!period || !items.length) return
		downloadReportCsv(items, period)
	}, [items, period])

	return {
		items,
		period, // Передаем это в ReportHeader для заголовка!
		isFetching,
		isUninitialized,
		canGenerate: Boolean(realmId) && !isFetching, // Блокируем кнопку, пока идет запрос
		canDownload: Boolean(items.length) && !isFetching,
		generate,
		download,
	}
}
