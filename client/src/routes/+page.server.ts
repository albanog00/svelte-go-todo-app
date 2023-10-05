import type { Task } from "$lib/store/tasks";
import type { PageServerLoad } from "./$types";

export const ssr = true;

export const load: PageServerLoad = async ({ fetch }) => {
    async function fetchTasks(): Promise<Task[]> {
        const data: Task[] = await fetch('/api/tasks', {
            cache: 'no-cache',
            credentials: "same-origin"
        })
            .then(async (data) => await data.json())
            .catch((error) => {
                console.error(error);
                return [];
            });
        return data;
    }

    const tasksInfo = await fetchTasks()

    return {
        tasks: tasksInfo
    }
}