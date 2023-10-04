<script lang="ts">
	import SignIn from '$lib/components/SignIn.svelte';
	import '$lib/index.css';
	import type { LayoutData } from './$types';

	export let data: LayoutData;
</script>

<svelte:head>
	<title>Todo App</title>
</svelte:head>

{#if !data.user}
	<SignIn />
{:else}
	<div></div>
	<header class="h-24"></header>
	<div class="m-auto max-w-7xl text-center">
		<slot />
	</div>
{/if}
