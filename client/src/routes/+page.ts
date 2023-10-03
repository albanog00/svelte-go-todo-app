import { tasks } from '$lib';
import type { Task } from '$lib/store/tasks';
import type { PageLoad } from './$types';

export const prerender = false;

export const load: PageLoad = async ({ fetch }) => {
	async function fetchTasks(): Promise<Task[]> {
		const data: Task[] = await fetch('http://localhost:3001/tasks', {
			cache: 'no-cache'
		})
			.then(async (data) => await data.json())
			.catch((error) => {
				console.error(error);
				return [];
			});
		return data;
	}
	tasks.set(await fetchTasks());
};
