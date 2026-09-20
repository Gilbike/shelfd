export type ApiError = {
	status: number;
	code: string;
	message: string;
	details?: Record<string, unknown>;
};

export type User = {
	id: number;
	username: string;
	display_name: string;
	created_at: string;
	updated_at: string;
};

export type AuthResponse = {
	user: User;
};

export type UserCurrentResponse = User;
