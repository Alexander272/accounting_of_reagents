import { Titles } from './constants/titles'
import type { IColumn, IHeadColumn } from './types/table'

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
				key: 'incomingControl_receiptData',
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
		key: 'place_closet',
		label: Titles.Place.Closet,
	},
	{
		key: 'place_shelf',
		label: Titles.Place.Shelf,
	},
	{
		key: 'incomingControl_receiptData',
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
	{ key: 'spending', label: Titles.Spending },
	{
		key: 'extending_date',
		label: Titles.Extending.Date,
	},
	{
		key: 'extending_period',
		label: Titles.Extending.Period,
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
