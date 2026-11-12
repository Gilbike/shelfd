<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { api, apiRoute } from '$lib/api';
	import type { Book } from '$lib/api/types';
	import BookForm from '$lib/components/books/BookForm.svelte';
	import Button from '$lib/components/shared/Button.svelte';
	import { _ } from '$lib/i18n/index.svelte';
	import { ImageUp } from '@lucide/svelte';

	let uploadForm: HTMLFormElement | undefined = $state();
	let fileInput: HTMLInputElement | undefined = $state();
	let coverFile: File | undefined = $state();
	let previewUrl: string | null = $state(null);

	let dropArea: HTMLDivElement | undefined = $state();

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

	async function handleFileSelect(
		event: Event & { currentTarget: EventTarget & HTMLInputElement }
	) {
		const file = event.currentTarget.files?.[0];
		if (!file) return;

		if (previewUrl) {
			URL.revokeObjectURL(previewUrl);
		}

		previewUrl = URL.createObjectURL(file);
		coverFile = file;
	}

	function handleUploadClick() {
		fileInput?.click();
	}

	function handleDrop(e: DragEvent) {
		e.preventDefault();
		if (e.dataTransfer == null) return;
		const files = [...e.dataTransfer.items].map((item) => item.getAsFile()).filter((file) => file);
		if (files.length == 0) return;
		if (files[0] == null) return;

		if (previewUrl) {
			URL.revokeObjectURL(previewUrl);
		}

		previewUrl = URL.createObjectURL(files[0]);
		coverFile = files[0];
	}

	function handleDragOver(e: DragEvent) {
		if (e.dataTransfer == null) return;
		e.preventDefault();
		e.dataTransfer.dropEffect = 'copy';
	}

	$effect(() => {
		dropArea?.addEventListener('drop', handleDrop);
		dropArea?.addEventListener('dragover', handleDragOver);

		return () => {
			dropArea?.removeEventListener('drop', handleDrop);
			dropArea?.removeEventListener('dragover', handleDragOver);
		};
	});

	$effect(() => {
		return () => {
			if (previewUrl) {
				URL.revokeObjectURL(previewUrl);
			}
		};
	});
</script>

<div class="p-4">
	<h1 class="text-2xl font-bold">{_('books.add_book')}</h1>
	<div class="flex gap-4 not-md:flex-col">
		<form
			bind:this={uploadForm}
			class="w-full not-md:mx-auto sm:max-w-1/3 sm:min-w-1/3 md:max-w-1/4 md:min-w-1/4 lg:max-w-1/5 lg:min-w-1/5 xl:max-w-1/6"
			onsubmit={(e) => e.preventDefault()}
		>
			<div
				bind:this={dropArea}
				class="flex aspect-1/1.5 flex-col items-center justify-center gap-2 rounded border border-dashed border-primary p-2 not-md:mx-auto"
			>
				{#if previewUrl == null}
					<ImageUp class="mb-2 text-foreground/60" size={32} />
					<Button onclick={handleUploadClick} secondary class="w-fit! px-2">Upload image</Button>
					<p class="text-sm text-foreground/60">Or drop the image here</p>
				{:else}
					<img
						alt="cover"
						src={previewUrl}
						class="w-full rounded border border-border not-md:mx-auto"
					/>
					<Button onclick={handleUploadClick} secondary class=" px-2">Change image</Button>
				{/if}
			</div>

			<input
				type="file"
				class="hidden"
				name="cover"
				id="cover"
				onchange={handleFileSelect}
				bind:this={fileInput}
			/>
		</form>
		<BookForm route={apiRoute('books.create')} onSuccess={handleAddSuccess} />
	</div>
</div>
