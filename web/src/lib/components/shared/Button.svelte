<script lang="ts">
	import { Button as BitsButton } from 'bits-ui';
	import type { Snippet } from 'svelte';
	import type { HTMLButtonAttributes } from 'svelte/elements';

	interface ButtonProps {
		children: Snippet;
	}

	type Props = ButtonProps & HTMLButtonAttributes;

	const { children, ...rest }: Props = $props();
</script>

<BitsButton.Root {...rest} class="interactive-primary w-full rounded p-1">
	{@render children()}
</BitsButton.Root>
