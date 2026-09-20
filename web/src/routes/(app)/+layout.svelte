<script lang="ts">
	import Navbar from '$lib/components/shared/Navbar.svelte';

	const { children } = $props();
</script>

<Navbar />
<main class="flex-1">
	{@render children()}
</main>
