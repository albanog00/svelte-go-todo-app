<script lang="ts">
	import '../index.css';

	import TaskItem from '$lib/components/TaskItem.svelte';
	import AddTask from '$lib/components/AddTask.svelte';

	import { tasks } from '$lib/store';
</script>

<div class="flex flex-col justify-between gap-4 sm:flex-row sm:gap-0">
	<AddTask />
	<div class="flex w-full flex-col items-center gap-2">
		<h1 class="text-4xl font-bold">Your tasks</h1>
		<div class="flex flex-col gap-1">
			{#if !$tasks.length}
				<span>No task scheduled yet.</span>
			{:else}
				{#each $tasks as task}
					<TaskItem {task} />
				{/each}
			{/if}
		</div>
	</div>
</div>
