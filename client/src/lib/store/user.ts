import { writable } from "svelte/store";

interface SignedUser {
    username: string;
}

export type User = SignedUser | undefined;

export interface SignUser {
    username: string;
    password: string
}

export function createUser() {
    const { subscribe, set } = writable<User | undefined>();

    return {
        subscribe,
        set,
        signin: async (user: SignUser) => {
            const signedUser: User = await fetch("/api/auth/signin", {
                body: JSON.stringify(user),
                cache: "no-cache",
                method: "POST",
                credentials: "include"
            })
                .then(async data => (await data.json()).data)
                .catch(error => console.error(error))

            if (signedUser) {
                set(signedUser)
            }
        },
        signout: async () => {
            const data = await fetch('/api/auth/signout', {
                cache: "no-cache",
                credentials: 'include'
            })
                .then((data) => data)
                .catch((error) => console.log(error));
            if (data && data.status === 200) set(undefined);
        }
    }
}
