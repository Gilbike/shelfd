<script lang="ts">
	import Navbar from '$lib/components/shared/Navbar.svelte';

	const { children } = $props();
</script>

<div class="flex h-dvh w-full flex-col-reverse overflow-hidden md:flex-row">
	<Navbar />
	<main class="min-h-0 flex-1 overflow-y-auto">
		{@render children()}
	</main>
</div>
