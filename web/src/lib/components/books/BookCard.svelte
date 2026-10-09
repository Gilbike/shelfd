<script lang="ts">
	import { resolve } from '$app/paths';
	import type { Book } from '$lib/api/types';
	import { _ } from '$lib/i18n/index.svelte';

	const { id, title, authors, pages, cover_url }: Book = $props();
</script>

<a
	href={resolve('/(app)/books/[bid]', { bid: id.toString() })}
	class="flex basis-11/12 flex-col overflow-hidden rounded border border-border bg-surface sm:basis-1/3 md:basis-1/4 lg:basis-1/5 xl:basis-1/6"
>
	<div class="flex-1 bg-secondary p-10">
		<div class="flex aspect-1/1.5 w-full items-center justify-center">
			{#if cover_url == null}
				<div
					class="flex h-full w-full items-center justify-center rounded border border-dashed border-primary text-sm font-light text-foreground/60"
				>
					{_('books.no_cover')}
				</div>
			{:else}
				<img alt={title} src={cover_url} class=" w-full rounded border border-border" />
			{/if}
		</div>
	</div>
	<div class="h-fit border-t border-border p-2">
		<p>{title}</p>
		<p class="text-sm text-foreground/60">{authors.join(', ')}</p>
		<p>{_('books.pages', { pages })}</p>
	</div>
</a>
