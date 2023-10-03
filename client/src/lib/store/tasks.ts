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
		const options = {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json; charset=utf-8'
			},
			body: JSON.stringify(task)
		};
		const response: Task = await fetch('http://localhost:3001/tasks', options)
			.then(async (data) => await data.json())
			.catch((error) => {
				console.error(error);
				return {};
			});
		return response;
	}

	async function deleteTask(id: string) {
		if (!id) return;
		const options = {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json; charset=utf-8'
			}
		};
		const response = await fetch(`http://localhost:3001/tasks/${id}`, options)
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
