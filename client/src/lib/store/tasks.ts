import { writable } from "svelte/store";

export interface Task {
	id: string;
	description: string;
	date: Date;
}

type FetchTasksResponse = {
	tasks: Task[],
	count: number
}

export type TasksStore = {
	values: Task[],
	count: number
}

async function fetchTasks(page: number = 0) {
	const data: FetchTasksResponse = await fetch(`/api/tasks?page=${page}`, {
		cache: 'no-cache',
		credentials: 'same-origin',
	})
		.then(async (data) => await data.json())
		.catch((error) => {
			console.error(error);
			return [];
		});
	return { values: [...data.tasks.map((d) => ({ ...d, date: new Date(d.date) }))], count: data.count };
}

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
		method: 'DELETE',
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

export function createTask() {
	const { subscribe, set } = writable<TasksStore>();

	return {
		subscribe,
		set,
		fetch: async (page: number) => set(await fetchTasks(page)),
		add: async (task: Task, page: number = 0) => {
			const newTask = await addTask(task);
			if (newTask) {
				set(await fetchTasks(page))
			}
		},
		delete: async (id: string, page: number = 0) => {
			const deleted = await deleteTask(id);
			if (deleted) {
				set(await fetchTasks(page))
			}
		}
	}
}