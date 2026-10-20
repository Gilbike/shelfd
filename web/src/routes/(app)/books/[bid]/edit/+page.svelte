<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { apiRoute } from '$lib/api';
	import type { Book } from '$lib/api/types';
	import BookForm from '$lib/components/books/BookForm.svelte';
	import { _ } from '$lib/i18n/index.svelte';
	import type { LayoutProps } from '../$types';

	const { data }: LayoutProps = $props();

	function handleSuccess(_book: Book) {
		goto(resolve('/(app)/books/[bid]', { bid: data.book.id.toString() }));
	}
</script>

<div class="p-4">
	<BookForm
		buttonLabel={_('actions.edit')}
		route={apiRoute('books.update', { id: data.book.id })}
		base={data.book}
		onSuccess={handleSuccess}
	/>
</div>
