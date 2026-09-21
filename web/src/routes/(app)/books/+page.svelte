<script lang="ts">
	import BookCard from '$lib/components/books/BookCard.svelte';

	const { data } = $props();
</script>

<div class="max-h-screen flex-1 overflow-hidden">
	<div class="flex max-h-screen flex-row flex-wrap justify-evenly gap-2 overflow-auto p-4">
		{#each data.books as book (book.id)}
			<BookCard {...book} />
		{/each}
	</div>
</div>
