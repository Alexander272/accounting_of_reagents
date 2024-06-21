import dayjs from 'dayjs'

import type { IColumn, IHeadColumn } from './types/table'
import { DateFormat } from '@/constants/date'
import { Titles } from './constants/titles'
import { useGetReagentTypesQuery } from './modules/Types/typesApiSlice'

export const HeaderColumns: IHeadColumn[] = [
	{
		key: 'type',
		label: Titles.Type,
		allowsSorting: true,
	},
	{
		key: 'name',
		label: Titles.Name,
		allowsSorting: true,
	},
	{
		key: 'uname',
		label: Titles.UName,
		allowsSorting: true,
	},
	{
		key: 'document',
		label: Titles.Doc,
		allowsSorting: true,
	},
	{
		key: 'purity',
		label: Titles.Purity,
		allowsSorting: true,
	},
	{
		key: 'dateOfManufacture',
		label: Titles.DateOfManufacture,
		allowsSorting: true,
	},
	{
		key: 'consignment',
		label: Titles.Consignment,
		allowsSorting: true,
	},
	{
		key: 'manufacturer',
		label: Titles.Manufacturer,
		allowsSorting: true,
	},
	{
		key: 'shelfLife',
		label: Titles.ShelfLife,
		allowsSorting: true,
	},
	{
		key: 'place',
		label: Titles.Place.Main,
		children: [
			{
				key: 'place_closet',
				label: Titles.Place.Closet,
				allowsSorting: true,
			},
			{
				key: 'place_shelf',
				label: Titles.Place.Shelf,
				allowsSorting: true,
			},
		],
	},
	{
		key: 'incomingControl',
		label: Titles.IncomingControl.Main,
		children: [
			{
				key: 'incomingControl_receiptDate',
				label: Titles.IncomingControl.ReceiptDate,
				allowsSorting: true,
			},
			{
				key: 'incomingControl_amount',
				label: Titles.IncomingControl.Amount,
				allowsSorting: true,
			},
			{
				key: 'incomingControl_date',
				label: Titles.IncomingControl.Date,
				allowsSorting: true,
			},
			{
				key: 'incomingControl_protocol',
				label: Titles.IncomingControl.Protocol,
				allowsSorting: true,
			},
			{
				key: 'incomingControl_result',
				label: Titles.IncomingControl.Result,
				allowsSorting: true,
			},
		],
	},
	{ key: 'spending', label: Titles.Spending },
	{
		key: 'extending',
		label: Titles.Extending.Main,
		children: [
			{
				key: 'extending_date',
				label: Titles.Extending.Date,
			},
			{
				key: 'extending_period',
				label: Titles.Extending.Period,
			},
		],
	},
	{
		key: 'seizureInformation',
		label: Titles.SeizureInformation,
	},
	{
		key: 'disposalInformation',
		label: Titles.DisposalInformation,
	},
	{
		key: 'comments',
		label: Titles.Comments,
	},
	{
		key: 'notes',
		label: Titles.Notes,
	},
]

export const Columns: IColumn[] = [
	{
		key: 'type',
		label: Titles.Type,
		filter: {
			type: 'list',
			getOptions: useGetReagentTypesQuery,
		},
	},
	{
		key: 'name',
		label: Titles.Name,
		allowSearch: true,
		filter: 'string',
	},
	{
		key: 'uname',
		label: Titles.UName,
		allowSearch: true,
		filter: 'string',
	},
	{
		key: 'document',
		label: Titles.Doc,
		allowSearch: true,
		filter: 'string',
	},
	{
		key: 'purity',
		label: Titles.Purity,
		allowSearch: true,
		filter: 'string',
	},
	{
		key: 'dateOfManufacture',
		label: Titles.DateOfManufacture,
		filter: 'date',
		formatter: value => {
			if (value == 0) return '-'
			else return dayjs((value as number) * 1000).format(DateFormat)
		},
	},
	{
		key: 'consignment',
		label: Titles.Consignment,
		filter: 'string',
	},
	{
		key: 'manufacturer',
		label: Titles.Manufacturer,
		filter: 'string',
	},
	{
		key: 'shelfLife',
		label: Titles.ShelfLife,
		filter: 'number',
	},
	{
		key: 'place_closet',
		label: Titles.Place.Closet,
		filter: 'string',
	},
	{
		key: 'place_shelf',
		label: Titles.Place.Shelf,
		formatter: value => {
			if (value == 0) return '-'
			else return (value as number).toString()
		},
	},
	{
		key: 'incomingControl_receiptDate',
		label: Titles.IncomingControl.ReceiptDate,
		filter: 'date',
		formatter: value => {
			if (value == 0) return '-'
			else return dayjs((value as number) * 1000).format(DateFormat)
		},
	},
	{
		key: 'incomingControl_amount',
		label: Titles.IncomingControl.Amount,
	},
	{
		key: 'incomingControl_date',
		label: Titles.IncomingControl.Date,
		filter: 'date',
		formatter: value => {
			if (value == 0) return '-'
			else return dayjs((value as number) * 1000).format(DateFormat)
		},
	},
	{
		key: 'incomingControl_protocol',
		label: Titles.IncomingControl.Protocol,
		filter: 'string',
	},
	{
		key: 'incomingControl_result',
		label: Titles.IncomingControl.Result,
		formatter: value => {
			if (value) return 'Соответствует'
			else return 'Несоответствует'
		},
		filter: {
			type: 'switch',
			options: {
				true: 'Соответствует',
				false: 'Несоответствует',
			},
		},
	},
	{
		key: 'spending',
		label: Titles.Spending,
		formatter: value => {
			if (value == 0) return '-'
			else return (value as number).toString()
		},
	},
	{
		key: 'extending_date',
		label: Titles.Extending.Date,
		formatter: value => {
			if (value == 0) return '-'
			else return dayjs((value as number) * 1000).format(DateFormat)
		},
	},
	{
		key: 'extending_period',
		label: Titles.Extending.Period,
		formatter: value => {
			if (value == 0) return '-'
			else return (value as number).toString()
		},
	},
	{
		key: 'seizureInformation',
		label: Titles.SeizureInformation,
	},
	{
		key: 'disposalInformation',
		label: Titles.DisposalInformation,
		filter: 'string',
	},
	{
		key: 'comments',
		label: Titles.Comments,
		allowSearch: true,
		filter: 'string',
	},
	{
		key: 'notes',
		label: Titles.Notes,
		allowSearch: true,
		filter: 'string',
	},
]
