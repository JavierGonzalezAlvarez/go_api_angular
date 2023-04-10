export interface Invoices {
    idheader: number;
    companyname: string | null;
    address: string | null;
}

export interface serverResponse  {
    count: number;
    invoices: Invoices[]
  };
  