import { user } from '$lib';
import type { User } from '$lib/store/user';
import type { LayoutServerLoad } from './$types';

export const ssr = true;

export const load: LayoutServerLoad = async ({ fetch, cookies }) => {
    const jwt = cookies.get("auth-jwt");

    async function fetchUserInfo(): Promise<User> {
        const userInfo = await fetch('http://localhost:3001/users', {
            cache: "no-cache",
            credentials: "include",
        })
            .then(async (response) => (await response.json()).data)
            .catch((error) => {
                console.error(error);
                return undefined;
            });
        return userInfo ? { ...userInfo, jwt } : undefined;
    }

    let userInfo: User | undefined;
    if (jwt) {
        userInfo = await fetchUserInfo()
        if (userInfo) {
            user.set(userInfo)
        } else {
            cookies.delete("auth-jwt")
        }
    }

    return {
        user: userInfo,
    }
};
