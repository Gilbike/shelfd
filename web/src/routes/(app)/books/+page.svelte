<script lang="ts">
	const { data } = $props();
</script>

<div class="max-h-screen flex-1 overflow-hidden">
	<div class="max-h-screen overflow-auto">
		{#each data.books as book (book.id)}
			<div>
				<div>{book.title} by {book.authors.join(', ')}</div>
				<div>{book.pages} pages</div>
			</div>
		{/each}
	</div>
</div>
