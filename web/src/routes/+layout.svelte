<script lang="ts">
	import './layout.css';

	let { children } = $props();
</script>

<svelte:head>
	<title>Shelfd</title>
</svelte:head>

{@render children()}
