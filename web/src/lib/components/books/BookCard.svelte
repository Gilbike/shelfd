<script lang="ts">
	import type { Book } from '$lib/api/types';
	import { _ } from '$lib/i18n/index.svelte';

	const { title, authors, pages, cover_url }: Book = $props();
</script>

<div
	class="basis-11/12 overflow-hidden rounded border border-border bg-surface sm:basis-1/3 md:basis-1/4 lg:basis-1/5 xl:basis-1/6"
>
	<div class="p-10">
		{#if cover_url === undefined}
			<div class="flex aspect-1/1.5 w-full items-center justify-center bg-primary">
				No cover image
			</div>
		{:else}
			<img alt={title} src={cover_url} class="aspect-1/1.5 w-full" />
		{/if}
	</div>
	<div class="border-t border-border p-2">
		<p>{title}</p>
		<p class="text-sm text-foreground/60">{authors.join(', ')}</p>
		<p>{_('books.pages', { pages })}</p>
	</div>
</div>
