<script lang="ts">
	import { _, type ErrorKeys } from '$lib/i18n/index.svelte';
	import { Label } from 'bits-ui';
	import clsx from 'clsx';
	import type { HTMLInputAttributes } from 'svelte/elements';

	interface FormInputProps {
		label: string;
		errors?: ErrorKeys[];
		name: string;
		optional?: boolean;
	}

	type Props = FormInputProps & HTMLInputAttributes;

	let { label, errors, value = $bindable(), optional = false, ...rest }: Props = $props();

	const styles = $derived(
		clsx(
			'rounded border border-border bg-surface px-2 py-1 outline-primary',
			errors != undefined && errors.length > 0 && 'border-red-600'
		)
	);
</script>

<div class="flex flex-col">
	<Label.Root id={`${rest.name}-label`} for={rest.name}>
		{label}
		{#if optional}
		<span>({_("common.optional")})</span>
		{/if}
	</Label.Root>
	<input aria-labelledby={`${rest.name}-label`} bind:value class={styles} {...rest} />
	{#if errors}
		<ul>
			{#each errors as error, index (index)}
				<li class="text-sm leading-tight font-light text-red-600">{_(error)}</li>
			{/each}
		</ul>
	{/if}
</div>
