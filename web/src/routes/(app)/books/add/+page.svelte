<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { ERROR_CODES } from '$lib/api/error';
	import type { ApiError, ApiValidationError } from '$lib/api/types';
	import Button from '$lib/components/shared/Button.svelte';
	import FormInput from '$lib/components/shared/FormInput.svelte';
	import { enhance } from '$lib/enhance';
	import { _, type ErrorKeys } from '$lib/i18n/index.svelte';
	import { Plus, Trash } from '@lucide/svelte';
	import { Label } from 'bits-ui';

	let authors: string[] = $state([]);

	let errors: Record<string, ErrorKeys[]> = $state({});

	function handleError(error: ApiError) {
		if (error.code === ERROR_CODES.INVALID_INPUT) {
			errors = Object.fromEntries(
				Object.entries(error.details!).map((err) => [
					err[0],
					(err[1] as ApiValidationError[]).map((e) => e.code)
				])
			);
		}
	}
</script>

<div class="p-4">
	<h1 class="text-2xl font-bold">{_('books.add_book')}</h1>
	<form
		use:enhance={{
			route: 'books.create',
			onError: handleError,
			onSuccess: () => goto(resolve('/(app)/books'))
		}}
		class="flex flex-col gap-1"
	>
		<FormInput label={_('books.attributes.title')} name="title" required aria-required errors={errors['title']} />
		<div class="flex flex-col gap-1">
			<p>{_('books.attributes.authors')} (min. 1 required)</p>
			{#each authors as author, index (index)}
				<div class="flex flex-row gap-2">
					<!-- TODO: extract input from forminput -->
					<input
						type="text"
						name="authors[]"
						bind:value={authors[index]}
						required
						aria-required="true"
						class="flex-1 rounded border border-border bg-surface px-2 py-1 outline-primary"
					/>
					<Button
						onclick={() => authors.splice(index, 1)}
						secondary
						class="flex size-8! items-center justify-center"
					>
						<Trash size={16} />
					</Button>
				</div>
			{/each}
			<!-- TODO: make button centering into class -->
			<Button
				onclick={() => authors.push('')}
				secondary
				type="button"
				class="flex flex-row items-center justify-center gap-1"
			>
				<Plus size={16} />{_('actions.add')}
			</Button>
			{#if errors['authors']}
				<ul>
					{#each errors['authors'] as error, index (index)}
						<li class="text-sm leading-tight font-light text-red-600">{_(error)}</li>
					{/each}
				</ul>
			{/if}
		</div>
		<FormInput
			label={_('books.attributes.pages')}
			type="number"
			name="pages"
			required
			min="1"
			errors={errors['pages']}
		/>
		<FormInput label={_('books.attributes.isbn')} optional name="isbn" />
		<FormInput
			label={_('books.attributes.published_year')}
			optional
			type="number"
			name="published_year"
		/>
		<Label.Root id="description_label" for="description">Description</Label.Root>
		<textarea
			name="description"
			id="description"
			aria-labelledby="description_label"
			class="flex-1 rounded border border-border bg-surface px-2 py-1 outline-primary"
			rows="3"></textarea>
		<Button>
			{_('actions.add')}
		</Button>
	</form>
</div>
