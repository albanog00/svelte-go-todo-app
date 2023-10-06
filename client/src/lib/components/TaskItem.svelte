<script lang="ts">
	import { tasks } from '$lib';
	import type { Task } from '$lib/store/tasks';
	import { CalendarDays, Clock, Edit, X } from 'lucide-svelte';

	export let task: Task;
	export let index: number;
	export let handleDeleteTask: (cb: (page?: number) => Promise<void>) => Promise<void>;

	const handleDelete = async () =>
		await handleDeleteTask(async function (page?: number) {
			await tasks.delete(task.id, page);
		});
</script>

<div class="flex flex-row items-center justify-center gap-1">
	<span class="px-2">
		{index}
	</span>
	<span
		class="w-full max-w-full rounded-lg border border-gray-400 p-1 transition hover:bg-gray-200 hover:drop-shadow-lg"
	>
		<div class="flex flex-col items-center">
			<span>{task.description}</span>
			<small class="flex flex-row justify-center gap-1">
				<CalendarDays size={16} />
				{task.date.toLocaleDateString('it-IT')}
			</small>
			<small class="flex flex-row justify-center gap-1">
				<Clock size={16} />
				{task.date.toLocaleTimeString('it-IT').slice(0, -3)}
			</small>
		</div>
	</span>
	<div class="flex flex-row gap-1">
		<button on:click={handleDelete} class="text-red-500" type="submit">
			<X />
		</button>
		<button type="submit">
			<Edit />
		</button>
	</div>
</div>
