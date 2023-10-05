<script lang="ts">
	import { user } from '$lib';

	const handleLogout = async () => await user.signout();
</script>

<button
	on:click={handleLogout}
	class="rounded-lg border border-gray-400 bg-gray-200 p-2 transition hover:focus:bg-gray-400 enabled:hover:bg-gray-300 enabled:hover:drop-shadow-lg enabled:hover:focus:drop-shadow-lg disabled:opacity-60"
>
	SignOut
</button>
