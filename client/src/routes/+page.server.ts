import { tasks } from "$lib";
import type { Task } from "$lib/store/tasks";
import type { PageServerLoad } from "./$types";

export const ssr = true;

export const load: PageServerLoad = async ({ fetch }) => {
    async function fetchTasks(): Promise<Task[]> {
        const data: Task[] = await fetch('http://localhost:3001/tasks', {
            cache: 'no-cache',
            credentials: "include"
        })
            .then(async (data) => await data.json())
            .catch((error) => {
                console.error(error);
                return [];
            });
        return data;
    }

    const tasksInfo = await fetchTasks()
    tasks.set(tasksInfo);

    return {
        tasks: tasksInfo
    }
}