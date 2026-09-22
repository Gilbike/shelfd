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

export type Book = {
	id: number;
	title: string;
	authors: string[];
	pages: number;
	isbn?: string;
	cover_url?: string;
	published_year?: number;
	description?: string;
	created_at: string;
	updated_at: string;
};

export type AuthResponse = {
	user: User;
};

export type UserCurrentResponse = User;

export type BookListResponse = {
	page: number;
	per_page: number;
	total_pages: number;
	total_books: number;
	data: Book[];
};
