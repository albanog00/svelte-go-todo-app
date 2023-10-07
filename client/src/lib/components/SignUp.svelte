<script lang="ts">
	import { user } from '$lib';

	let username: string;
	let password: string;
	let confirmPassword: string;

	const handleSignup = async () => {
		if (password !== confirmPassword) {
			alert("Password don't match");
			return;
		}
		const res = await user.signup({ username, password });
		if (res) {
			location.replace('/');
		} else {
			alert('username already taken');
		}
	};
</script>

<div class="flex h-screen w-screen items-center justify-center p-4">
	<div class="flex max-w-7xl flex-col items-center justify-center gap-2">
		<span class="text-3xl">SignUp to continue</span>
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
			<div class="grid-col grid">
				<label for="password">Confirm Password</label>
				<input
					type="password"
					bind:value={confirmPassword}
					class="rounded-lg border border-gray-400 px-4 py-2 focus:drop-shadow-lg"
				/>
			</div>
			<button
				type="submit"
				disabled={!username || !password || !confirmPassword}
				on:click={handleSignup}
				class="rounded-lg border border-gray-400 bg-gray-200 p-2 transition hover:focus:bg-gray-400 enabled:hover:bg-gray-300 enabled:hover:drop-shadow-lg enabled:hover:focus:drop-shadow-lg disabled:opacity-60"
			>
				SignUp
			</button>
		</div>
		<a class="text-sm" href="/sign-in">SignIn</a>
	</div>
</div>
