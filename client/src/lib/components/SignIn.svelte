<script lang="ts">
	import { user } from '$lib';

	let username: string;
	let password: string;

	async function handleLogin() {
		await user.signin({ username, password });
	}
</script>

<div class="flex h-screen w-screen items-center justify-center p-4">
	<div class="flex max-w-7xl flex-col items-center justify-center gap-2">
		<span class="text-3xl">Login to continue</span>
		<div
			class="grid-row grid w-full max-w-5xl grid-cols-1 gap-2 rounded-xl border-2 border-gray-200 p-4 shadow-lg"
		>
			<div class="grid-col grid">
				<label for="username">Username</label>
				<input
					type="text"
					bind:value={username}
					class="rounded-lg border border-gray-400 px-4 py-2 focus:drop-shadow-lg"
				/>
			</div>
			<div class="grid-col grid">
				<label for="password">Password</label>
				<input
					type="password"
					bind:value={password}
					class="rounded-lg border border-gray-400 px-4 py-2 focus:drop-shadow-lg"
				/>
			</div>
			<button
				type="submit"
				disabled={!username || !password}
				on:click={handleLogin}
				class="rounded-lg border border-gray-400 bg-gray-200 p-2 transition hover:focus:bg-gray-400 enabled:hover:bg-gray-300 enabled:hover:drop-shadow-lg enabled:hover:focus:drop-shadow-lg disabled:opacity-60"
			>
				SignIn
			</button>
		</div>
	</div>
</div>
