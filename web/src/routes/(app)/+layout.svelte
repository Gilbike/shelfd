<script lang="ts">
	import Navbar from '$lib/components/shared/Navbar.svelte';

	const { children } = $props();
</script>

<Navbar />
<main class="flex-1">
	{@render children()}
</main>

<style lang="postcss">
	@reference '../layout.css';

	:global(body) {
		@apply flex w-full flex-row;
	}
</style>
