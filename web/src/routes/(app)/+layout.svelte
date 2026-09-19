<script lang="ts">
	import Button from '$lib/components/shared/Button.svelte';
	import { User } from '@lucide/svelte';

	const { children } = $props();
</script>

<nav class="flex flex-col justify-between border-r border-r-border bg-surface p-4">
	<div></div>
	<Button secondary class="p-2">
		<User />
	</Button>
</nav>
<main class="flex-1">
	{@render children()}
</main>

<style lang="postcss">
	@reference '../layout.css';

	:global(body) {
		@apply flex w-full flex-row;
	}
</style>
