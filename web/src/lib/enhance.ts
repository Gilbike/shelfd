// src/lib/actions/enhance.ts
import { api } from '$lib/api';
import type { Route } from './api';
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
		const method = (node.method || 'GET').toUpperCase();

		options.onPending?.();

		try {
			const data = await api(options.route, {
				method: method,
				body: method === 'GET' ? undefined : JSON.stringify(payload)
			});
			options.onSuccess?.(data);
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

		// Típus szerinti konverzió
		if (input.type === 'number' || input.type === 'range') {
			data[input.name] = input.value === '' ? null : Number(input.value);
		} else if (input.type === 'checkbox') {
			data[input.name] = input.checked;
		} else if (input.type === 'radio') {
			if (input.checked) data[input.name] = input.value;
		} else {
			// Sima szöveg, email, dátum, select, textarea
			data[input.name] = input.value;
		}
	}

	return data;
}
