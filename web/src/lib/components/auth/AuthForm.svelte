<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { ERROR_CODES } from '$lib/api/error';
	import type { ApiError } from '$lib/api/types';
	import { enhance } from '$lib/enhance';
	import { _ } from '$lib/i18n/index.svelte';
	import Button from '../shared/Button.svelte';
	import FormInput from '../shared/FormInput.svelte';

	let isCredentialsOk = $state(true);

	function handleError(error: ApiError) {
		if (error.code === ERROR_CODES.INVALID_CREDENTIALS) {
			isCredentialsOk = false;
		}
	}

	function handleSuccess() {
		goto(resolve('/(app)/books'));
	}
</script>

<h1 class="font-semibold">{_('auth.signin')}</h1>
<form
	use:enhance={{
		route: 'user.auth',
		onError: handleError,
		onSuccess: () => handleSuccess()
	}}
	class="flex flex-col gap-1"
>
	<FormInput
		id="username"
		name="username"
		label={_('auth.username')}
		type="text"
		placeholder={_('auth.username')}
		autocomplete="username"
		required
		aria-required
		errors={isCredentialsOk ? undefined : ['errors.invalid_credentials']}
	/>
	<FormInput
		id="password"
		name="password"
		label={_('auth.password')}
		type="password"
		placeholder={_('auth.password')}
		autocomplete="current-password"
		required
		aria-required
		errors={isCredentialsOk ? undefined : ['errors.invalid_credentials']}
	/>
	<Button type="submit">{_('auth.signin')}</Button>
</form>
