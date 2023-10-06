<script lang="ts">
	import TaskItem from '$lib/components/TaskItem.svelte';
	import AddTask from '$lib/components/AddTask.svelte';
	import { tasks } from '$lib';
	import type { Task } from '$lib/store/tasks';
	import { onMount } from 'svelte';

	const enum TaskListState {
		Loading,
		Loaded,
		Error
	}

	async function fetchTasks(): Promise<Task[]> {
		const data: Task[] = await fetch('/api/tasks', {
			cache: 'no-cache',
			credentials: 'same-origin'
		})
			.then(async (data) => await data.json())
			.catch((error) => {
				console.error(error);
				return [];
			});
		return [...data.map((d) => ({ ...d, date: new Date(d.date) }))];
	}

	let taskState = TaskListState.Loading;
	onMount(async () => {
		const tasksInfo = await fetchTasks();
		tasks.set(tasksInfo);
		taskState = TaskListState.Loaded;
	});
</script>

<div class="flex flex-col justify-between gap-4 sm:flex-row sm:gap-0">
	<AddTask />
	<div class="flex w-full flex-col items-center gap-2">
		<h1 class="text-4xl font-bold">Your tasks</h1>
		<div class="flex flex-col gap-1">
			{#if taskState == TaskListState.Loading}
				<span>Retrieving data...</span>
			{:else if !$tasks}
				<span>No task scheduled yet.</span>
			{:else}
				{#each $tasks as task}
					<TaskItem {task} />
				{/each}
			{/if}
		</div>
	</div>
</div>
