import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
	server: {
		https: {
			cert: "./app.localhost.com.crt",
			key: "./app.localhost.com.key",
		},
		proxy: {
			"^/api/.*": {
				target: "https://localhost:3001",
				rewrite: path => path.replace(/^\/api/, ''),
				changeOrigin: false,
				secure: false,
			}
		},
	}
});
