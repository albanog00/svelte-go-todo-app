<script lang="ts">
	import '$lib/index.css';

	import SignIn from '$lib/components/SignIn.svelte';
	import SignOut from '$lib/components/SignOut.svelte';
	import { UserCircle2 } from 'lucide-svelte';
	import { user } from '$lib';
	import type { LayoutServerData } from './$types';

	export let data: LayoutServerData;
	user.set(data.user);
</script>

<svelte:head>
	<title>Todo App</title>
</svelte:head>

{#if !$user}
	<SignIn />
{:else}
	<header class="h-24">
		<div class="flex justify-end p-4">
			<div class="flex flex-col items-center justify-center gap-1">
				<div class="flex gap-1">
					<UserCircle2 />
					{$user.username}
				</div>
				<SignOut />
			</div>
		</div>
	</header>
	<div class="m-auto max-w-7xl text-center">
		<slot />
	</div>
{/if}
