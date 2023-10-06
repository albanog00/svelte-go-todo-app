import type { User } from '$lib/store/user';
import type { LayoutServerLoad } from './$types';

export const ssr = true;

export const load: LayoutServerLoad = async ({ fetch, cookies }) => {
    const jwt = cookies.get("auth-jwt");

    async function fetchUserInfo(): Promise<User> {
        const userInfo: User = await fetch('/api/users', {
            cache: "no-cache",
            credentials: "same-origin",
        })
            .then(async (response) => (await response.json()).data)
            .catch((error) => {
                console.error(error);
                return undefined;
            });
        return userInfo;
    }

    let userInfo: User;
    if (jwt) {
        userInfo = await fetchUserInfo()
        if (!userInfo) {
            cookies.delete("auth-jwt", {
                path: "/"
            })
        }
    }

    return {
        user: userInfo
    }
};
