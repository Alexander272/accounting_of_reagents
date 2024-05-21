import dayjs from 'dayjs'

import type { IColumn, IHeadColumn } from './types/table'
import { DateFormat } from '@/constants/date'
import { Titles } from './constants/titles'

export const HeaderColumns: IHeadColumn[] = [
	{
		key: 'type',
		label: Titles.Type,
	},
	{
		key: 'name',
		label: Titles.Name,
	},
	{
		key: 'uname',
		label: Titles.UName,
	},
	{
		key: 'document',
		label: Titles.Doc,
	},
	{
		key: 'purity',
		label: Titles.Purity,
	},
	{
		key: 'dateOfManufacture',
		label: Titles.DateOfManufacture,
	},
	{
		key: 'consignment',
		label: Titles.Consignment,
	},
	{
		key: 'manufacturer',
		label: Titles.Manufacturer,
	},
	{
		key: 'shelfLife',
		label: Titles.ShelfLife,
	},
	{
		key: 'place',
		label: Titles.Place.Main,
		children: [
			{
				key: 'place_closet',
				label: Titles.Place.Closet,
			},
			{
				key: 'place_shelf',
				label: Titles.Place.Shelf,
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
			},
			{
				key: 'incomingControl_amount',
				label: Titles.IncomingControl.Amount,
			},
			{
				key: 'incomingControl_date',
				label: Titles.IncomingControl.Date,
			},
			{
				key: 'incomingControl_protocol',
				label: Titles.IncomingControl.Protocol,
			},
			{
				key: 'incomingControl_result',
				label: Titles.IncomingControl.Result,
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
