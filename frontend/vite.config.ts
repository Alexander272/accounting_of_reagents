import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [
		react({
			jsxImportSource: '@emotion/react',
			plugins: [['@swc/plugin-emotion', {}]],
		}),
	],
	resolve: {
		alias: [
			{
				find: '@',
				replacement: path.resolve(__dirname, 'src'),
			},
		],
	},
	server: {
		proxy: {
			'/api': 'http://localhost:9000',
		},
	},
	build: {
		target: 'es2015',
	},
})
