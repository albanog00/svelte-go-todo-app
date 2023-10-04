<script lang="ts">
	import TaskItem from '$lib/components/TaskItem.svelte';
	import AddTask from '$lib/components/AddTask.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<div class="flex flex-col justify-between gap-4 sm:flex-row sm:gap-0">
	<AddTask />
	<div class="flex w-full flex-col items-center gap-2">
		<h1 class="text-4xl font-bold">Your tasks</h1>
		<div class="flex flex-col gap-1">
			{#if !data.tasks.length}
				<span>No task scheduled yet.</span>
			{:else}
				{#each data.tasks as task}
					<TaskItem {task} />
				{/each}
			{/if}
		</div>
	</div>
</div>
