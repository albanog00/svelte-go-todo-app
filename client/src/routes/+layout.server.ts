import type { User } from "$lib/store/user";
import { redirect } from "@sveltejs/kit";
import type { LayoutServerLoad } from "./$types";

export const ssr = true;

export const load: LayoutServerLoad = async ({ fetch, cookies, route }) => {
    async function fetchUserInfo(): Promise<User> {
        const userInfo: User = await fetch('/api/users', {
            cache: "no-cache",
            credentials: "same-origin",
            mode: "same-origin"
        })
            .then(async (response) => (await response.json()).data)
            .catch((error) => {
                console.error(error);
                return {};
            });
        return userInfo;
    }

    const user = cookies.get("auth-jwt") ? await fetchUserInfo() : undefined;
    if (!user && (route.id !== "/sign-in" && route.id !== "/sign-up")) {
        cookies.delete("auth-jwt", {
            path: "/"
        })
        throw redirect(307, `/sign-in`);
    }

    return {
        user
    }
}
