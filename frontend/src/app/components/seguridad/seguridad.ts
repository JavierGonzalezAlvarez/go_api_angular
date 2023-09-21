export interface userCredentials {
  username: string;
  email: string;
  password: string;
  role: string;
}

export interface responseAuthentication {
  token: string;
  expiracion: Date;
  user: string;
  role: string;
}

export interface user {
  id: string;
  email: string;
}
