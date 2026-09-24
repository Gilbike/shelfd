<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { ERROR_CODES } from '$lib/api/error';
	import type { ApiError } from '$lib/api/types';
	import Button from '$lib/components/shared/Button.svelte';
	import FormInput from '$lib/components/shared/FormInput.svelte';
	import { enhance } from '$lib/enhance';
	import { _ } from '$lib/i18n/index.svelte';

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

<div class="panel mx-auto w-11/12 self-center sm:w-8/12 md:w-6/12 lg:w-1/4">
	<h1 class="font-semibold">{_('auth.login')}</h1>
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
			label={_("auth.username")}
			type="text"
			placeholder={_("auth.username")}
			autocomplete="username"
			errors={isCredentialsOk ? undefined : ['errors.invalid_credentials']}
		/>
		<FormInput
			id="password"
			name="password"
			label={_("auth.password")}
			type="password"
			placeholder={_("auth.password")}
			autocomplete="current-password"
			errors={isCredentialsOk ? undefined : ['errors.invalid_credentials']}
		/>
		<Button type="submit">{_('auth.login')}</Button>
	</form>
</div>
