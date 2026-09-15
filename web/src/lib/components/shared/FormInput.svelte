<script lang="ts">
	import { _, type ErrorKeys } from '$lib/i18n/index.svelte';
	import { Label } from 'bits-ui';
	import type { HTMLInputAttributes } from 'svelte/elements';

	interface FormInputProps {
		label: string;
		errors?: ErrorKeys[];
		name: string;
	}

	type Props = FormInputProps & HTMLInputAttributes;

	let { label, errors, value = $bindable(), ...rest }: Props = $props();
</script>

<div class="flex flex-col">
	<Label.Root id={`${rest.name}-label`} for={rest.name}>{label}</Label.Root>
	<input
		aria-labelledby={`${rest.name}-label`}
		bind:value
		class="rounded border border-border bg-background px-2 py-1 outline-primary"
		{...rest}
	/>
	{#if errors}
		<ul>
			{#each errors as error, index (index)}
				<li class="text-sm leading-tight font-light text-red-600">{_(error)}</li>
			{/each}
		</ul>
	{/if}
</div>
