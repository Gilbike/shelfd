<script>
	import Button from '$lib/components/shared/Button.svelte';
	import FormInput from '$lib/components/shared/FormInput.svelte';
	import { _ } from '$lib/i18n/index.svelte';
</script>

<div class="panel flex w-11/12 flex-col gap-1 sm:w-8/12 md:w-6/12 lg:w-1/4">
	<h1 class="font-semibold">{_('auth.login')}</h1>
	<FormInput id="username" name="username" label="Username" type="text" />
	<FormInput id="password" name="password" label="Password" type="password" />
	<Button>{_('auth.login')}</Button>
</div>

<style lang="postcss">
	:global(body) {
		@apply flex items-center justify-center;
	}
</style>
