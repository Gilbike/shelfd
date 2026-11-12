<script lang="ts">
	import { ImageUp } from '@lucide/svelte';
	import { untrack } from 'svelte';
	import Button from '../shared/Button.svelte';

	interface BookCoverFormProps {
		formRef: HTMLFormElement | undefined;
		file: File | undefined;
	}

	let fileInput: HTMLInputElement | undefined = $state();
	let previewUrl: string | null = $state(null);

	let dropArea: HTMLDivElement | undefined = $state();

	let { formRef = $bindable(), file = $bindable() }: BookCoverFormProps = $props();

	async function handleFileSelect(
		event: Event & { currentTarget: EventTarget & HTMLInputElement }
	) {
		const selectedFile = event.currentTarget.files?.[0];
		if (!selectedFile) return;

		file = selectedFile;
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

		file = files[0];
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
		let changedFile = file;

		untrack(() => {
			if (previewUrl) {
				URL.revokeObjectURL(previewUrl);
			}

			if (changedFile) {
				previewUrl = URL.createObjectURL(changedFile);
			}
		});
	});

	$effect(() => {
		return () => {
			if (previewUrl) {
				URL.revokeObjectURL(previewUrl);
			}
		};
	});
</script>

<form
	bind:this={formRef}
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
