<script lang="ts">
	import { user } from '$lib';

	const handleLogout = async () => await user.signout();
</script>

<button on:click={handleLogout} class="hover:text-gray-400 hover:focus:text-black">
	SignOut
</button>
