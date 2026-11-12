<script lang="ts">
	import { goto, invalidateAll } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { api, apiRoute } from '$lib/api';
	import type { Book } from '$lib/api/types';
	import BookCoverForm from '$lib/components/books/BookCoverForm.svelte';
	import BookForm from '$lib/components/books/BookForm.svelte';
	import { _ } from '$lib/i18n/index.svelte';
	import { onMount } from 'svelte';
	import type { LayoutProps } from '../$types';

	let originalFile: File | undefined;
	let uploadForm: HTMLFormElement | undefined = $state();
	let coverFile: File | undefined = $state.raw();

	const { data }: LayoutProps = $props();

	async function handleSuccess(_book: Book) {
		const isOriginalFile = coverFile === originalFile;
		const coverNeedsUpdate = Boolean(uploadForm && coverFile && !isOriginalFile);

		if (coverNeedsUpdate && coverFile !== undefined) {
			const body = new FormData();
			body.append('cover', coverFile);

			await api(apiRoute('books.cover.update', { id: data.book.id }), {
				body
			});
		}

		await invalidateAll(); // TODO: find better way, it wasnt working with invalidate() as intended
		goto(resolve('/(app)/books/[bid]', { bid: data.book.id.toString() }));
	}

	onMount(async () => {
		if (data.book.cover_url == undefined) return;
		const response = await fetch(data.book.cover_url);
		if (!response.ok) return;

		const imageData = await response.blob();
		const image = new File([imageData], `b${data.book.id}`);
		originalFile = image;
		coverFile = image;
	});
</script>

<div class="p-4">
	<h1 class="text-2xl font-bold">
		{_('actions.edit')} <span class="font-normal">{data.book.title}</span>
	</h1>
	<div class="flex gap-4 not-md:flex-col">
		<BookCoverForm bind:formRef={uploadForm} bind:file={coverFile} />
		<BookForm
			buttonLabel={_('actions.save')}
			route={apiRoute('books.update', { id: data.book.id })}
			base={data.book}
			onSuccess={handleSuccess}
		/>
	</div>
</div>
