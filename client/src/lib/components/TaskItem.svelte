<script lang="ts">
	import { tasks } from '$lib/store';
	import type { Task } from '$lib/store/tasks';
	import { CalendarDays, Clock, Edit, X } from 'lucide-svelte';

	export let task: Task;

	async function deleteTask() {
		await tasks.remove(task.id);
	}
</script>

<div class="flex flex-row items-center justify-center gap-1">
	<span
		class="w-full max-w-full rounded-lg border border-gray-400 p-1 transition hover:bg-gray-200 hover:drop-shadow-lg"
	>
		<div class="flex flex-col items-center">
			<span>{task.description}</span>
			<small class="flex flex-row justify-center gap-1">
				<CalendarDays size={16} />
				{task.date}
			</small>
			<small class="flex flex-row justify-center gap-1">
				<Clock size={16} />
				{task.time}
			</small>
		</div>
	</span>
	<div class="flex flex-row gap-1">
		<button on:click={deleteTask} class="text-red-500" type="submit">
			<X />
		</button>
		<button type="submit">
			<Edit />
		</button>
	</div>
</div>
