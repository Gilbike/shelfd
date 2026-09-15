<script lang="ts">
	import { Button as BitsButton } from 'bits-ui';

	const { children } = $props();
</script>

<BitsButton.Root class="interactive-primary w-full rounded p-1">
	{@render children()}
</BitsButton.Root>
