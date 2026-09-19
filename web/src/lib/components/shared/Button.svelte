<script lang="ts">
	import { Button as BitsButton } from 'bits-ui';
	import clsx from 'clsx';
	import type { Snippet } from 'svelte';
	import type { HTMLButtonAttributes } from 'svelte/elements';

	interface ButtonProps {
		children: Snippet;
		secondary?: boolean;
	}

	type Props = ButtonProps & HTMLButtonAttributes;

	const { children, secondary = false, ...rest }: Props = $props();

	const styles = $derived(
		clsx(
			'rounded w-full p-1',
			secondary ? 'interactive-secondary' : 'interactive-primary',
			rest.class
		)
	);
</script>

<BitsButton.Root {...rest} class={styles}>
	{@render children()}
</BitsButton.Root>
