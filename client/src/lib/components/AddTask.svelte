<script lang="ts">
	import { tasks } from '$lib';
	import type { TaskHandler } from '../../routes/+page.svelte';

	export let addTaskCallback: (cb: TaskHandler) => Promise<void>;

	let description: string;
	let date: string;
	let time: string;

	async function onAdd() {
		if (!description || !date || !time) return;
		const dateTime = new Date(`${date} ${time}:00`);
		let newTask = {
			id: '',
			description,
			date: dateTime
		};

		addTaskCallback(async (page?: number) => await tasks.add(newTask, page));

		description = '';
		date = '';
		time = '';
	}
</script>

<div class="flex w-full flex-col items-center gap-2">
	<div class="flex items-center justify-center">
		<h1 class="text-4xl font-bold">Todo list app</h1>
	</div>
	<div class="flex w-full flex-col items-center justify-center gap-1">
		<input
			type="text"
			bind:value={description}
			placeholder="Write you task"
			class="rounded-lg border border-gray-400 px-4 py-2 focus:drop-shadow-lg"
		/>
		<div class="flex w-full flex-row items-center justify-center gap-1">
			<input
				type="date"
				bind:value={date}
				class="rounded-lg border border-gray-400 px-4 py-2 focus:drop-shadow-lg"
			/>
			<input
				type="time"
				bind:value={time}
				class="rounded-lg border border-gray-400 px-4 py-2 focus:drop-shadow-lg"
			/>
		</div>
		<button
			type="submit"
			disabled={!description || !date || !time}
			on:click={onAdd}
			class="rounded-lg border border-gray-400 bg-gray-200 p-2 transition hover:focus:bg-gray-400 enabled:hover:bg-gray-300 enabled:hover:drop-shadow-lg enabled:hover:focus:drop-shadow-lg disabled:opacity-60"
		>
			Add task
		</button>
	</div>
</div>
