import type { TableCellAlign } from '@/components/Table/types'
import { Titles } from '@/features/table/constants/titles'

export const ReportHeaders = [Titles.Name, Titles.UName, Titles.Consignment, 'Количество', 'Расход', 'Остаток'] as const

export const ReportColumnWidths = [500, 500, 220, 220, 220, 220] as const

export const ReportColumnAligns: TableCellAlign[] = ['left', 'left', 'left', 'right', 'right', 'right']
