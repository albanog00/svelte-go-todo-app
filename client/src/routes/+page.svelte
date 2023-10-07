<script lang="ts" context="module">
	export type TaskHandler = (page?: number) => Promise<void>;

	const enum TaskListState {
		Loading,
		Loaded,
		Error
	}
</script>

<script lang="ts">
	import TaskItem from '$lib/components/TaskItem.svelte';
	import AddTask from '$lib/components/AddTask.svelte';
	import { onMount } from 'svelte';
	import { tasks, user } from '$lib';

	let taskState = TaskListState.Loading;
	let selectedPage = 0;
	let numberOfPages = 0;
	let start = 0;

	async function changePage(newPage: number) {
		taskState = TaskListState.Loading;

		selectedPage = newPage >= 0 ? newPage : selectedPage;
		start = selectedPage * 5;
		await tasks.fetch(selectedPage);
		numberOfPages = Math.floor(($tasks.count - 1) / 5) + 1;

		taskState = TaskListState.Loaded;
	}

	async function handleDeleteTask(cb: TaskHandler) {
		// taskState = TaskListState.Loading;

		const tempNumberOfPages = Math.floor(($tasks.count - 1) / 5);
		if (selectedPage === tempNumberOfPages && (selectedPage * ($tasks.count - 1)) % 5 === 0)
			selectedPage -= 1;
		await cb(selectedPage);
		numberOfPages = Math.floor(($tasks.count - 1) / 5) + 1;
		start = selectedPage * 5;

		// taskState = TaskListState.Loaded;
	}

	async function addTaskCallback(cb: TaskHandler) {
		// taskState = TaskListState.Loading;

		await cb(selectedPage);
		numberOfPages = Math.floor(($tasks.count - 1) / 5) + 1;

		// taskState = TaskListState.Loaded;
	}

	onMount(async () => {
		await tasks.fetch(selectedPage);
		numberOfPages = Math.floor(($tasks.count - 1) / 5) + 1;
		start = selectedPage * 5;

		taskState = TaskListState.Loaded;
	});
</script>

<div class="m-auto max-w-7xl py-4 text-center">
	<div class="flex flex-col justify-between gap-4 sm:flex-row sm:gap-0">
		<AddTask {addTaskCallback} />
		<div class="flex w-full flex-col items-center gap-2">
			<h1 class="text-4xl font-bold">Your tasks</h1>
			<div class="flex flex-col gap-1">
				<div class="flex h-96 flex-col gap-1">
					{#if taskState == TaskListState.Loading}
						<span>Retrieving data...</span>
					{:else if !$tasks.count}
						<span>No task scheduled yet.</span>
					{:else}
						{#each $tasks.values as task, index}
							<TaskItem {task} index={start + index + 1} {handleDeleteTask} />
						{/each}
					{/if}
				</div>
				<div class="grid grid-flow-col gap-1 text-sm">
					<button
						disabled={selectedPage <= 0}
						on:click={async () => await changePage(selectedPage - 1)}
						class={`${
							!numberOfPages && 'hidden'
						} rounded-md border border-gray-400 px-3 py-1 transition hover:focus:bg-gray-400 enabled:hover:bg-gray-300 enabled:hover:drop-shadow-lg enabled:hover:focus:drop-shadow-lg disabled:opacity-60`}
						>Prev</button
					>
					{#each Array(numberOfPages) as _, pageNumber}
						<button
							disabled={selectedPage === pageNumber}
							on:click={async () => await changePage(pageNumber)}
							class={`${
								(pageNumber > selectedPage + 2 || pageNumber < selectedPage - 2) && 'hidden'
							} rounded-md border border-gray-400 px-3 py-1 transition hover:focus:bg-gray-400 enabled:hover:bg-gray-300 enabled:hover:drop-shadow-lg enabled:hover:focus:drop-shadow-lg disabled:bg-gray-300`}
							>{pageNumber + 1}</button
						>
					{/each}
					<button
						on:click={async () => await changePage(selectedPage + 1)}
						disabled={selectedPage >= numberOfPages - 1}
						class={`${
							!numberOfPages && 'hidden'
						} rounded-md border border-gray-400 px-3 py-1 transition hover:focus:bg-gray-400 enabled:hover:bg-gray-300 enabled:hover:drop-shadow-lg enabled:hover:focus:drop-shadow-lg disabled:opacity-60`}
						>Next</button
					>
				</div>
			</div>
		</div>
	</div>
</div>
