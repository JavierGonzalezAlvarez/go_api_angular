export interface userCredentials {
  email: string;
  password: string;
}

export interface responseAuthentication {
  token: string;
  expiracion: Date;
  user: string;
}

export interface user {
  id: string;
  email: string;
}
