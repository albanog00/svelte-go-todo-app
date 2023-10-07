import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const ssr = true;
export const prerender = true;

export const load: PageServerLoad = ({ cookies }) => {
    if (cookies.get("auth-jwt"))
        throw redirect(301, "/")
}