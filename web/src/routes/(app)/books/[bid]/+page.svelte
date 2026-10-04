<script lang="ts">
	import { _ } from '$lib/i18n/index.svelte';
	import { EllipsisVertical } from '@lucide/svelte';
	import { DropdownMenu } from 'bits-ui';
	import { fly } from 'svelte/transition';
	import type { LayoutProps } from './$types';

	const { data }: LayoutProps = $props();
	const { title, cover_url, pages, authors, description, isbn, published_year } = $derived(
		data.book
	);
</script>

<div class="flex gap-4 p-4 not-md:flex-col">
	{#if cover_url == null}
		<div
			class="flex aspect-1/1.5 w-full items-center justify-center rounded bg-primary/60 not-md:mx-auto sm:min-w-1/2 md:min-w-1/3 lg:min-w-1/4 xl:min-w-1/5"
		>
			{_('books.no_cover')}
		</div>
	{:else}
		<img
			alt={title}
			src={cover_url}
			class="aspect-1/1.5 w-full rounded border border-border not-md:mx-auto sm:min-w-1/2 md:min-w-1/3 lg:min-w-1/4 xl:min-w-1/5"
		/>
	{/if}
	<div>
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
								<div {...props} class="w-48" transition:fly={{ duration: 160, opacity: 0, x: 20 }}>
									!!Book actions!!
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
