export interface User {
  email: string;
  name: string;
  token: string;
}

export function getToken(): string | null;
