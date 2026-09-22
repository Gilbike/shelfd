<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import BookCard from '$lib/components/books/BookCard.svelte';
	import Button from '$lib/components/shared/Button.svelte';
	import { _ } from '$lib/i18n/index.svelte.js';
	import { ChevronLeft, ChevronRight, Plus } from '@lucide/svelte';
	import { Pagination } from 'bits-ui';

	const { data } = $props();

	function handlePageChange(page: number) {
		goto(resolve(`/(app)/books?page=${page}`), {
			replaceState: true,
			keepFocus: true,
			noScroll: true
		});
	}

	const firstIndex = $derived((data.metadata.page - 1) * 20 + 1);
	const lastIndex = $derived(data.metadata.page * 20);
</script>

<div class="flex h-dvh flex-col overflow-hidden p-4">
	<div class="mb-4 flex flex-row items-center justify-between">
		<div>
			<h1 class="text-2xl font-bold">{_('books.books')} ({data.metadata.totalBooks})</h1>
			<p class="text-sm text-foreground/60">
				Showing {firstIndex} - {Math.min(lastIndex, data.metadata.totalBooks)}
			</p>
		</div>
		<Button class="flex h-fit w-fit! flex-row items-center gap-1 px-2 py-1">
			<Plus size={16} /> Add
		</Button>
	</div>
	<div
		class="grid min-h-0 flex-1 auto-rows-max grid-cols-1 gap-2 overflow-y-auto sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5"
	>
		{#each data.books as book (book.id)}
			<BookCard {...book} />
		{/each}
	</div>
	<Pagination.Root
		page={data.metadata.page}
		class="mt-4 self-center"
		count={data.metadata.totalPages}
		onPageChange={handlePageChange}
	>
		{#snippet children({ pages })}
			<div class="flex items-center">
				<Pagination.PrevButton
					class="text-foreground/75 hover:text-foreground/90 disabled:text-foreground/25"
				>
					<ChevronLeft />
				</Pagination.PrevButton>
				<div class="mx-4 flex items-center gap-2">
					{#each pages as page (page.key)}
						{#if page.type === 'ellipsis'}
							<div class="text-foreground-alt text-[15px] font-medium select-none">...</div>
						{:else}
							<Pagination.Page
								{page}
								class="inline-flex size-8 items-center justify-center rounded bg-transparent text-[15px] select-none hover:bg-primary/25 active:scale-[0.98] disabled:cursor-not-allowed disabled:opacity-50 hover:disabled:bg-transparent data-selected:bg-primary data-selected:text-background"
							>
								{page.value}
							</Pagination.Page>
						{/if}
					{/each}
				</div>
				<Pagination.NextButton
					class="text-foreground/75 hover:text-foreground/90 disabled:text-foreground/25"
				>
					<ChevronRight />
				</Pagination.NextButton>
			</div>
		{/snippet}
	</Pagination.Root>
</div>
