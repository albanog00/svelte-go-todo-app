import { writable } from 'svelte/store';

export interface Task {
	id: string;
	description: string;
	date: string;
	time: string;
}

export function createTasks() {
	const { subscribe, set, update } = writable<Task[]>([]);

	async function addTask(task: Task) {
		if (!task) return;
		const response: Task = await fetch('/api/tasks', {
			method: 'POST',
			cache: "no-cache",
			headers: {
				'Content-Type': 'application/json; charset=utf-8'
			},
			body: JSON.stringify(task),
			credentials: "same-origin"
		})
			.then(async (data) => await data.json())
			.catch((error) => {
				console.error(error);
				return {};
			});
		return response;
	}

	async function deleteTask(id: string) {
		if (!id) return;
		const response = await fetch(`/api/tasks/${id}`, {
			method: 'PUT',
			cache: "no-cache",
			headers: {
				'Content-Type': 'application/json; charset=utf-8'
			},
			credentials: "same-origin"
		})
			.then((data) => data)
			.catch((error) => {
				console.error(error);
				return;
			});
		return response && response.ok;
	}

	return {
		subscribe,
		set,
		create: async (task: Task) => {
			const newTask = await addTask(task);
			if (newTask) {
				update((t) => [...t, newTask]);
			}
		},
		delete: async (id: string) => {
			if (await deleteTask(id)) {
				update((t) => t.filter((x) => x.id !== id));
			}
		}
	};
}
