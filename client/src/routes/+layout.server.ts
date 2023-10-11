import type { User } from "$lib/store/user";
import { redirect } from "@sveltejs/kit";
import type { LayoutServerLoad } from "./$types";

export const ssr = true;

export const load: LayoutServerLoad = async ({ fetch, cookies, route }) => {
    async function fetchUserInfo(): Promise<User> {
        const userInfo: User = await fetch('/api/users', {
            cache: "no-store",
            credentials: "same-origin",
        })
            .then(async (response) => (await response.json()).data)
            .catch((error) => {
                console.error(error);
                return {};
            });
        return userInfo;
    }

    const user = cookies.get("auth-jwt") ? await fetchUserInfo() : undefined;
    const notAuthRoute = (route.id !== "/sign-in" && route.id !== "/sign-up");
    if (!user && notAuthRoute) {
        cookies.delete("auth-jwt", {
            path: "/"
        })
        throw redirect(307, `/sign-in`);
    }

    if (user && !notAuthRoute) {
        throw redirect(301, '/');
    }

    return {
        user
    }
}
