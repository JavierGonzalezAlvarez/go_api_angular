export interface userCredentials {
  email: string;
  password: string;
}

export interface responseAuthentication {
  token: string;
  expiracion: Date;
}

export interface user {
  id: string;
  email: string;
}
