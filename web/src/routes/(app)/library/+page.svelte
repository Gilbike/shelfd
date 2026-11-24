<script lang="ts">
	import BookCard from '$lib/components/books/BookCard.svelte';
	import Title from '$lib/components/shared/Title.svelte';
	import { _ } from '$lib/i18n/index.svelte';
	import {
		BookBookmark,
		BookDashed,
		BookExclamationPoint,
		BookOpen,
		BookOpenCheck
	} from '@lucide/svelte';
	import { Tabs } from 'bits-ui';
	import type { PageProps } from './$types';

	const tabs = ['want_to_read', 'currently_reading', 'completed', 'backlog'];

	const { data }: PageProps = $props();
</script>

<Title text={_('profile.library')} />
<div class="flex h-full flex-col p-4">
	<h1 class="mb-2 text-2xl font-bold">{_('profile.library')} ({data.metadata.totalBooks})</h1>
	<Tabs.Root value="want_to_read">
		<Tabs.List class="relative mb-2 flex w-full max-w-full flex-row gap-2 overflow-x-auto">
			{#each tabs as tab (tab)}
				<Tabs.Trigger
					value={tab}
					class="flex shrink-0 flex-row items-center gap-2 rounded border-b-2 border-b-transparent px-4 py-2 data-[state=active]:bg-secondary"
				>
					{@render collectionBadge(tab)}
				</Tabs.Trigger>
			{/each}
		</Tabs.List>
		{#each tabs as tab (tab)}
			<Tabs.Content value={tab}>
				<div
					class="grid min-h-0 flex-1 auto-rows-max grid-cols-1 gap-2 overflow-y-auto sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5"
				>
					{#each data.books.filter((x) => x.shelf === tab) as book (book.id)}
						<BookCard {...book} />
					{/each}
				</div>
			</Tabs.Content>
		{/each}
	</Tabs.Root>
</div>

{#snippet collectionBadge(collection: string)}
	{#if collection == 'want_to_read'}
		<BookBookmark size={16} />
		{_('books.shelfs.want_to_read')}
	{:else if collection == 'currently_reading'}
		<BookOpen size={16} />
		{_('books.shelfs.currently_reading')}
	{:else if collection == 'backlog'}
		<BookExclamationPoint size={16} />
		{_('books.shelfs.backlog')}
	{:else if collection == 'completed'}
		<BookOpenCheck size={16} />
		{_('books.shelfs.completed')}
	{:else if collection == 'not_added'}
		<BookDashed size={16} />
		{_('books.shelfs.not_added')}
	{/if}
{/snippet}
