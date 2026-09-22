// src/lib/actions/enhance.ts
import { api } from '$lib/api';
import type { Route } from './api';
import { isApiError } from './api/error';
import type { ApiError } from './api/types';

type EnhanceOptions = {
	route: Route;
	onPending?: () => void;
	onSuccess?: (data: unknown) => void;
	onError?: (err: ApiError) => void;
};

export function enhance(node: HTMLFormElement, options: EnhanceOptions) {
	const handleSubmit = async (e: SubmitEvent) => {
		e.preventDefault();

		const payload = serializeForm(node);

		options.onPending?.();

		try {
			const data = await api(options.route, {
				body: JSON.stringify(payload)
			});
			if (isApiError(data)) {
				options.onError?.(data as ApiError);
			} else {
				options.onSuccess?.(data);
			}
		} catch (err) {
			options.onError?.(err as ApiError);
		}
	};

	node.addEventListener('submit', handleSubmit);

	return {
		destroy() {
			node.removeEventListener('submit', handleSubmit);
		}
	};
}

function serializeForm(form: HTMLFormElement): Record<string, unknown> {
	const data: Record<string, unknown> = {};

	for (const element of form.elements) {
		const input = element as HTMLInputElement;
		if (!input.name || input.disabled) continue;

		let value;

		if (input.type === 'number' || input.type === 'range') {
			value = input.value === '' ? null : Number(input.value);
		} else if (input.type === 'checkbox') {
			value = input.checked;
		} else if (input.type === 'radio') {
			if (input.checked) value = input.value;
		} else if (input.type === 'text' || input.type === 'password') {
			if (input.value === '') value = null;
			else value = input.value;
		} else {
			value = input.value;
		}

		if (input.name.endsWith('[]')) {
			const key = input.name.slice(0, -2);
			if (!(key in data)) {
				data[key] = [];
			}
			(data[key] as unknown[]).push(value);
		} else {
			data[input.name] = value;
		}
	}

	return data;
}
