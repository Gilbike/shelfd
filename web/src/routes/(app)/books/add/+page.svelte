<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { api, apiRoute } from '$lib/api';
	import type { Book } from '$lib/api/types';
	import BookCoverForm from '$lib/components/books/BookCoverForm.svelte';
	import BookForm from '$lib/components/books/BookForm.svelte';
	import { _ } from '$lib/i18n/index.svelte';

	let uploadForm: HTMLFormElement | undefined = $state();
	let coverFile: File | undefined = $state();

	async function handleAddSuccess(book: Book) {
		if (!uploadForm) {
			goto(resolve('/(app)/books/[bid]', { bid: book.id.toString() }));
			return;
		}
		if (!coverFile) {
			goto(resolve('/(app)/books/[bid]', { bid: book.id.toString() }));
			return;
		}

		const body = new FormData();
		body.append('cover', coverFile);

		await api(apiRoute('books.cover.update', { id: book.id }), {
			body
		});

		goto(resolve('/(app)/books/[bid]', { bid: book.id.toString() }));
	}
</script>

<div class="p-4">
	<h1 class="text-2xl font-bold">{_('books.add_book')}</h1>
	<div class="flex gap-4 not-md:flex-col">
		<BookCoverForm bind:formRef={uploadForm} bind:file={coverFile} />
		<BookForm route={apiRoute('books.create')} onSuccess={handleAddSuccess} />
	</div>
</div>
