<script lang="ts">
	import { goto, invalidate } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { api, apiRoute } from '$lib/api';
	import { isApiError } from '$lib/api/error';
	import Button from '$lib/components/shared/Button.svelte';
	import Title from '$lib/components/shared/Title.svelte';
	import { _ } from '$lib/i18n/index.svelte';
	import {
		BookBookmark,
		BookDashed,
		BookExclamationPoint,
		BookOpen,
		BookOpenCheck,
		ChevronDown,
		EllipsisVertical,
		Pencil,
		Trash
	} from '@lucide/svelte';
	import { DropdownMenu } from 'bits-ui';
	import { fly } from 'svelte/transition';
	import type { LayoutProps } from './$types';

	const { data }: LayoutProps = $props();
	const { id, title, cover_url, pages, authors, description, isbn, published_year, collection } =
		$derived(data.book);

	let collectionName: string = $derived(collection !== undefined ? collection : 'not_added');

	async function handleDeleteClick() {
		const result = await api(apiRoute('books.delete', { id }));
		if (isApiError(result)) {
			// TODO: implement feedback
			return;
		}
		goto(resolve('/(app)/books'));
	}

	async function handleCollectionClick(collection: string) {
		const result = await api(apiRoute('books.collection.set', { id }), {
			body: JSON.stringify({ collection })
		});
		if (isApiError(result)) {
			// TODO: implment feedback
			return;
		}
		await invalidate('data:book');
	}
</script>

<Title text={title} />
<div class="flex gap-4 p-4 not-md:flex-col">
	{#if cover_url == null}
		<div
			class="flex aspect-1/1.5 w-full items-center justify-center rounded border border-dashed border-primary text-sm font-light text-foreground/60 not-md:mx-auto sm:max-w-1/2 sm:min-w-1/2 md:max-w-1/3 md:min-w-1/3 lg:max-w-1/4 lg:min-w-1/4 xl:max-w-1/5"
		>
			{_('books.no_cover')}
		</div>
	{:else}
		<img
			alt={title}
			src={cover_url}
			class="w-full rounded border border-border not-md:mx-auto sm:max-w-1/2 sm:min-w-1/2 md:max-w-1/3 md:min-w-1/3 lg:max-w-1/4 lg:min-w-1/4 xl:max-w-1/5"
		/>
	{/if}
	<div class="flex-1">
		<div class="flex flex-row items-center justify-between">
			<div>
				<h1 class="text-2xl font-bold">{title}</h1>
				<p class="text-sm text-foreground/60">{authors.join(', ')}</p>
			</div>
			<DropdownMenu.Root>
				<DropdownMenu.Trigger class="cursor-pointer rounded p-1">
					<EllipsisVertical />
				</DropdownMenu.Trigger>

				<DropdownMenu.Content side="left" align="start" sideOffset={8} forceMount>
					{#snippet child({ wrapperProps, props, open })}
						{#if open}
							<div {...wrapperProps}>
								<div
									{...props}
									class="w-48 *:not-last:mb-1"
									transition:fly={{ duration: 160, opacity: 0, x: 20 }}
								>
									<DropdownMenu.Item>
										{#snippet child({ props })}
											<a
												{...props}
												href={resolve('/(app)/books/[bid]/edit', { bid: id.toString() })}
												class="interactive-secondary flex flex-row items-center gap-1 rounded p-1 px-2"
											>
												<Pencil size={16} />
												{_('actions.edit')}
											</a>
										{/snippet}
									</DropdownMenu.Item>
									<DropdownMenu.Item onSelect={handleDeleteClick}>
										{#snippet child({ props })}
											<Button {...props} secondary class="flex flex-row items-center gap-1 px-2">
												<Trash size={16} />
												{_('actions.delete')}
											</Button>
										{/snippet}
									</DropdownMenu.Item>
								</div>
							</div>
						{/if}
					{/snippet}
				</DropdownMenu.Content>
			</DropdownMenu.Root>
		</div>
		<div
			class="items-centerrounded mt-2 flex w-fit flex-row overflow-hidden rounded border border-border"
		>
			<p class="flex flex-row items-center gap-2 border-r border-r-border px-2 py-2">
				{@render collectionBadge(collectionName)}
			</p>
			<DropdownMenu.Root>
				<DropdownMenu.Trigger class="interactive-secondary cursor-pointer px-2">
					<ChevronDown />
				</DropdownMenu.Trigger>

				<DropdownMenu.Content side="bottom" align="end" sideOffset={8} forceMount>
					{#snippet child({ wrapperProps, props, open })}
						{#if open}
							<div {...wrapperProps}>
								<div
									{...props}
									class="w-48 *:not-last:mb-1"
									transition:fly={{ duration: 160, opacity: 0, y: -20 }}
								>
									<DropdownMenu.Item
										class="interactive-secondary flex cursor-pointer flex-row items-center gap-1 rounded p-1 px-2"
										onSelect={() => handleCollectionClick('want_to_read')}
									>
										{@render collectionBadge('want_to_read')}
									</DropdownMenu.Item>
									<DropdownMenu.Item
										class="interactive-secondary flex cursor-pointer flex-row items-center gap-1 rounded p-1 px-2"
										onSelect={() => handleCollectionClick('currently_reading')}
									>
										{@render collectionBadge('currently_reading')}
									</DropdownMenu.Item>
									<DropdownMenu.Item
										class="interactive-secondary flex cursor-pointer flex-row items-center gap-1 rounded p-1 px-2"
										onSelect={() => handleCollectionClick('backlog')}
									>
										{@render collectionBadge('backlog')}
									</DropdownMenu.Item>
									<DropdownMenu.Item
										class="interactive-secondary flex cursor-pointer flex-row items-center gap-1 rounded p-1 px-2"
										onSelect={() => handleCollectionClick('completed')}
									>
										{@render collectionBadge('completed')}
									</DropdownMenu.Item>
								</div>
							</div>
						{/if}
					{/snippet}
				</DropdownMenu.Content>
			</DropdownMenu.Root>
		</div>
		{#if description}
			<p class="my-4 text-justify">{description}</p>
		{/if}
		<p>{_('books.pages', { pages })}</p>
		{#if isbn}
			<p><span class="font-semibold">ISBN:</span> {isbn}</p>
		{/if}
		{#if published_year}
			<p>
				<span class="font-semibold">{_('books.attributes.published_year')}:</span>
				{published_year}
			</p>
		{/if}
	</div>
</div>

{#snippet collectionBadge(collection: string)}
	{#if collection == 'want_to_read'}
		<BookBookmark size={16} />
		{_('books.collections.want_to_read')}
	{:else if collection == 'currently_reading'}
		<BookOpen size={16} />
		{_('books.collections.currently_reading')}
	{:else if collection == 'backlog'}
		<BookExclamationPoint size={16} />
		{_('books.collections.backlog')}
	{:else if collection == 'completed'}
		<BookOpenCheck size={16} />
		{_('books.collections.completed')}
	{:else if collection == 'not_added'}
		<BookDashed size={16} />
		{_('books.collections.not_added')}
	{/if}
{/snippet}
