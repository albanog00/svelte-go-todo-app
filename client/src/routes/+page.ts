import { tasks } from "$lib/store";

export const prerender = true;

export async function load() {
    tasks.fetchData();
}