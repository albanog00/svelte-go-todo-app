import { writable } from "svelte/store";

export interface User {
    username: string;
    jwt: string;
}

export interface SignUser {
    username: string;
    password: string
}

export function createUser() {
    const { subscribe, set } = writable<User>();

    return {
        subscribe,
        set,
        signin: async (user: SignUser) => {
            const signedUser: User = await fetch("http://localhost:3001/auth/signin", {
                body: JSON.stringify(user),
                method: "POST",
                credentials: "include"
            })
                .then(async data => (await data.json()).data)
                .catch(error => console.error(error))

            if (signedUser) {
                set(signedUser)
            }
            window.location.reload()
        },
    }
}
