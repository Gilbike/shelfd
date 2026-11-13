<script lang="ts">
	interface TitleProps {
		text?: string;
	}

	const { text }: TitleProps = $props();
</script>

<svelte:head>
	{#if text !== undefined}
		<title>{text} - Shelfd</title>
	{:else}
		<title>Shelfd</title>
	{/if}
</svelte:head>
